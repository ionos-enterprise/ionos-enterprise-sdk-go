// loadbalancer_test.go
package profitbricks

import (
	"fmt"
	"testing"
	"time"
)

var lbal_dcid string
var lbalid string

func TestCreateLoadbalancer(t *testing.T) {
	//t.Parallel()
	want := 202
	lbal_dcid = mkdcid()
	var jason = []byte(`{"properties": {"name":"Goat"}}`)
	resp := CreateLoadbalancer(lbal_dcid, jason)
	lbalid = resp.Id
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

	time.Sleep(20 * time.Second)
}

func TestListLoadbalancers(t *testing.T) {
	//t.Parallel()
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

	resp := GetLoadbalancer(lbal_dcid, lbalid)

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestPatchLoadbalancer(t *testing.T) {
	//t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed Loadbalancer"
					}`)
	resp := PatchLoadbalancer(lbal_dcid, lbalid, jason_patch)
	if resp.Resp.StatusCode != want {
		fmt.Println(resp.Resp.Body)
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestDeleteLoadbalancer(t *testing.T) {
	//t.Parallel()
	want := 202
	resp := DeleteLoadbalancer(lbal_dcid, lbalid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestLoadBalancerCleanup(t *testing.T) {
	DeleteDatacenter(lbal_dcid)
}
