package domeneshop

import (
	"net/http"
)

// BasicAuthTransport is the struct holding the login
type BasicAuthTransport struct {
	Token  string
	Secret string

	Transport http.RoundTripper
}

// RoundTrip creates an auth client
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Respons, error) {
	req2 := cloneRequest(req)

	req2.SetBasicAuth(t.Token, t.Token)
	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client
func (t *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *BasicAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
