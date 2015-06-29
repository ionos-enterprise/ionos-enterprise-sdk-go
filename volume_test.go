package profitbricks

import (
	"fmt"
	"testing"
	"time"
)

const vol_sleep = 30

var (
	volumeId string
	vol_srv_id string
)


func TestCreateVolume(t *testing.T) {
	want := 202
	var jason = []byte(`{
    "properties": {
         "size": "2",
        "name": "volume-name",
		"licenceType" : "LINUX"
    }
	}`)
	
	// reuse server DC
	dcID = srv_dc_id
	resp := CreateVolume(dcID, jason)

	volumeId = resp.Id

	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

	time.Sleep(vol_sleep * time.Second)
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

func TestAttachVolume(t *testing.T) {
	want := 202
	
	vol_srv_id = setupCreateServer(dcID)
	
	resp := AttachVolume(dcID, vol_srv_id, volumeId)
	if resp.Resp.StatusCode != want {
		t.Error(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestGetAttachedVolume(t *testing.T) {
	want := 200
	shouldbe := "volume"
	
	// wait for volume to attach
	time.Sleep(time.Second * vol_sleep)
	
	resp := GetAttachedVolume(dcID, vol_srv_id, volumeId)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Error(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestDetachVolume(t *testing.T) {
	want := 202
	
	resp := DetachVolume(dcID, vol_srv_id, volumeId)
	
	if resp.StatusCode != want {
		t.Error(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}	
}

func TestPatchVolume(t *testing.T) {
	want := 202
	
	resp := PatchVolume(dcID, volumeId, []byte(`{"name": "volume-name1234"}`))
	if resp.Resp.StatusCode != want {
		t.Error(string(resp.Resp.Body))
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
