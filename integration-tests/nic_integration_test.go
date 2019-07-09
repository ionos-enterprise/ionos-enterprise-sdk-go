package integration_tests

import (
	"fmt"
	"sync"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

var (
	onceNicDc     sync.Once
	onceNicServer sync.Once
)

func TestCreateNic(t *testing.T) {
	fmt.Println("Nic tests")
	onceNicDc.Do(createDataCenter)
	onceNicServer.Do(createServer)
	onceNicNic.Do(createNic)

	assert.Equal(t, nic.Properties.Name, "GO SDK Test")
	assert.Equal(t, nic.Properties.Lan, 1)
}

func TestCreateNicFailure(t *testing.T) {
	c := setupTestEnv()
	onceNicDc.Do(createDataCenter)
	onceNicServer.Do(createServer)

	var request = sdk.Nic{
		Properties: &sdk.NicProperties{
			Name:           "GO SDK Test",
			Nat:            boolAddr(false),
			FirewallActive: boolAddr(true),
			Ips:            []string{"10.0.0.1"},
		},
	}

	_, err := c.CreateNic(dataCenter.ID, server.ID, request)

	assert.NotNil(t, err)
}

func TestListNics(t *testing.T) {
	c := setupTestEnv()
	onceNicDc.Do(createDataCenter)
	onceNicServer.Do(createServer)
	onceNicNic.Do(createNic)

	resp, err := c.ListNics(dataCenter.ID, server.ID)

	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetNic(t *testing.T) {
	c := setupTestEnv()
	onceNicDc.Do(createDataCenter)
	onceNicServer.Do(createServer)
	onceNicNic.Do(createNic)

	resp, err := c.GetNic(dataCenter.ID, server.ID, nic.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, nic.ID)
	assert.Equal(t, resp.PBType, "nic")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Lan, 1)
}

func TestGetNicFailure(t *testing.T) {
	c := setupTestEnv()
	onceNicDc.Do(createDataCenter)
	onceNicServer.Do(createServer)
	onceNicNic.Do(createNic)

	_, err := c.GetNic(dataCenter.ID, server.ID, "00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, err)
}

func TestUpdateNic(t *testing.T) {
	c := setupTestEnv()
	onceNicDc.Do(createDataCenter)
	onceNicServer.Do(createServer)
	onceNicNic.Do(createNic)

	obj := sdk.NicProperties{Name: "GO SDK Test - RENAME"}

	resp, err := c.UpdateNic(dataCenter.ID, server.ID, nic.ID, obj)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.Properties.Name, "GO SDK Test - RENAME")
}
func TestDeleteNic(t *testing.T) {
	c := setupTestEnv()
	onceNicDc.Do(createDataCenter)
	onceNicServer.Do(createServer)
	onceNicNic.Do(createNic)

	resp, err := c.DeleteNic(dataCenter.ID, server.ID, nic.ID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	resp, err = c.DeleteServer(dataCenter.ID, server.ID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	c.DeleteDatacenter(dataCenter.ID)
}
