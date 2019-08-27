package profitbricks

import "strconv"

//Client object
type Client struct {
	client *client
}

//NewClient is a constructor for Client object
func NewClient(username, password string) *Client {
	c := newPBRestClient(username, password, "", "", "", true)
	return &Client{
		client: c,
	}
}

// NewClientbyToken is a constructor for Client object using bearer tokens for
// authentication instead of username, password
func NewClientbyToken(token string) *Client {
	c := newPBRestClientbyToken(token, "", "", "", true)
	return &Client{
		client: c,
	}
}

// SetDepth sets depth parameter for api calls
func (c *Client) SetDepth(depth int) {
	c.client.depth = strconv.Itoa(depth)
}

//SetUserAgent sets User-Agent request header for all API calls
func (c *Client) SetUserAgent(agent string) {
	c.client.agentHeader = agent
}

//GetUserAgent gets User-Agent header
func (c *Client) GetUserAgent() string {
	return c.client.agentHeader
}

//SetCloudApiUrl sets Cloud API URL
func (c *Client) SetCloudApiUrl(url string) {
	c.client.cloudApiUrl = url
}

//SetAuthApiUrl sets Auth API URL
func (c *Client) SetAuthApiUrl(url string) {
	c.client.authApiUrl = url
}
