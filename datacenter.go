package profitbricks

import (
	"net/http"
)

const (
	// Resource is being provisioned
	StateBusy = "BUSY"
	// Resource is ready to be used
	StateAvailable = "AVAILABLE"
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
	url := datacentersPath()
	ret := &Datacenters{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
}

// CreateDatacenter creates a data center
func (c *Client) CreateDatacenter(dc Datacenter) (*Datacenter, error) {
	url := datacentersPath()
	ret := &Datacenter{}
	err := c.Post(url, dc, ret, http.StatusAccepted)
	return ret, err
}

// GetDatacenter gets a datacenter
func (c *Client) GetDatacenter(dcid string) (*Datacenter, error) {
	url := datacenterPath(dcid)
	ret := &Datacenter{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
}

// UpdateDataCenter updates a data center
func (c *Client) UpdateDataCenter(dcid string, obj DatacenterProperties) (*Datacenter, error) {
	url := datacenterPath(dcid)
	ret := &Datacenter{}
	err := c.Patch(url, obj, ret, http.StatusAccepted)
	return ret, err
}

// DeleteDatacenter deletes a data center
func (c *Client) DeleteDatacenter(dcid string) (*http.Header, error) {
	url := datacenterPath(dcid)
	ret := &http.Header{}
	return ret, c.Delete(url, ret, http.StatusAccepted)
}
