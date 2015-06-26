// server_test.go
package profitbricks

import (
	"testing"
	//"sync"
	"time"
	"fmt"
)





// TODO Tests 
// AttachCdrom
// GetAttachedCdrom
// DetachCdrom
// AttachVolume


func TestAttachVolume(t *testing.T) {
	once_dc.Do(setupDataCenter)
	once_srv.Do(setupServer)
	
	want := 202

	var jason = []byte(`{
	    "properties": {
	        "size": "2",
	        "name": "volume-name",
			"licenceType" : "LINUX"
	    }
		}`)

	vol := CreateVolume(srv_dc_id, jason)
	vol_prop := GetVolume(srv_dc_id, vol.Id)
	t.Log(string(vol_prop.Resp.Body))
	num_tries := 120
	seconds := 0
	for seconds < num_tries && vol_prop.Resp.StatusCode == 404  {
		time.Sleep(time.Second)
		vol_prop = GetVolume(srv_dc_id, vol.Id)
		t.Log(string(vol_prop.Resp.Body))
		seconds++
	}
	if num_tries == 0 {
		fmt.Errorf("Timeout! Server not running in 120 secs")
	} else {
		fmt.Printf("Server %s created in %d seconds\n", string(vol.Properties["name"].(string)), seconds)
	}
	srv_vol_id = vol.Id
	t.Log("VolumeId: ", vol.Id, " , Server Id: ", srv_srvid, " ,DC id: ", srv_dc_id)
	time.Sleep(time.Second*20)
	vol_prop = GetVolume(srv_dc_id, vol.Id)
	t.Log(string(vol_prop.Resp.Body))
	
	resp := AttachVolume(srv_dc_id, srv_srvid, srv_vol_id)
	if resp.Resp.StatusCode != want {
		t.Error(string(resp.Resp.Body))
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}	

}

// GetAttachedVolume
// DetachVolume

