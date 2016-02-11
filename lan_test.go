// lan_test.go
package profitbricks

import (
	"testing"
	"time"
)

var lan_dcid string
var lanid string

func TestCreateLan(t *testing.T) {
	lan_dcid = mkdcid("GO SDK LAN DC")
	want := 202
	var request = CreateLanRequest{
		LanProperties: LanProperties{
			Public: true,
			Name:   "Lan Test",
		},
	}
	lan := CreateLan(lan_dcid, request)
	lanid = lan.Id
	if lan.Resp.StatusCode != want {
		t.Errorf(bad_status(want, lan.Resp.StatusCode))
	}
	time.Sleep(20 * time.Second)
}

func TestListLans(t *testing.T) {
	shouldbe := "collection"
	want := 200
	lans := ListLans(lan_dcid)

	if lans.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, lans.Type))
	}
	if lans.Resp.StatusCode != want {
		t.Errorf(bad_status(want, lans.Resp.StatusCode))
	}
}

func TestGetLan(t *testing.T) {
	shouldbe := "lan"
	want := 200
	lan := GetLan(lan_dcid, lanid)
	if lan.Type != shouldbe {
		t.Errorf(bad_status(want, lan.Resp.StatusCode))
	}
	if lan.Resp.StatusCode != want {
		t.Errorf(bad_status(want, lan.Resp.StatusCode))
	}
}

func TestPatchLan(t *testing.T) {
	want := 202
	obj := map[string]string{"public": "false"}

	lan := PatchLan(lan_dcid, lanid, obj)
	if lan.Resp.StatusCode != want {
		t.Errorf(bad_status(want, lan.Resp.StatusCode))
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
	DeleteDatacenter(lan_dcid)
}
