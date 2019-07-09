package integration_tests

import (
	"fmt"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

var lanid string
var lanfailoverid string
var lannicsrvid string
var lannicid string
var ipblockID2 string

func TestCreateLan(t *testing.T) {
	fmt.Println("Lan tests")
	c := setupTestEnv()
	onceLan.Do(createDataCenter)

	var request = sdk.Lan{
		Properties: sdk.LanProperties{
			Public: true,
			Name:   "GO SDK Test",
		},
	}

	lan, err := c.CreateLan(dataCenter.ID, request)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(lan.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	lanid = lan.ID

	assert.Equal(t, lan.PBType, "lan")
	assert.Equal(t, lan.Properties.Name, "GO SDK Test")
	assert.True(t, lan.Properties.Public)
}

func TestCreateLanFailure(t *testing.T) {
	c := setupTestEnv()
	var request = sdk.Lan{
		Properties: sdk.LanProperties{
			Public: true,
		},
	}
	_, err := c.CreateLan("00000000-0000-0000-0000-000000000000", request)
	assert.NotNil(t, err)
}

func TestCreateCompositeLan(t *testing.T) {
	c := setupTestEnv()
	onceLan.Do(createDataCenter)
	onceLanServer.Do(createCompositeServerFW)

	var obj = sdk.IPBlock{
		Properties: sdk.IPBlockProperties{
			Name:     "test",
			Size:     1,
			Location: "us/las",
		},
	}

	ipResponse, err := c.ReserveIPBlock(obj)
	ipblockID2 = ipResponse.ID

	lannicsrvid = nic.ID

	var nicRequest = sdk.Nic{
		Properties: &sdk.NicProperties{
			Lan:  1,
			Name: "Test NIC with failover",
			Nat:  boolAddr(false),
		},
	}

	nicResponse, err := c.CreateNic(dataCenter.ID, server.ID, nicRequest)
	if err != nil {
		t.Error(err)
	}
	err = c.WaitTillProvisioned(nicResponse.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	lannicid = nicResponse.ID
	lanNics := sdk.LanNics{
		Items: []sdk.Nic{{ID: nic.ID}},
	}

	var request = sdk.Lan{
		Properties: sdk.LanProperties{
			Public: true,
			Name:   "GO SDK Test with failover",
		},
		Entities: &sdk.LanEntities{
			Nics: &lanNics,
		},
	}
	lan, err := c.CreateLan(dataCenter.ID, request)
	if err != nil {
		t.Error(err)
	}
	lanfailoverid = lan.ID

	err = c.WaitTillProvisioned(lan.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, lan.PBType, "lan")
	assert.Equal(t, lan.Properties.Name, "GO SDK Test with failover")
	assert.True(t, lan.Properties.Public)
}

func TestListLans(t *testing.T) {
	c := setupTestEnv()
	onceLan.Do(createDataCenter)
	onceLanServer.Do(createCompositeServerFW)
	lans, err := c.ListLans(dataCenter.ID)

	if err != nil {
		t.Error(err)
	}
	assert.True(t, len(lans.Items) > 0)
}

func TestGetLan(t *testing.T) {
	c := setupTestEnv()
	onceLan.Do(createDataCenter)
	onceLanLan.Do(createLan)
	lan, err := c.GetLan(dataCenter.ID, lan.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, lan.ID, lan.ID)
	assert.Equal(t, lan.PBType, "lan")
}

func TestUpdateLan(t *testing.T) {
	c := setupTestEnv()
	onceLan.Do(createDataCenter)
	onceLanLan.Do(createLan)
	obj := sdk.LanProperties{
		Name: "newName",
	}

	lan, err := c.UpdateLan(dataCenter.ID, lan.ID, obj)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(lan.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, lan.ID, lan.ID)
	assert.Equal(t, lan.PBType, "lan")
	assert.Equal(t, lan.Properties.Name, obj.Name)
}

func TestDeleteLan(t *testing.T) {
	onceLan.Do(createDataCenter)
	onceLanLan.Do(createLan)
	c := setupTestEnv()
	_, err := c.DeleteLan(dataCenter.ID, lan.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestLanCleanup(t *testing.T) {
	c := setupTestEnv()

	deleted, err := c.DeleteDatacenter(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}
	c.WaitTillProvisioned(deleted.Get("Location"))

	c.ReleaseIPBlock(ipblockID2)
}
