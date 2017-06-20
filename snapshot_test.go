package profitbricks

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

var snapshotId string
var volume Volume

var snapshotname string="GO SDK TEST"
var snapshotdescription string="GO SDK test snapshot"

func createVolume() {
	setupTestEnv()
	want := 202
	var request = Volume{
		Properties: VolumeProperties{
			Size:          5,
			Name:          "Volume Test",
			Image:         image,
			Type:          "HDD",
			ImagePassword: "test1234",

		},
	}

	dcID = mkdcid("GO SDK snapshot DC")
	resp := CreateVolume(dcID, request)

	volume = resp
	waitTillProvisioned(resp.Headers.Get("Location"))
	volumeId = resp.Id

	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
	}
}

func TestCreateSnapshots(t *testing.T) {
	createVolume()
	resp := CreateSnapshot(dcID, volumeId, snapshotname, snapshotdescription)
	waitTillProvisioned(resp.Headers.Get("Location"))
	snapshotId = resp.Id

	assert.Equal(t, resp.Type_, "snapshot")
	assert.Equal(t, resp.Properties.Name, snapshotname)
	assert.Equal(t, resp.Properties.Description, snapshotdescription)
}

func TestCreateSnapshotFailure(t *testing.T) {
	resp := CreateSnapshot(dcID, volumeId, "*&*&^*&^*&^*&^*", snapshotdescription)
	assert.True(t, strings.Contains(resp.Response, "Attribute 'name' is required"))
}

func TestGetSnapshot(t *testing.T) {
	want := 200

	resp := GetSnapshot(snapshotId)
	volume = GetVolume(dcID, volumeId)

	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.Equal(t, resp.Id, snapshotId)
	assert.Equal(t, resp.Properties.Size, volume.Properties.Size)
	assert.Equal(t, resp.Properties.CpuHotPlug, volume.Properties.CpuHotPlug)
	assert.Equal(t, resp.Properties.CpuHotUnplug, volume.Properties.CpuHotUnplug)
	assert.Equal(t, resp.Properties.RamHotPlug, volume.Properties.RamHotPlug)
	assert.Equal(t, resp.Properties.RamHotUnplug, volume.Properties.RamHotUnplug)
	assert.Equal(t, resp.Properties.NicHotPlug, volume.Properties.NicHotPlug)
	assert.Equal(t, resp.Properties.NicHotUnplug, volume.Properties.NicHotUnplug)
	assert.Equal(t, resp.Properties.DiscScsiHotPlug, volume.Properties.DiscScsiHotPlug)
	assert.Equal(t, resp.Properties.DiscScsiHotUnplug, volume.Properties.DiscScsiHotUnplug)
	assert.Equal(t, resp.Properties.DiscVirtioHotPlug, volume.Properties.DiscVirtioHotPlug)
	assert.Equal(t, resp.Properties.DiscVirtioHotUnplug, volume.Properties.DiscVirtioHotUnplug)
	assert.Equal(t, resp.Properties.LicenceType, volume.Properties.LicenceType)
}

func TestGetSnapshotFailure(t *testing.T) {
	want := 404

	resp := GetSnapshot("00000000-0000-0000-0000-000000000000")

	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.True(t, strings.Contains(resp.Response, "Resource does not exist"))
}

func TestListSnapshot(t *testing.T) {
	want := 200

	resp := ListSnapshots()

	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestUpdateSnapshot(t *testing.T) {
	want := 202
	newValue := "GO SDK Test - RENAME"
	resp := UpdateSnapshot(snapshotId, SnapshotProperties{Name: newValue})

	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	if newValue != resp.Properties.Name {
		t.Errorf("Snapshot wasn't updated.")
	}
}

func TestDeleteSnapshot(t *testing.T) {
	want := 202

	resp := DeleteSnapshot(snapshotId)

	if resp.StatusCode != want {
		fmt.Println(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	resp = DeleteDatacenter(dcID)

	if resp.StatusCode != want {
		fmt.Println(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
