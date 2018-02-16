package profitbricks

import (
	"net/http"
	"strconv"
)

//FirewallRule object
type FirewallRule struct {
	ID         string                 `json:"id,omitempty"`
	PBType     string                 `json:"type,omitempty"`
	Href       string                 `json:"href,omitempty"`
	Metadata   *Metadata              `json:"metadata,omitempty"`
	Properties FirewallruleProperties `json:"properties,omitempty"`
	Response   string                 `json:"Response,omitempty"`
	Headers    *http.Header           `json:"headers,omitempty"`
	StatusCode int                    `json:"statuscode,omitempty"`
}

//FirewallruleProperties object
type FirewallruleProperties struct {
	Name           string  `json:"name,omitempty"`
	Protocol       string  `json:"protocol,omitempty"`
	SourceMac      *string `json:"sourceMac,omitempty"`
	SourceIP       *string `json:"sourceIp,omitempty"`
	TargetIP       *string `json:"targetIp,omitempty"`
	IcmpCode       *int    `json:"icmpCode,omitempty"`
	IcmpType       *int    `json:"icmpType,omitempty"`
	PortRangeStart *int    `json:"portRangeStart,omitempty"`
	PortRangeEnd   *int    `json:"portRangeEnd,omitempty"`
}

//FirewallRules object
type FirewallRules struct {
	ID         string         `json:"id,omitempty"`
	PBType     string         `json:"type,omitempty"`
	Href       string         `json:"href,omitempty"`
	Items      []FirewallRule `json:"items,omitempty"`
	Response   string         `json:"Response,omitempty"`
	Headers    *http.Header   `json:"headers,omitempty"`
	StatusCode int            `json:"statuscode,omitempty"`
}

//ListFirewallRules lists all firewall rules
func (c *Client) ListFirewallRules(dcID string, serverID string, nicID string) (*FirewallRules, error) {
	url := fwruleColPath(dcID, serverID, nicID) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &FirewallRules{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//GetFirewallRule gets a firewall rule
func (c *Client) GetFirewallRule(dcID string, serverID string, nicID string, fwID string) (*FirewallRule, error) {
	url := fwrulePath(dcID, serverID, nicID, fwID) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &FirewallRule{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//CreateFirewallRule creates a firewall rule
func (c *Client) CreateFirewallRule(dcID string, serverID string, nicID string, fw FirewallRule) (*FirewallRule, error) {
	url := fwruleColPath(dcID, serverID, nicID) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &FirewallRule{}
	err := c.client.Post(url, fw, ret, http.StatusAccepted)
	return ret, err
}

//UpdateFirewallRule updates a firewall rule
func (c *Client) UpdateFirewallRule(dcID string, serverID string, nicID string, fwID string, obj FirewallruleProperties) (*FirewallRule, error) {
	url := fwrulePath(dcID, serverID, nicID, fwID) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &FirewallRule{}
	err := c.client.Patch(url, obj, ret, http.StatusAccepted)
	return ret, err
}

//DeleteFirewallRule deletes a firewall rule
func (c *Client) DeleteFirewallRule(dcID string, serverID string, nicID string, fwID string) (*http.Header, error) {
	url := fwrulePath(dcID, serverID, nicID, fwID) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}
