package profitbricks

import (
	"github.com/ionos-cloud/sdk-go/v5"
	"net/http"
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

	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	rsp, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsGet(ctx, dcid, srvid).Execute()
	ret := Nics{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
		url := nicsPath(dcid, srvid)
		ret := &Nics{}
		err := c.Get(url, ret, http.StatusOK)
		return ret, err
	*/
}

// CreateNic creates a nic on a server
func (c *Client) CreateNic(dcid string, srvid string, nic Nic) (*Nic, error) {

	input := ionoscloud.Nic{}
	if errConvert := convertToCore(&nic, &input); errConvert != nil {
		return nil, errConvert
	}
	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	rsp, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsPost(ctx, dcid, srvid).Nic(input).Execute()
	ret := Nic{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
		url := nicsPath(dcid, srvid)
		ret := &Nic{}
		err := c.Post(url, nic, ret, http.StatusAccepted)

		return ret, err
	*/
}

// GetNic pulls data for the nic where id = srvid returns a Instance struct
func (c *Client) GetNic(dcid, srvid, nicid string) (*Nic, error) {

	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	rsp, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsFindById(ctx, dcid, srvid, nicid).Execute()
	ret := Nic{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
		url := nicPath(dcid, srvid, nicid)
		ret := &Nic{}
		err := c.Get(url, ret, http.StatusOK)
		return ret, err
	*/
}

// UpdateNic partial update of nic properties
func (c *Client) UpdateNic(dcid string, srvid string, nicid string, obj NicProperties) (*Nic, error) {

	input := ionoscloud.NicProperties{}
	if errConvert := convertToCore(&obj, &input); errConvert != nil {
		return nil, errConvert
	}
	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	rsp, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsPatch(ctx, dcid, srvid, nicid).Nic(input).Execute()
	ret := Nic{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
		url := nicPath(dcid, srvid, nicid)
		ret := &Nic{}
		err := c.Patch(url, obj, ret, http.StatusAccepted)
		return ret, err
	*/
}

// DeleteNic deletes the nic where id=nicid and returns a Resp struct
func (c *Client) DeleteNic(dcid, srvid, nicid string) (*http.Header, error) {

	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	_, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsDelete(ctx, dcid, srvid, nicid).Execute()
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
	/*
		url := nicPath(dcid, srvid, nicid)
		ret := &http.Header{}
		err := c.Delete(url, ret, http.StatusAccepted)
		return ret, err
	*/
}
