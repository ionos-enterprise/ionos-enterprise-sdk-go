package integration_tests

import (
	"fmt"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateFirewallRule(t *testing.T) {
	fmt.Println("FirewallRule tests")
	c := setupTestEnv()
	onceDC.Do(createDataCenter)
	onceFw.Do(createCompositeServerFW)
	sm := "01:23:45:67:89:11"
	start := 23
	end := 23
	fw := &sdk.FirewallRule{
		Properties: sdk.FirewallruleProperties{
			Name:           "SSH",
			Protocol:       "TCP",
			SourceMac:      &sm,
			PortRangeStart: &start,
			PortRangeEnd:   &end,
		},
	}

	fw, err := c.CreateFirewallRule(dataCenter.ID, server.ID, nic.ID, *fw)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(fw.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, fw.Properties.Name, "SSH")
	assert.Equal(t, fw.Properties.Protocol, "TCP")
	assert.Equal(t, *fw.Properties.SourceMac, sm)
	assert.Equal(t, *fw.Properties.PortRangeStart, start)
	assert.Equal(t, *fw.Properties.PortRangeEnd, end)
	assert.Nil(t, fw.Properties.SourceIP)
	assert.Nil(t, fw.Properties.IcmpCode)
	assert.Nil(t, fw.Properties.IcmpType)

}

func TestCreateFirewallRuleFailure(t *testing.T) {
	c := setupTestEnv()
	onceDC.Do(createDataCenter)
	onceFw.Do(createCompositeServerFW)
	fw := sdk.FirewallRule{
		Properties: sdk.FirewallruleProperties{
			Name:           "SSH",
			SourceMac:      &sourceMac,
			PortRangeStart: &portRangeStart,
			PortRangeEnd:   &portRangeEnd,
		},
	}

	_, err := c.CreateFirewallRule(dataCenter.ID, server.ID, nic.ID, fw)

	assert.NotNil(t, err)
}

func TestGetFirewallRule(t *testing.T) {
	c := setupTestEnv()
	onceDC.Do(createDataCenter)
	onceFw.Do(createCompositeServerFW)

	resp, err := c.GetFirewallRule(dataCenter.ID, server.ID, nic.ID, fw.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, fw.ID)
	assert.Equal(t, resp.Properties.Name, "SSH")
	assert.Equal(t, resp.Properties.Protocol, "TCP")
	assert.Equal(t, *resp.Properties.SourceMac, sourceMac)
	assert.Equal(t, *resp.Properties.PortRangeStart, portRangeStart)
	assert.Equal(t, *resp.Properties.PortRangeEnd, portRangeEnd)
	assert.Nil(t, resp.Properties.SourceIP)
	assert.Nil(t, resp.Properties.IcmpCode)
	assert.Nil(t, resp.Properties.IcmpType)
}

func TestGetFirewallRuleFailure(t *testing.T) {
	c := setupTestEnv()

	_, err := c.GetFirewallRule("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, err)
}

func TestListFirewallRules(t *testing.T) {
	c := setupTestEnv()
	onceDC.Do(createDataCenter)
	onceFw.Do(createCompositeServerFW)
	fws, err := c.ListFirewallRules(dataCenter.ID, server.ID, nic.ID)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(fws.Items) > 0)
}

func TestUpdateFirewallRule(t *testing.T) {
	c := setupTestEnv()
	onceDC.Do(createDataCenter)
	onceFw.Do(createCompositeServerFW)

	props := sdk.FirewallruleProperties{
		Name: "SSH - RENAME",
	}
	resp, err := c.UpdateFirewallRule(dataCenter.ID, server.ID, nic.ID, fw.ID, props)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, fw.ID)
	assert.Equal(t, resp.Properties.Name, "SSH - RENAME")
}

func TestDeleteFirewallRule(t *testing.T) {
	c := setupTestEnv()
	onceDC.Do(createDataCenter)
	onceFw.Do(createCompositeServerFW)

	resp, err := c.DeleteFirewallRule(dataCenter.ID, server.ID, nic.ID, fw.ID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	c.DeleteDatacenter(dataCenter.ID)
}
