package domeneshop

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	// Version defines a version
	Version = "0.0.0"

	defaultBaseURL = "https://api.domeneshop.no"

	defaultUserAgent = "domeneshop.go/" + Version

	apiVersion = "v0"
)

// Client represent Domenshop API client
type Client struct {
	httpClient *http.Client

	BaseURL string

	UserAgent string

	Domains *DomainsService
	// Records / DNS
	// Redirects
	// Invoice

	Debug bool
}

// NewClient returns a new client
func NewClient(httpClient *http.Client) *Client {
	c := &Client{httpClient: httpClient, BaseURL: defaultBaseURL}
	c.Domains = &DomainsService{client: c}
	return c
}

// NewRequest creates an API request
func (c *Client) NewRequest(method, path string, payload interface{}) (*http.Request, error) {
	url := c.BaseURL + path

	body := new(bytes.Buffer)
	if payload != nil {
		err := json.NewDecoder(body).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

}
