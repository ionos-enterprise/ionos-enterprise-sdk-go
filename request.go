package profitbricks

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// RequestStatus object
type RequestStatus struct {
	ID         string                `json:"id,omitempty"`
	PBType     string                `json:"type,omitempty"`
	Href       string                `json:"href,omitempty"`
	Metadata   RequestStatusMetadata `json:"metadata,omitempty"`
	Response   string                `json:"Response,omitempty"`
	Headers    *http.Header          `json:"headers,omitempty"`
	StatusCode int                   `json:"statuscode,omitempty"`
}

// RequestStatusMetadata object
type RequestStatusMetadata struct {
	Status  string          `json:"status,omitempty"`
	Message string          `json:"message,omitempty"`
	Etag    string          `json:"etag,omitempty"`
	Targets []RequestTarget `json:"targets,omitempty"`
}

// RequestTarget object
type RequestTarget struct {
	Target ResourceReference `json:"target,omitempty"`
	Status string            `json:"status,omitempty"`
}

// Requests object
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
	CreatedDate   time.Time     `json:"createdDate"`
	CreatedBy     string        `json:"createdBy"`
	Etag          string        `json:"etag"`
	RequestStatus RequestStatus `json:"requestStatus"`
}

type RequestProperties struct {
	Method  string      `json:"method"`
	Headers interface{} `json:"headers"`
	Body    string      `json:"body"`
	URL     string      `json:"url"`
}

// Request object
type Request struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Href       string            `json:"href"`
	Metadata   RequestMetadata   `json:"metadata"`
	Properties RequestProperties `json:"properties"`
	Response   string            `json:"Response,omitempty"`
	Headers    *http.Header      `json:"headers,omitempty"`
	StatusCode int               `json:"statuscode,omitempty"`
}

type RequestListFilter struct {
	url.Values
}

func NewRequestListFilter() *RequestListFilter {
	return &RequestListFilter{Values: url.Values{}}
}

func (f *RequestListFilter) AddUrl(url string) {
	f.WithUrl(url)
}

func (f *RequestListFilter) WithUrl(url string) *RequestListFilter {
	f.Add("filter.url", url)
	return f
}

func (f *RequestListFilter) AddCreatedDate(createdDate string) {
	f.WithCreatedDate(createdDate)
}

func (f *RequestListFilter) WithCreatedDate(createdDate string) *RequestListFilter {
	f.Add("filter.createdDate", createdDate)
	return f
}

func (f *RequestListFilter) AddMethod(method string) {
	f.WithMethod(method)
}

func (f *RequestListFilter) WithMethod(method string) *RequestListFilter {
	f.Add("filter.method", method)
	return f
}

func (f *RequestListFilter) AddBody(body string) {
	f.WithBody(body)
}

func (f *RequestListFilter) WithBody(body string) *RequestListFilter {
	f.Add("filter.body", body)
	return f
}

func (f *RequestListFilter) AddRequestStatus(requestStatus string) {
	f.WithRequestStatus(requestStatus)
}

func (f *RequestListFilter) WithRequestStatus(requestStatus string) *RequestListFilter {
	f.Add("filter.requestStatus", requestStatus)
	return f
}

// ListRequests lists all requests
func (c *Client) ListRequests() (*Requests, error) {
	url := "/requests" + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Requests{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// ListRequestsWithFilter lists all requests that matches the given filters
// Available filters are:
// * url
// * createdDate
// * method
// * body
// * requestStatus
func (c *Client) ListRequestsWithFilter(filter *RequestListFilter) (*Requests, error) {
	path := "/requests"
	query := url.Values{
		"depth": []string{"10"},
	}
	if filter != nil {
		for k, v := range filter.Values {
			for _, i := range v {
				query.Add(k, i)
			}
		}
	}
	path += "?" + query.Encode()
	ret := &Requests{}
	err := c.client.Get(path, ret, http.StatusOK)
	return ret, err

}

// GetRequest gets a specific request
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

func (c *Client) IsRequestDone(path string) (bool, error) {
	request, err := c.GetRequestStatus(path)
	if err != nil {
		return false, err
	}
	switch request.Metadata.Status {
	case "DONE":
		return true, nil
	case "FAILED":
		return true, NewClientError(
			RequestFailed,
			fmt.Sprintf("Request %s failed: %s", request.ID, request.Metadata.Message),
		)
	}
	return false, nil
}

// WaitTillProvisionedOrCanceled waits for a request to be completed.
// It returns an error if the request status could not be fetched, the request
// failed or the given context is canceled.
func (c *Client) WaitTillProvisionedOrCanceled(ctx context.Context, path string) error {
	done, err := c.IsRequestDone(path)
	if err != nil {
		return err
	} else if done {
		return nil
	}
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for !done {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		select {
		case <-ticker.C:
			done, err = c.IsRequestDone(path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// WaitTillProvisioned waits for a request to be completed.
// It returns an error if the request status could not be fetched, the request
// failed or a timeout of 2.5 minutes is exceeded.
func (c *Client) WaitTillProvisioned(path string) (err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 150*time.Second)
	defer cancel()
	if err = c.WaitTillProvisionedOrCanceled(ctx, path); err != nil {
		if err == context.DeadlineExceeded {
			return errors.New("timeout expired while waiting for request to complete")
		}
	}
	return
}

type RequestSelector func(Request) bool

func IsFinished(r Request) bool {
	switch r.Metadata.RequestStatus.Metadata.Status {
	case "QUEUED", "RUNNING":
		return false
	}
	return true
}

// WaitTillRequestsFinished will wait until list requests with applied filters does not
// return requests which are not finished.
func (c *Client) WaitTillRequestsFinished(ctx context.Context, filter *RequestListFilter) error {
	return c.WaitTillMatchingRequestsFinished(ctx, filter, func(r Request) bool { return !IsFinished(r) })
}

// WaitTillMatchingRequestsFinished gets open requests with given filters and will
// wait for each request that is selected by the selector. The selector
// should consider filtering out requests that are finished. (e.g. using IsFinished)
func (c *Client) WaitTillMatchingRequestsFinished(
	ctx context.Context, filter *RequestListFilter, selector RequestSelector) error {

	waited := true
	for waited && ctx.Err() == nil {

		waited = false
		requests, err := c.ListRequestsWithFilter(filter)
		if err != nil {
			return err
		}
		for _, r := range requests.Items {
			if selector(r) {
				waited = true
				if err := c.WaitTillProvisionedOrCanceled(ctx, r.Metadata.RequestStatus.Href); err != nil {
					if !IsRequestFailed(err) {
						return err
					}

				}
			}
		}
		if !waited {
			break
		}
	}
	return nil
}
