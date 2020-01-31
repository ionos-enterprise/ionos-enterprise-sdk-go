package profitbricks

import (
	"fmt"
	"net/http"
	"reflect"

	resty "github.com/go-resty/resty/v2"
)

func (c *Client) Do(url, method string, body, result interface{}, expectedStatus int) error {
	req := c.R()
	if body != nil {
		req.SetBody(body)
	}
	if result != nil {
		req.SetResult(result)
	}
	return c.DoWithRequest(req, method, url, expectedStatus)
}

func (c *Client) DoWithRequest(request *resty.Request, method, url string, expectedStatus int) error {
	rsp, err := request.SetError(ApiError{}).Execute(method, url)
	if err != nil {
		return NewClientError(HttpClientError, fmt.Sprintf("[%s] %s: Client error %s", method, url, err))
	}
	if result := rsp.Result(); result != nil {
		if val := reflect.ValueOf(result).Elem().FieldByName("Headers"); val.IsValid() {
			h := rsp.Header()
			val.Set(reflect.ValueOf(&h))
		}
	}
	return validateResponse(rsp, expectedStatus)
}
func (c *Client) Get(url string, item interface{}, expectedStatus int) error {
	req := c.R().SetResult(item)
	var cached cachable
	if _, ok := item.(cachable); ok {

		// if someone messed with depth, we simply do not use the cached entries
		if depth, err := c.GetDepth(); err == nil {
			cached = c.cache.Get(url, depth)
		}
		if cached != nil {
			if md := cached.GetMetadata(); md != nil {
				req.SetHeader("If-None-Match", md.Etag)
			}
		}
	}

	rsp, err := req.Get(url)
	if err != nil {
		return NewClientError(HttpClientError, fmt.Sprintf("[%s] %s: Client error %s", "GET", url, err))
	}
	if rsp.StatusCode() == http.StatusNotModified {
		reflect.ValueOf(item).Elem().Set(reflect.ValueOf(cached).Elem())
		c.cacheHits++
		return nil
	}

	if err := validateResponse(rsp, expectedStatus); err != nil {
		return err
	}
	// if someone messed with depth, we do not cache the result
	if depth, err := c.GetDepth(); err == nil {
		c.cache.Add(url, depth, item.(cachable))
	}

	return nil
}
func (c *Client) GetOK(url string, result interface{}) error {
	return c.Get(url, result, http.StatusOK)
}

func (c *Client) Post(
	url string, body interface{}, result interface{}, expectedStatus int) error {
	return c.Do(url, resty.MethodPost, body, result, expectedStatus)
}

func (c *Client) PostAcc(url string, body, result interface{}) error {
	return c.Do(url, resty.MethodPost, body, result, http.StatusAccepted)
}

func (c *Client) PatchAcc(url string, body, result interface{}) error {
	return c.Do(url, resty.MethodPatch, body, result, http.StatusAccepted)
}

func (c *Client) Patch(url string, body, result interface{}, expectedStatus int) error {
	return c.Do(url, resty.MethodPatch, body, result, expectedStatus)
}

func (c *Client) PutAcc(url string, body, result interface{}) error {
	return c.Do(url, resty.MethodPut, body, result, http.StatusAccepted)
}

func (c *Client) Put(url string, body, result interface{}, expectedStatus int) error {
	return c.Do(url, resty.MethodPut, body, result, expectedStatus)
}

func (c *Client) DeleteAcc(url string) (*http.Header, error) {
	h := &http.Header{}
	return h, c.Delete(url, h, http.StatusAccepted)
}

func (c *Client) Delete(url string, responseHeader *http.Header, expectedStatus int) error {
	rsp, err := c.R().SetError(ApiError{}).Delete(url)
	if err != nil {
		return NewClientError(HttpClientError, fmt.Sprintf("[DELETE] %s: Client error: %s", url, err))
	}
	if responseHeader != nil {
		*responseHeader = rsp.Header()
	}
	return validateResponse(rsp, expectedStatus)
}

func validateResponse(rsp *resty.Response, expectedStatus ...int) error {
	for _, exp := range expectedStatus {
		if rsp.StatusCode() == exp {
			return nil
		}
	}
	if rsp.StatusCode() >= 400 {
		e := rsp.Error().(*ApiError)
		e.RawBody = rsp.Body()
		return *e

	}
	return NewClientError(UnexpectedResponse, fmt.Sprintf("[%s] %s: Unexpected status %s",
		rsp.Request.Method, rsp.Request.URL, rsp.Status()))
}
