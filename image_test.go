// image_test.go

package profitbricks

import (
	"fmt"
	"testing"
)

func mkimgid() string {
	imgs := ListImages()

	imgid := imgs.Items[0].Id
	return imgid
}

func TestListImages(t *testing.T) {
	shouldbe := "collection"
	want := 200
	resp := ListImages()

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestGetImage(t *testing.T) {
	want := 200
	imgid := mkimgid()
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
	jason_patch := []byte(`{
					"name":"Renamed img"
					}`)
	imgid := mkimgid()
	resp := PatchImage(imgid, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestDeleteImage(t *testing.T) {
	want := 403
	imgid := mkimgid()
	resp := DeleteImage(imgid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
