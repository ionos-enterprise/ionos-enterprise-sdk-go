package profitbricks

import "encoding/json"
import "fmt"

type CreateVolumeRequest struct {
	VolumeProperties `json:"properties"`
}

type VolumeProperties struct {
	Name          string   `json:"name,omitempty"`
	Size          int      `json:"size,omitempty"`
	Bus           string   `json:",bus,omitempty"`
	Image         string   `json:"image,omitempty"`
	Type          string   `json:"type,omitempty"`
	LicenceType   string   `json:"licenceType,omitempty"`
	ImagePassword string   `json:"imagePassword,omitempty"`
	SshKey        []string `json:"sshKeys,omitempty"`
}

// ListVolumes returns a Collection struct for volumes in the Datacenter
func ListVolumes(dcid string) Collection {
	path := volume_col_path(dcid)
	return is_list(path)
}

func GetVolume(dcid string, volumeId string) Instance {
	path := volume_path(dcid, volumeId)
	return is_get(path)
}

func PatchVolume(dcid string, volid string, request VolumeProperties) Instance {
	obj, _ := json.Marshal(request)
	path := volume_path(dcid, volid)
	return is_patch(path, obj)
}

func CreateVolume(dcid string, request CreateVolumeRequest) Instance {
	obj, _ := json.Marshal(request)
	path := volume_col_path(dcid)
	return is_post(path, obj)
}

func DeleteVolume(dcid, volid string) Resp {
	path := volume_path(dcid, volid)
	return is_delete(path)
}

func CreateSnapshot(dcid string, volid string, jason []byte) Resp {

	empty := `
		{}
		`
	var path = volume_path(dcid, volid)
	path = path + "/create-snapshot"
	var data StringMap
	json.Unmarshal(jason, &data)
	for key, value := range data {
		path += ("&" + key + "=" + value)
		fmt.Println(path)
	}
	return is_command(path, empty)
}

/**



	restoreSnapshot : function(dc_id,volume_id,jason,callback){
		pbreq.is_post([ "datacenters",dc_id,"volumes",volume_id,"restore-snapshot" ],str,callback)
	}
**/
