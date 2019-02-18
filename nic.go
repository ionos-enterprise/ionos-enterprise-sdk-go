package profitbricks

import (
	"net/http"
	"strconv"
)

// Nic object
type Nic struct {
	ID         string         `json:"id,omitempty"`
	PBType     string         `json:"type,omitempty"`
	Href       string         `json:"href,omitempty"`
	Metadata   *Metadata      `json:"metadata,omitempty"`
	Properties *NicProperties `json:"properties,omitempty"`
	Entities   *NicEntities   `json:"entities,omitempty"`
	Response   string         `json:"Response,omitempty"`
	Headers    *http.Header   `json:"headers,omitempty"`
	StatusCode int            `json:"statuscode,omitempty"`
}

// NicProperties object
type NicProperties struct {
	Name           string   `json:"name,omitempty"`
	Mac            string   `json:"mac,omitempty"`
	Ips            []string `json:"ips,omitempty"`
	Dhcp           *bool    `json:"dhcp,omitempty"`
	Lan            int      `json:"lan,omitempty"`
	FirewallActive *bool    `json:"firewallActive,omitempty"`
	Nat            *bool    `json:"nat,omitempty"`
}

// NicEntities object
type NicEntities struct {
	FirewallRules *FirewallRules `json:"firewallrules,omitempty"`
}

// Nics object
type Nics struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Nic        `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// ListNics returns a Nics struct collection
func (c *Client) ListNics(dcid, srvid string) (*Nics, error) {
	url := nicColPath(dcid, srvid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Nics{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// CreateNic creates a nic on a server
func (c *Client) CreateNic(dcid string, srvid string, nic Nic) (*Nic, error) {

	url := nicColPath(dcid, srvid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Nic{}
	err := c.client.Post(url, nic, ret, http.StatusAccepted)

	return ret, err
}

// GetNic pulls data for the nic where id = srvid returns a Instance struct
func (c *Client) GetNic(dcid, srvid, nicid string) (*Nic, error) {

	url := nicPath(dcid, srvid, nicid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Nic{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// UpdateNic partial update of nic properties
func (c *Client) UpdateNic(dcid string, srvid string, nicid string, obj NicProperties) (*Nic, error) {

	url := nicPath(dcid, srvid, nicid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Nic{}
	err := c.client.Patch(url, obj, ret, http.StatusAccepted)
	return ret, err

}

// DeleteNic deletes the nic where id=nicid and returns a Resp struct
func (c *Client) DeleteNic(dcid, srvid, nicid string) (*http.Header, error) {
	url := nicPath(dcid, srvid, nicid)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}
