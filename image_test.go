// image_test.go

package profitbricks

import (
	"fmt"
	"testing"
)

var imgid string

func TestListImages(t *testing.T) {
	shouldbe := "collection"
	want := 200
	resp := ListImages()
	imgid = resp.Items[0].Id
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestGetImage(t *testing.T) {
	want := 200
	resp := GetImage(imgid)

	if resp.Resp.StatusCode != want {
		if resp.Resp.StatusCode == 403 {
			fmt.Println(bad_status(want, resp.Resp.StatusCode))
			fmt.Println("This error might be due to user's permission level ")
		}
	}
}

func TestPatchImage(t *testing.T) {
	want := 403
	obj := map[string]string{"name": "Renamed img"}
	resp := PatchImage(imgid, obj)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestDeleteImage(t *testing.T) {
	want := 403
	resp := DeleteImage(imgid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
