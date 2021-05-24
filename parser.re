package uri

import (
	"fmt"
)

// Re2GoParse sip URI
func Re2GoParse(str string) (*URI, error) {
	var cursor, marker int
	limit := len(str)
	var ts, te int
	/*!stags:re2c format = 'var @@ int'; separator = "\n\t"; */
	var parseError error

	err := func(msg string) { parseError = fmt.Errorf("Invalid URI '%s': %s", str, msg) }
	peek := func(str string, cursor, limit int) byte {
		if cursor >= limit {
			return 0
		}
		return str[cursor]
	}

	uri := &URI{}
	/*!re2c
	re2c:flags:tags         = 1;
	re2c:yyfill:enable      = 0;
	re2c:eof                = 0;
	re2c:define:YYCTYPE     = byte;
	re2c:define:YYLESSTHAN  = "cursor >= limit";
	re2c:define:YYPEEK      = "peek(str, cursor, limit)";
	re2c:define:YYSKIP      = "cursor += 1";
	re2c:define:YYSHIFT     = "cursor += @@{shift}";
	re2c:define:YYBACKUP    = "marker = cursor";
	re2c:define:YYRESTORE   = "cursor = marker";
	re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";
	re2c:define:YYSTAGP     = "@@{tag} = cursor";
	re2c:define:YYSTAGN     = "@@{tag} = -1";

	alpha	    	= [a-zA-Z];
	digit	    	= [0-9];
	alphanum	  = alpha | digit;
	hexdig      = [0-9a-fA-F];

	unreserved  = alphanum | [-_.!~*'()];
	escaped	    = "%" hexdig{2};
	user_unreserved = [&=+$,;?/];

	domainlabel = alphanum | (alphanum ( alphanum | "-" )* alphanum);
	toplabel	  = alpha | (alpha ( alphanum | "-" )* alphanum);
	hostname	  = (domainlabel ".")* toplabel "."?;
	ipv4addr	  = digit{1,3} "." digit{1,3} "." digit{1,3} "." digit{1,3};
	hexseq	    = hexdig{1,4} (":" hexdig{1,4})*;
	hexpart     = hexseq | (hexseq "::" hexseq?) | ("::" hexseq?);
	ipv6addr	  = hexpart (":" ipv4addr)?;
	ipv6ref	    = "[" ipv6addr "]";
	paramchar   = [[\]/:&+$] | unreserved | escaped;
	hdrchar     = [[\]/?:+$] | unreserved | escaped;

	user	    = (unreserved | escaped | user_unreserved)+;
	password  = (unreserved | escaped | [&=+$,])*;
	host      = hostname | ipv4addr | ipv6ref;
	port      = digit+;
	uri_param = paramchar+ ("=" paramchar+)?;
	header    = hdrchar+ "=" hdrchar*;

	*       { err("invalid scheme"); goto fail }
	$       { err("invalid scheme"); goto fail }
	"sip:"  { uri.scheme = SIP; goto userinfo }
	"sips:" { uri.scheme = SIPS; goto userinfo }
	*/

userinfo:
	/*!re2c
	*    { cursor--; goto hostport }
	$    { err("invalid userinfo"); goto fail }
	@ts user (":" password)? @te "@" {
		uri.userinfo = str[ts:te]
		goto hostport
	}
	*/
hostport:
	/*!re2c
	*    { err("invalid host or port"); goto fail }
	$    { err("invalid host or port"); goto fail }
	@ts host (":" port)? @te {
		uri.hostport = str[ts:te]
		goto params
	}
	*/
params:
	/*!re2c
	*    { err("invalid params"); goto fail }
	$    { goto done }
	";" @ts uri_param (";" uri_param)* @te {
		uri.params = str[ts:te]
		goto headers
	}
	"?"  { cursor--; goto headers }
	*/
headers:
	/*!re2c
	*    { err("invalid headers"); goto fail }
	$    { goto done }
	"?" @ts header ("&" header)* @te {
		uri.headers = str[ts:te]
		goto done
	}
	*/

fail:
	return nil, parseError

done:
	return uri, nil
}

/* vim: set filetype=go : */
