// image_test.go

package profitbricks

import "testing"

func mkimgid() string {
	imgs := ListImages()

	imgid := imgs.Items[0].Id
	return imgid
}

func TestListImages(t *testing.T) {
	//t.Parallel()
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
	//t.Parallel()
	shouldbe := "image"
	want := 200
	imgid := mkimgid()
	resp := GetImage(imgid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestPatchImage(t *testing.T) {
	//t.Parallel()
	want := 202
	jason_patch := []byte(`{
					"name":"Renamed img",
					}`)
	imgid := mkimgid()
	resp := PatchImage(imgid, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestUpdateImage(t *testing.T) {
	//t.Parallel()
	want := 202
	jason_update := []byte(`{
					"size":80,
					
					}`)

	imgid := mkimgid()
	resp := UpdateImage(imgid, jason_update)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestDeleteImage(t *testing.T) {
	//t.Parallel()
	want := 202
	imgid := mkimgid()
	resp := DeleteImage(imgid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
