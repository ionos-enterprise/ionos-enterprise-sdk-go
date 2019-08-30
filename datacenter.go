package profitbricks

import (
	"net/http"
	"time"
)

// Datacenter represents Virtual Data Center
type Datacenter struct {
	BaseResource `json:",inline"`
	ID           string               `json:"id,omitempty"`
	PBType       string               `json:"type,omitempty"`
	Href         string               `json:"href,omitempty"`
	Metadata     *Metadata            `json:"metadata,omitempty"`
	Properties   DatacenterProperties `json:"properties,omitempty"`
	Entities     DatacenterEntities   `json:"entities,omitempty"`
	Response     string               `json:"Response,omitempty"`
}

// Metadata represents metadata received from Cloud API
type Metadata struct {
	CreatedDate      time.Time `json:"createdDate,omitempty"`
	CreatedBy        string    `json:"createdBy,omitempty"`
	Etag             string    `json:"etag,omitempty"`
	LastModifiedDate time.Time `json:"lastModifiedDate,omitempty"`
	LastModifiedBy   string    `json:"lastModifiedBy,omitempty"`
	State            string    `json:"state,omitempty"`
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
	BaseResource `json:",inline"`
	ID           string       `json:"id,omitempty"`
	PBType       string       `json:"type,omitempty"`
	Href         string       `json:"href,omitempty"`
	Items        []Datacenter `json:"items,omitempty"`
	Response     string       `json:"Response,omitempty"`
}

// ListDatacenters lists all data centers
func (c *Client) ListDatacenters() (*Datacenters, error) {
	ret := &Datacenters{}
	return ret, c.GetOK(datacentersPath(), ret)
}

// CreateDatacenter creates a data center
func (c *Client) CreateDatacenter(dc Datacenter) (*Datacenter, error) {
	ret := &Datacenter{}
	return ret, c.PostAcc(datacentersPath(), dc, ret)
}

// GetDatacenter gets a datacenter
func (c *Client) GetDatacenter(dcid string) (*Datacenter, error) {
	ret := &Datacenter{}
	return ret, c.GetOK(datacenterPath(dcid), ret)
}

// UpdateDataCenter updates a data center
func (c *Client) UpdateDataCenter(dcid string, obj DatacenterProperties) (*Datacenter, error) {
	ret := &Datacenter{}
	return ret, c.PatchAcc(datacenterPath(dcid), obj, ret)
}

// DeleteDatacenter deletes a data center
func (c *Client) DeleteDatacenter(dcid string) (*http.Header, error) {
	return c.DeleteAcc(datacenterPath(dcid))
}
