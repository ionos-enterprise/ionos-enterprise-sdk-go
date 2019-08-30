package profitbricks

import (
	"net/http"
)

// FirewallRule object
type FirewallRule struct {
	BaseResource `json:",inline"`
	ID           string                 `json:"id,omitempty"`
	PBType       string                 `json:"type,omitempty"`
	Href         string                 `json:"href,omitempty"`
	Metadata     *Metadata              `json:"metadata,omitempty"`
	Properties   FirewallruleProperties `json:"properties,omitempty"`
	Response     string                 `json:"Response,omitempty"`
	StatusCode   int                    `json:"statuscode,omitempty"`
}

// FirewallruleProperties object
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

// FirewallRules object
type FirewallRules struct {
	BaseResource `json:",inline"`
	ID           string         `json:"id,omitempty"`
	PBType       string         `json:"type,omitempty"`
	Href         string         `json:"href,omitempty"`
	Items        []FirewallRule `json:"items,omitempty"`
	Response     string         `json:"Response,omitempty"`
	StatusCode   int            `json:"statuscode,omitempty"`
}

// ListFirewallRules lists all firewall rules
func (c *Client) ListFirewallRules(dcID string, serverID string, nicID string) (*FirewallRules, error) {
	ret := &FirewallRules{}
	return ret, c.GetOK(firewallRulesPath(dcID, serverID, nicID), ret)
}

// GetFirewallRule gets a firewall rule
func (c *Client) GetFirewallRule(dcID string, serverID string, nicID string, fwID string) (*FirewallRule, error) {
	ret := &FirewallRule{}
	return ret, c.GetOK(firewallRulePath(dcID, serverID, nicID, fwID), ret)
}

// CreateFirewallRule creates a firewall rule
func (c *Client) CreateFirewallRule(
	dcID string, serverID string, nicID string, fw FirewallRule) (*FirewallRule, error) {
	ret := &FirewallRule{}
	return ret, c.PostAcc(firewallRulesPath(dcID, serverID, nicID), fw, ret)

}

// UpdateFirewallRule updates a firewall rule
func (c *Client) UpdateFirewallRule(
	dcID string, serverID string, nicID string, fwID string, obj FirewallruleProperties) (*FirewallRule, error) {
	ret := &FirewallRule{}
	return ret, c.PatchAcc(firewallRulePath(dcID, serverID, nicID, fwID), obj, ret)
}

// DeleteFirewallRule deletes a firewall rule
func (c *Client) DeleteFirewallRule(dcID string, serverID string, nicID string, fwID string) (*http.Header, error) {
	return c.DeleteAcc(firewallRulePath(dcID, serverID, nicID, fwID))
}
