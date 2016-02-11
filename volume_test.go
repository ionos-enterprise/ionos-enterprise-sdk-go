package profitbricks

import (
	"fmt"
	"testing"
	"time"
)

var volumeId string

func TestCreateVolume(t *testing.T) {
	want := 202

	var request = CreateVolumeRequest{
		VolumeProperties: VolumeProperties{
			Size:   4,
			Name:   "Volume Test",
			Image:  "fbaae2b2-c899-11e5-aa10-52540005ab80",
			Type:   "HDD",
			SshKey: []string{"hQGOEJeFL91EG3+l9TtRbWNjzhDVHeLuL3NWee6bekA="},
		},
	}

	dcID = mkdcid("GO SDK VOLUME DC")
	fmt.Println("*****************")
	fmt.Println(dcID)
	fmt.Println("*****************")
	resp := CreateVolume(dcID, request)

	volumeId = resp.Id

	if resp.Resp.StatusCode != want {
		fmt.Println(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

	time.Sleep(30 * time.Second)
}

func TestListVolumes(t *testing.T) {
	shouldbe := "collection"
	want := 200
	resp := ListVolumes(dcID)

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestGetVolume(t *testing.T) {
	want := 200

	resp := GetVolume(dcID, volumeId)
	fmt.Println(dcID)
	fmt.Println(volumeId)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestPatchVolume(t *testing.T) {
	want := 202
	obj := VolumeProperties{
		Name: "Renamed Volume",
		Size: 2,
	}

	resp := PatchVolume(dcID, volumeId, obj)

	if resp.Resp.StatusCode != want {
		fmt.Println(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestCleanup(t *testing.T) {
	fmt.Println("CLEANING UP AFTER VOLUMES")
	resp := DeleteVolume(dcID, volumeId)
	fmt.Println(resp.StatusCode)
	resp = DeleteDatacenter(dcID)
	fmt.Println(resp.StatusCode)
}
