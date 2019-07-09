package integration_tests

import (
	"fmt"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateServer(t *testing.T) {
	fmt.Println("Server tests")
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)

	assert.Equal(t, server.PBType, "server")
	assert.Equal(t, server.Properties.Name, "GO SDK Test")
	assert.Equal(t, server.Properties.RAM, 1024)
	assert.Equal(t, server.Properties.Cores, 1)
	assert.Equal(t, server.Properties.AvailabilityZone, "ZONE_1")
	assert.Equal(t, server.Properties.CPUFamily, "INTEL_XEON")
}

func TestCreateServerFailure(t *testing.T) {
	c := setupTestEnv()

	var req = sdk.Server{
		Properties: sdk.ServerProperties{
			Name:             "GO SDK Test",
			RAM:              1024,
			AvailabilityZone: "ZONE_1",
			CPUFamily:        "INTEL_XEON",
		},
	}
	_, err := c.CreateServer(dataCenter.ID, req)
	if err == nil {
		t.Errorf("no error has been returned")
	}
}

func TestGetServer(t *testing.T) {
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	c := setupTestEnv()
	resp, err := c.GetServer(dataCenter.ID, server.ID)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.ID, server.ID)
	assert.Equal(t, resp.PBType, "server")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.RAM, 1024)
	assert.Equal(t, resp.Properties.Cores, 1)
	assert.Equal(t, resp.Properties.AvailabilityZone, "ZONE_1")
	assert.Equal(t, resp.Properties.CPUFamily, "INTEL_XEON")
}

func TestGetServerFailure(t *testing.T) {
	c := setupTestEnv()
	_, err := c.GetServer(dataCenter.ID, "00000000-0000-0000-0000-000000000000")

	if err == nil {
		t.Errorf("no error has been returned")
	}
}

func TestListServers(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)

	resp, err := c.ListServers(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, len(resp.Items) > 0)
}

func TestUpdateServer(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)

	req := sdk.ServerProperties{
		Name: "GO SDK Test RENAME",
	}
	resp, err := c.UpdateServer(dataCenter.ID, server.ID, req)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, server.ID)
	assert.Equal(t, resp.Properties.Name, "GO SDK Test RENAME")
}

func TestStopServer(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	resp, err := c.StopServer(dataCenter.ID, server.ID)
	if err != nil {
		t.Error(err)
	}
	c.WaitTillProvisioned(resp.Get("Location"))
}

func TestStartServer(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)

	resp, err := c.StartServer(dataCenter.ID, server.ID)
	if err != nil {
		t.Error(err)
	}
	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}
}

func TestRebootServer(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)

	resp, err := c.RebootServer(dataCenter.ID, server.ID)
	if err != nil {
		t.Error(err)
	}
	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}
}

func TestAttachImage(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	onceServerVolume.Do(setupVolume)

	resp, err := c.AttachVolume(dataCenter.ID, server.ID, volume.ID)

	if err != nil {
		t.Error(err)
	}
	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, volume.ID)
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Size, 2)
	assert.Equal(t, resp.Properties.Type, "HDD")
	assert.Equal(t, resp.Properties.LicenceType, "UNKNOWN")
}

func TestListAttachedVolumes(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	onceServerVolume.Do(setupVolumeAttached)

	resp, err := c.ListAttachedVolumes(dataCenter.ID, server.ID)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetAttachedVolume(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	onceServerVolume.Do(setupVolumeAttached)

	resp, err := c.GetAttachedVolume(dataCenter.ID, server.ID, volume.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, volume.ID)
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Size, 2)
	assert.Equal(t, resp.Properties.Bus, "VIRTIO")
	assert.Equal(t, resp.Properties.Type, "HDD")
	assert.Equal(t, resp.Properties.LicenceType, "UNKNOWN")
}

func TestDetachVolume(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	onceServerVolume.Do(setupVolumeAttached)

	resp, err := c.DetachVolume(dataCenter.ID, server.ID, volume.ID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}
}

func TestAttachCdrom(t *testing.T) {
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	onceServerVolume.Do(setupCDAttached)
}

func TestListAttachedCdroms(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	onceCD.Do(setupCDAttached)

	_, err := c.ListAttachedCdroms(dataCenter.ID, server.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAttachedCdrom(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	onceCD.Do(setupCDAttached)

	resp, err := c.GetAttachedCdrom(dataCenter.ID, server.ID, image.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, image.ID)
}

func TestDetachCdrom(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)
	onceCD.Do(setupCDAttached)

	_, err := c.DetachCdrom(dataCenter.ID, server.ID, image.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateCompositeServer(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
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

	if err != nil {
		t.Error(err)
	}
	err = c.WaitTillProvisioned(srv.Headers.Get("Location"))

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, srv.PBType, "server")
	assert.Equal(t, srv.Properties.Name, "GO SDK Test")
	assert.Equal(t, srv.Properties.RAM, 1024)
	assert.Equal(t, srv.Properties.Cores, 1)
	assert.True(t, len(srv.Entities.Nics.Items) > 0)
	assert.True(t, len(srv.Entities.Nics.Items[0].Entities.FirewallRules.Items) > 0)
	assert.True(t, len(srv.Entities.Volumes.Items) > 0)
}

func TestDeleteServer(t *testing.T) {
	c := setupTestEnv()
	onceServerDC.Do(createDataCenter)
	onceServer.Do(createServer)

	_, err := c.DeleteServer(dataCenter.ID, server.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = c.DeleteDatacenter(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}
}
