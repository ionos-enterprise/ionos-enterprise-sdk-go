package profitbricks

import (
	"net/http"
	"strings"
)

// Location object
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

// Locations object
type Locations struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Location   `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// LocationProperties object
type LocationProperties struct {
	Name         string   `json:"name,omitempty"`
	Features     []string `json:"features,omitempty"`
	ImageAliases []string `json:"imageAliases,omitempty"`
}

// ListLocations returns location collection data
func (c *Client) ListLocations() (*Locations, error) {
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.LocationApi.LocationsGet(ctx).Execute()
	ret := Locations{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
}

// GetRegionalLocations returns a list of available locations in a specific region
func (c *Client) GetRegionalLocations(regid string) (*Locations, error) {


    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.LocationApi.LocationsFindByRegionId(ctx, regid).Execute()
	ret := Locations{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret,err
}

// GetLocation returns location data
func (c *Client) GetLocation(locid string) (*Location, error) {

	parts := strings.SplitN(locid, "/", 2)
	if len(parts) != 2 {
		return nil, NewClientError(InvalidInput, "Invalid location id")
	}

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.LocationApi.LocationsFindByRegionIdAndId(ctx, parts[0], parts[1]).Execute()
	ret := Location{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
}
