package profitbricks

import (
	"net/http"
	"strconv"
)

//Lan object
type Lan struct {
	ID         string        `json:"id,omitempty"`
	PBType     string        `json:"type,omitempty"`
	Href       string        `json:"href,omitempty"`
	Metadata   *Metadata     `json:"metadata,omitempty"`
	Properties LanProperties `json:"properties,omitempty"`
	Entities   *LanEntities  `json:"entities,omitempty"`
	Response   string        `json:"Response,omitempty"`
	Headers    *http.Header  `json:"headers,omitempty"`
	StatusCode int           `json:"statuscode,omitempty"`
}

//LanProperties object
type LanProperties struct {
	Name       string        `json:"name,omitempty"`
	Public     bool          `json:"public,omitempty"`
	IPFailover *[]IPFailover `json:"ipFailover,omitempty"`
}

//LanEntities object
type LanEntities struct {
	Nics *LanNics `json:"nics,omitempty"`
}

//IPFailover object
type IPFailover struct {
	NicUUID string `json:"nicUuid,omitempty"`
	IP      string `json:"ip,omitempty"`
}

//LanNics object
type LanNics struct {
	ID     string `json:"id,omitempty"`
	PBType string `json:"type,omitempty"`
	Href   string `json:"href,omitempty"`
	Items  []Nic  `json:"items,omitempty"`
}

//Lans object
type Lans struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Lan        `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// ListLans returns a Collection for lans in the Datacenter
func (c *Client) ListLans(dcid string) (*Lans, error) {
	url := lanColPath(dcid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Lans{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// CreateLan creates a lan in the datacenter
// from a jason []byte and returns a Instance struct
func (c *Client) CreateLan(dcid string, request Lan) (*Lan, error) {
	url := lanColPath(dcid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Lan{}
	err := c.client.Post(url, request, ret, http.StatusAccepted)
	return ret, err
}

// GetLan pulls data for the lan where id = lanid returns an Instance struct
func (c *Client) GetLan(dcid, lanid string) (*Lan, error) {
	url := lanPath(dcid, lanid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Lan{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// UpdateLan does a partial update to a lan using json from []byte json returns a Instance struct
func (c *Client) UpdateLan(dcid string, lanid string, obj LanProperties) (*Lan, error) {
	url := lanPath(dcid, lanid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Lan{}
	err := c.client.Patch(url, obj, ret, http.StatusAccepted)
	return ret, err
}

// DeleteLan deletes a lan where id == lanid
func (c *Client) DeleteLan(dcid, lanid string) (*http.Header, error) {
	url := lanPath(dcid, lanid)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}
