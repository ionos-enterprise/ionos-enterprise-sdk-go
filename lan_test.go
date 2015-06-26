// lan_test.go
package profitbricks

import (
	"testing"
	"time"
)
import "fmt"

var lan_dcid string
var lanid string

func TestCreateLan(t *testing.T) {
	lan_dcid = mkdcid()
	want := 202
	var jason = []byte(`{
					  "properties": {
         			   "public": "true"
        			}
					}`)
	fmt.Println(lan_dcid)

	lan := CreateLan(lan_dcid, jason)
	lanid = lan.Id
	if lan.Resp.StatusCode != want {
		t.Errorf(bad_status(want, lan.Resp.StatusCode))
	}
	time.Sleep(2 * time.Second)
}

func TestListLans(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	lans := ListLans(lan_dcid)

	if lans.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, lans.Type))
	}
	if lans.Resp.StatusCode != want {
		t.Errorf(bad_status(want, lans.Resp.StatusCode))
	}
}

func TestGetLan(t *testing.T) {
	shouldbe := "lan"
	want := 200
	lan := GetLan(lan_dcid, lanid)
	if lan.Type != shouldbe {
		t.Errorf(bad_status(want, lan.Resp.StatusCode))
	}
	if lan.Resp.StatusCode != want {
		t.Errorf(bad_status(want, lan.Resp.StatusCode))
	}
}

func TestPatchLan(t *testing.T) {
	want := 202
	jason_patch := []byte(`{
					 "public": "false"
					}`)
	lan := PatchLan(lan_dcid, lanid, jason_patch)
	if lan.Resp.StatusCode != want {
		t.Errorf(bad_status(want, lan.Resp.StatusCode))
	}
}

func TestDeleteLan(t *testing.T) {
	want := 202
	resp := DeleteLan(lan_dcid, lanid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
