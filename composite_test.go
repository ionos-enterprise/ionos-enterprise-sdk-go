package profitbricks

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var ipId string

func TestCompositeCreate(t *testing.T) {
	setupTestEnv()
	location := "us/las"

	ipblockreq := IpBlock{
		Properties: IpBlockProperties{
			Size:     1,
			Location: location,
		},
	}

	ipblockresp := ReserveIpBlock(ipblockreq)
	ipId = ipblockresp.Id
	fmt.Println(ipId)
	datacenter := Datacenter{
		Properties: DatacenterProperties{
			Name:     "composite test",
			Location: location,
		},
		Entities: DatacenterEntities{
			Lans: &Lans{
				Items: []Lan{
					Lan{
						Properties: LanProperties{
							Public: true,
						},
					},
				},
			},
			Loadbalancers: &Loadbalancers{
				Items: []Loadbalancer{
					Loadbalancer{
						Properties: LoadbalancerProperties{
							Name: "test",
						},
					},
				},
			},
		},
	}

	fmt.Println(ipblockresp.Properties.Ips)
	dc := CompositeCreateDatacenter(datacenter)
	dcID = dc.Id
	waitTillProvisioned(dc.Headers.Get("Location"))
	SetDepth("5")

	lan_id, _ := strconv.Atoi(dc.Entities.Lans.Items[0].Id)
	server := Server{
		Properties: ServerProperties{
			Name:  "server1",
			Ram:   2048,
			Cores: 1,
		},
		Entities: &ServerEntities{
			Volumes: &Volumes{
				Items: []Volume{
					Volume{
						Properties: VolumeProperties{
							Type:    "HDD",
							Size:    10,
							Name:    "volume1",
							Image:   "1f46a4a3-3f47-11e6-91c6-52540005ab80",
							Bus:     "VIRTIO",
							SshKeys: []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoLVLHON4BSK3D8L4H79aFo+0cj7VM2NiRR/K9wrfkK/XiTc7FlEU4Bs8WLZcsIOxbCGWn2zKZmrLaxYlY+/3aJrxDxXYCy8lRUMnqcQ2JCFY6tpZt/DylPhS9L6qYNpJ0F4FlqRsWxsjpF8TDdJi64k2JFJ8TkvX36P2/kqyFfI+N0/axgjhqV3BgNgApvMt9jxWB5gi8LgDpw9b+bHeMS7TrAVDE7bzT86dmfbTugtiME8cIday8YcRb4xAFgRH8XJVOcE3cs390V/dhgCKy1P5+TjQMjKbFIy2LJoxb7bd38kAl1yafZUIhI7F77i7eoRidKV71BpOZsaPEbWUP jasmin@Jasmins-MBP"},
						},
					},
				},
			},
			Nics: &Nics{
				Items: []Nic{
					Nic{
						Properties: NicProperties{
							Name: "nic",
							Lan:  lan_id,
							Ips:  ipblockresp.Properties.Ips,
						},
					},
				},
			},
		},
	}

	server = CreateServer(dcID, server)
	waitTillProvisioned(server.Headers.Get("Location"))
	//lanrequest := CreateLanRequest{
	//	LanProperties: LanProperties{
	//		Public: true,
	//	},
	//}
	//
	//lan := CreateLan(dcID, lanrequest)
	//
	//obj := PatchNic(dcID, dc.Entities.Servers.Items[0].Id, dc.Entities.Servers.Items[0].Entities.Nics.Items[0].Id, NicProperties{Lan: lan.Id})
	//
	//waitTillProvisioned(obj.Headers.Get("Location"))

	for i := 0; i < 10; i++ {
		request := GetDatacenter(dcID)
		if request.Metadata.State == "AVAILABLE" {
			fmt.Println("DC operational")
			break
		}
		time.Sleep(10 * time.Second)
		i++
	}
}

func TestDeleteDC(t *testing.T) {
	want := 202

	resp := DeleteDatacenter(dcID)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	time.Sleep(60 * time.Second)

	resp = ReleaseIpBlock(ipId)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

}
