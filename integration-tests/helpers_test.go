package integration_tests

import (
	"fmt"
	"os"
	"strings"
	"sync"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
)

var (
	syncDC              sync.Once
	syncCDC             sync.Once
	dataCenter          *sdk.Datacenter
	compositeDataCenter *sdk.Datacenter
	server              *sdk.Server
	volume              *sdk.Volume
	lan                 *sdk.Lan
	location            = "us/las"
	image               *sdk.Image
	fw                  *sdk.FirewallRule
	nic                 *sdk.Nic
	sourceMac           = "01:23:45:67:89:00"
	portRangeStart      = 22
	portRangeEnd        = 22
	onceDC              sync.Once
	onceServerDC        sync.Once
	onceServer          sync.Once
	onceFw              sync.Once
	onceServerVolume    sync.Once
	onceCD              sync.Once
	onceLan             sync.Once
	onceLanServer       sync.Once
	onceLanLan          sync.Once
	onceLB              sync.Once
	onceLBDC            sync.Once
	onceLBServer        sync.Once
	onceLBNic           sync.Once
	onceNicNic          sync.Once
	ipBlock             *sdk.IPBlock
	loadBalancer        *sdk.Loadbalancer
	snapshot            *sdk.Snapshot
	snapshotname        = "GO SDK TEST"
	snapshotdescription = "GO SDK test snapshot"
)

func boolAddr(v bool) *bool {
	return &v
}

// Setup creds for single running tests
func setupTestEnv() sdk.Client {
	client := *sdk.NewClient(os.Getenv("PROFITBRICKS_USERNAME"), os.Getenv("PROFITBRICKS_PASSWORD"))
	if val, ok := os.LookupEnv("PROFITBRICKS_API_URL"); ok {
		client.SetCloudApiUrl(val)
	}

	return client
}

func createDataCenter() {
	c := setupTestEnv()

	var obj = sdk.Datacenter{
		Properties: sdk.DatacenterProperties{
			Name:        "GO SDK Test",
			Description: "GO SDK test datacenter",
			Location:    location,
		},
	}
	resp, _ := c.CreateDatacenter(obj)

	dataCenter = resp
}

func createLan() {
	c := setupTestEnv()

	var obj = sdk.Lan{
		Properties: sdk.LanProperties{
			Name:   "GO SDK Test",
			Public: true,
		},
	}
	resp, _ := c.CreateLan(dataCenter.ID, obj)

	c.WaitTillProvisioned(resp.Headers.Get("Location"))
	lan = resp
}

func createCompositeDataCenter() {
	c := setupTestEnv()
	var obj = sdk.Datacenter{
		Properties: sdk.DatacenterProperties{
			Name:        "GO SDK Test Composite",
			Description: "GO SDK test composite datacenter",
			Location:    location,
		},
		Entities: sdk.DatacenterEntities{
			Servers: &sdk.Servers{
				Items: []sdk.Server{
					{
						Properties: sdk.ServerProperties{
							Name:  "GO SDK Test",
							RAM:   1024,
							Cores: 1,
						},
					},
				},
			},
			Volumes: &sdk.Volumes{
				Items: []sdk.Volume{
					{
						Properties: sdk.VolumeProperties{
							Type:             "HDD",
							Size:             2,
							Name:             "GO SDK Test",
							Bus:              "VIRTIO",
							LicenceType:      "UNKNOWN",
							AvailabilityZone: "ZONE_3",
						},
					},
				},
			},
		},
	}
	resp, err := c.CreateDatacenter(obj)
	if err != nil {
		fmt.Println("error while creating", err)
	}
	compositeDataCenter = resp

	err = c.WaitTillProvisioned(compositeDataCenter.Headers.Get("Location"))
	if err != nil {
		fmt.Println("error while waiting", err)
	}
}

