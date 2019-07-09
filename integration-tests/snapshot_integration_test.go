package integration_tests

import (
	"fmt"
	"sync"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

var (
	onceSnapshotVolume sync.Once
	onceSnapshotDC     sync.Once
	onceSnapshot       sync.Once
)

func TestCreateSnapshots(t *testing.T) {
	fmt.Println("Snapshot test")
	onceSnapshotDC.Do(createDataCenter)
	onceSnapshotVolume.Do(createVolume)
	onceSnapshot.Do(createSnapshot)

	assert.Equal(t, snapshot.PBType, "snapshot")
	assert.Equal(t, snapshot.Properties.Name, snapshotname)
	assert.Equal(t, snapshot.Properties.Description, snapshotdescription)
}

func TestCreateSnapshotFailure(t *testing.T) {
	c := setupTestEnv()
	_, err := c.CreateSnapshot("00000000-0000-0000-0000-000000000000", "volumeId", "fail", snapshotdescription)
	assert.NotNil(t, err)
}

func TestGetSnapshot(t *testing.T) {
	c := setupTestEnv()
	onceSnapshotDC.Do(createDataCenter)
	onceSnapshotVolume.Do(createVolume)
	onceSnapshot.Do(createSnapshot)

	resp, err := c.GetSnapshot(snapshot.ID)

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.ID, snapshot.ID)
	assert.Equal(t, resp.Properties.Size, volume.Properties.Size)
	assert.Equal(t, resp.Properties.CPUHotPlug, volume.Properties.CPUHotPlug)
	assert.Equal(t, resp.Properties.CPUHotUnplug, volume.Properties.CPUHotUnplug)
	assert.Equal(t, resp.Properties.RAMHotPlug, volume.Properties.RAMHotPlug)
	assert.Equal(t, resp.Properties.RAMHotUnplug, volume.Properties.RAMHotUnplug)
	assert.Equal(t, resp.Properties.NicHotPlug, volume.Properties.NicHotPlug)
	assert.Equal(t, resp.Properties.NicHotUnplug, volume.Properties.NicHotUnplug)
	assert.Equal(t, resp.Properties.DiscScsiHotPlug, volume.Properties.DiscScsiHotPlug)
	assert.Equal(t, resp.Properties.DiscScsiHotUnplug, volume.Properties.DiscScsiHotUnplug)
	assert.Equal(t, resp.Properties.DiscVirtioHotPlug, volume.Properties.DiscVirtioHotPlug)
	assert.Equal(t, resp.Properties.DiscVirtioHotUnplug, volume.Properties.DiscVirtioHotUnplug)
	assert.Equal(t, resp.Properties.LicenceType, volume.Properties.LicenceType)
}

func TestGetSnapshotFailure(t *testing.T) {
	c := setupTestEnv()

	_, err := c.GetSnapshot("00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, err)
}

func TestListSnapshot(t *testing.T) {
	c := setupTestEnv()

	resp, err := c.ListSnapshots()
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestUpdateSnapshot(t *testing.T) {
	c := setupTestEnv()
	onceSnapshotDC.Do(createDataCenter)
	onceSnapshotVolume.Do(createVolume)
	onceSnapshot.Do(createSnapshot)

	newValue := "GO SDK Test - RENAME"
	resp, err := c.UpdateSnapshot(snapshot.ID, sdk.SnapshotProperties{Name: newValue})
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, newValue, resp.Properties.Name)
}

func TestDeleteSnapshot(t *testing.T) {
	c := setupTestEnv()
	onceSnapshotDC.Do(createDataCenter)
	onceSnapshotVolume.Do(createVolume)
	onceSnapshot.Do(createSnapshot)

	resp, err := c.DeleteSnapshot(snapshot.ID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		t.Error(err)
	}

	_, err = c.DeleteDatacenter(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}
}
