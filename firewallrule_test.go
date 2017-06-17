package profitbricks

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var fwId string

func setup() {
	datacenter := Datacenter{
		Properties: DatacenterProperties{
			Name:     "composite test",
			Location: location,
		},
		Entities: DatacenterEntities{
			Servers: &Servers{
				Items: []Server{
					Server{
						Properties: ServerProperties{
							Name:  "server1",
							Ram:   2048,
							Cores: 1,
						},
						Entities: &ServerEntities{
							Nics: &Nics{
								Items: []Nic{
									Nic{
										Properties: &NicProperties{
											Name: "SSH",
											Lan:  1,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	datacenter = CompositeCreateDatacenter(datacenter)
	waitTillProvisioned(datacenter.Headers.Get("Location"))

	dcID = datacenter.Id
	srv_srvid = datacenter.Entities.Servers.Items[0].Id
	nicid = datacenter.Entities.Servers.Items[0].Entities.Nics.Items[0].Id
}
func TestCreateFirewallRule(t *testing.T) {
	setupTestEnv()
	want := 202
	setup()

	fw := FirewallRule{
		Properties: FirewallruleProperties{
			Name:           "SSH",
			Protocol:       "TCP",
			SourceMac:      &sourceMac,
			PortRangeStart: &portRangeStart,
			PortRangeEnd:   &portRangeEnd,
		},
	}

	fw = CreateFirewallRule(dcID, srv_srvid, nicid, fw)

	waitTillProvisioned(fw.Headers.Get("Location"))

	if fw.StatusCode != want {
		t.Error(fw.Response)
		t.Errorf(bad_status(want, fw.StatusCode))
	}
	fwId = fw.Id

	assert.Equal(t, fw.Properties.Name, "SSH")
	assert.Equal(t, fw.Properties.Protocol, "TCP")
	assert.Equal(t, *fw.Properties.SourceMac, sourceMac)
	assert.Equal(t, *fw.Properties.PortRangeStart, portRangeStart)
	assert.Equal(t, *fw.Properties.PortRangeEnd, portRangeEnd)
	assert.Nil(t, fw.Properties.SourceIp)
	assert.Nil(t, fw.Properties.IcmpCode)
	assert.Nil(t, fw.Properties.IcmpType)

}

func TestCreateFirewallRuleFailure(t *testing.T) {
	want := 422

	fw := FirewallRule{
		Properties: FirewallruleProperties{
			Name:           "SSH",
			SourceMac:      &sourceMac,
			PortRangeStart: &portRangeStart,
			PortRangeEnd:   &portRangeEnd,
		},
	}

	fw = CreateFirewallRule(dcID, srv_srvid, nicid, fw)

	if fw.StatusCode != want {
		fmt.Println(string(fw.Response))
		t.Errorf(bad_status(want, fw.StatusCode))
	}

	assert.True(t, strings.Contains(fw.Response, "Attribute 'protocol' is required"))
}

func TestGetFirewallRule(t *testing.T) {
	want := 200

	fw := GetFirewallRule(dcID, srv_srvid, nicid, fwId)
	if fw.StatusCode != want {
		t.Error(fw.Response)
		t.Errorf(bad_status(want, fw.StatusCode))
	}

	assert.Equal(t, fw.Id, fwId)
	assert.Equal(t, fw.Properties.Name, "SSH")
	assert.Equal(t, fw.Properties.Protocol, "TCP")
	assert.Equal(t, *fw.Properties.SourceMac, sourceMac)
	assert.Equal(t, *fw.Properties.PortRangeStart, portRangeStart)
	assert.Equal(t, *fw.Properties.PortRangeEnd, portRangeEnd)
	assert.Nil(t, fw.Properties.SourceIp)
	assert.Nil(t, fw.Properties.IcmpCode)
	assert.Nil(t, fw.Properties.IcmpType)
}

func TestGetFirewallRuleFailure(t *testing.T) {
	want := 404

	fw := GetFirewallRule(dcID, srv_srvid, nicid, "00000000-0000-0000-0000-000000000000")
	if fw.StatusCode != want {
		t.Errorf(bad_status(want, fw.StatusCode))
	}

	assert.True(t, strings.Contains(fw.Response, "Resource does not exist"))
}

func TestListFirewallRules(t *testing.T) {
	want := 200
	fws := ListFirewallRules(dcID, srv_srvid, nicid)
	if fws.StatusCode != want {
		t.Error(fws.Response)
		t.Errorf(bad_status(want, fws.StatusCode))
	}

	assert.True(t, len(fws.Items) > 0)
}

func TestPatchFirewallRule(t *testing.T) {
	want := 202
	props := FirewallruleProperties{
		Name: "SSH - RENAME",
	}
	fw := PatchFirewallRule(dcID, srv_srvid, nicid, fwId, props)
	if fw.StatusCode != want {
		t.Error(fw.Response)
		t.Errorf(bad_status(want, fw.StatusCode))
	}

	assert.Equal(t, fw.Id, fwId)
	assert.Equal(t, fw.Properties.Name, "SSH - RENAME")
	assert.Equal(t, fw.Type_, "firewall-rule")
}

func TestDeleteFirewallRule(t *testing.T) {
	want := 202
	resp := DeleteFirewallRule(dcID, srv_srvid, nicid, fwId)

	if resp.StatusCode != want {
		t.Error(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	DeleteDatacenter(dcID)
}
