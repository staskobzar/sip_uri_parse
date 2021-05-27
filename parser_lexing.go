package uri

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Token uint8

const (
	tUnknown Token = iota
	tSip
	tSips
	tUserinfo
	tHost
	tPort
	tParams
	tHeader
	tEOF
)

const eof = -1

type Item struct {
	token Token
	value string
}

type lexer struct {
	input  string
	limit  int
	cursor int
	marker int
	items  chan Item
	err    chan string
}

type lexFunc func() lexFunc

func LexerParse(data string) (*URI, error) {
	uri := &URI{}

	l := newLexer(data)
	go l.run()
mainLoop:
	for {
		select {
		case item := <-l.items:
			if item.token == tEOF {
				break mainLoop
			}
			uri.setSegment(item)
		case err := <-l.err:
			return nil, fmt.Errorf("Invalid uri '%s': %s", data, err)
		case <-time.After(time.Millisecond * 10):
			return nil, fmt.Errorf("Invalid uri '%s': timeout parsing", data)
		}
	}
	return uri, nil
}

func (uri *URI) setSegment(item Item) {
	switch item.token {
	case tSip:
		uri.scheme = SIP
	case tSips:
		uri.scheme = SIPS
	case tUserinfo:
		uri.userinfo = item.value
	case tHost:
		uri.hostport = item.value
	case tPort:
		uri.hostport = fmt.Sprintf("%s:%s", uri.hostport, item.value)
	case tParams:
		uri.params = item.value
	case tHeader:
		uri.headers = item.value
	}
}

func newLexer(data string) *lexer {
	return &lexer{
		input:  data,
		limit:  len(data),
		cursor: 0,
		marker: 0,
		items:  make(chan Item),
		err:    make(chan string),
	}
}

func (l *lexer) run() {
	for state := l.lexScheme; state != nil; {
		state = state()
	}
}

func (l *lexer) emit(tk Token) {
	l.items <- Item{tk, l.input[l.marker:l.cursor]}
}

func (l *lexer) errorf(pattern string, v ...interface{}) {
	l.err <- fmt.Sprintf(pattern, v...)
}

func (l *lexer) next() int {
	if l.cursor >= l.limit {
		return eof
	}
	cur := l.cursor
	l.cursor++
	return int(l.input[cur])
}

func (l *lexer) current() int {
	if l.cursor >= l.limit {
		return eof
	}
	return int(l.input[l.cursor])
}

func (l *lexer) lexScheme() lexFunc {
	if strings.HasPrefix(l.input, "sip:") {
		l.cursor = 4
		l.emit(tSip)
	} else if strings.HasPrefix(l.input, "sips:") {
		l.cursor = 5
		l.emit(tSips)
	} else {
		l.errorf("invalid scheme")
		return nil
	}
	return l.lexUserinfo
}

// rfc3261  #25.1 Basic Rules
// userinfo         =  ( user / telephone-subscriber ) [ ":" password ] "@"
// user             =  1*( unreserved / escaped / user-unreserved )
// user-unreserved  =  "&" / "=" / "+" / "$" / "," / ";" / "?" / "/"
// password         =  *( unreserved / escaped / "&" / "=" / "+" / "$" / "," )
func (l *lexer) lexUserinfo() lexFunc {
	l.marker = l.cursor

	atIndex := strings.IndexByte(l.input, '@')
	if atIndex == -1 {
		// continue with host:port part
		return l.lexHostport
	}

	for ; l.cursor < atIndex; l.cursor++ {
		c := l.input[l.cursor]
		if isUnreserved(c) {
			continue
		}
		// user unreserved and escape and ':'
		switch c {
		case '&', '=', '+', '$', ',', ';', '?', '/', '%', ':':
			continue
		}
		l.errorf("invalid userinfo ...%s", l.input[l.cursor:])
		return nil
	}

	l.emit(tUserinfo)
	l.cursor++ // skip '@'
	return l.lexHostport
}

// rfc3261  #25.1 Basic Rules
// hostport         =  host [ ":" port ]
// host             =  hostname / IPv4address / IPv6reference
// hostname         =  *( domainlabel "." ) toplabel [ "." ]
// domainlabel      =  alphanum
//                     / alphanum *( alphanum / "-" ) alphanum
// toplabel         =  ALPHA / ALPHA *( alphanum / "-" ) alphanum
// IPv4address      =  1*3DIGIT "." 1*3DIGIT "." 1*3DIGIT "." 1*3DIGIT
// IPv6reference    =  "[" IPv6address "]"
// IPv6address      =  hexpart [ ":" IPv4address ]
// hexpart          =  hexseq / hexseq "::" [ hexseq ] / "::" [ hexseq ]
// hexseq           =  hex4 *( ":" hex4)
// hex4             =  1*4HEXDIG
// port             =  1*DIGIT
func (l *lexer) lexHostport() lexFunc {
	l.marker = l.cursor

	c := l.current()
	if c == eof {
		l.errorf("invalid host/port")
		return nil
	}

	if c == '[' {
		l.cursor++
		return l.lexIPv6
	}

	if length, match := parseIPv4(l.input[l.marker:]); match {
		l.cursor += length
		l.emit(tHost)
		return l.lexPort
	}

	if isAlphaNum(byte(c)) {
		return l.lexHostname
	}

	l.errorf("invalid host part")
	return nil
}