func createCompositeServerFW() {
	c := setupTestEnv()
	var req = sdk.Server{
		Properties: sdk.ServerProperties{
			Name:             "GO SDK Test",
			RAM:              1024,
			Cores:            1,
			AvailabilityZone: "ZONE_1",
			CPUFamily:        "INTEL_XEON",
		},
		Entities: &sdk.ServerEntities{
			Volumes: &sdk.Volumes{
				Items: []sdk.Volume{
					{
						Properties: sdk.VolumeProperties{
							Type:          "HDD",
							Size:          5,
							Name:          "volume1",
							ImageAlias:    "ubuntu:latest",
							ImagePassword: "JWXuXR9CMghXAc6v",
						},
					},
				},
			},
			Nics: &sdk.Nics{
				Items: []sdk.Nic{
					{
						Properties: &sdk.NicProperties{
							Name: "nic",
							Lan:  1,
						},
						Entities: &sdk.NicEntities{
							FirewallRules: &sdk.FirewallRules{
								Items: []sdk.FirewallRule{
									{
										Properties: sdk.FirewallruleProperties{
											Name:           "SSH",
											Protocol:       "TCP",
											SourceMac:      &sourceMac,
											PortRangeStart: &portRangeStart,
											PortRangeEnd:   &portRangeEnd,
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

	srv, err := c.CreateServer(dataCenter.ID, req)

	server = srv
	nic = &srv.Entities.Nics.Items[0]
	fw = &nic.Entities.FirewallRules.Items[0]

	if err != nil {
		fmt.Println("error while creating a server", err)
	}
	err = c.WaitTillProvisioned(srv.Headers.Get("Location"))

	if err != nil {
		fmt.Println("error while waiting on", err)
	}
}

func createNic() {
	c := setupTestEnv()
	obj := sdk.Nic{
		Properties: &sdk.NicProperties{
			Name: "GO SDK Test",
			Lan:  1,
		},
	}

	resp, _ := c.CreateNic(dataCenter.ID, server.ID, obj)
	c.WaitTillProvisioned(resp.Headers.Get("Location"))

	nic = resp
}

func createLoadBalancerWithIP() {
	c := setupTestEnv()
	var obj = sdk.IPBlock{
		Properties: sdk.IPBlockProperties{
			Name:     "GO SDK Test",
			Size:     1,
			Location: "us/las",
		},
	}
	resp, err := c.ReserveIPBlock(obj)
	if err != nil {
		fmt.Println("Error while reserving an IP block", err)
	}

	c.WaitTillProvisioned(resp.Headers.Get("Location"))
	ipBlock = resp
	var request = sdk.Loadbalancer{
		Properties: sdk.LoadbalancerProperties{
			Name: "GO SDK Test",
			IP:   resp.Properties.IPs[0],
			Dhcp: true,
		},
		Entities: sdk.LoadbalancerEntities{
			Balancednics: &sdk.BalancedNics{
				Items: []sdk.Nic{
					{
						ID: nic.ID,
					},
				},
			},
		},
	}

	resp1, _ := c.CreateLoadbalancer(dataCenter.ID, request)
	c.WaitTillProvisioned(resp1.Headers.Get("Location"))
	loadBalancer = resp1
	nic = &loadBalancer.Entities.Balancednics.Items[0]
}

func createVolume() {
	c := setupTestEnv()
	var request = sdk.Volume{
		Properties: sdk.VolumeProperties{
			Size:        2,
			Name:        "GO SDK Test",
			LicenceType: "OTHER",
			Type:        "HDD",
		},
	}

	resp, _ := c.CreateVolume(dataCenter.ID, request)

	volume = resp
	c.WaitTillProvisioned(resp.Headers.Get("Location"))
}

func createSnapshot() {
	c := setupTestEnv()
	resp, _ := c.CreateSnapshot(dataCenter.ID, volume.ID, snapshotname, snapshotdescription)
	snapshot = resp
	c.WaitTillProvisioned(snapshot.Headers.Get("Location"))
}

func mknicCustom(client sdk.Client, dcid, serverid string, lanid int, ips []string) string {
	var request = sdk.Nic{
		Properties: &sdk.NicProperties{
			Lan:            lanid,
			Name:           "GO SDK Test",
			Nat:            boolAddr(false),
			FirewallActive: boolAddr(true),
			Ips:            ips,
		},
	}

	resp, err := client.CreateNic(dcid, serverid, request)
	if err != nil {
		return ""
	}
	err = client.WaitTillProvisioned(resp.Headers.Get("Location"))
	if err != nil {
		return ""
	}
	return resp.ID
}

func createServer() {
	server = setupCreateServer(dataCenter.ID)
	if server == nil {
		panic("Server not created")
	}
}

func setupCreateServer(srvDc string) *sdk.Server {
	c := setupTestEnv()

	var req = sdk.Server{
		Properties: sdk.ServerProperties{
			Name:             "GO SDK Test",
			RAM:              1024,
			Cores:            1,
			AvailabilityZone: "ZONE_1",
			CPUFamily:        "INTEL_XEON",
		},
	}
	srv, err := c.CreateServer(srvDc, req)
	if err != nil {
		return nil
	}

	err = c.WaitTillProvisioned(srv.Headers.Get("Location"))
	if err != nil {
		return nil
	}
	return srv
}

func setupVolume() {
	c := setupTestEnv()

	vol := sdk.Volume{
		Properties: sdk.VolumeProperties{
			Type:        "HDD",
			Size:        2,
			Name:        "GO SDK Test",
			Bus:         "VIRTIO",
			LicenceType: "UNKNOWN",
		},
	}
	resp, err := c.CreateVolume(dataCenter.ID, vol)
	if err != nil {
		fmt.Println("create volume failed")
	}
	volume = resp

	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))
	if err != nil {
		fmt.Println("failed while waiting on volume to finish")
	}

}

func setupVolumeAttached() {
	c := setupTestEnv()

	vol := sdk.Volume{
		Properties: sdk.VolumeProperties{
			Type:        "HDD",
			Size:        2,
			Name:        "GO SDK Test",
			Bus:         "VIRTIO",
			LicenceType: "UNKNOWN",
		},
	}
	resp, err := c.CreateVolume(dataCenter.ID, vol)
	if err != nil {
		fmt.Println("create volume failed")
	}

	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))
	if err != nil {
		fmt.Println("failed while waiting on volume to finish")
	}
	volume = resp

	volume, err = c.AttachVolume(dataCenter.ID, server.ID, volume.ID)
	if err != nil {
		fmt.Println("attach volume failed", err)
	}

	err = c.WaitTillProvisioned(volume.Headers.Get("Location"))

	if err != nil {
		fmt.Println("failed while waiting on volume to finish")
	}
}

func setupCDAttached() {
	c := setupTestEnv()

	var imageID string
	images, err := c.ListImages()
	for _, img := range images.Items {
		if img.Properties.ImageType == "CDROM" && img.Properties.Location == "us/las" && img.Properties.Public == true {
			imageID = img.ID
			break
		}
	}

	resp, err := c.AttachCdrom(dataCenter.ID, server.ID, imageID)
	if err != nil {
		fmt.Println("attach CD failed", err)
	}

	image = resp

	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))

	if err != nil {
		fmt.Println("failed while waiting on volume to finish")
	}
}

func reserveIP() {
	c := setupTestEnv()
	var obj = sdk.IPBlock{
		Properties: sdk.IPBlockProperties{
			Name:     "GO SDK Test",
			Size:     1,
			Location: location,
		},
	}
	resp, _ := c.ReserveIPBlock(obj)
	ipBlock = resp
}

func getImageID(location string, imageName string, imageType string) string {
	if imageName == "" {
		return ""
	}

	c := setupTestEnv()

	images, err := c.ListImages()
	if err != nil {
		return ""
	}

	if len(images.Items) > 0 {
		for _, i := range images.Items {
			imgName := ""
			if i.Properties.Name != "" {
				imgName = i.Properties.Name
			}

			if imageType == "SSD" {
				imageType = "HDD"
			}
			if imgName != "" && strings.Contains(strings.ToLower(imgName), strings.ToLower(imageName)) && i.Properties.ImageType == imageType && i.Properties.Location == location && i.Properties.Public == true {
				return i.ID
			}
		}
	}
	return ""
}
