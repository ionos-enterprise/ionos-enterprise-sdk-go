// loadbalancer_test.go
package profitbricks

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var lbal_dcid string
var lbalid string
var lbal_srvid string
var lbal_ipid string
var lbal_nic string
var ips []string

func TestCreateLoadbalancer(t *testing.T) {
	setupTestEnv()
	want := 202
	lbal_dcid = mkdcid("GO SDK Test")
	lbal_srvid = mksrvid(lbal_dcid)
	lbal_nic = mknic(lbal_dcid, lbal_srvid)
	var obj = IpBlock{
		Properties: IpBlockProperties{
			Size:     1,
			Location: "us/las",
		},
	}
	resp := ReserveIpBlock(obj)
	ips=resp.Properties.Ips
	waitTillProvisioned(resp.Headers.Get("Location"))
	lbal_ipid = resp.Id
	var request = Loadbalancer{
		Properties: LoadbalancerProperties{
			Name: "GO SDK Test",
			Ip:   resp.Properties.Ips[0],
			Dhcp: true,
		},
		Entities: LoadbalancerEntities{
			Balancednics: &BalancedNics{
				Items: []Nic{
					{
						Id: lbal_nic,
					},
				},

			},
		},
	}

	resp1 := CreateLoadbalancer(lbal_dcid, request)
	waitTillProvisioned(resp1.Headers.Get("Location"))
	lbalid = resp1.Id
	fmt.Println("Loadbalancer ID", lbalid)
	if resp1.StatusCode != want {
		t.Errorf(bad_status(want, resp1.StatusCode))
	}

	assert.Equal(t, resp1.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp1.Properties.Dhcp, true)
	resp1 = GetLoadbalancer(lbal_dcid, lbalid)
	assert.True(t, len(resp1.Entities.Balancednics.Items) > 0)
}

func TestCreateLoadbalancerFailure(t *testing.T) {
	want := 422
	var request = Loadbalancer{
		Properties: LoadbalancerProperties{
			Dhcp: true,
		},
	}

	resp := CreateLoadbalancer(lbal_dcid, request)

	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, strings.Contains(resp.Response, "Attribute 'name' is required"))
}

func TestListLoadbalancers(t *testing.T) {
	want := 200
	resp := ListLoadbalancers(lbal_dcid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetLoadbalancer(t *testing.T) {
	want := 200
	fmt.Println("TestGetLoadbalancer", lbalid)

	resp := GetLoadbalancer(lbal_dcid, lbalid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, lbalid)
	assert.Equal(t, resp.Type_, "loadbalancer")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Dhcp, true)
	assert.True(t, len(resp.Entities.Balancednics.Items) > 0)
}

func TestGetLoadbalancerFailure(t *testing.T) {
	want := 404
	fmt.Println("TestGetLoadbalancer", "00000000-0000-0000-0000-000000000000")

	resp := GetLoadbalancer(lbal_dcid, "00000000-0000-0000-0000-000000000000")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, strings.Contains(resp.Response, "Resource does not exist"))
}

func TestPatchLoadbalancer(t *testing.T) {
	want := 202

	obj := LoadbalancerProperties{Name: "GO SDK Test - RENAME"}

	resp := PatchLoadbalancer(lbal_dcid, lbalid, obj)
	waitTillProvisioned(resp.Headers.Get("Location"))
	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, lbalid)
	assert.Equal(t, resp.Type_, "loadbalancer")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test - RENAME")
}

func TestAssociateNic(t *testing.T) {
	want := 202

	nicid = mknic(lbal_dcid, lbal_srvid)
	fmt.Println("AssociateNic params ", lbal_dcid, lbalid, nicid)
	resp := AssociateNic(lbal_dcid, lbalid, nicid)
	waitTillProvisioned(resp.Headers.Get("Location"))
	nicid = resp.Id
	if resp.StatusCode != want {
		t.Error(resp.Response)
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
}

func TestGetBalancedNics(t *testing.T) {
	want := 200
	resp := ListBalancedNics(lbal_dcid, lbalid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetBalancedNic(t *testing.T) {
	want := 200
	resp := GetBalancedNic(lbal_dcid, lbalid, lbal_nic)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, lbal_nic)
	assert.Equal(t, resp.Type_, "nic")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Lan, 1)
	assert.Equal(t, resp.Properties.Nat, false)
	assert.Equal(t, resp.Properties.Dhcp, true)
	assert.Equal(t, resp.Properties.FirewallActive, true)
}

func TestDeleteBalancedNic(t *testing.T) {
	want := 202

	resp := DeleteBalancedNic(lbal_dcid, lbalid, lbal_nic)
	waitTillProvisioned(resp.Headers.Get("Location"))

	if resp.StatusCode != want {
		t.Error(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestDeleteLoadbalancer(t *testing.T) {
	want := 202
	resp := DeleteLoadbalancer(lbal_dcid, lbalid)
	waitTillProvisioned(resp.Headers.Get("Location"))
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestLoadBalancerCleanup(t *testing.T) {
	resp := DeleteDatacenter(lbal_dcid)
	waitTillProvisioned(resp.Headers.Get("Location"))
	DeleteDatacenter(dcID)
	ReleaseIpBlock(lbal_ipid)

}
