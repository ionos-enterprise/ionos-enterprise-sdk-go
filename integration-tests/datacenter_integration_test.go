package integration_tests

import (
	"fmt"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateDataCenter(t *testing.T) {
	fmt.Println("DataCenter tests")
	syncDC.Do(createDataCenter)

	assert.Equal(t, dataCenter.PBType, "datacenter")
	assert.Equal(t, dataCenter.Properties.Name, "GO SDK Test")
	assert.Equal(t, dataCenter.Properties.Description, "GO SDK test datacenter")
	assert.Equal(t, dataCenter.Properties.Location, location)
}

func TestListDatacenters(t *testing.T) {
	c := setupTestEnv()

	resp, err := c.ListDatacenters()

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.True(t, len(resp.Items) > 0)
}

func TestGetDatacenterFail(t *testing.T) {
	c := setupTestEnv()
	_, err := c.GetDatacenter("231")
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateFailure(t *testing.T) {
	c := setupTestEnv()
	var obj = sdk.Datacenter{
		Properties: sdk.DatacenterProperties{
			Name:        "GO SDK Test",
			Description: "GO SDK test datacenter",
		},
	}
	_, err := c.CreateDatacenter(obj)
	if err == nil {
		t.Error(err)
		t.Fail()
	}

}

func TestCreateComposite(t *testing.T) {
	syncCDC.Do(createCompositeDataCenter)

	assert.Equal(t, compositeDataCenter.PBType, "datacenter")
	assert.Equal(t, compositeDataCenter.Properties.Name, "GO SDK Test Composite")
	assert.Equal(t, compositeDataCenter.Properties.Description, "GO SDK test composite datacenter")
	assert.Equal(t, compositeDataCenter.Properties.Location, location)
	assert.True(t, len(compositeDataCenter.Entities.Servers.Items) > 0)
	assert.True(t, len(compositeDataCenter.Entities.Volumes.Items) > 0)
}

func TestGetDatacenter(t *testing.T) {
	syncDC.Do(createDataCenter)
	c := setupTestEnv()
	resp, err := c.GetDatacenter(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.ID, dataCenter.ID)
	assert.Equal(t, resp.PBType, "datacenter")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Description, "GO SDK test datacenter")
	assert.Equal(t, resp.Properties.Location, location)
}

func TestUpdateDatacenter(t *testing.T) {
	syncDC.Do(createDataCenter)
	c := setupTestEnv()
	newName := "GO SDK Test - RENAME"
	obj := sdk.DatacenterProperties{Name: newName}

	resp, err := c.UpdateDataCenter(dataCenter.ID, obj)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, dataCenter.ID)
	assert.Equal(t, resp.Properties.Name, newName)
}

func TestDeleteDatacenter(t *testing.T) {
	syncDC.Do(createDataCenter)
	syncCDC.Do(createCompositeDataCenter)
	c := setupTestEnv()
	_, err := c.DeleteDatacenter(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = c.DeleteDatacenter(compositeDataCenter.ID)
	if err != nil {
		t.Error(err)
	}
}
