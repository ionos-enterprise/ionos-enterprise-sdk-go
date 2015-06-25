// server_test.go
package profitbricks

import (
	"testing"
	"sync"
	"time"
)



var (
	once 	sync.Once
	dc_id   string
	srv_id  string
)

func setup(){
	dc_id = mkdcid()
}


func mksrvid(dc_id string) string {
	var jason = []byte(`{"properties":{
						"name":"GoServer",
						"cores":4,
						"ram": 4096}
					}`)
	srv := CreateServer(dc_id, jason)

	srvid := srv.Id
	return srvid
}

func delsrvid(dc_id string, srv_id string) int {

	resp := DeleteServer(dc_id, srv_id)
	
	return resp.StatusCode
}


func TestCreateServer(t *testing.T) {
	once.Do(setup)
	//
	//
	//
	want := 202
	var jason = []byte(`{"properties":{
			"name":"go01",
			"cores":2,
			"ram": 1024
			}}`)
	t.Logf("Creating server in DC: %s", dc_id)
	srv := CreateServer(dc_id, jason)
	if srv.Resp.StatusCode != want {
		t.Errorf(bad_status(want, srv.Resp.StatusCode))
	}
	
	srv_id = srv.Id
    //t.Logf("Server ...... %s\n", string(srv.Resp.Body))

	
	// wait for server creation
	srv_prop := GetServer(dc_id, srv.Id)
	t.Logf("Server ...... %s\n", string(srv_prop.Resp.Body))
	num_tries := 120
	for num_tries > 0 && srv_prop.Resp.StatusCode == 404  {
		time.Sleep(time.Second)
		srv_prop = GetServer(dc_id, srv.Id)
		//t.Logf("Server ...... %s\n", string(srv_prop.Resp.Body))
		t.Logf("StatusCode ...... %d\n", srv_prop.Resp.StatusCode)
	}
	if num_tries == 0 {
		t.Errorf("Timeout! Server not created in 120 secs")
	} else {
		t.Logf("Server %s created in %d seconds\n", string(srv.Properties["name"].(string)), num_tries)
	}

	//cleanup
	//t.Log("Performing cleanup...")
	//delsrvid(dc_id, srv.Id)
	
	if srv.Resp.StatusCode != want {
		t.Errorf(bad_status(want, srv.Resp.StatusCode))
	}
}
func TestGetServer(t *testing.T) {
	once.Do(setup)
	
	shouldbe := "server"
	want := 200
	//srvid := mksrvid(dc_id)
	resp := GetServer(dc_id, srv_id)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestListServers(t *testing.T) {
	once.Do(setup)
	
	shouldbe := "collection"
	want := 200
	
	//
	// List Servers
	//
	resp := ListServers(dc_id)
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
	jason_patch := []byte(` {
				"cores": 1
			}`)
	//srvid := mksrvid(dc_id)
	resp := PatchServer(dc_id, srv_id, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestDeleteServer(t *testing.T) {
	once.Do(setup)
	
	want := 202

	resp := DeleteServer(dc_id, srv_id)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	
	//cleanup
	DeleteDatacenter(dc_id)
}

