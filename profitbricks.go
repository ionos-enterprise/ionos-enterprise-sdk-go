package profitbricks

import "strconv"

//Client object
type Client struct {
	baseURL  string
	userName string
	password string
	client   *client
}

//NewClient is a constructor for Client object
func NewClient(username, password string) *Client {

	c := newPBRestClient(username, password, "", "", true)
	return &Client{
		userName: c.username,
		password: c.password,
		client:   c,
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

//SetURL sets Cloud API url
func (c *Client) SetURL(url string) {
	c.client.apiURL = url
}
