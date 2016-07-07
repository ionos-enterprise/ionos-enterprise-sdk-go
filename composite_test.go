package profitbricks

import (
	"testing"
	"os"
	"fmt"
	"github.com/profitbricks/profitbricks-sdk-go/model"
	"time"
)

var ipId string

func TestCompositeCreate(t *testing.T) {
	SetAuth(os.Getenv("PROFITBRICKS_USERNAME"), os.Getenv("PROFITBRICKS_PASSWORD"))
	location := "us/las"

	ipblockreq := IPBlockReserveRequest{
		IPBlockProperties: IPBlockProperties{
			Size:     1,
			Location: location,
		},
	}

	ipblockresp := ReserveIpBlock(ipblockreq)
	ipId = ipblockresp.Id
	fmt.Println(ipId)
	datacenter := model.Datacenter{
		Properties: model.DatacenterProperties{
			Name: "composite test",
			Location:location,
		},
		Entities:model.DatacenterEntities{
			Servers: &model.Servers{
				Items:[]model.Server{
					model.Server{
						Properties: model.ServerProperties{
							Name : "server1",
							Ram: 2048,
							Cores: 1,
						},
						Entities:model.ServerEntities{
							Volumes: &model.AttachedVolumes{
								Items:[]model.Volume{
									model.Volume{
										Properties: model.VolumeProperties{
											Type_:"HDD",
											Size:10,
											Name:"volume1",
											Image:"1f46a4a3-3f47-11e6-91c6-52540005ab80",
											Bus:"VIRTIO",
											ImagePassword:"test1234",
											SshKeys: []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoLVLHON4BSK3D8L4H79aFo+0cj7VM2NiRR/K9wrfkK/XiTc7FlEU4Bs8WLZcsIOxbCGWn2zKZmrLaxYlY+/3aJrxDxXYCy8lRUMnqcQ2JCFY6tpZt/DylPhS9L6qYNpJ0F4FlqRsWxsjpF8TDdJi64k2JFJ8TkvX36P2/kqyFfI+N0/axgjhqV3BgNgApvMt9jxWB5gi8LgDpw9b+bHeMS7TrAVDE7bzT86dmfbTugtiME8cIday8YcRb4xAFgRH8XJVOcE3cs390V/dhgCKy1P5+TjQMjKbFIy2LJoxb7bd38kAl1yafZUIhI7F77i7eoRidKV71BpOZsaPEbWUP jasmin@Jasmins-MBP"},
										},
									},
								},
							},
							Nics: &model.Nics{
								Items: []model.Nic{
									model.Nic{
										Properties: model.NicProperties{
											Name : "nic",
											Lan : "1",
											Ips: []string{ipblockresp.Properties["ips"].([]interface{})[0].(string)},
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

	fmt.Println(ipblockresp.Properties["ips"].([]interface{})[0].(string))
	dc := CompositeCreateDatacenter(datacenter)
	dcID = dc.Id

	waitTillProvisioned(dc.Headers.Get("Location"))
	SetDepth("5")

	lanrequest := CreateLanRequest{
		LanProperties: LanProperties{
			Public: true,
		},
	}

	lan := CreateLan(dcID, lanrequest)

	obj := PatchNic(dcID, dc.Entities.Servers.Items[0].Id, dc.Entities.Servers.Items[0].Entities.Nics.Items[0].Id, map[string]string{"lan": lan.Id})

	waitTillProvisioned(obj.Resp.Headers.Get("Location"))

	for i := 0; i < 10; i++ {
		request := GetDatacenter(dcID)
		if request.MetaData["state"] == "AVAILABLE" {
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

func waitTillProvisioned(path string) {
	//d.setPB()
	for i := 0; i < 5; i++ {
		request := GetRequestStatus(path)
		if request.MetaData["status"] == "DONE" {
			break
		}
		time.Sleep(10 * time.Second)
		i++
	}
}