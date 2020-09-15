package profitbricks

import (
	"context"
	ionossdk "github.com/ionos-cloud/ionos-cloud-sdk-go/v5"
	"net/http"
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
	SourceIp       string `json:"sourceIp,omitempty"`
	TargetIp       string `json:"targetIp,omitempty"`
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

	rsp, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsFirewallrulesGet(context.TODO(), dcID, serverID, nicID, nil)

	ret := FirewallRules{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := firewallRulesPath(dcID, serverID, nicID)
	ret := &FirewallRules{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//GetFirewallRule gets a firewall rule
func (c *Client) GetFirewallRule(dcID string, serverID string, nicID string, fwID string) (*FirewallRule, error) {

	rsp, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsFirewallrulesFindById(context.TODO(), dcID, serverID, nicID, fwID, nil)
	ret := FirewallRule{}

	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}

	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := firewallRulePath(dcID, serverID, nicID, fwID)
	ret := &FirewallRule{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//CreateFirewallRule creates a firewall rule
func (c *Client) CreateFirewallRule(dcID string, serverID string, nicID string, fw FirewallRule) (*FirewallRule, error) {

	input := ionossdk.FirewallRule{}
	if errConvert := convertToCore(&fw, &input); errConvert != nil {
		return nil, errConvert
	}
	rsp, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsFirewallrulesPost(context.TODO(), dcID, serverID, nicID, input, nil)

	ret := FirewallRule{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}

	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := firewallRulesPath(dcID, serverID, nicID)
	ret := &FirewallRule{}
	err := c.Post(url, fw, ret, http.StatusAccepted)
	return ret, err
	 */
}

//UpdateFirewallRule updates a firewall rule
func (c *Client) UpdateFirewallRule(dcID string, serverID string, nicID string, fwID string, obj FirewallruleProperties) (*FirewallRule, error) {

	input := ionossdk.FirewallruleProperties{}
	if errConvert := convertToCore(&obj, &input); errConvert != nil {
		return nil, errConvert
	}
	rsp, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsFirewallrulesPatch(context.TODO(), dcID, serverID, nicID, fwID, input, nil)

	ret := FirewallRule{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := firewallRulePath(dcID, serverID, nicID, fwID)
	ret := &FirewallRule{}
	err := c.Patch(url, obj, ret, http.StatusAccepted)
	return ret, err

	 */
}

//DeleteFirewallRule deletes a firewall rule
func (c *Client) DeleteFirewallRule(dcID string, serverID string, nicID string, fwID string) (*http.Header, error) {

	_, apiResponse, err := c.CoreSdk.NicApi.DatacentersServersNicsFirewallrulesDelete(context.TODO(), dcID, serverID, nicID, fwID, nil)
	return &apiResponse.Header, err

	/*
	url := firewallRulePath(dcID, serverID, nicID, fwID)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}
