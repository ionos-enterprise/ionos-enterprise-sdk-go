package profitbricks

import (
	"context"
	ionossdk "github.com/ionos-cloud/sdk-go/v5"
	"net/http"
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

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.LoadBalancerApi.DatacentersLoadbalancersGet(ctx, dcid).Execute()
	ret := Loadbalancers{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := loadbalancersPath(dcid)
	ret := &Loadbalancers{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//CreateLoadbalancer creates a loadbalancer in the datacenter from a jason []byte and returns a Instance struct
func (c *Client) CreateLoadbalancer(dcid string, request Loadbalancer) (*Loadbalancer, error) {

	input := ionossdk.Loadbalancer{}
	if errConvert := convertToCore(&request, &input); errConvert != nil {
		return nil, errConvert
	}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.LoadBalancerApi.DatacentersLoadbalancersPost(ctx, dcid).Loadbalancer(input).Execute()
	ret := Loadbalancer{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := loadbalancersPath(dcid)
	ret := &Loadbalancer{}
	err := c.Post(url, request, ret, http.StatusAccepted)

	return ret, err
	 */
}

//GetLoadbalancer pulls data for the Loadbalancer  where id = lbalid returns a Instance struct
func (c *Client) GetLoadbalancer(dcid, lbalid string) (*Loadbalancer, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.LoadBalancerApi.DatacentersLoadbalancersFindById(ctx, dcid, lbalid).Execute()
	ret := Loadbalancer{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := loadbalancerPath(dcid, lbalid)
	ret := &Loadbalancer{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//UpdateLoadbalancer updates a load balancer
func (c *Client) UpdateLoadbalancer(dcid string, lbalid string, obj LoadbalancerProperties) (*Loadbalancer, error) {

	input := ionossdk.LoadbalancerProperties{}
	if errConvert := convertToCore(&obj, &input); errConvert != nil {
		return nil, errConvert
	}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.LoadBalancerApi.DatacentersLoadbalancersPatch(ctx, dcid, lbalid).Loadbalancer(input).Execute()
	ret := Loadbalancer{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := loadbalancerPath(dcid, lbalid)
	ret := &Loadbalancer{}
	err := c.Patch(url, obj, ret, http.StatusAccepted)
	return ret, err
	 */
}

//DeleteLoadbalancer deletes a load balancer
func (c *Client) DeleteLoadbalancer(dcid, lbalid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.LoadBalancerApi.DatacentersLoadbalancersDelete(ctx, dcid, lbalid).Execute()
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}

	/*
	url := loadbalancerPath(dcid, lbalid)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}

//ListBalancedNics lists balanced nics
func (c *Client) ListBalancedNics(dcid, lbalid string) (*Nics, error) {

	rsp, apiResponse, err := c.CoreSdk.LoadBalancerApi.DatacentersLoadbalancersBalancednicsGet(
		context.TODO(), dcid, lbalid).Execute()
	ret := Nics{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := balancedNicsPath(dcid, lbalid)
	ret := &Nics{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//AssociateNic attach a nic to load balancer
func (c *Client) AssociateNic(dcid string, lbalid string, nicid string) (*Nic, error) {

	input := ionossdk.Nic{
		Id: &nicid,
	}

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.LoadBalancerApi.DatacentersLoadbalancersBalancednicsPost(ctx, dcid, lbalid).Nic(input).Execute()
	ret := Nic{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	sm := map[string]string{"id": nicid}
	url := balancedNicsPath(dcid, lbalid)
	ret := &Nic{}
	err := c.Post(url, sm, ret, http.StatusAccepted)
	return ret, err
	 */
}

//GetBalancedNic gets a balanced nic
func (c *Client) GetBalancedNic(dcid, lbalid, balnicid string) (*Nic, error) {

	rsp, apiResponse, err := c.CoreSdk.LoadBalancerApi.DatacentersLoadbalancersBalancednicsFindByNicId(
		context.TODO(), dcid, lbalid, balnicid).Execute()
	ret := Nic{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := balancedNicPath(dcid, lbalid, balnicid)
	ret := &Nic{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//DeleteBalancedNic removes a balanced nic
func (c *Client) DeleteBalancedNic(dcid, lbalid, balnicid string) (*http.Header, error) {

	_, apiResponse, err := c.CoreSdk.LoadBalancerApi.DatacentersLoadbalancersBalancednicsDelete(
		context.TODO(), dcid, lbalid, balnicid).Execute()
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}

	/*
	url := balancedNicPath(dcid, lbalid, balnicid)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}
