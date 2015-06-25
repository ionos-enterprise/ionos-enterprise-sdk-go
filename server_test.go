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
	dc_id  		string
	srv_id		string
)

func SetupDataCenter(){
	SetupCredentials()
	dc_id = mkdcid()
	if len(dc_id) == 0 { 
		//panic("DataCenter not created")
		fmt.Errorf("DataCenter not created")
	}
}

func SetupServer(){
	srv_id = SetupCreateServer(dc_id)
	if len(srv_id) == 0 { 
		fmt.Errorf("DataCenter not created")
	}
}

func SetupCreateServer(dc_id string) string {
	var jason = []byte(`{"properties":{
						"name":"GoServer",
						"cores":4,
						"ram": 4096}
					}`)
	fmt.Println("----------------- Creating server.... ----------")
	srv := CreateServer(dc_id, jason)
	// wait for server to be running
	fmt.Println("Waiting for server to start....")
	srv_prop := GetServer(dc_id, srv.Id)
	num_tries := 120
	seconds := 0
	for seconds < num_tries && srv_prop.Resp.StatusCode == 404  {
		time.Sleep(time.Second)
		srv_prop = GetServer(dc_id, srv.Id)
		seconds++
	}
	fmt.Println("------ Server available ----------")

	srvid := srv.Id
	return srvid
}

func TestCreateServer(t *testing.T) {
	once_dc.Do(SetupDataCenter)
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
    //t.Logf("Server ...... %s\n", string(srv.Resp.Body))

	
	// wait for server creation
	/*srv_prop := GetServer(dc_id, srv.Id)
	t.Logf("Server ...... %s\n", string(srv_prop.Resp.Body))
	num_tries := 120
	seconds := 0
	for seconds < num_tries && srv_prop.Resp.StatusCode == 404  {
		time.Sleep(time.Second)
		srv_prop = GetServer(dc_id, srv.Id)
		//t.Logf("Server ...... %s\n", string(srv_prop.Resp.Body))
		t.Logf("StatusCode ...... %d\n", srv_prop.Resp.StatusCode)
		seconds++
	}
	if num_tries == 0 {
		t.Errorf("Timeout! Server not running in 120 secs")
	} else {
		t.Logf("Server %s created in %d seconds\n", string(srv.Properties["name"].(string)), seconds)
	}*/

	//cleanup
	t.Log("Performing cleanup...")
	DeleteServer(dc_id, srv.Id)
}

func TestGetServer(t *testing.T) {
	once_dc.Do(SetupDataCenter)
	once_srv.Do(SetupServer)
	
	shouldbe := "server"
	want := 200
	resp := GetServer(dc_id, srv_id)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestListServers(t *testing.T) {
	once_dc.Do(SetupDataCenter)
	once_srv.Do(SetupServer)
	
	shouldbe := "collection"
	want := 200
	
	//
	// List Servers
	//
	resp := ListServers(dc_id)
	//t.Logf("ListServers ...... %s\n", string(resp.Resp.Body))
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

}

func TestPatchServer(t *testing.T) {
	once_dc.Do(SetupDataCenter)
	once_srv.Do(SetupServer)
	
	want := 202
	jason_patch := []byte(` {
				"name": "NewName",
				"cores": 1
			}`)
	//srvid := CreateServer(dc_id)
	resp := PatchServer(dc_id, srv_id, jason_patch)
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestStopServer(t *testing.T){
	once_dc.Do(SetupDataCenter)
	once_srv.Do(SetupServer)
	
	want := 202
	resp := StopServer(dc_id, srv_id)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	
}

func TestStartServer(t *testing.T){
	once_dc.Do(SetupDataCenter)
	once_srv.Do(SetupServer)
	
	want := 202
	resp := StartServer(dc_id, srv_id)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	
}

func TestDeleteServer(t *testing.T) {
	once_dc.Do(SetupDataCenter)
	once_srv.Do(SetupServer)
	
	want := 202

	resp := DeleteServer(dc_id, srv_id)
	if resp.StatusCode != want {
		t.Error(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	
	//cleanup
	DeleteDatacenter(dc_id)
}

