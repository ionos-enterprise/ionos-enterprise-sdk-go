package integration_tests

import (
	"fmt"
	"strings"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

var ipblkid string

func TestReserveIpBlock(t *testing.T) {
	fmt.Println("IP block tests")
	c := setupTestEnv()
	var obj = sdk.IPBlock{
		Properties: sdk.IPBlockProperties{
			Name:     "GO SDK Test",
			Size:     1,
			Location: location,
		},
	}

	resp, err := c.ReserveIPBlock(obj)
	ipblkid = resp.ID
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Size, 1)
	assert.Equal(t, resp.Properties.Location, location)

}

func TestReserveIpBlockFailure(t *testing.T) {
	c := setupTestEnv()
	var obj = sdk.IPBlock{
		Properties: sdk.IPBlockProperties{
			Name: "GO SDK Test",
			Size: 2,
		},
	}

	_, err := c.ReserveIPBlock(obj)
	if err == nil {
		t.Errorf("reserve IP block did not fail.")
	}
	assert.True(t, strings.Contains(err.Error(), "422"))
}

func TestListIpBlocks(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.ListIPBlocks()
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetIpBlock(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.GetIPBlock(ipblkid)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, ipblkid)
	assert.Equal(t, resp.PBType, "ipblock")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Size, 1)
	assert.Equal(t, resp.Properties.Location, location)
	assert.Equal(t, len(resp.Properties.IPs), 1)
}

func TestUpdateIpBlock(t *testing.T) {
	c := setupTestEnv()
	ipblock, err := c.GetIPBlock(ipblkid)
	if err != nil {
		t.Error(err)
	}

	req := sdk.IPBlockProperties{
		Name: "GO SDK Test RENAME",
	}
	resp, err := c.UpdateIPBlock(ipblkid, req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, ipblock.ID)
	assert.Equal(t, resp.Properties.Name, "GO SDK Test RENAME")
}

func TestGetIpBlockFailure(t *testing.T) {
	c := setupTestEnv()
	_, err := c.GetIPBlock("00000000-0000-0000-0000-000000000000")
	if err == nil {
		t.Errorf("get ip block did not fail")
	}
	assert.True(t, strings.Contains(err.Error(), "404"))
}

func TestReleaseIpBlock(t *testing.T) {
	c := setupTestEnv()
	_, err := c.ReleaseIPBlock(ipblkid)
	if err != nil {
		t.Error(err)
	}
}
