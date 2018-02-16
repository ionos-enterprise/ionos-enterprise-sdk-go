package profitbricks

import (
	"net/http"
	"strconv"
)

//Loadbalancer object
type Loadbalancer struct {
	ID         string                 `json:"id,omitempty"`
	PBType     string                 `json:"type,omitempty"`
	Href       string                 `json:"href,omitempty"`
	Metadata   *Metadata              `json:"metadata,omitempty"`
	Properties LoadbalancerProperties `json:"properties,omitempty"`
	Entities   LoadbalancerEntities   `json:"entities,omitempty"`
	Response   string                 `json:"Response,omitempty"`
	Headers    *http.Header           `json:"headers,omitempty"`
	StatusCode int                    `json:"statuscode,omitempty"`
}

//LoadbalancerProperties object
type LoadbalancerProperties struct {
	Name string `json:"name,omitempty"`
	IP   string `json:"ip,omitempty"`
	Dhcp bool   `json:"dhcp,omitempty"`
}

//LoadbalancerEntities object
type LoadbalancerEntities struct {
	Balancednics *BalancedNics `json:"balancednics,omitempty"`
}

//BalancedNics object
type BalancedNics struct {
	ID     string `json:"id,omitempty"`
	PBType string `json:"type,omitempty"`
	Href   string `json:"href,omitempty"`
	Items  []Nic  `json:"items,omitempty"`
}

//Loadbalancers object
type Loadbalancers struct {
	ID     string         `json:"id,omitempty"`
	PBType string         `json:"type,omitempty"`
	Href   string         `json:"href,omitempty"`
	Items  []Loadbalancer `json:"items,omitempty"`

	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

//ListLoadbalancers returns a Collection struct for loadbalancers in the Datacenter
func (c *Client) ListLoadbalancers(dcid string) (*Loadbalancers, error) {

	url := lbalColPath(dcid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Loadbalancers{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//CreateLoadbalancer creates a loadbalancer in the datacenter from a jason []byte and returns a Instance struct
func (c *Client) CreateLoadbalancer(dcid string, request Loadbalancer) (*Loadbalancer, error) {
	url := lbalColPath(dcid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Loadbalancer{}
	err := c.client.Post(url, request, ret, http.StatusAccepted)

	return ret, err
}

//GetLoadbalancer pulls data for the Loadbalancer  where id = lbalid returns a Instance struct
func (c *Client) GetLoadbalancer(dcid, lbalid string) (*Loadbalancer, error) {
	url := lbalPath(dcid, lbalid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Loadbalancer{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//UpdateLoadbalancer updates a load balancer
func (c *Client) UpdateLoadbalancer(dcid string, lbalid string, obj LoadbalancerProperties) (*Loadbalancer, error) {
	url := lbalPath(dcid, lbalid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Loadbalancer{}
	err := c.client.Patch(url, obj, ret, http.StatusAccepted)
	return ret, err
}

//DeleteLoadbalancer deletes a load balancer
func (c *Client) DeleteLoadbalancer(dcid, lbalid string) (*http.Header, error) {
	url := lbalPath(dcid, lbalid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}

//ListBalancedNics lists balanced nics
func (c *Client) ListBalancedNics(dcid, lbalid string) (*Nics, error) {
	url := balnicColPath(dcid, lbalid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Nics{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//AssociateNic attach a nic to load balancer
func (c *Client) AssociateNic(dcid string, lbalid string, nicid string) (*Nic, error) {
	sm := map[string]string{"id": nicid}
	url := balnicColPath(dcid, lbalid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Nic{}
	err := c.client.Post(url, sm, ret, http.StatusAccepted)
	return ret, err
}

//GetBalancedNic gets a balanced nic
func (c *Client) GetBalancedNic(dcid, lbalid, balnicid string) (*Nic, error) {
	url := balnicPath(dcid, lbalid, balnicid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Nic{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//DeleteBalancedNic removes a balanced nic
func (c *Client) DeleteBalancedNic(dcid, lbalid, balnicid string) (*http.Header, error) {
	url := balnicPath(dcid, lbalid, balnicid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}
