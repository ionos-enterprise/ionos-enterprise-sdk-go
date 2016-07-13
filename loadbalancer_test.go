// loadbalancer_test.go
package profitbricks

import (
	"fmt"
	"testing"
)

var lbal_dcid string
var lbalid string
var lbal_srvid string

func TestCreateLoadbalancer(t *testing.T) {
	setupCredentials()
	want := 202
	lbal_dcid = mkdcid("GO SDK LB DC")
	lbal_srvid = mksrvid(lbal_dcid)
	var request = Loadbalancer{
		Properties: LoadbalancerProperties{
			Name: "test",
			Ip:   "127.0.0.0",
			Dhcp: true,
		},
	}
	resp := CreateLoadbalancer(lbal_dcid, request)
	waitTillProvisioned(resp.Headers.Get("Location"))
	lbalid = resp.Id
	fmt.Println("Loadbalancer ID", lbalid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

}

func TestListLoadbalancers(t *testing.T) {
	want := 200
	resp := ListLoadbalancers(lbal_dcid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestGetLoadbalancer(t *testing.T) {
	want := 200
	fmt.Println("TestGetLoadbalancer", lbalid)

	resp := GetLoadbalancer(lbal_dcid, lbalid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestPatchLoadbalancer(t *testing.T) {
	want := 202

	obj := LoadbalancerProperties{Name: "Renamed Loadbalancer"}

	resp := PatchLoadbalancer(lbal_dcid, lbalid, obj)
	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestAssociateNic(t *testing.T) {
	want := 202

	nicid = mknic(lbal_dcid, lbal_srvid)

	resp := AssociateNic(lbal_dcid, lbalid, nicid)
	waitTillProvisioned(resp.Headers.Get("Location"))
	nicid = resp.Id
	if resp.StatusCode != want {
		t.Error(resp.Response)
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestGetBalancedNics(t *testing.T) {
	want := 200
	resp := ListBalancedNics(lbal_dcid, lbalid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestGetBalancedNic(t *testing.T) {
	want := 200
	resp := GetBalancedNic(lbal_dcid, lbalid, nicid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestDeleteBalancedNic(t *testing.T) {
	want := 202

	resp := DeleteBalancedNic(lbal_dcid, lbalid, nicid)

	if resp.StatusCode != want {
		t.Error(string(resp.Body))
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
	DeleteDatacenter(dcID)
}
