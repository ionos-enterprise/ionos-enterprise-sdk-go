// lan_test.go
package profitbricks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var lan_dcid string
var lanid string
var lan_nic_srvid string
var lan_nic_id string
var reservedIp string

func TestCreateLan(t *testing.T) {
	setupTestEnv()
	lan_dcid = mkdcid("GO SDK Test")
	want := 202
	var request = Lan{
		Properties: LanProperties{
			Public: true,
			Name:   "GO SDK Test",
		},
	}
	lan := CreateLan(lan_dcid, request)
	waitTillProvisioned(lan.Headers.Get("Location"))
	lanid = lan.Id
	if lan.StatusCode != want {
		t.Errorf(bad_status(want, lan.StatusCode))
	}

	assert.Equal(t, lan.Type_, "lan")
	assert.Equal(t, lan.Properties.Name, "GO SDK Test")
	assert.True(t, lan.Properties.Public)
}

func TestCreateLanFailure(t *testing.T) {
	want := 422
	var request = Lan{
		Properties: LanProperties{
			Public: true,
		},
	}
	lan := CreateLan(lan_dcid, request)
	assert.Equal(t, lan.StatusCode, want)
}

func TestCreateCompositeLan(t *testing.T) {

	var obj = IpBlock{
		Properties: IpBlockProperties{
			Name:     "test",
			Size:     1,
			Location: "us/las",
		},
	}

	ipResponse := ReserveIpBlock(obj)
	reservedIp = ipResponse.Id

	lan_nic_srvid = mksrvid(lan_dcid)

	var nicRequest = Nic{
		Properties: &NicProperties{
			Lan:  1,
			Name: "Test NIC with failover",
			Nat:  false,
			Ips:  ipResponse.Properties.Ips,
		},
	}

	nicResponse := CreateNic(lan_dcid, lan_nic_srvid, nicRequest)
	waitTillProvisioned(nicResponse.Headers.Get("Location"))
	lan_nic_id = nicResponse.Id
	lanNics := LanNics{
		Items: []Nic{Nic{Id: lan_nic_id}},
	}

	want := 202
	var request = Lan{
		Properties: LanProperties{
			Public: true,
			Name:   "Lan Test with failover",
		},
		Entities: &LanEntities{
			Nics: &lanNics,
		},
	}
	lan := CreateLan(lan_dcid, request)
	waitTillProvisioned(lan.Headers.Get("Location"))

	lanUpdate := LanProperties{IpFailover: []IpFailover{IpFailover{
		Ip:    ipResponse.Id,
		nicId: nicResponse.Id,
	}}, }

	lanPatch := PatchLan(lan_dcid, lan.Id, lanUpdate)
	if lanPatch.StatusCode != want {
		t.Errorf(bad_status(want, lanPatch.StatusCode))
	}
	if lan.StatusCode != want {
		t.Errorf(bad_status(want, lan.StatusCode))
	}

	assert.Equal(t, lan.Type_, "lan")
	assert.Equal(t, lan.Properties.Name, "GO SDK Test")
	assert.True(t, lan.Properties.Public)
}

func TestListLans(t *testing.T) {
	want := 200
	lans := ListLans(lan_dcid)

	if lans.StatusCode != want {
		t.Errorf(bad_status(want, lans.StatusCode))
	}
	assert.True(t, len(lans.Items) > 0)
}

func TestGetLan(t *testing.T) {
	want := 200
	lan := GetLan(lan_dcid, lanid)

	if lan.StatusCode != want {
		t.Errorf(bad_status(want, lan.StatusCode))
	}

	assert.Equal(t, lan.Id, lanid)
	assert.Equal(t, lan.Type_, "lan")
	assert.Equal(t, lan.Properties.Name, "GO SDK Test")
	assert.True(t, lan.Properties.Public)
}

func TestPatchLan(t *testing.T) {
	want := 202
	obj := LanProperties{Name:"GO SDK Test - RENAME",Public: false}

	lan := PatchLan(lan_dcid, lanid, obj)
	if lan.StatusCode != want {
		t.Errorf(bad_status(want, lan.StatusCode))
	}

	assert.Equal(t, lan.Id, lanid)
	assert.Equal(t, lan.Type_, "lan")
	assert.Equal(t, lan.Properties.Name, "GO SDK Test - RENAME")
	assert.False(t, lan.Properties.Public)
}

func TestDeleteLan(t *testing.T) {
	want := 202
	resp := DeleteLan(lan_dcid, lanid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestLanCleanup(t *testing.T) {
	DeleteLan(lan_dcid, lanid)
	DeleteDatacenter(lan_dcid)
	ReleaseIpBlock(reservedIp)
}