// hostname         =  *( domainlabel "." ) toplabel [ "." ]
// domainlabel      =  alphanum / alphanum *( alphanum / "-" ) alphanum
// toplabel         =  ALPHA / ALPHA *( alphanum / "-" ) alphanum
func (l *lexer) lexHostname() lexFunc {
	for _, c := range l.input[l.marker:] {
		if isAlphaNum(byte(c)) || c == '-' || c == '.' {
			l.cursor++
			continue
		}
		if c == ':' {
			l.emit(tHost)
			return l.lexPort
		}
		if c == ';' {
			l.emit(tHost)
			return l.lexParams
		}
		if c == '?' {
			l.emit(tHost)
			return l.lexHeaders
		}
		l.errorf("invalid hostname")
		return nil
	}
	l.emit(tHost)
	l.emit(tEOF)
	return nil
}

// IPv6reference  =  "[" IPv6address "]"
// IPv6address    =  hexpart [ ":" IPv4address ]
// hexpart        =  hexseq / hexseq "::" [ hexseq ] / "::" [ hexseq ]
// hexseq         =  hex4 *( ":" hex4)
// hex4           =  1*4HEXDIG
func (l *lexer) lexIPv6() lexFunc {
	for _, c := range l.input[l.cursor:] {
		l.cursor++
		if c == ']' {
			l.emit(tHost)
			return l.lexPort
		}
		if isHex(byte(c)) || c == ':' || c == '.' {
			continue
		}
		l.errorf("invalid host part ipv6 >%c<", c)
		return nil
	}
	l.errorf("invalid host part ipv6")
	return nil
}

func (l *lexer) lexPort() lexFunc {
	c := l.current()
	if c == eof {
		l.emit(tEOF)
		return nil
	}
	if c != ':' {
		return l.lexParams
	}
	l.cursor++
	l.marker = l.cursor
	n, c, ok := dtoi(l.input[l.marker:])
	if !ok || n > 0xFFFF {
		l.errorf("invalid port %s", l.input[l.marker:l.marker+c])
	}

	l.cursor += c
	l.emit(tPort)
	return l.lexParams
}

// paramchar  = [[\]/:&+$] | unreserved | escaped;
// uriparam   = paramchar+ ("=" paramchar+)?;
// params     = (";" uriparam)* >sm %prms;
func (l *lexer) lexParams() lexFunc {
	l.marker = l.cursor
	c := l.current()
	if c == eof {
		l.emit(tEOF)
		return nil
	}
	if c == '?' {
		return l.lexHeaders
	}
	if c != ';' {
		l.errorf("invalid params")
		return nil
	}
	l.cursor++
	l.marker = l.cursor
	for _, c := range l.input[l.cursor:] {
		if c == '?' {
			l.emit(tParams)
			return l.lexHeaders
		}
		l.cursor++
		if isUnreserved(byte(c)) || c == '=' || c == ';' || c == '%' {
			continue
		}
		switch c {
		case '[', ']', '/', ':', '&', '+', '$':
			continue
		}
		l.errorf("invalid param ...>%s", l.input[l.cursor:])
		return nil
	}
	l.emit(tParams)
	l.emit(tEOF)
	return nil
}

// hdrchar = [[\]/?:+$] | unreserved | escaped;
// header  = hdrchar+ "=" hdrchar*;
// headers  = "?" (header ("&" header)*) >sm %hdrs;
func (l *lexer) lexHeaders() lexFunc {
	l.marker = l.cursor
	c := l.current()
	if c == eof {
		l.emit(tEOF)
		return nil
	}
	if c != '?' {
		l.errorf("invalid headers")
		return nil
	}
	l.cursor++
	l.marker = l.cursor
	for _, c := range l.input[l.marker:] {
		l.cursor++
		if isUnreserved(byte(c)) || c == '=' || c == '%' || c == '&' {
			continue
		}
		switch c {
		case '[', ']', '/', '?', ':', '+', '$':
			continue
		}
		l.errorf("invalid header ...>%s", l.input[l.cursor:])
	}
	l.emit(tHeader)
	l.emit(tEOF)
	return nil
}

