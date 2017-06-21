// lan_test.go
package profitbricks

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
)

var lan_dcid string
var lanid string
var lanfailoverid string
var lan_nic_srvid string
var lan_nic_id string
var reservedIp []string
var ipblockId2 string

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
	//lan gets created even if no paramters are passed
	//want := 422
	//var request = Lan{
	//	Properties: LanProperties{
	//		Public: true,
	//	},
	//}
	//lan := CreateLan(lan_dcid, request)
	////assert.Equal(t, lan.StatusCode, want)
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
	ipblockId2 = ipResponse.Id
	reservedIp = ipResponse.Properties.Ips

	lan_nic_srvid = mksrvid(lan_dcid)

	var nicRequest = Nic{
		Properties: &NicProperties{
			Lan:  1,
			Name: "Test NIC with failover",
			Nat:  false,
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
			Name:   "GO SDK Test with failover",
		},
		Entities: &LanEntities{
			Nics: &lanNics,
		},
	}
	lan := CreateLan(lan_dcid, request)
	lanfailoverid = lan.Id
	waitTillProvisioned(lan.Headers.Get("Location"))

	if lan.StatusCode != want {
		t.Errorf(bad_status(want, lan.StatusCode))
	}

	assert.Equal(t, lan.Type_, "lan")
	assert.Equal(t, lan.Properties.Name, "GO SDK Test with failover")
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

	ip := reservedIp[0]
	//pareparing to add the ipfailover feature
	//creating a lan
	var request = Lan{
		Properties: LanProperties{
			Public: true,
			Name:   "GO SDK Test with failover",
		},
	}
	failoverLan := CreateLan(lan_dcid, request)
	waitTillProvisioned(failoverLan.Headers.Get("Location"))
	//creating a server
	failover_server := mksrvid(lan_dcid)

	failoverlanid, err := strconv.Atoi(failoverLan.Id)
	if err != nil {
		//do error
	}

	//creating an nic attached to the failover server
	failover_nic := mknic_custom(lan_dcid, failover_server, failoverlanid, []string{ip})

	obj := LanProperties{
		IpFailover: []IpFailover{IpFailover{
			Ip:      ip,
			NicUuid: failover_nic,
		}}, }

	lan := PatchLan(lan_dcid, failoverLan.Id, obj)
	if lan.StatusCode != want {
		t.Errorf(bad_status(want, lan.StatusCode))
	}

	assert.Equal(t, lan.Id, failoverLan.Id)
	assert.Equal(t, lan.Type_, "lan")
	assert.Equal(t, lan.Properties.Name, "GO SDK Test with failover")
	assert.True(t, len(lan.Properties.IpFailover) > 0)
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
	deleted := DeleteDatacenter(lan_dcid)
	waitTillProvisioned(deleted.Headers.Get("Location"))
	ReleaseIpBlock(ipblockId2)
}
