package profitbricks

import (
	"fmt"
	"testing"
)

var dcID string

func TestListDatacenters(t *testing.T) {
	shouldbe := "collection"
	want := 200
	resp := ListDatacenters()

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestCreateDatacenter(t *testing.T) {
	want := 202
	var obj = CreateDatacenterRequest{
		DCProperties: DCProperties{
			Name:        "GO SDK",
			Description: "description",
			Location:    "us/lasdev",
		},
	}
	resp := CreateDatacenter(obj)
	dcID = resp.Id
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestGetDatacenter(t *testing.T) {
	shouldbe := "datacenter"
	want := 200

	fmt.Println(dcID)
	resp := GetDatacenter(dcID)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestPatchDatacenter(t *testing.T) {
	want := 202
	obj := map[string]string{"name": "Renamed DC"}

	resp := PatchDatacenter(dcID, obj)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

}

func TestDeleteDatacenter(t *testing.T) {
	want := 202
	resp := DeleteDatacenter(dcID)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
