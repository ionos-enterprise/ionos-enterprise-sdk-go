package profitbricks

import (
	"fmt"
	"os"
	"testing"
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
							Volumes: &Volumes{
								Items: []Volume{
									Volume{
										Properties: VolumeProperties{
											Type:          "HDD",
											Size:          10,
											Name:          "volume1",
											Image:         "1f46a4a3-3f47-11e6-91c6-52540005ab80",
											Bus:           "VIRTIO",
											ImagePassword: "test1234",
										},
									},
								},
							},
							Nics: &Nics{
								Items: []Nic{
									Nic{
										Properties: NicProperties{
											Name: "nic",
											Lan:  "1",
											Ips:  ipblockresp.Properties.Ips,
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

	fmt.Println(ipblockresp.Properties.Ips)
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

	obj := PatchNic(dcID, dc.Entities.Servers.Items[0].Id, dc.Entities.Servers.Items[0].Entities.Nics.Items[0].Id, NicProperties{Lan: lan.Id})

	waitTillProvisioned(obj.Headers.Get("Location"))

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
