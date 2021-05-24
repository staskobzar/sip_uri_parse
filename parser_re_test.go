package uri

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRe3GoParseURISuccess(t *testing.T) {
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
		uri, err := Re2GoParse(tc.input)
		assert.Nil(t, err)
		assert.Equal(t, tc.scheme, uri.scheme)
		assert.Equal(t, tc.userinfo, uri.userinfo)
		assert.Equal(t, tc.hostport, uri.hostport)
		assert.Equal(t, tc.params, uri.params)
		assert.Equal(t, tc.headers, uri.headers)
	}
}

func TestRe2GoParseURIFail(t *testing.T) {
	tests := []struct {
		input string
		err   string
	}{
		{"", "invalid scheme"},
		{"foo", "invalid scheme"},
		{"sipsfoo", "invalid scheme"},
		{"sip:1.1.1.1:a22", "invalid"},
		{"sip:atlanta.com;foo\"", "invalid"},
		{"sip:atlanta.com;foo?bar", "invalid headers"},
		{"sip:;foo?bar", "invalid host"},
		{"sip:?foo", "invalid host"},
	}

	for _, tc := range tests {
		uri, err := Re2GoParse(tc.input)
		assert.NotNil(t, err)
		assert.Nil(t, uri)
		assert.Contains(t, err.Error(), tc.err)
	}
}

func BenchmarkRe2GoParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Re2GoParse("sips:bob:pa55w0rd@example.com:8080;user=phone?X-t=foo")
	}
}
