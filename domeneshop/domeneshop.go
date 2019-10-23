package domeneshop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.domeneshop.no/v0/"

// Client is the struct holding the login
type Client struct {
	Token  string
	Secret string
}

// BasicAuth creates an auth client
func BasicAuth(token, secret string) *Client {
	return &Client{
		Token:  token,
		Secret: secret,
	}
}

// Domain hold the basic info struct we want
type Domain struct {
	ID     int    `json:"id"`
	Name   string `json:"domain"`
	Status string `json:"status"`
}

// Domains is the array we seek
type Domains []Domain

// doRequest handles the requests to the API
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.Token, c.Secret)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// GetDomains puils down all domains
func (c *Client) GetDomains() (*Domains, error) {
	url := fmt.Sprintf(baseURL + "/domains")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var domains Domains
	err = json.Unmarshal(bytes, &domains)
	if err != nil {
		return nil, err
	}

	return &domains, nil
}
