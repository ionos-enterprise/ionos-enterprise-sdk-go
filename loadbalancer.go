package profitbricks

import (
	"net/http"
)

// Loadbalancer object
type Loadbalancer struct {
	BaseResource `json:",inline"`
	ID           string                 `json:"id,omitempty"`
	PBType       string                 `json:"type,omitempty"`
	Href         string                 `json:"href,omitempty"`
	Metadata     *Metadata              `json:"metadata,omitempty"`
	Properties   LoadbalancerProperties `json:"properties,omitempty"`
	Entities     LoadbalancerEntities   `json:"entities,omitempty"`
	Response     string                 `json:"Response,omitempty"`
	StatusCode   int                    `json:"statuscode,omitempty"`
}

// LoadbalancerProperties object
type LoadbalancerProperties struct {
	Name string `json:"name,omitempty"`
	IP   string `json:"ip,omitempty"`
	Dhcp bool   `json:"dhcp,omitempty"`
}

// LoadbalancerEntities object
type LoadbalancerEntities struct {
	Balancednics *BalancedNics `json:"balancednics,omitempty"`
}

// BalancedNics object
type BalancedNics struct {
	ID     string `json:"id,omitempty"`
	PBType string `json:"type,omitempty"`
	Href   string `json:"href,omitempty"`
	Items  []Nic  `json:"items,omitempty"`
}

// Loadbalancers object
type Loadbalancers struct {
	BaseResource `json:",inline"`
	ID           string         `json:"id,omitempty"`
	PBType       string         `json:"type,omitempty"`
	Href         string         `json:"href,omitempty"`
	Items        []Loadbalancer `json:"items,omitempty"`

	Response   string `json:"Response,omitempty"`
	StatusCode int    `json:"statuscode,omitempty"`
}

// ListLoadbalancers returns a Collection struct for loadbalancers in the Datacenter
func (c *Client) ListLoadbalancers(dcid string) (*Loadbalancers, error) {
	ret := &Loadbalancers{}
	return ret, c.GetOK(loadbalancersPath(dcid), ret)
}

// CreateLoadbalancer creates a loadbalancer in the datacenter from a jason []byte and returns a Instance struct
func (c *Client) CreateLoadbalancer(dcid string, request Loadbalancer) (*Loadbalancer, error) {
	ret := &Loadbalancer{}
	return ret, c.PostAcc(loadbalancersPath(dcid), request, ret)
}

// GetLoadbalancer pulls data for the Loadbalancer  where id = lbalid returns a Instance struct
func (c *Client) GetLoadbalancer(dcid, lbalid string) (*Loadbalancer, error) {
	ret := &Loadbalancer{}
	return ret, c.GetOK(loadbalancerPath(dcid, lbalid), ret)
}

// UpdateLoadbalancer updates a load balancer
func (c *Client) UpdateLoadbalancer(dcid string, lbalid string, obj LoadbalancerProperties) (*Loadbalancer, error) {
	ret := &Loadbalancer{}
	return ret, c.PatchAcc(loadbalancerPath(dcid, lbalid), obj, ret)
}

// DeleteLoadbalancer deletes a load balancer
func (c *Client) DeleteLoadbalancer(dcid, lbalid string) (*http.Header, error) {
	return c.DeleteAcc(loadbalancerPath(dcid, lbalid))
}

// ListBalancedNics lists balanced nics
func (c *Client) ListBalancedNics(dcid, lbalid string) (*Nics, error) {
	ret := &Nics{}
	return ret, c.GetOK(balancedNicsPath(dcid, lbalid), ret)

}

// AssociateNic attach a nic to load balancer
func (c *Client) AssociateNic(dcid string, lbalid string, nicid string) (*Nic, error) {
	ret := &Nic{}
	return ret, c.PostAcc(balancedNicsPath(dcid, lbalid), map[string]string{"id": nicid}, ret)
}

// GetBalancedNic gets a balanced nic
func (c *Client) GetBalancedNic(dcid, lbalid, balnicid string) (*Nic, error) {
	ret := &Nic{}
	return ret, c.GetOK(balancedNicPath(dcid, lbalid, balnicid), ret)

}

// DeleteBalancedNic removes a balanced nic
func (c *Client) DeleteBalancedNic(dcid, lbalid, balnicid string) (*http.Header, error) {
	return c.DeleteAcc(balancedNicPath(dcid, lbalid, balnicid))
}
