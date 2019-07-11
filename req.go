package profitbricks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	u "net/url"
	"reflect"
	"strings"
	"time"
)

const (
	//FullHeader is the standard header to include with all http requests except is_patch and is_command
	FullHeader = "application/json"

	//AgentHeader is used for user agent request header
	AgentHeader   = "profitbricks-sdk-go/5.0.0"
	DefaultApiUrl = "https://api.profitbricks.com/cloudapi/v4"
)

type client struct {
	username    string
	password    string
	depth       string
	pretty      bool
	apiURL      string
	agentHeader string
	token       string
}

func newPBRestClient(username string, password string, apiURL string, depth string, pretty bool) *client {
	client := new(client)
	client.username = username
	client.password = password
	client.agentHeader = AgentHeader
	if apiURL == "" {
		client.apiURL = DefaultApiUrl
	} else {
		client.apiURL = apiURL
	}

	if depth == "" {
		client.depth = "5"
	} else {
		client.depth = depth
	}

	return client
}

func newPBRestClientbyToken(token, apiURL, depth string, pretty bool) *client {
	client := new(client)
	client.token = token
	client.agentHeader = "profitbricks-sdk-go/5.0.0"
	if apiURL == "" {
		client.apiURL = "https://api.profitbricks.com/cloudapi/v4"
	} else {
		client.apiURL = apiURL
	}

	if depth == "" {
		client.depth = "5"
	} else {
		client.depth = depth
	}

	return client
}

func (c *client) mkURL(path string) string {
	url := c.apiURL + path

	return url
}

func (c *client) do(url string, method string, requestBody interface{}, result interface{}, expectedStatus int) error {
	var bodyData io.Reader
	if requestBody != nil {
		if method == "POST" && (strings.HasSuffix(url, "create-snapshot") || strings.HasSuffix(url, "restore-snapshot")) {
			data := requestBody.(u.Values)
			bodyData = bytes.NewBufferString(data.Encode())
		} else {
			data, err := json.Marshal(requestBody)
			if err != nil {
				return err
			}
			bodyData = bytes.NewBuffer(data)
		}
	}

	r, err := http.NewRequest(method, url, bodyData)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(url, "stop") && !strings.HasSuffix(url, "start") && !strings.HasSuffix(url, "reboot") && !strings.HasSuffix(url, "create-snapshot") && !strings.HasSuffix(url, "restore-snapshot") {
		r.Header.Add("Content-Type", FullHeader)
	}

	r.Header.Add("User-Agent", c.agentHeader)

	var br *bytes.Reader
	if r.Body != nil {
		buf, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			return err
		}
		br = bytes.NewReader(buf)
	}

	for {
		if br != nil {
			_, err := br.Seek(0, 0)
			if err != nil {
				return err
			}

			r.Body = ioutil.NopCloser(br)
		}

		client := &http.Client{}
		if c.token != "" {
			r.Header.Add("Authorization", "Bearer "+c.token)
		} else {
			r.SetBasicAuth(c.username, c.password)
		}
		resp, err := client.Do(r)
		if err != nil {
			return err
		}

		if resp != nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			if resp.StatusCode == http.StatusTooManyRequests {
				retryAfter := resp.Header.Get("Retry-After")
				if retryAfter == "" {
					return err
				}

				sleep, err := time.ParseDuration(retryAfter + "s")
				if err != nil {
					return err
				}

				time.Sleep(sleep)
			} else if resp.StatusCode != expectedStatus {
				erResp := &errorResponse{}

				err = json.Unmarshal(body, erResp)
				if err != nil {
					return err
				}
				return ApiError{*erResp}
			} else {

				if string(body) != "" {
					err = json.Unmarshal(body, result)
					val := reflect.ValueOf(result).Elem().FieldByName("Headers")
					val.Set(reflect.ValueOf(&resp.Header))
				} else {
					raw, err := json.Marshal(resp.Header)
					if err != nil {
						return err
					}
					json.Unmarshal(raw, result)
				}

				return err
			}
		}
	}
}

func (c *client) Get(url string, result interface{}, expectedStatus int) error {
	return c.do(c.mkURL(url), "GET", nil, result, expectedStatus)
}

func (c *client) GetRequestStatus(url string, result interface{}, expectedStatus int) error {
	return c.do(url, "GET", nil, result, expectedStatus)
}

func (c *client) Delete(url string, result interface{}, expectedStatus int) error {
	return c.do(c.mkURL(url), "DELETE", nil, result, expectedStatus)
}

func (c *client) Post(url string, requestBody interface{}, result interface{}, expectedStatus int) error {
	return c.do(c.mkURL(url), "POST", requestBody, result, expectedStatus)
}

func (c *client) Put(url string, requestBody interface{}, result interface{}, expectedStatus int) error {
	return c.do(c.mkURL(url), "PUT", requestBody, result, expectedStatus)
}
func (c *client) Patch(url string, requestBody interface{}, result interface{}, expectedStatus int) error {
	return c.do(c.mkURL(url), "PATCH", requestBody, result, expectedStatus)
}

type errorResponse struct {
	HTTPStatus int `json:"httpStatus"`
	Messages   []struct {
		ErrorCode string `json:"errorCode"`
		Message   string `json:"message"`
	} `json:"messages"`
}

type ApiError struct {
	response errorResponse
}

func (e ApiError) Error() string {
	return e.response.String()
}

func (e errorResponse) String() string {
	toReturn := fmt.Sprintf("HTTP Status: %s \n%s", fmt.Sprint(e.HTTPStatus), "Error Messages:")
	for _, m := range e.Messages {
		toReturn = toReturn + fmt.Sprintf("Error Code: %s Message: %s\n", m.ErrorCode, m.Message)
	}
	return toReturn
}

func (e ApiError) HttpStatusCode() int {
	return e.response.HTTPStatus
}
