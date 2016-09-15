package profitbricks

import (
	"fmt"
	"testing"
	"time"
)

var volumeId string

func TestCreateVolume(t *testing.T) {
	setupTestEnv()
	want := 202
	var request = Volume{
		Properties: VolumeProperties{
			Size:             5,
			Name:             "Volume Test",
			Image:            "6aa59ab7-3f45-11e6-91c6-52540005ab80",
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

	time.Sleep(30 * time.Second)
}

func TestListVolumes(t *testing.T) {
	want := 200
	resp := ListVolumes(dcID)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestGetVolume(t *testing.T) {
	want := 200

	resp := GetVolume(dcID, volumeId)
	fmt.Println(dcID)
	fmt.Println(volumeId)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestPatchVolume(t *testing.T) {
	want := 202
	obj := VolumeProperties{
		Name: "Renamed Volume",
		Size: 2,
	}

	resp := PatchVolume(dcID, volumeId, obj)

	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestCreateSnapshot(t *testing.T) {
	want := 202

	resp := CreateSnapshot(dcID, volumeId, "testSnapshot")
	waitTillProvisioned(resp.Headers.Get("Location"))
	if resp.StatusCode != want {
		fmt.Println(string(resp.Response))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	snapshotId = resp.Id

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
	fmt.Println("CLEANING UP AFTER VOLUMES")
	resp := DeleteVolume(dcID, volumeId)
	fmt.Println(resp.StatusCode)
	resp = DeleteDatacenter(dcID)
	fmt.Println(resp.StatusCode)
}
