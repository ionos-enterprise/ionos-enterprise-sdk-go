// loadbalancer_test.go
package profitbricks

import (
	"fmt"
	"testing"
	"time"
)

var lbal_dcid string
var lbalid string
var lbal_srvid string

func TestCreateLoadbalancer(t *testing.T) {
	fmt.Println("TestCreateLoadbalancer")
	want := 202
	lbal_dcid = mkdcid("GO SDK LB DC")
	lbal_srvid = mksrvid(lbal_dcid)
	var request = LoablanacerCreateRequest{
		LoablanacerProperties: LoablanacerProperties{
			Name: "test",
			Ip:   "127.0.0.0",
			Dhcp: true,
		},
	}
	time.Sleep(30 * time.Second)

	resp := CreateLoadbalancer(lbal_dcid, request)
	lbalid = resp.Id
	fmt.Println("Loadbalancer ID", lbalid)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

}

func TestListLoadbalancers(t *testing.T) {
	shouldbe := "collection"
	want := 200
	resp := ListLoadbalancers(lbal_dcid)

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestGetLoadbalancer(t *testing.T) {
	shouldbe := "loadbalancer"
	want := 200
	fmt.Println("TestGetLoadbalancer", lbalid)
	time.Sleep(120 * time.Second)

	resp := GetLoadbalancer(lbal_dcid, lbalid)

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestPatchLoadbalancer(t *testing.T) {
	want := 202

	obj := map[string]string{"name": "Renamed Loadbalancer"}

	resp := PatchLoadbalancer(lbal_dcid, lbalid, obj)
	if resp.Resp.StatusCode != want {
		fmt.Println(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestAssociateNic(t *testing.T) {
	want := 202

	nicid = mknic(lbal_dcid, lbal_srvid)

	time.Sleep(120 * time.Second)

	resp := AssociateNic(lbal_dcid, lbalid, nicid)
	nicid = resp.Id
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
	time.Sleep(45 * time.Second)
}

func TestGetBalancedNics(t *testing.T) {
	want := 200
	shouldbe := "collection"
	resp := ListBalancedNics(lbal_dcid, lbalid)

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestGetBalancedNic(t *testing.T) {
	want := 200
	shouldbe := "nic"
	resp := GetBalancedNic(lbal_dcid, lbalid, nicid)

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestDeleteBalancedNic(t *testing.T) {

	want := 202

	resp := DeleteBalancedNic(lbal_dcid, lbalid, nicid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestDeleteLoadbalancer(t *testing.T) {
	want := 202
	resp := DeleteLoadbalancer(lbal_dcid, lbalid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestLoadBalancerCleanup(t *testing.T) {
	DeleteDatacenter(lbal_dcid)
}
