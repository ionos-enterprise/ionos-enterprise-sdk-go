// server_test.go
package profitbricks

import (
	"testing"
	"sync"
	"time"
	"fmt"
)



var (
	once_dc 	sync.Once
	once_srv 	sync.Once
	srv_dc_id  	string
	srv_srvid	string
	srv_srv01	Instance
)

func setupDataCenter(){
	setupCredentials()
	srv_dc_id = mkdcid("SERVER DC")
	if len(srv_dc_id) == 0 { 
		//panic("DataCenter not created")
		fmt.Errorf("DataCenter not created")
	}
}

func setupServer(){
	srv_srvid = setupCreateServer(srv_dc_id)
	if len(srv_srvid) == 0 { 
		fmt.Errorf("DataCenter not created")
	}
}

// called from TestMain
func serverCleanup() {
	// TODO how to do cleanup, should use TestMain
	fmt.Println("Performing cleanup...")
	DeleteServer(srv_dc_id, srv_srv01.Id)
	DeleteDatacenter(srv_dc_id)
}

func setupCreateServer(srv_dc_id string) string {
	var jason = []byte(`{"properties":{
						"name":"GoServer",
						"cores":4,
						"ram": 4096}
					}`)
	fmt.Println("Creating server....")
	srv := CreateServer(srv_dc_id, jason)
	// wait for server to be running
	fmt.Println("Waiting for server to start....")
	srv_prop := GetServer(srv_dc_id, srv.Id)
	num_tries := 120
	seconds := 0
	for seconds < num_tries && srv_prop.Resp.StatusCode == 404  {
		time.Sleep(time.Second)
		srv_prop = GetServer(srv_dc_id, srv.Id)
		seconds++
	}
	if num_tries == 0 {
		fmt.Errorf("Timeout! Server not running in 120 secs")
	} else {
		fmt.Printf("Server %s created in %d seconds\n", string(srv.Properties["name"].(string)), seconds)
	}

	srvid := srv.Id
	return srvid
}


//
//  ----------------- Tests -------------------
//

func TestCreateServer(t *testing.T) {
	once_dc.Do(setupDataCenter)

	want := 202
	var jason = []byte(`{"properties":{
			"name":"go01",
			"cores":2,
			"ram": 1024
			}}`)
	t.Logf("Creating server in DC: %s", srv_dc_id)
	srv_srv01 = CreateServer(srv_dc_id, jason)
	if srv_srv01.Resp.StatusCode != want {
		t.Errorf(bad_status(want, srv_srv01.Resp.StatusCode))
	}
    //t.Logf("Server ...... %s\n", string(srv.Resp.Body))
}

func TestGetServer(t *testing.T) {
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	shouldbe := "server"
	want := 200
	resp := GetServer(srv_dc_id, srv_srvid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestListServers(t *testing.T) {
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	shouldbe := "collection"
	want := 200
	
	//
	// List Servers
	//
	resp := ListServers(srv_dc_id)
	//t.Logf("ListServers ...... %s\n", string(resp.Resp.Body))
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

}

func TestPatchServer(t *testing.T) {
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 202
	jason_patch := []byte(` {
				"name": "NewName",
				"cores": 1
			}`)
	resp := PatchServer(srv_dc_id, srv_srvid, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestStopServer(t *testing.T){
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 202
	resp := StopServer(srv_dc_id, srv_srvid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	
}

func TestStartServer(t *testing.T){
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 202
	resp := StartServer(srv_dc_id, srv_srvid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	
}

func TestRebootServer(t *testing.T){
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 202
	resp := RebootServer(srv_dc_id, srv_srvid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	
}

func TestListAttachedVolumes_NoItems(t *testing.T){
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 200
	shouldbe := "collection"

	resp := ListAttachedVolumes(srv_dc_id, srv_srvid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Error(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}


func TestListAttachedCdroms_NoItems(t *testing.T){
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 200
	shouldbe := "collection"
	
	resp := ListAttachedCdroms(srv_dc_id, srv_srvid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Error(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}


func TestDeleteServer(t *testing.T) {
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 202

	resp := DeleteServer(srv_dc_id, srv_srvid)
	if resp.StatusCode != want {
		t.Error(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

// TODO Tests 
// AttachCdrom
// GetAttachedCdrom
// DetachCdrom
// AttachVolume
// GetAttachedVolume
// DetachVolume