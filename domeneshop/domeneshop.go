package domeneshop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
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
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", formatUserAgent(c.UserAgent))

	return req, nil
}

func formatUserAgent(customUserAgent string) string {
	if customUserAgent == "" {
		return defaultUserAgent
	}

	return fmt.Sprintf("%s %s", defaultUserAgent, customUserAgent)
}

func versioned(path string) string {
	return fmt.Sprintf("/%s/%s", apiVersion, strings.Trim(path, "/"))

}

func (c *Client) get(path string, obj interface{}) (*http.Response, error) {
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, obj)
}

// Do sends API
func (c *Client) Do(req *http.Request, obj interface{}) (*http.Response, error) {
	if c.Debug {
		log.Printf("Executing request (%v): %#v", req.URL, req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c.Debug {
		log.Printf("Response recieved: %#v", resp)
	}

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	// DEBUG:  Must remove....
	if c.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q", dump)
	}

	if obj != nil {
		if w, ok := obj.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(obj)
		}
	}

	if c.Debug {
		// dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Hello")
		log.Printf("Body recieved: %#v", resp.Body)
	}

	return resp, err
}

// Response responds
type Response struct {
	HTTPResponse *http.Response
}

// ErrorResponse is an error
type ErrorResponse struct {
	Response
	Message string `json:"code"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %v %v",
		r.HTTPResponse.Request.Method, r.HTTPResponse.Request.URL,
		r.HTTPResponse.StatusCode, r.Message)
}

// CheckResponse check an response
func CheckResponse(resp *http.Response) error {

	if code := resp.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{}
	errorResponse.HTTPResponse = resp

	err := json.NewDecoder(resp.Body).Decode(errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

func addURLQueryOptions(path string, options interface{}) (string, error) {
	opt := reflect.ValueOf(options)

	if opt.Kind() == reflect.Ptr && opt.IsNil() {
		return path, nil
	}

	u, err := url.Parse(path)
	if err != nil {
		return path, err
	}

	qs, err := query.Values(options)
	if err != nil {
		return path, err
	}

	uqs := u.Query()
	for k := range qs {
		uqs.Set(k, qs.Get(k))
	}
	u.RawQuery = uqs.Encode()

	return u.String(), nil
}
