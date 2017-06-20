// ipblock_test.go
package profitbricks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var ipblkid string

func TestReserveIpBlock(t *testing.T) {
	setupTestEnv()
	want := 202
	var obj = IpBlock{
		Properties: IpBlockProperties{
			Name:     "GO SDK Test",
			Size:     2,
			Location: location,
		},
	}

	resp := ReserveIpBlock(obj)
	ipblkid = resp.Id
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Size, 2)
	assert.Equal(t, resp.Properties.Location, location)

}

func TestListIpBlocks(t *testing.T) {
	want := 200
	resp := ListIpBlocks()
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetIpBlock(t *testing.T) {
	want := 200
	resp := GetIpBlock(ipblkid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, ipblkid)
	assert.Equal(t, resp.Type_, "ipblock")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Size, 2)
	assert.Equal(t, resp.Properties.Location, location)
	assert.Equal(t, len(resp.Properties.Ips), 2)
}

func TestReleaseIpBlock(t *testing.T) {
	want := 202
	resp := ReleaseIpBlock(ipblkid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
