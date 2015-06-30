package profitbricks

import (
	"testing"
	"time"
)

var nic_dcid string
var nic_srvid string
var nicid string

func TestCreateNic(t *testing.T) {
	nic_dcid = mkdcid("NIC DC")
	nic_srvid = mksrvid(nic_dcid)
	time.Sleep(15 * time.Second)
	want := 202
	var jason = []byte(`{"properties": {"name":"Original Nic","lan":1}}`)

	resp := CreateNic(nic_dcid, nic_srvid, jason)
	nicid = resp.Id
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
	time.Sleep(20 * time.Second)
}

func TestListNics(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListNics(nic_dcid, nic_srvid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestGetNic(t *testing.T) {
	//t.Parallel()
	shouldbe := "nic"
	want := 200
	resp := GetNic(nic_dcid, nic_srvid, nicid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestPatchNic(t *testing.T) {
	//t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Patched Nic",
					"lan":1
					}`)

	resp := PatchNic(nic_dcid, nic_srvid, nicid, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
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
