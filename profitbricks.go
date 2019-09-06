package profitbricks

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty"
)

var DebugHTTP = false

type Client struct {
	*resty.Client
	// AuthApiUrl will be used by methods talking to the auth api by sending absolute urls
	AuthApiUrl  string
	CloudApiUrl string
}

var Version = "5.0.2"

const (
	DefaultApiUrl  = "https://api.ionos.com/cloudapi/v5"
	DefaultAuthUrl = "https://api.ionos.com/auth/v1"
)

func RestyClient(username, password, token string) *Client {
	c := &Client{
		Client:      resty.New(),
		AuthApiUrl:  DefaultAuthUrl,
		CloudApiUrl: DefaultApiUrl,
	}
	if token == "" {
		c.SetBasicAuth(username, password)
	} else {
		c.SetAuthToken(token)
	}
	c.SetDebug(DebugHTTP)
	c.SetHostURL(DefaultApiUrl)
	c.SetDepth(5)
	c.SetTimeout(3 * time.Minute)
	c.SetUserAgent("ionos-enterprise-sdk-go " + Version)
	c.SetRetryCount(1)
	c.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			if r.StatusCode() == http.StatusTooManyRequests {
				retryAfter := r.Header().Get("Retry-After")
				dur := 1 * time.Second
				var err error
				if retryAfter != "" {
					dur, err = time.ParseDuration(retryAfter + "s")
					if err != nil {
						return false
					}
				}
				c.SetRetryWaitTime(dur)
				c.SetRetryCount(1)
				return true
			}
			return false
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
	c.AuthApiUrl = url
}
