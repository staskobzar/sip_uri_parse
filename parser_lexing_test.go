package uri

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerParseURISuccess(t *testing.T) {
	tests := []struct {
		input                               string
		scheme                              Scheme
		userinfo, hostport, params, headers string
	}{
		{
			"sip:alice@atlanta.com",
			SIP, "alice", "atlanta.com", "", "",
		}, {
			"sip:alice:secretword@atlanta.com;transport=tcp",
			SIP, "alice:secretword", "atlanta.com", "transport=tcp", "",
		}, {
			"sips:alice@atlanta.com?subject=project%20x&priority=urgent",
			SIPS, "alice", "atlanta.com", "", "subject=project%20x&priority=urgent",
		}, {
			"sip:+1-212-555-1212:1234@gateway.com;user=phone",
			SIP, "+1-212-555-1212:1234", "gateway.com", "user=phone", "",
		}, {
			"sips:gateway.com",
			SIPS, "", "gateway.com", "", "",
		}, {
			"sip:alice@192.0.2.4:8899",
			SIP, "alice", "192.0.2.4:8899", "", "",
		}, {
			"sip:atlanta.com;method=REGISTER?to=alice%40atlanta.com",
			SIP, "", "atlanta.com", "method=REGISTER", "to=alice%40atlanta.com",
		}, {
			"sips:alice;day=tuesday@atlanta.com",
			SIPS, "alice;day=tuesday", "atlanta.com", "", "",
		},
	}

	for _, tc := range tests {
		uri, err := LexerParse(tc.input)
		assert.Nil(t, err)
		assert.Equal(t, tc.scheme, uri.scheme)
		assert.Equal(t, tc.userinfo, uri.userinfo)
		assert.Equal(t, tc.hostport, uri.hostport)
		assert.Equal(t, tc.params, uri.params)
		assert.Equal(t, tc.headers, uri.headers)
	}
}

func TestLexerParseURIFail(t *testing.T) {
	tests := []struct {
		input string
		err   string
	}{
		{"", "invalid scheme"},
		{"foo", "invalid scheme"},
		{"sipsfoo", "invalid scheme"},
		{"sip:1.1.1.1:a22", "invalid"},
		{"sip:atlanta.com;foo\"", "invalid"},
		// {"sip:atlanta.com;foo?bar", "invalid headers"},
		{"sip:;foo?bar", "invalid host"},
		{"sip:?foo", "invalid host"},
	}

	for _, tc := range tests {
		uri, err := LexerParse(tc.input)
		assert.NotNil(t, err)
		assert.Nil(t, uri)
		assert.Contains(t, err.Error(), tc.err)
	}
}

func BenchmarkLexerParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LexerParse("sips:bob:pa55w0rd@example.com:8080;user=phone?X-t=foo")
	}
}

func BenchmarkNetURLParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url.Parse("sips://bob:pa55w0rd@example.com:8080/hello/path?X-t=foo&hello")
	}
}

func BenchmarkDummyParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DummyParser("sips://bob:pa55w0rd@example.com:8080/hello/path?X-t=foo&hello")
	}
}

func BenchmarkRegexParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RegexParse("sips://bob:pa55w0rd@example.com:8080/hello/path?X-t=foo&hello")
	}
}
func TestDummyParser(t *testing.T) {
	tests := []struct {
		input                               string
		scheme                              Scheme
		userinfo, hostport, params, headers string
	}{
		{
			"sip:alice@atlanta.com",
			SIP, "alice", "atlanta.com", "", "",
		}, {
			"sip:alice:secretword@atlanta.com;transport=tcp",
			SIP, "alice:secretword", "atlanta.com", "transport=tcp", "",
		}, {
			"sips:alice@atlanta.com?subject=project%20x&priority=urgent",
			SIPS, "alice", "atlanta.com", "", "subject=project%20x&priority=urgent",
		}, {
			"sip:+1-212-555-1212:1234@gateway.com;user=phone",
			SIP, "+1-212-555-1212:1234", "gateway.com", "user=phone", "",
		}, {
			"sips:gateway.com",
			SIPS, "", "gateway.com", "", "",
		}, {
			"sip:alice@192.0.2.4:8899",
			SIP, "alice", "192.0.2.4:8899", "", "",
		}, {
			"sip:atlanta.com;method=REGISTER?to=alice%40atlanta.com",
			SIP, "", "atlanta.com", "method=REGISTER", "to=alice%40atlanta.com",
		}, {
			"sips:alice;day=tuesday@atlanta.com",
			SIPS, "alice;day=tuesday", "atlanta.com", "", "",
		},
	}

	for _, tc := range tests {
		uri, err := DummyParser(tc.input)
		assert.Nil(t, err, fmt.Sprintf("%+v", tc))
		assert.Equal(t, tc.scheme, uri.scheme)
		assert.Equal(t, tc.userinfo, uri.userinfo)
		assert.Equal(t, tc.hostport, uri.hostport)
		assert.Equal(t, tc.params, uri.params)
		assert.Equal(t, tc.headers, uri.headers)
	}
}

func TestRegexParse(t *testing.T) {
	tests := []struct {
		input                               string
		scheme                              Scheme
		userinfo, hostport, params, headers string
	}{
		{
			"sip:alice@atlanta.com",
			SIP, "alice", "atlanta.com", "", "",
		}, {
			"sip:alice:secretword@atlanta.com;transport=tcp",
			SIP, "alice:secretword", "atlanta.com", "transport=tcp", "",
		}, {
			"sips:alice@atlanta.com?subject=project%20x&priority=urgent",
			SIPS, "alice", "atlanta.com", "", "subject=project%20x&priority=urgent",
		}, {
			"sip:+1-212-555-1212:1234@gateway.com;user=phone",
			SIP, "+1-212-555-1212:1234", "gateway.com", "user=phone", "",
		}, {
			"sips:gateway.com",
			SIPS, "", "gateway.com", "", "",
		}, {
			"sip:alice@192.0.2.4:8899",
			SIP, "alice", "192.0.2.4:8899", "", "",
		}, {
			"sip:atlanta.com;method=REGISTER?to=alice%40atlanta.com",
			SIP, "", "atlanta.com", "method=REGISTER", "to=alice%40atlanta.com",
		}, {
			"sips:alice;day=tuesday@atlanta.com",
			SIPS, "alice;day=tuesday", "atlanta.com", "", "",
		},
	}

	for _, tc := range tests {
		uri, err := RegexParse(tc.input)
		assert.Nil(t, err, fmt.Sprintf("%+v", tc))
		assert.Equal(t, tc.scheme, uri.scheme)
		assert.Equal(t, tc.userinfo, uri.userinfo)
		assert.Equal(t, tc.hostport, uri.hostport)
		assert.Equal(t, tc.params, uri.params)
		assert.Equal(t, tc.headers, uri.headers)
	}
}

func TestLexerScanIPv4(t *testing.T) {
	tests := []struct {
		input  string
		length int
		ok     bool
	}{
		{"8.8.8.8", 7, true},
		{"18.8.88.8", 9, true},
		{"199.0.17.255", 12, true},
		{"199.199.177.255", 15, true},
		{"8.8.8.8:999", 7, true},
		{"10.0.0.121?header=foo", 10, true},
		{"77.123.0.11;header=foo", 11, true},
		{"8.8d.8.8", 0, false},
		{"8.8.8.256", 0, false},
		{"8000.8.8.56", 0, false},
		{"", 0, false},
		{"foo", 0, false},
	}

	for _, tc := range tests {
		l, ok := parseIPv4(tc.input)
		assert.Equal(t, tc.length, l)
		assert.Equal(t, tc.ok, ok)
	}
}
