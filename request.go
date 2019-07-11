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

// RequestListFilter is a wrapper around url.Values to provide a common
// interface to make use of the filters that the ionos API provides for the
// requests endpoint.
// Example:
//   filter := NewRequestListFilter().WithUrl("volumes").WithBody("de/fra") will create a api call
//   with query args like: /requests?filter.url=volumes&filter.body=de%2Ffra
type RequestListFilter struct {
	url.Values
}

// NewRequestListFilter creates a new RequestListFilter
func NewRequestListFilter() *RequestListFilter {
	return &RequestListFilter{Values: url.Values{}}
}

// AddUrl adds an url filter to the request list filter
func (f *RequestListFilter) AddUrl(url string) {
	f.WithUrl(url)
}

// WithUrl adds an url filter to the request list filter returning the filter for chaining
func (f *RequestListFilter) WithUrl(url string) *RequestListFilter {
	f.Add("filter.url", url)
	return f
}

// AddCreatedDate adds a createdDate filter to the request list filter
func (f *RequestListFilter) AddCreatedDate(createdDate string) {
	f.WithCreatedDate(createdDate)
}

// WithCreatedDate adds a createdDate filter to the request list filter returning the filter for chaining
func (f *RequestListFilter) WithCreatedDate(createdDate string) *RequestListFilter {
	f.Add("filter.createdDate", createdDate)
	return f
}

// AddMethod adds a method filter to the request list filter
func (f *RequestListFilter) AddMethod(method string) {
	f.WithMethod(method)
}

// WithMethod adds a method filter to the request list filter returning the filter for chaining
func (f *RequestListFilter) WithMethod(method string) *RequestListFilter {
	f.Add("filter.method", method)
	return f
}

// AddBody adds a body filter to the request list filter
func (f *RequestListFilter) AddBody(body string) {
	f.WithBody(body)
}

// WithBody adds a body filter to the request list filter returning the filter for chaining
func (f *RequestListFilter) WithBody(body string) *RequestListFilter {
	f.Add("filter.body", body)
	return f
}

// AddRequestStatus adds a requestStatus filter to the request list filter
func (f *RequestListFilter) AddRequestStatus(requestStatus string) {
	f.WithRequestStatus(requestStatus)
}

// WithRequestStatus adds a requestStatus filter to the request list filter returning the filter for chaining
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

// ListRequestsWithFilter lists all requests that match the given filters
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

// GetRequestStatus returns status of the request
func (c *Client) GetRequestStatus(path string) (*RequestStatus, error) {
	url := path + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &RequestStatus{}
	err := c.client.GetRequestStatus(url, ret, http.StatusOK)
	return ret, err
}

// IsRequestFinished checks the given path to a request status resource. The request is considered "done"
// if its status won't change, which is true for status FAILED and DONE. Since Failed is obviously not done,
// the method returns true and RequestFailed error in that case.
func (c *Client) IsRequestFinished(path string) (bool, error) {
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
	done, err := c.IsRequestFinished(path)
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
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			done, err = c.IsRequestFinished(path)
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

// IsRequestStatusFinished is true if the requests Status is neither QUEUED or RUNNING.
func IsRequestStatusFinished(r Request) bool {
	switch r.Metadata.RequestStatus.Metadata.Status {
	case "QUEUED", "RUNNING":
		return false
	}
	return true
}

// WaitTillRequestsFinished will wait until there are no more unfinished requests matching the given filter
func (c *Client) WaitTillRequestsFinished(ctx context.Context, filter *RequestListFilter) error {
	return c.WaitTillMatchingRequestsFinished(ctx, filter, func(r Request) bool { return !IsRequestStatusFinished(r) })
}

// WaitTillMatchingRequestsFinished gets open requests with given filters and will
// wait for each request that is selected by the selector. The selector
// should consider filtering out requests that are finished. (e.g. using IsRequestStatusFinished)
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
