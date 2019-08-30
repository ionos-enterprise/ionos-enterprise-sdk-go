package profitbricks

import (
	"net/http"
	"regexp"

	"gopkg.in/resty.v1"
)

var rPathIds = regexp.MustCompile(`\{\s*([^\s]+?)\s*\}`)

type BaseResource struct {
	Headers *http.Header `json:"headers,omitempty"`
}

type HeadersHaver interface {
	SetHeaders(http.Header)
	GetHeaders() *http.Header
}

func (b *BaseResource) SetHeaders(hd http.Header) {
	cp := make(http.Header, len(hd))
	for k, v := range hd {
		values := make([]string, len(v))
		copy(values, v)
		cp[k] = values
	}
	b.Headers = &cp
}

func (b *BaseResource) GetHeaders() *http.Header {
	return b.Headers
}

func getPathParams(url string) []string {
	res := rPathIds.FindAllStringSubmatch(url, -1)
	if res == nil {
		return nil
	}
	results := make([]string, len(res))
	for i, r := range res {
		results[i] = r[1]
	}
	return results
}

func (c *Client) Do(
	url, method string, body interface{}, result HeadersHaver, expectedStatus int) error {
	request := c.R().SetResult(result).SetBody(body)
	return c.DoWithRequest(request, method, url, expectedStatus)
}

func (c *Client) DoWithRequest(request *resty.Request, method, url string, expectedStatus int) error {
	rsp, err := request.SetError(&ApiErrorResponse{}).Execute(method, url)
	if err != nil {
		return ClientError.Wrapf(err, "[%s] %s: Client error", method, url)
	}
	if result := rsp.Result(); result != nil {
		if r, ok := result.(HeadersHaver); ok {
			r.SetHeaders(rsp.Header())
		}
	}
	return validateResponse(rsp, expectedStatus)
}

func (c *Client) GetOK(url string, result HeadersHaver) error {
	return c.Do(url, resty.MethodGet, nil, result, http.StatusOK)
}

func (c *Client) Get(url string, result HeadersHaver, expectedStatus int) error {
	return c.Do(url, resty.MethodGet, nil, result, expectedStatus)
}

func (c *Client) Post(
	url string, body interface{}, result HeadersHaver, expectedStatus int) error {
	return c.Do(url, resty.MethodPost, body, result, expectedStatus)
}

func (c *Client) PostAcc(url string, body interface{}, result HeadersHaver) error {
	return c.Do(url, resty.MethodPost, body, result, http.StatusAccepted)
}

func (c *Client) PatchAcc(url string, body interface{}, result HeadersHaver) error {
	return c.Do(url, resty.MethodPatch, body, result, http.StatusAccepted)
}

func (c *Client) Patch(
	url string, body interface{}, result HeadersHaver, expectedStatus int) error {
	return c.Do(url, resty.MethodPatch, body, result, expectedStatus)
}

func (c *Client) PutAcc(url string, body interface{}, result HeadersHaver) error {
	return c.Do(url, resty.MethodPut, body, result, http.StatusAccepted)
}

func (c *Client) Put(
	url string, body interface{}, result HeadersHaver, expectedStatus int, pathParams ...string) error {
	return c.Do(url, resty.MethodPut, body, result, expectedStatus)
}

func (c *Client) DeleteAcc(url string) (*http.Header, error) {
	return c.Delete(url, http.StatusAccepted)
}

func (c *Client) Delete(url string, expectedStatus int) (*http.Header, error) {
	rsp, err := c.R().Delete(url)
	if err != nil {
		return nil, ClientError.Wrapf(err, "[DELETE] %s: Client error", url)
	}
	h := rsp.Header()
	return &h, validateResponse(rsp, expectedStatus)
}

func validateResponse(rsp *resty.Response, expectedStatus ...int) error {
	for _, exp := range expectedStatus {
		if rsp.StatusCode() == exp {
			return nil
		}
	}
	if rsp.StatusCode() >= 400 {
		return ErrorType(rsp.StatusCode()).Wrapf(rsp.Error().(*ApiErrorResponse), "[%s] %s: got error %s",
			rsp.Request.Method, rsp.Request.URL, rsp.Status())

	}
	return ErrorType(rsp.StatusCode()).Newf("[%s] %s: Unexpected status %s",
		rsp.Request.Method, rsp.Request.URL, rsp.Status())
}
