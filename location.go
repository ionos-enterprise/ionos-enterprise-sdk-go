package profitbricks

import (
	"strings"
)

// Location object
type Location struct {
	BaseResource `json:",inline"`
	ID           string             `json:"id,omitempty"`
	PBType       string             `json:"type,omitempty"`
	Href         string             `json:"href,omitempty"`
	Properties   LocationProperties `json:"properties,omitempty"`
	Response     string             `json:"Response,omitempty"`
	StatusCode   int                `json:"statuscode,omitempty"`
}

// Locations object
type Locations struct {
	BaseResource `json:",inline"`
	ID           string     `json:"id,omitempty"`
	PBType       string     `json:"type,omitempty"`
	Href         string     `json:"href,omitempty"`
	Items        []Location `json:"items,omitempty"`
	Response     string     `json:"Response,omitempty"`
	StatusCode   int        `json:"statuscode,omitempty"`
}

// LocationProperties object
type LocationProperties struct {
	Name         string   `json:"name,omitempty"`
	Features     []string `json:"features,omitempty"`
	ImageAliases []string `json:"imageAliases,omitempty"`
}

// ListLocations returns location collection data
func (c *Client) ListLocations() (*Locations, error) {
	ret := &Locations{}
	return ret, c.GetOK(locationsPath(), ret)
}

// GetRegionalLocations returns a list of available locations in a specific region
func (c *Client) GetRegionalLocations(regid string) (*Locations, error) {
	ret := &Locations{}
	return ret, c.GetOK(locationRegionPath(regid), ret)

}

// GetLocation returns location data
func (c *Client) GetLocation(locid string) (*Location, error) {
	ret := &Location{}
	parts := strings.SplitN(locid, "/", 2)
	if len(parts) != 2 {
		return nil, ClientError.New("Invalid location id")
	}
	return ret, c.GetOK(locationPath(parts[0], parts[1]), ret)
}
