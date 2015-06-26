package profitbricks

import (
	"fmt"
	"testing"
	"time"
)

var volumeId string

func TestCreateVolume(t *testing.T) {
	want := 202
	var jason = []byte(`{
    "properties": {
         "size": "2",
        "name": "volume-name",
		"licenceType" : "LINUX"
    }
	}`)
	dcID = mkdcid()
	resp := CreateVolume(dcID, jason)

	volumeId = resp.Id

	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

	time.Sleep(2 * time.Second)
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

	resp := PatchVolume(dcID, volumeId, []byte(`{"name": "volume-name1234"}`))

	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func Cleanup(t *testing.T) {
	DeleteVolume(dcID, volumeId)
	DeleteDatacenter(dcID)
}
