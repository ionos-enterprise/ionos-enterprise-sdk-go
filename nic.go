package profitbricks

import (
	"net/http"
)

// Nic object
type Nic struct {
	BaseResource `json:",inline"`
	ID           string         `json:"id,omitempty"`
	PBType       string         `json:"type,omitempty"`
	Href         string         `json:"href,omitempty"`
	Metadata     *Metadata      `json:"metadata,omitempty"`
	Properties   *NicProperties `json:"properties,omitempty"`
	Entities     *NicEntities   `json:"entities,omitempty"`
	Response     string         `json:"Response,omitempty"`
	StatusCode   int            `json:"statuscode,omitempty"`
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
	BaseResource `json:",inline"`
	ID           string `json:"id,omitempty"`
	PBType       string `json:"type,omitempty"`
	Href         string `json:"href,omitempty"`
	Items        []Nic  `json:"items,omitempty"`
	Response     string `json:"Response,omitempty"`
	StatusCode   int    `json:"statuscode,omitempty"`
}

// ListNics returns a Nics struct collection
func (c *Client) ListNics(dcid, srvid string) (*Nics, error) {
	ret := &Nics{}
	return ret, c.GetOK(nicsPath(dcid, srvid), ret)
}

// CreateNic creates a nic on a server
func (c *Client) CreateNic(dcid string, srvid string, nic Nic) (*Nic, error) {
	ret := &Nic{}
	return ret, c.PostAcc(nicsPath(dcid, srvid), nic, ret)
}

// GetNic pulls data for the nic where id = srvid returns a Instance struct
func (c *Client) GetNic(dcid, srvid, nicid string) (*Nic, error) {
	ret := &Nic{}
	return ret, c.GetOK(nicPath(dcid, srvid, nicid), ret)

}

// UpdateNic partial update of nic properties
func (c *Client) UpdateNic(dcid string, srvid string, nicid string, obj NicProperties) (*Nic, error) {
	ret := &Nic{}
	return ret, c.PatchAcc(nicPath(dcid, srvid, nicid), obj, ret)
}

// DeleteNic deletes the nic where id=nicid and returns a Resp struct
func (c *Client) DeleteNic(dcid, srvid, nicid string) (*http.Header, error) {
	return c.DeleteAcc(nicPath(dcid, srvid, nicid))
}
