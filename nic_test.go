package profitbricks

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var nic_dcid string
var nic_srvid string
var nicid string

func TestCreateNic(t *testing.T) {
	setupTestEnv()
	nic_dcid = mkdcid("GO SDK NIC DC")
	nic_srvid = mksrvid(nic_dcid)

	want := 202
	var request = Nic{
		Properties: &NicProperties{
			Lan:            1,
			Name:           "GO SDK Test",
			Nat:            false,
			FirewallActive: true,
			Ips:            []string{"10.0.0.1"},
		},
	}

	resp := CreateNic(nic_dcid, nic_srvid, request)
	waitTillProvisioned(resp.Headers.Get("Location"))
	nicid = resp.Id
	if resp.StatusCode != want {
		t.Error(resp.Response)
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Lan, 1)
	assert.Equal(t, resp.Properties.Nat, false)
	assert.Equal(t, resp.Properties.Dhcp, request.Properties.Dhcp)
	assert.Equal(t, resp.Properties.FirewallActive, true)
	assert.Equal(t, resp.Properties.Ips, []string{"10.0.0.1"})
}

func TestCreateNicFailure(t *testing.T) {
	want := 422
	var request = Nic{
		Properties: &NicProperties{
			Name:           "GO SDK Test",
			Nat:            false,
			FirewallActive: true,
			Ips:            []string{"10.0.0.1"},
		},
	}

	resp := CreateNic(nic_dcid, nic_srvid, request)
	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, strings.Contains(resp.Response, "Attribute 'lan' is required"))
}

func TestListNics(t *testing.T) {
	//t.Parallel()
	want := 200
	resp := ListNics(nic_dcid, nic_srvid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetNic(t *testing.T) {
	want := 200
	resp := GetNic(nic_dcid, nic_srvid, nicid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, nicid)
	assert.Equal(t, resp.Type_, "nic")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Lan, 1)
	assert.Equal(t, resp.Properties.Nat, false)
	assert.Equal(t, *resp.Properties.Dhcp, true)
	assert.Equal(t, resp.Properties.FirewallActive, true)
	assert.Equal(t, resp.Properties.Ips, []string{"10.0.0.1"})
}

func TestGetNicFailure(t *testing.T) {
	want := 404
	resp := GetNic(nic_dcid, nic_srvid, "00000000-0000-0000-0000-000000000000")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, strings.Contains(resp.Response, "Resource does not exist"))
}

func TestPatchNic(t *testing.T) {
	want := 202
	obj := NicProperties{Name: "GO SDK Test - RENAME", Lan: 1}

	resp := PatchNic(nic_dcid, nic_srvid, nicid, obj)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, nicid)
	assert.Equal(t, resp.Type_, "nic")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test - RENAME")
}
func TestDeleteNic(t *testing.T) {
	want := 202
	resp := DeleteNic(nic_dcid, nic_srvid, nicid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestNicCleanup(t *testing.T) {
	DeleteServer(nic_dcid, nic_srvid)
	DeleteDatacenter(nic_dcid)
}
