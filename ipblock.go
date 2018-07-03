package profitbricks

import (
	"net/http"
	"strconv"
)

//IPBlock object
type IPBlock struct {
	ID         string            `json:"id,omitempty"`
	PBType     string            `json:"type,omitempty"`
	Href       string            `json:"href,omitempty"`
	Metadata   *Metadata         `json:"metadata,omitempty"`
	Properties IPBlockProperties `json:"properties,omitempty"`
	Response   string            `json:"Response,omitempty"`
	Headers    *http.Header      `json:"headers,omitempty"`
	StatusCode int               `json:"statuscode,omitempty"`
}

//IPBlockProperties object
type IPBlockProperties struct {
	Name     string   `json:"name,omitempty"`
	IPs      []string `json:"ips,omitempty"`
	Location string   `json:"location,omitempty"`
	Size     int      `json:"size,omitempty"`
}

//IPBlocks object
type IPBlocks struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []IPBlock    `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

//ListIPBlocks lists all IP blocks
func (c *Client) ListIPBlocks() (*IPBlocks, error) {
	url := ipblockColPath() + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &IPBlocks{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//ReserveIPBlock creates an IP block
func (c *Client) ReserveIPBlock(request IPBlock) (*IPBlock, error) {
	url := ipblockColPath() + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &IPBlock{}
	err := c.client.Post(url, request, ret, http.StatusAccepted)
	return ret, err
}

//GetIPBlock gets an IP blocks
func (c *Client) GetIPBlock(ipblockid string) (*IPBlock, error) {
	url := ipblockPath(ipblockid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &IPBlock{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// UpdateIPBlock partial update of ipblock properties
func (c *Client) UpdateIPBlock(ipblockid string, props IPBlockProperties) (*IPBlock, error) {
	url := ipblockPath(ipblockid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &IPBlock{}
	err := c.client.Patch(url, props, ret, http.StatusAccepted)
	return ret, err
}

//ReleaseIPBlock deletes an IP block
func (c *Client) ReleaseIPBlock(ipblockid string) (*http.Header, error) {
	url := ipblockPath(ipblockid)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}
