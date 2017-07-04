package profitbricks

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

var volumeId string

func TestCreateVolume(t *testing.T) {
	setupTestEnv()
	want := 202
	var request = Volume{
		Properties: VolumeProperties{
			Size:             2,
			Name:             "GO SDK Test",
			ImageAlias:       "ubuntu:latest",
			Bus:              "VIRTIO",
			SshKeys:          []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoLVLHON4BSK3D8L4H79aFo..."},
			Type:             "HDD",
			ImagePassword:    "test1234",
			AvailabilityZone: "ZONE_3",
		},
	}

	dcID = mkdcid("GO SDK VOLUME DC")
	resp := CreateVolume(dcID, request)

	waitTillProvisioned(resp.Headers.Get("Location"))
	volumeId = resp.Id
	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Type_, "volume")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Size, 2)
	assert.Equal(t, resp.Properties.Bus, "VIRTIO")
	assert.Equal(t, resp.Properties.AvailabilityZone, "ZONE_3")
	assert.Equal(t, resp.Properties.Type, "HDD")
	assert.Equal(t, resp.Properties.SshKeys, []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoLVLHON4BSK3D8L4H79aFo..."})
}

func TestCreateVolumeFail(t *testing.T) {
	want := 422
	var request = Volume{
		Properties: VolumeProperties{
			Name:             "Volume Test",
			Image:            "rewar",
			Type:             "HDD",
			ImagePassword:    "test1234",
			AvailabilityZone: "ZONE_3",
		},
	}

	resp := CreateVolume(dcID, request)

	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, strings.Contains(resp.Response, "Attribute 'size' is required"))
}

func TestListVolumes(t *testing.T) {
	want := 200
	resp := ListVolumes(dcID)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.True(t, len(resp.Items) > 0)
}

func TestGetVolume(t *testing.T) {
	want := 200

	time.Sleep(5000)
	resp := GetVolume(dcID, volumeId)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, volumeId)
	assert.Equal(t, resp.Type_, "volume")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Size, 2)
	//assert.Equal(t, resp.Properties.Bus, "VIRTIO")
	assert.Equal(t, resp.Properties.AvailabilityZone, "ZONE_3")
	assert.Equal(t, resp.Properties.Type, "HDD")
}

func TestGetVolumeFailure(t *testing.T) {
	want := 404

	resp := GetVolume(dcID, "00000000-0000-0000-0000-000000000000")
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, strings.Contains(resp.Response, "Resource does not exist"))
}

func TestPatchVolume(t *testing.T) {
	want := 202
	obj := VolumeProperties{
		Name: "GO SDK Test - RENAME",
		Size: 5,
	}

	resp := PatchVolume(dcID, volumeId, obj)

	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	waitTillProvisioned(resp.Headers.Get("Location"))
	assert.Equal(t, resp.Id, volumeId)
	assert.Equal(t, resp.Properties.Name, "GO SDK Test - RENAME")
	assert.Equal(t, resp.Properties.Size, 5)
}

func TestCreateSnapshot(t *testing.T) {
	want := 202

	resp := CreateSnapshot(dcID, volumeId, snapshotname, snapshotdescription)
	waitTillProvisioned(resp.Headers.Get("Location"))
	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	time.Sleep(30 * time.Second)
	snapshotId = resp.Id

	assert.Equal(t, resp.Properties.Name, snapshotname)
	assert.Equal(t, resp.Type_, "snapshot")
}

func TestRestoreSnapshot(t *testing.T) {
	want := 202

	resp := RestoreSnapshot(dcID, volumeId, snapshotId)

	waitTillProvisioned(resp.Headers.Get("Location"))
	if resp.StatusCode != want {
		fmt.Println(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestCleanup(t *testing.T) {
	DeleteSnapshot(snapshotId)
	DeleteVolume(dcID, volumeId)
	DeleteDatacenter(dcID)
}
