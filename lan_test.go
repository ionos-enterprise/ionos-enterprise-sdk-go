// lan_test.go
package profitbricks

import (
	"testing"
)

var lan_dcid string
var lanid string
var lan_nic_srvid string
var lan_nic_id string
var reservedIp string

func TestCreateLan(t *testing.T) {
	setupTestEnv()
	lan_dcid = mkdcid("GO SDK LAN DC")
	want := 202
	var request = Lan{
		Properties: LanProperties{
			Public: true,
			Name:   "Lan Test",
		},
	}
	lan := CreateLan(lan_dcid, request)
	waitTillProvisioned(lan.Headers.Get("Location"))
	lanid = lan.Id
	if lan.StatusCode != want {
		t.Errorf(bad_status(want, lan.StatusCode))
	}
}

func TestCreateLanWithIpFailover(t *testing.T) {

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
}

func TestListLans(t *testing.T) {
	want := 200
	lans := ListLans(lan_dcid)

	if lans.StatusCode != want {
		t.Errorf(bad_status(want, lans.StatusCode))
	}
}

func TestGetLan(t *testing.T) {
	want := 200
	lan := GetLan(lan_dcid, lanid)

	if lan.StatusCode != want {
		t.Errorf(bad_status(want, lan.StatusCode))
	}
}

func TestPatchLan(t *testing.T) {
	want := 202
	obj := LanProperties{Public: false}

	lan := PatchLan(lan_dcid, lanid, obj)
	if lan.StatusCode != want {
		t.Errorf(bad_status(want, lan.StatusCode))
	}
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
