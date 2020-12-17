package profitbricks

import (
	ionossdk "github.com/ionos-cloud/sdk-go/v5"
	"net/http"
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
	Name        string       `json:"name,omitempty"`
	IPs         []string     `json:"ips,omitempty"`
	IPConsumers []IPConsumer `json:"ipConsumers,omitempty"`
	Location    string       `json:"location,omitempty"`
	Size        int          `json:"size,omitempty"`
}

type IPConsumer struct {
	IP             string `json:"ip,omitempty"`
	Mac            string `json:"mac,omitempty"`
	NicID          string `json:"nicId,omitempty"`
	ServerID       string `json:"serverId,omitempty"`
	ServerName     string `json:"serverName,omitempty"`
	DatacenterID   string `json:"datacenterId,omitempty"`
	DatacenterName string `json:"datacenterName,omitempty"`
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

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.IPBlocksApi.IpblocksGet(ctx).Execute()
	ret := IPBlocks{}

	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := ipblocksPath()
	ret := &IPBlocks{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err

	 */
}

//ReserveIPBlock creates an IP block
func (c *Client) ReserveIPBlock(request IPBlock) (*IPBlock, error) {

	input := ionossdk.IpBlock{}
	if errConvert := convertToCore(&request, &input); errConvert != nil {
		return nil, errConvert
	}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.IPBlocksApi.IpblocksPost(ctx).Ipblock(input).Execute()
	ret := IPBlock{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := ipblocksPath()
	ret := &IPBlock{}
	err := c.Post(url, request, ret, http.StatusAccepted)
	return ret, err
	*/
}

//GetIPBlock gets an IP blocks
func (c *Client) GetIPBlock(ipblockid string) (*IPBlock, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.IPBlocksApi.IpblocksFindById(ctx, ipblockid).Execute()
	ret := IPBlock{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := ipblockPath(ipblockid)
	ret := &IPBlock{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// UpdateIPBlock partial update of ipblock properties
func (c *Client) UpdateIPBlock(ipblockid string, props IPBlockProperties) (*IPBlock, error) {

	input := ionossdk.IpBlockProperties{}
	if errConvert := convertToCore(&props, &input); errConvert != nil {
		return nil, errConvert
	}

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.IPBlocksApi.IpblocksPatch(ctx, ipblockid).Ipblock(input).Execute()

	ret := IPBlock{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}

	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := ipblockPath(ipblockid)
	ret := &IPBlock{}
	err := c.Patch(url, props, ret, http.StatusAccepted)
	return ret, err

	 */
}

//ReleaseIPBlock deletes an IP block
func (c *Client) ReleaseIPBlock(ipblockid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.IPBlocksApi.IpblocksDelete(ctx, ipblockid).Execute()

	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}

	/*
	url := ipblockPath(ipblockid)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}
