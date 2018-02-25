package profitbricks

import (
	"net/http"
	"strconv"
)

//Location object
type Location struct {
	ID         string             `json:"id,omitempty"`
	PBType     string             `json:"type,omitempty"`
	Href       string             `json:"href,omitempty"`
	Metadata   Metadata           `json:"metadata,omitempty"`
	Properties LocationProperties `json:"properties,omitempty"`
	Response   string             `json:"Response,omitempty"`
	Headers    *http.Header       `json:"headers,omitempty"`
	StatusCode int                `json:"statuscode,omitempty"`
}

//Locations object
type Locations struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Location   `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

//LocationProperties object
type LocationProperties struct {
	Name         string   `json:"name,omitempty"`
	Features     []string `json:"features,omitempty"`
	ImageAliases []string `json:"imageAliases,omitempty"`
}

// ListLocations returns location collection data
func (c *Client) ListLocations() (*Locations, error) {
	url := locationColPath() + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Locations{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// GetRegionalLocations returns a list of available locations in a specific region
func (c *Client) GetRegionalLocations(regid string) (*Locations, error) {
	url := locationRegPath(regid) + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Locations{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// GetLocation returns location data
func (c *Client) GetLocation(locid string) (*Location, error) {
	url := locationPath(locid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Location{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}