// IPv4address    =  1*3DIGIT "." 1*3DIGIT "." 1*3DIGIT "." 1*3DIGIT
// modified copy from go source net/ip.go
func parseIPv4(data string) (int, bool) {
	input := data
	l := 0
	for i := 0; i < 4; i++ {
		if len(input) == 0 {
			return 0, false
		}
		if i > 0 {
			if input[0] != '.' {
				return 0, false
			}
			input = input[1:]
			l++
		}
		n, c, ok := dtoi(input)
		if !ok || n > 0xFF {
			return 0, false
		}
		input = input[c:]
		l += c
	}
	return l, true
}

func isAlphaNum(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || isNum(c)
}

func isHex(c byte) bool {
	return ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F') || isNum(c)
}

func isNum(c byte) bool {
	return ('0' <= c && c <= '9')
}

func isUnreserved(c byte) bool {
	if isAlphaNum(c) {
		return true
	}
	switch c {
	case '-', '_', '.', '!', '~', '*', '\'', '(', ')':
		return true
	}
	return false
}

// copy from go source net/parser.go
// Decimal to integer.
// Returns number, characters consumed, success.
func dtoi(s string) (n int, i int, ok bool) {
	n = 0
	big := 0xFFFFF
	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		n = n*10 + int(s[i]-'0')
		if n >= big {
			return big, i, false
		}
	}
	if i == 0 {
		return 0, 0, false
	}
	return n, i, true
}

type URIRegex struct {
	input string
	uri   *URI
}

func DummyParser(data string) (*URI, error) {
	u := &URIRegex{uri: &URI{}, input: data}

	if err := u.parseScheme(); err != nil {
		return nil, err
	}

	if err := u.parseUserinfo(); err != nil {
		return nil, err
	}

	if err := u.parseHostport(); err != nil {
		return nil, err
	}

	if err := u.parseParams(); err != nil {
		return nil, err
	}

	if err := u.parseHeaders(); err != nil {
		return nil, err
	}

	return u.uri, nil
}

func (u *URIRegex) parseHeaders() error {
	if u.input[0] != '?' {
		return nil
	}
	u.uri.headers = u.input[1:]
	return nil
}

func (u *URIRegex) parseParams() error {
	if u.input[0] != ';' {
		return nil
	}
	if idx := strings.IndexByte(u.input, '?'); idx >= 0 {
		u.uri.params = u.input[1:idx]
		u.input = u.input[idx:]
		return nil
	}
	u.uri.params = u.input[1:]
	return nil
}

func (u *URIRegex) parseHostport() error {
	if idx := strings.IndexByte(u.input, ';'); idx >= 0 {
		u.uri.hostport = u.input[:idx]
		u.input = u.input[idx:]
		return nil
	}

	if idx := strings.IndexByte(u.input, '?'); idx >= 0 {
		u.uri.hostport = u.input[:idx]
		u.input = u.input[idx:]
		return nil
	}

	u.uri.hostport = u.input
	return nil
}

func (u *URIRegex) parseUserinfo() error {
	idx := strings.IndexByte(u.input, '@')
	if idx == -1 {
		return nil
	}
	usrinf := u.input[:idx]
	u.input = u.input[idx+1:]

	u.uri.userinfo = usrinf
	return nil
}

func (u *URIRegex) parseScheme() error {
	if strings.HasPrefix(u.input, "sip:") {
		u.uri.scheme = SIP
		u.input = u.input[4:]
		return nil
	}
	if strings.HasPrefix(u.input, "sips:") {
		u.uri.scheme = SIPS
		u.input = u.input[5:]
		return nil
	}
	return fmt.Errorf("Invalid URI scheme")
}

func RegexParse(input string) (*URI, error) {
	uri := &URI{}
	exp := "^(?P<scheme>sips?):" +
		"(:?(?P<userinfo>[^@]+)@)?" +
		"(?P<hostport>[^;?]+)" +
		"(:?;(?P<params>[^?]+))?" +
		"(:?[?](?P<headers>.*))?$"

	re, err := regexp.Compile(exp)
	if err != nil {
		return nil, err
	}
	matches := re.FindStringSubmatch(input)
	if len(matches) == 0 {
		return nil, fmt.Errorf("Invalid URI")
	}
	idx := re.SubexpIndex("scheme")
	if idx == -1 {
		return nil, fmt.Errorf("Failed to parse URI")
	}
	switch matches[idx] {
	case "sip":
		uri.scheme = SIP
	case "sips":
		uri.scheme = SIPS
	}

	if idx := re.SubexpIndex("userinfo"); idx > 0 {
		uri.userinfo = matches[idx]
	}
	idx = re.SubexpIndex("hostport")
	uri.hostport = matches[idx]

	if idx := re.SubexpIndex("params"); idx > 0 {
		uri.params = matches[idx]
	}

	if idx := re.SubexpIndex("headers"); idx > 0 {
		uri.headers = matches[idx]
	}
	return uri, nil
}
