package integration_tests

import (
	"fmt"
	"sync"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

var (
	onceVolume         sync.Once
	onceVolumeDC       sync.Once
	onceVolumeSnapshot sync.Once
)

func TestCreateVolume(t *testing.T) {
	fmt.Println("Volume tests")
	onceVolumeDC.Do(createDataCenter)
	onceVolume.Do(createVolume)
	assert.Equal(t, volume.PBType, "volume")
	assert.Equal(t, volume.Properties.Name, "GO SDK Test")
	assert.Equal(t, volume.Properties.Size, 2)
	assert.Equal(t, volume.Properties.Type, "HDD")
	assert.Equal(t, volume.Properties.LicenceType, "OTHER")
}

func TestCreateVolumeFail(t *testing.T) {
	c := setupTestEnv()
	onceVolumeDC.Do(createDataCenter)
	var request = sdk.Volume{
		Properties: sdk.VolumeProperties{
			Name:             "Volume Test",
			Image:            "rewar",
			Type:             "HDD",
			ImagePassword:    "test1234",
			AvailabilityZone: "ZONE_3",
		},
	}

	_, err := c.CreateVolume(dataCenter.ID, request)

	assert.NotNil(t, err)
}

func TestListVolumes(t *testing.T) {
	c := setupTestEnv()
	onceVolumeDC.Do(createDataCenter)
	onceVolume.Do(createVolume)
	resp, err := c.ListVolumes(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetVolume(t *testing.T) {
	c := setupTestEnv()
	onceVolumeDC.Do(createDataCenter)
	onceVolume.Do(createVolume)

	resp, err := c.GetVolume(dataCenter.ID, volume.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.PBType, "volume")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Size, 2)
	assert.Equal(t, resp.Properties.Type, "HDD")
	assert.Equal(t, resp.Properties.LicenceType, "OTHER")
}

func TestGetVolumeFailure(t *testing.T) {
	c := setupTestEnv()
	onceVolumeDC.Do(createDataCenter)
	onceVolume.Do(createVolume)
	_, err := c.GetVolume(dataCenter.ID, "00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, err)
}

func TestUpdateVolume(t *testing.T) {
	c := setupTestEnv()
	onceVolumeDC.Do(createDataCenter)
	onceVolume.Do(createVolume)
	newName := "GO SDK Test - RENAME"
	obj := sdk.VolumeProperties{
		Name: newName,
		Size: 5,
	}

	resp, err := c.UpdateVolume(dataCenter.ID, volume.ID, obj)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.ID, volume.ID)
	assert.Equal(t, resp.Properties.Name, "GO SDK Test - RENAME")
	assert.Equal(t, resp.Properties.Size, 5)
}

func TestCreateSnapshot(t *testing.T) {
	onceVolumeDC.Do(createDataCenter)
	onceVolume.Do(createVolume)
	onceVolumeSnapshot.Do(createSnapshot)

	assert.Equal(t, snapshot.Properties.Name, snapshotname)
	assert.Equal(t, snapshot.PBType, "snapshot")
}

func TestRestoreSnapshot(t *testing.T) {
	c := setupTestEnv()
	onceVolumeDC.Do(createDataCenter)
	onceVolume.Do(createVolume)
	onceVolumeSnapshot.Do(createSnapshot)

	resp, err := c.RestoreSnapshot(dataCenter.ID, volume.ID, snapshot.ID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}
}

func TestCleanup(t *testing.T) {
	c := setupTestEnv()
	c.DeleteSnapshot(snapshot.ID)
	c.DeleteVolume(dataCenter.ID, volume.ID)
	c.DeleteDatacenter(dataCenter.ID)
}
