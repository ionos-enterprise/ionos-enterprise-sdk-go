package profitbricks

import (
	"context"
	"github.com/ionos-cloud/sdk-go/v5"
	"net/http"
)

const (
	// Resource state is unknown
	StateUnknown = "UNKNOWN"
	// Resource is being provisioned
	StateBusy = "BUSY"
	// Resource is ready to be used
	StateAvailable = "AVAILABLE"
	// Resource has been de-provisioned
	StateInactive = "INACTIVE"
)

// Datacenter represents Virtual Data Center
type Datacenter struct {
	ID         string               `json:"id,omitempty"`
	PBType     string               `json:"type,omitempty"`
	Href       string               `json:"href,omitempty"`
	Metadata   *Metadata            `json:"metadata,omitempty"`
	Properties DatacenterProperties `json:"properties,omitempty"`
	Entities   DatacenterEntities   `json:"entities,omitempty"`
	Response   string               `json:"Response,omitempty"`
	Headers    *http.Header         `json:"headers,omitempty"`
}

// Metadata represents metadata recieved from Cloud API
type Metadata struct {
	CreatedDate          string `json:"createdDate,omitempty"`
	CreatedBy            string `json:"createdBy,omitempty"`
	CreatedByUserID      string `json:"createdByUserId,omitempty"`
	Etag                 string `json:"etag,omitempty"`
	LastModifiedDate     string `json:"lastModifiedDate,omitempty"`
	LastModifiedBy       string `json:"lastModifiedBy,omitempty"`
	LastModifiedByUserID string `json:"lastModifiedByUserId,omitempty"`
	State                string `json:"state,omitempty"`
}

// DatacenterProperties represents Virtual Data Center properties
type DatacenterProperties struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
	Version     int32  `json:"version,omitempty"`
}

// DatacenterEntities represents Virtual Data Center entities
type DatacenterEntities struct {
	Servers       *Servers       `json:"servers,omitempty"`
	Volumes       *Volumes       `json:"volumes,omitempty"`
	Loadbalancers *Loadbalancers `json:"loadbalancers,omitempty"`
	Lans          *Lans          `json:"lans,omitempty"`
}

// Datacenters is a list of Virtual Data Centers
type Datacenters struct {
	ID       string       `json:"id,omitempty"`
	PBType   string       `json:"type,omitempty"`
	Href     string       `json:"href,omitempty"`
	Items    []Datacenter `json:"items,omitempty"`
	Response string       `json:"Response,omitempty"`
	Headers  *http.Header `json:"headers,omitempty"`
}

// ListDatacenters lists all data centers
func (c *Client) ListDatacenters() (*Datacenters, error) {

	/*
	       ctx, cancel := c.GetContext()
	       if cancel != nil { defer cancel() }
	   	resp, apiResponse, err := c.CoreSdk.DataCenterApi.DatacentersGet(ctx, nil)
	   	ret := coreDataCentersConvertor(&resp)
	   	fillInResponse(&ret, apiResponse)
	   	return &ret, err
	*/

	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	resp, apiResponse, err := c.CoreSdk.DataCenterApi.DatacentersGet(ctx).Execute()
	ret := Datacenters{}
	if err != nil {
		return &ret, err
	}
	err2 := convertToCompat(&resp, &ret)
	if err2 != nil {
		return &ret, err2
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
		ret := &Datacenters{}
		url := datacentersPath()
		err := c.Get(url, ret, http.StatusOK)
		return ret, err
	*/
}

// CreateDatacenter creates a data center
func (c *Client) CreateDatacenter(dc Datacenter) (*Datacenter, error) {
	// url := datacentersPath()
	// ret := &Datacenter{}
	// err := c.Post(url, dc, ret, http.StatusAccepted)
	input := ionoscloud.Datacenter{}
	err := convertToCore(&dc, &input)
	if err != nil {
		return nil, err
	}

	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	resp, apiResponse, err := c.CoreSdk.DataCenterApi.DatacentersPost(ctx).Datacenter(input).Execute()
	ret := Datacenter{}
	err2 := convertToCompat(&resp, &ret)
	if err2 != nil {
		return nil, err2
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

}

// CreateDatacenterAndWait creates a data center, waits for the request to finish and returns a refreshed
// result.
// Note that an error does not necessarily means that the resource has not been created.
// If err & res are not nil, a resource with res.ID exists, but an error occurred either while waiting for
// the request or when refreshing the resource.
func (c *Client) CreateDatacenterAndWait(ctx context.Context, dc Datacenter) (res *Datacenter, err error) {
	res, err = c.CreateDatacenter(dc)
	if err != nil {
		return
	}
	if err = c.WaitTillProvisionedOrCanceled(ctx, res.Headers.Get("location")); err != nil {
		return
	}
	var rdc *Datacenter
	rdc, err = c.GetDatacenter(res.ID)
	if err != nil {
		return
	}
	return rdc, nil
}

// GetDatacenter gets a datacenter
func (c *Client) GetDatacenter(dcid string) (*Datacenter, error) {

	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	dc, response, err := c.CoreSdk.DataCenterApi.DatacentersFindById(ctx, dcid).Execute()
	ret := Datacenter{}
	err2 := convertToCompat(&dc, &ret)
	if err2 != nil {
		return nil, err2
	}
	fillInResponse(&ret, response)
	return &ret, err

}

// UpdateDataCenter updates a data center
func (c *Client) UpdateDataCenter(dcid string, obj DatacenterProperties) (*Datacenter, error) {

	input := ionoscloud.DatacenterProperties{}
	err := convertToCore(&obj, &input)
	if err != nil {
		return nil, err
	}

	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	resp, apiResponse, err := c.CoreSdk.DataCenterApi.DatacentersPatch(ctx, dcid).Datacenter(input).Execute()
	ret := Datacenter{}
	err2 := convertToCompat(&resp, &ret)
	if err2 != nil {
		return nil, err2
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

}

// UpdateDatacenter updates a data center, waits for the request to finish and returns a refreshed result.
// Note that an error does not necessarily means that the resource has not been updated.
// If err & res are not nil, a resource with res.ID exists, but an error occurred either while waiting for
// the request or when refreshing the resource.
func (c *Client) UpdateDatacenterAndWait(ctx context.Context, dcid string, obj DatacenterProperties) (res *Datacenter, err error) {
	res, err = c.UpdateDataCenter(dcid, obj)
	if err != nil {
		return
	}
	if err = c.WaitTillProvisionedOrCanceled(ctx, res.Headers.Get("location")); err != nil {
		return
	}
	var rdc *Datacenter
	if rdc, err = c.GetDatacenter(res.ID); err != nil {
		return
	} else {
		return rdc, nil
	}
}

// DeleteDatacenter deletes a data center
func (c *Client) DeleteDatacenter(dcid string) (*http.Header, error) {

	ctx, cancel := c.GetContext()
	if cancel != nil {
		defer cancel()
	}
	_, apiResponse, err := c.CoreSdk.DataCenterApi.DatacentersDelete(ctx, dcid).Execute()
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}

}

// DeleteDatacenterAndWait deletes given datacenter and waits for the request to finish
func (c *Client) DeleteDatacenterAndWait(ctx context.Context, dcid string) error {
	rsp, err := c.DeleteDatacenter(dcid)
	if err != nil {
		return err
	}
	return c.WaitTillProvisionedOrCanceled(ctx, rsp.Get("location"))
}
