package profitbricks

import (
	"net/http"
)

// IPBlock object
type IPBlock struct {
	BaseResource `json:",inline"`
	ID           string            `json:"id,omitempty"`
	PBType       string            `json:"type,omitempty"`
	Href         string            `json:"href,omitempty"`
	Metadata     *Metadata         `json:"metadata,omitempty"`
	Properties   IPBlockProperties `json:"properties,omitempty"`
	Response     string            `json:"Response,omitempty"`
	StatusCode   int               `json:"statuscode,omitempty"`
}

// IPBlockProperties object
type IPBlockProperties struct {
	Name     string   `json:"name,omitempty"`
	IPs      []string `json:"ips,omitempty"`
	Location string   `json:"location,omitempty"`
	Size     int      `json:"size,omitempty"`
}

// IPBlocks object
type IPBlocks struct {
	BaseResource `json:",inline"`
	ID           string    `json:"id,omitempty"`
	PBType       string    `json:"type,omitempty"`
	Href         string    `json:"href,omitempty"`
	Items        []IPBlock `json:"items,omitempty"`
	Response     string    `json:"Response,omitempty"`
	StatusCode   int       `json:"statuscode,omitempty"`
}

// ListIPBlocks lists all IP blocks
func (c *Client) ListIPBlocks() (*IPBlocks, error) {
	ret := &IPBlocks{}
	return ret, c.GetOK(ipblocksPath(), ret)

}

// ReserveIPBlock creates an IP block
func (c *Client) ReserveIPBlock(request IPBlock) (*IPBlock, error) {
	ret := &IPBlock{}
	return ret, c.PostAcc(ipblocksPath(), request, ret)
}

// GetIPBlock gets an IP blocks
func (c *Client) GetIPBlock(ipblockid string) (*IPBlock, error) {
	ret := &IPBlock{}
	return ret, c.GetOK(ipblockPath(ipblockid), ret)
}

// UpdateIPBlock partial update of ipblock properties
func (c *Client) UpdateIPBlock(ipblockid string, props IPBlockProperties) (*IPBlock, error) {
	ret := &IPBlock{}
	return ret, c.PatchAcc(ipblockPath(ipblockid), props, ret)
}

// ReleaseIPBlock deletes an IP block
func (c *Client) ReleaseIPBlock(ipblockid string) (*http.Header, error) {
	return c.DeleteAcc(ipblockPath(ipblockid))
}
