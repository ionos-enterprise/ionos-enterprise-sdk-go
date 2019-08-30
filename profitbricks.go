package profitbricks

import (
	"net/http"
	"strconv"
	"time"
)
import "gopkg.in/resty.v1"

type Client struct {
	*resty.Client
	// AuthApi will be used by methods talking to the auth api by sending absolute urls
	AuthApi string
}

const (
	DefaultApiUrl = "https://api.ionos.com/cloudapi/v5"
	DefaultAuthUrl = "https://api.ionos.com/auth/v1"
)

func RestyClient(username, password, token string) *Client {
	c := &Client{
		Client: resty.New(),
		AuthApi: DefaultAuthUrl,
	}
	if token == "" {
		c.SetBasicAuth(username, password)
	} else {
		c.SetAuthToken(token)
	}
	c.SetDebug(true)
	c.SetHostURL(DefaultApiUrl)
	c.SetDepth(5)
	c.SetUserAgent("profitbricks-sdk-go " + Version)
	c.SetRetryCount(1)
	c.AddRetryCondition(
		func(r *resty.Response) (bool, error) {
			if r.StatusCode() == http.StatusTooManyRequests {
				dur, err := time.ParseDuration(r.Header().Get("Retry-After") + "s")
				if err != nil {
					return false, err
				}
				c.SetRetryWaitTime(dur)
				c.SetRetryCount(1)
				return true, nil
			}
			return false, nil
		})
	return c
}

func (c *Client) SetDepth(depth int) {
	c.Client.SetQueryParam("depth", strconv.Itoa(depth))
}

func (c *Client) SetPretty(pretty bool) {
	c.Client.SetQueryParam("pretty", strconv.FormatBool(pretty))
}

// NewClient is a constructor for Client object
func NewClient(username, password string) *Client {
	return RestyClient(username, password, "")
}

// NewClientbyToken is a constructor for Client object using bearer tokens for
// authentication instead of username, password
func NewClientbyToken(token string) *Client {
	return RestyClient("", "", token)
}

// SetUserAgent sets User-Agent request header for all API calls
func (c *Client) SetUserAgent(agent string) {
	c.Client.SetHeader("User-Agent", agent)
}

// GetUserAgent gets User-Agent header
func (c *Client) GetUserAgent() string {
	return c.Client.Header.Get("User-Agent")
}

// SetCloudApiURL sets Cloud API url
func (c *Client) SetCloudApiURL(url string) {
	c.Client.SetHostURL(url)
}

func (c *Client) SetAuthApiUrl(url string) {
	c.AuthApi = url
}
