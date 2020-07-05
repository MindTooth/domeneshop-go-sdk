/*
 * Domeneshop API Documentation
 *
 * # Overview  Domeneshop offers a simple, REST-based API, which currently supports the following features:  - List domains - List invoices - Create, read, update and delete DNS records for domains - Create, read, update and delete HTTP forwards (\"WWW forwarding\") for domains - Dynamic DNS (DDNS) update endpoints for use in consumer routers  More features are planned, including:  - Web hosting administration - Email address and email user/account administration  # Testing period  The API service is in version 0, which means it is possible that the interface will change rapidly during the testing period. For that reason, **the documentation on these pages may sometimes be outdated.**  Additionally, we make no guarantees about the stability of the API service during this testing period, and therefore ask customers to be careful with using the service for business critical purposes.  # Authentication  The Domeneshop API currently supports only one method of authentication, **HTTP Basic Auth**. More authentication methods may be added in the future.  To generate credentials, visit <a href=\"https://www.domeneshop.no/login\" target=\"_blank\">this page</a> after logging in to the control panel on our website:  <a href=\"https://www.domeneshop.no/login\" target=\"_blank\">https://www.domeneshop.no/login</a>  # Libraries  Domeneshop maintains multiple API libraries to simplify using the API. Please note that these libraries have the same stability guarantees to the API while the API is in version 0.  The libraries may be found in our [Github repository](https://github.com/domeneshop/).  Domeneshop also maintains a plugin for [EFF's Certbot](https://certbot.eff.org/), which automates issuance and renewal of SSL-certificates on your own servers for domains that use Domeneshop's DNS service. This plugin is found in our Github repository [here](https://github.com/domeneshop/certbot-dns-domeneshop).  <SecurityDefinitions /> 
 *
 * API version: v0
 * Contact: kundeservice@domeneshop.no
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package domeneshop
// Srv SRV records yo!
type Srv struct {
	// ID of DNS record
	Id int32 `json:"id"`
	// The host/subdomain the DNS record applies to
	Host string `json:"host"`
	// TTL of DNS record in seconds. Must be a multiple of 60.
	Ttl int32 `json:"ttl,omitempty"`
	Type string `json:"type"`
	// The target hostname
	Data string `json:"data"`
	// SRV record priority, also known as preference. Lower values are usually preferred first
	Priority int32 `json:"priority"`
	// SRV record weight. Relevant if multiple records have same preference
	Weight int32 `json:"weight"`
	// SRV record port. The port where the service is found.
	Port int32 `json:"port"`
}