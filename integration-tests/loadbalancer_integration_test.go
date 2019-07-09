package integration_tests

import (
	"fmt"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateLoadbalancer(t *testing.T) {
	fmt.Println("Load balancer tests")
	onceLBDC.Do(createDataCenter)
	onceLBServer.Do(createServer)
	onceLBNic.Do(createNic)
	onceLB.Do(createLoadBalancerWithIP)

	assert.Equal(t, loadBalancer.Properties.Name, "GO SDK Test")
	assert.Equal(t, loadBalancer.Properties.Dhcp, true)
}

func TestCreateLoadbalancerFailure(t *testing.T) {
	c := setupTestEnv()
	var request = sdk.Loadbalancer{
		Properties: sdk.LoadbalancerProperties{
			Dhcp: true,
		},
	}

	_, err := c.CreateLoadbalancer("00000000-0000-0000-0000-000000000000", request)

	assert.NotNil(t, err)
}

func TestListLoadbalancers(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.ListLoadbalancers(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetLoadbalancer(t *testing.T) {
	c := setupTestEnv()
	onceLBDC.Do(createDataCenter)
	onceLBServer.Do(createServer)
	onceLBNic.Do(createNic)
	onceLB.Do(createLoadBalancerWithIP)
	resp, err := c.GetLoadbalancer(dataCenter.ID, loadBalancer.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, loadBalancer.ID)
	assert.Equal(t, resp.PBType, "loadbalancer")
	assert.Equal(t, resp.Properties.Name, loadBalancer.Properties.Name)
	assert.Equal(t, resp.Properties.Dhcp, true)
	assert.True(t, len(resp.Entities.Balancednics.Items) > 0)
}

func TestGetLoadbalancerFailure(t *testing.T) {
	c := setupTestEnv()

	_, err := c.GetLoadbalancer("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, err)
}

func TestPatchLoadbalancer(t *testing.T) {
	c := setupTestEnv()
	onceLBDC.Do(createDataCenter)
	onceLBServer.Do(createServer)
	onceLBNic.Do(createNic)
	onceLB.Do(createLoadBalancerWithIP)

	obj := sdk.LoadbalancerProperties{Name: "GO SDK Test - RENAME"}

	resp, err := c.UpdateLoadbalancer(dataCenter.ID, loadBalancer.ID, obj)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, loadBalancer.ID)
	assert.Equal(t, resp.PBType, "loadbalancer")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test - RENAME")
}

func TestAssociateNic(t *testing.T) {
	c := setupTestEnv()
	onceLBDC.Do(createDataCenter)
	onceLBServer.Do(createServer)
	onceLBNic.Do(createNic)

	resp, err := c.AssociateNic(dataCenter.ID, loadBalancer.ID, nic.ID)
	if err != nil {
		t.Error(err)
	}
	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.Properties.Name, loadBalancer.Properties.Name)
}

func TestGetBalancedNics(t *testing.T) {
	c := setupTestEnv()
	onceLBDC.Do(createDataCenter)
	onceLBServer.Do(createServer)
	onceLBNic.Do(createNic)
	onceLB.Do(createLoadBalancerWithIP)

	resp, err := c.ListBalancedNics(dataCenter.ID, loadBalancer.ID)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetBalancedNic(t *testing.T) {
	c := setupTestEnv()
	onceLBDC.Do(createDataCenter)
	onceLBServer.Do(createServer)
	onceLBNic.Do(createNic)
	onceLB.Do(createLoadBalancerWithIP)

	resp, err := c.GetBalancedNic(dataCenter.ID, loadBalancer.ID, nic.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, nic.ID)
	assert.Equal(t, resp.PBType, "nic")
	assert.Equal(t, resp.Properties.Lan, 2)
	assert.Equal(t, *resp.Properties.Nat, false)
	assert.Equal(t, *resp.Properties.Dhcp, true)
}

func TestDeleteBalancedNic(t *testing.T) {
	c := setupTestEnv()
	onceLBDC.Do(createDataCenter)
	onceLBServer.Do(createServer)
	onceLBNic.Do(createNic)
	onceLB.Do(createLoadBalancerWithIP)

	resp, err := c.DeleteBalancedNic(dataCenter.ID, loadBalancer.ID, nic.ID)
	if err != nil {
		t.Error(err)
	}
	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteLoadbalancer(t *testing.T) {
	onceLBDC.Do(createDataCenter)
	onceLBServer.Do(createServer)
	onceLBNic.Do(createNic)
	onceLB.Do(createLoadBalancerWithIP)
	c := setupTestEnv()
	resp, err := c.DeleteLoadbalancer(dataCenter.ID, loadBalancer.ID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	resp, err = c.ReleaseIPBlock(ipBlock.ID)
	if err != nil {
		t.Error(err)
	}

	resp, err = c.DeleteDatacenter(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}

}
