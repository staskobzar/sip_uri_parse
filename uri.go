package uri

// Scheme for sip URI
type Scheme uint8

// URI schemes sip/sips
const (
	UNKNOWN Scheme = iota
	SIPS
	SIP
)

// URI SIP struct
type URI struct {
	scheme   Scheme
	userinfo string
	hostport string
	params   string
	headers  string
}
