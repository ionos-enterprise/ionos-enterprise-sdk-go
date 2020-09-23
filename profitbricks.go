package profitbricks

import (
	"context"
	"github.com/ionos-cloud/ionos-cloud-sdk-go/v5"
	"strconv"
	"time"
)

type Client struct {
	// *resty.Client
	CoreSdk *ionossdk.APIClient
	// AuthApiUrl will be used by methods talking to the auth api by sending absolute urls
	AuthApiUrl  string
	CloudApiUrl string
	Config Configuration
}

type Configuration struct {
	Timeout time.Duration
}

const (
	DefaultApiUrl  = "https://api.ionos.com/cloudapi/v5"
	DefaultAuthUrl = "https://api.ionos.com/auth/v1"
	Version        = "5.1.0"
)

func RestyClient(username, password, token string) *Client {
	c := &Client{
		// Client:      resty.New(),
		CoreSdk:     ionossdk.NewAPIClient(ionossdk.NewConfiguration(username, password, token)),
		AuthApiUrl:  DefaultAuthUrl,
		CloudApiUrl: DefaultApiUrl,
	}

	/*
	if token == "" {
		c.SetBasicAuth(username, password)
	} else {
		c.SetAuthToken(token)
	}*/
	// c.SetHostURL(DefaultApiUrl)
	c.SetDepth(10)
	c.SetTimeout(3 * time.Minute)
	c.SetUserAgent("ionos-enterprise-sdk-go-compat " + Version)
	/*
	c.SetRetryCount(3)
	c.SetRetryMaxWaitTime(10 * time.Minute)
	c.SetRetryWaitTime(1 * time.Second)
	c.SetRetryAfter(func(cl *resty.Client, r *resty.Response) (time.Duration, error) {
		switch r.StatusCode() {
		case http.StatusTooManyRequests:
			if retryAfterSeconds := r.Header().Get("Retry-After"); retryAfterSeconds != "" {
				return time.ParseDuration(retryAfterSeconds + "s")
			}
		}
		return cl.RetryWaitTime, nil
	})
	c.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			switch r.StatusCode() {
			case http.StatusTooManyRequests,
				http.StatusServiceUnavailable,
				http.StatusGatewayTimeout,
				http.StatusBadGateway:
				return true
			}
			return false
		})
	 */
	return c
}

func (c *Client) SetTimeout(t time.Duration) {
	c.Config.Timeout = t
}

// SetDebug activates/deactivates resty's debug mode. For better readability
// the pretty print feature is also enabled.
func (c *Client) SetDebug(debug bool) {
	// c.Client.SetDebug(debug)
	c.SetPretty(debug)
}

// SetDepth sets the depth of information that will be retrieved by api calls. The
// API accepts values from 0 to 10, a low depth means mostly only IDs and hrefs will be
// returned. Therefore nested structures may be nil.
func (c *Client) SetDepth(depth int) {
	c.CoreSdk.GetConfig().AddDefaultQueryParam("depth", strconv.Itoa(depth))
}

// SetPretty toggles if the data retrieved from the api will be delivered pretty printed.
// Usually this does not make sense from an sdk perspective, but for debugging it's nice
// therefore it is also set to true, if debug is enabled.
func (c *Client) SetPretty(pretty bool) {
	c.CoreSdk.GetConfig().AddDefaultQueryParam("pretty", strconv.FormatBool(pretty))
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
	c.CoreSdk.GetConfig().UserAgent = agent
}

// GetUserAgent gets User-Agent header
func (c *Client) GetUserAgent() string {
	return c.CoreSdk.GetConfig().UserAgent
}

// SetCloudApiURL sets Cloud API url
func (c *Client) SetCloudApiURL(url string) {
	c.CoreSdk.GetConfig().BasePath = url
}

// SetAuthApiUrl sets the Auth API url
func (c *Client) SetAuthApiUrl(url string) {
	c.AuthApiUrl = url
}

func (c *Client) GetContext() (context.Context, context.CancelFunc) {
	if c.Config.Timeout > 0 {
		return context.WithTimeout(context.Background(), c.Config.Timeout)
	}

	return context.Background(), nil
}
