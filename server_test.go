// server_test.go
package goprofitbricks

import "testing"
import "fmt"

var srv_dcid = mkdcid()

func mksrvid(srv_dcid string) string {
	var jason = []byte(`{
					"name":"Original Server",
					"cores":4,
					"ram": 4096
					}`)
	srv := CreateServer(srv_dcid, jason)

	srvid := srv.Id
	return srvid
}

func ExampleListServers() {
	s := ListServers(srv_dcid)
	fmt.Println(s.Resp.StatusCode)
	// Output: 200

}

func TestListServers(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListServers(srv_dcid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestCreateServer(t *testing.T) {
	//t.Parallel()
	want := 202
	var jason = []byte(`{
			"name":"Goat",
			"cores":4,
			"ram": 4096
			}`)
	resp := CreateServer(srv_dcid, jason)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestGetServer(t *testing.T) {
	//t.Parallel()
	shouldbe := "server"
	want := 200
	srvid := mksrvid(srv_dcid)
	resp := GetServer(srv_dcid, srvid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestPatchServer(t *testing.T) {
	//t.Parallel()
	want := 202
	jason_patch := []byte(`{
			"name":"Renamed Server",
					}`)
	srvid := mksrvid(srv_dcid)
	resp := PatchServer(srv_dcid, srvid, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestUpdateServer(t *testing.T) {
	//t.Parallel()
	want := 202
	jason_update := []byte(`{
			"name":"Renamed Server",
			"cores":16,
			"ram": 8192
					}`)

	srvid := mksrvid(srv_dcid)
	resp := UpdateServer(srv_dcid, srvid, jason_update)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}
func TestDeleteServer(t *testing.T) {
	//t.Parallel()
	want := 202
	srvid := mksrvid(srv_dcid)
	resp := DeleteServer(srv_dcid, srvid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
