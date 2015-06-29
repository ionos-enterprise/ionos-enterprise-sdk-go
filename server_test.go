// server_test.go
package profitbricks

import (
	"testing"
	"sync"
	"time"
	"fmt"
	"strings"
)

var (
	once_dc 	sync.Once
	once_srv 	sync.Once
	srv_dc_id  	string
	srv_srvid	string
	srv_srv01	string
	srv_vol_id  string
	srv_cdrom	string
)


func setupDataCenter(){
	setupCredentials()
	srv_dc_id = mkdcid()
	if len(srv_dc_id) == 0 { 
		//panic("DataCenter not created")
		fmt.Errorf("DataCenter not created")
	}
}

func setupServer(){
	srv_srvid = setupCreateServer(srv_dc_id)
	fmt.Println("Server id: ", srv_srvid)
	if len(srv_srvid) == 0 { 
		fmt.Errorf("Server not created")
	}
}

// called from TestMain
func serverCleanup() {
	// TODO how to do cleanup, should use TestMain
	fmt.Println("Performing cleanup...")
	res := DeleteServer(srv_dc_id, srv_srvid)
	fmt.Println("DeleteServer: ", res.StatusCode)
	res = DeleteVolume(srv_dc_id, srv_vol_id)
	fmt.Println("DeleteVolume: ", res.StatusCode)
	res = DeleteDatacenter(srv_dc_id)
	fmt.Println("DeleteDatacenter: ", res.StatusCode)
}

func setupCreateServer(srv_dc_id string) string {
	var jason = []byte(`{"properties":{
						"name":"GoServer",
						"cores":1,
						"ram": 1024}
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
	srv := CreateServer(srv_dc_id, jason)
	srv_srv01 = srv.Id
	if srv.Resp.StatusCode != want {
		t.Errorf(bad_status(want, srv.Resp.StatusCode))
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

func TestAttachCdrom(t *testing.T) {
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 202


	// Setup -- Find appropriate image
	resp := ListImages()	
	for i := 0; i < len(resp.Items); i++ {
		name := resp.Items[i].Properties["name"].(string)
		region := resp.Items[i].Properties["location"].(string)
		img_type := resp.Items[i].Properties["imageType"].(string)
		
		if ( strings.HasPrefix(name, "ubuntu-") &&
			region == "us/lasdev" && img_type == "CDROM") {
				fmt.Println("Found volume: ", name)
				srv_cdrom = resp.Items[i].Id
				break
		}	
	}
	
	//
	// Test
	//
	resp_cdrom := AttachCdrom(srv_dc_id, srv_srvid, srv_cdrom)
 	if resp_cdrom.Resp.StatusCode != want {
		t.Error(string(resp_cdrom.Resp.Body))
		t.Errorf(bad_status(want, resp_cdrom.Resp.StatusCode))
	}
}

func TestListAttachedCdroms(t *testing.T){
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 200
	shouldbe := "collection"
	
	// wait for volume to attach
	time.Sleep(time.Second * 120)
	resp := ListAttachedCdroms(srv_dc_id, srv_srvid)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Error(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}


func TestGetAttachedCdrom(t *testing.T) {
	want := 200
	shouldbe := "volume"
	
	resp := GetAttachedCdrom(srv_dc_id, srv_srvid, srv_cdrom)
	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Error(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestDetachCdrom(t *testing.T) {
	want := 202
	
	resp := DetachCdrom(srv_dc_id, srv_srvid, srv_cdrom)
	if resp.StatusCode != want {
		t.Error(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}	
}

func TestDeleteServer(t *testing.T) {
	once_dc.Do(setupDataCenter)
	once_dc.Do(setupServer)
	
	want := 202

	resp := DeleteServer(srv_dc_id, srv_srv01)
	if resp.StatusCode != want {
		t.Error(string(resp.Body))
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

