// server_test.go
package profitbricks

import (
	"testing"
	//"time"
)



var srv_dcid = mkdcid()

func mksrvid(srv_dcid string) string {
	var jason = []byte(`{"properties":{
						"name":"Original Server",
						"cores":4,
						"ram": 4096}
					}`)
	srv := CreateServer(srv_dcid, jason)

	srvid := srv.Id
	return srvid
}

func delsrvid(srv_dcid string, srv_id string) int {

	resp := DeleteServer(srv_dcid, srv_id)
	
	return resp.StatusCode
}


func TestCreateServer(t *testing.T) {
	//Setup
	t.Log(srv_dcid)
	dc := CreateDatacenter([]byte(`{
    "properties": {
        "name": "DCtest001",
        "description": "datacenter-description",
        "location": "us/lasdev"
    }
	}`))
    if dc.Resp.StatusCode > 205 {
		t.Errorf(bad_status(202, dc.Resp.StatusCode))
	}
	dcid := dc.Id
	
	//
	//
	//
	want := 202
	var jason = []byte(`{"properties":{
			"name":"go01",
			"cores":4,
			"ram": 4096
			}}`)
	t.Logf("Creating server in DC: %s", dcid)
	srv := CreateServer(dcid, jason)
    t.Logf("Server ...... %s\n", string(srv.Resp.Body))
	t.Logf("Server request ...... %s\n", string(srv.Href))
	// wait for server creation
	/*srv_prop := GetServer(srv_dcid, srv.Id)
	t.Logf("Server ...... %s\n", string(srv_prop.Resp.Body))
	num_tries := 120
	for num_tries > 0 && srv_prop.Resp.StatusCode == 404  {
		time.Sleep(time.Second)
		srv_prop = GetServer(srv_dcid, srv.Id)
		//t.Logf("Server ...... %s\n", string(srv_prop.Resp.Body))
		t.Logf("StatusCode ...... %d\n", srv_prop.Resp.StatusCode)
	}
	if num_tries == 0 {
		t.Errorf("Timeout! Server not created in 120 secs")
	} else {
		t.Logf("Server Properties ...... %s\n", string(srv.Properties["name"].(string)))
	}*/

	//cleanup
	t.Log("Performing cleanup...")
	delsrvid(dcid, srv.Id)
	
	if srv.Resp.StatusCode != want {
		t.Errorf(bad_status(want, srv.Resp.StatusCode))
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

func TestDeleteServer(t *testing.T) {
	//t.Parallel()
	want := 202
	srvid := mksrvid(srv_dcid)
	resp := DeleteServer(srv_dcid, srvid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestListServers(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	//
	// Setup
	//
	srvid := mksrvid(srv_dcid)
	
	//
	// List Servers
	//
	resp := ListServers(srv_dcid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
	
	//cleanup
	delsrvid(srv_dcid, srvid)
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
	//cleanup
	delsrvid(srv_dcid, srvid)
}


