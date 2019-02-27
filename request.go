package profitbricks

import (
	"net/http"
	"strconv"
	"time"
)

//RequestStatus object
type RequestStatus struct {
	ID         string                `json:"id,omitempty"`
	PBType     string                `json:"type,omitempty"`
	Href       string                `json:"href,omitempty"`
	Metadata   RequestStatusMetadata `json:"metadata,omitempty"`
	Response   string                `json:"Response,omitempty"`
	Headers    *http.Header          `json:"headers,omitempty"`
	StatusCode int                   `json:"statuscode,omitempty"`
}

//RequestStatusMetadata object
type RequestStatusMetadata struct {
	Status  string          `json:"status,omitempty"`
	Message string          `json:"message,omitempty"`
	Etag    string          `json:"etag,omitempty"`
	Targets []RequestTarget `json:"targets,omitempty"`
}

//RequestTarget object
type RequestTarget struct {
	Target ResourceReference `json:"target,omitempty"`
	Status string            `json:"status,omitempty"`
}

//Requests object
type Requests struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Request    `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

type RequestMetadata struct {
	CreatedDate   time.Time `json:"createdDate"`
	CreatedBy     string    `json:"createdBy"`
	Etag          string    `json:"etag"`
	RequestStatus RequestStatus `json:"requestStatus"`
}


type RequestProperties struct {
	Method  string      `json:"method"`
	Headers interface{} `json:"headers"`
	Body    interface{} `json:"body"`
	URL     string      `json:"url"`
}

//Request object
type Request struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Href     string `json:"href"`
	Metadata RequestMetadata `json:"metadata"`
	Properties RequestProperties `json:"properties"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

//ListRequests lists all requests
func (c *Client) ListRequests() (*Requests, error) {
	url := "/requests" + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Requests{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//GetRequest gets a specific request
func (c *Client) GetRequest(reqID string) (*Request, error) {
	url := "/requests/" + reqID + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Request{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// GetRequestStatus retursn status of the request
func (c *Client) GetRequestStatus(path string) (*RequestStatus, error) {
	url := path + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &RequestStatus{}
	err := c.client.GetRequestStatus(url, ret, http.StatusOK)
	return ret, err
}
