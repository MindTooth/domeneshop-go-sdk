package domeneshop

import "fmt"

// DomainsService handles communication
type DomainsService struct {
	client *Client
}

// Domain main domain
type Domain struct {
	Domain         string   `json:"domain,omitempty"`
	ExpireDate     string   `json:"expire_date,omitempty"`
	ID             int64    `json:"id,omitempty"`
	Nameservers    []string `json:"nameservers,omitempty"`
	RegisteredDate string   `json:"registered_date,omitempty"`
	Registrant     string   `json:"registrant,omitempty"`
	Renew          bool     `json:"renew,omitempty"`
	Services       services `json:"services,omitempty"`
	Status         string   `json:"status,omitempty"`
}

type services struct {
	DNS       bool   `json:"dns,omitempty"`
	Email     bool   `json:"email,omitempty"`
	Registrar bool   `json:"registrar,omitempty"`
	Webhotel  string `json:"webhotel,omitempty"`
}

func domainPath(domainIdentifier string) (path string) {
	path = fmt.Sprintf("/domains")
	if domainIdentifier != "" {
		path += fmt.Sprintf("/%v", domainIdentifier)
	}
	return
}

type domainResponse struct {
	Response
	Data *Domain
}

type domainsResponse struct {
	Response
	Data []Domain
}

// DomainListOptions specifies extra options.
type DomainListOptions struct {
	Domain string `url:"domain,omitempty"`
}

// ListDomains Lists all domains
func (s *DomainsService) ListDomains(options *DomainListOptions) (*domainsResponse, error) {
	path := versioned(domainPath(""))
	domainsResponse := &domainsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(path, domainsResponse)
	if err != nil {
		return nil, err
	}

	domainsResponse.HTTPResponse = resp
	return domainsResponse, nil
}

// GetDomain retrieves info about a single domain
func (s *DomainsService) GetDomain(domainIdentifier string) (*domainResponse, error) {
	path := versioned(domainPath(domainIdentifier))
	domainResponse := &domainResponse{}

	resp, err := s.client.get(path, domainResponse)
	if err != nil {
		return nil, err
	}

	domainResponse.HTTPResponse = resp
	return domainResponse, nil
}
