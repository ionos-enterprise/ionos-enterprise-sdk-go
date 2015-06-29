package profitbricks

// ListNics returns a Nics struct collection
func ListNics(dcid, srvid string) Collection {
	path := nic_col_path(dcid, srvid)
	return is_list(path)
}

// CreateNic creates a nic on a server
// from a jason []byte and returns a Instance struct
func CreateNic(dcid string, srvid string, jason []byte) Instance {
	path := nic_col_path(dcid, srvid)
	return is_post(path, jason)
}

// GetNic pulls data for the nic where id = srvid returns a Instance struct
func GetNic(dcid, srvid, nicid string) Instance {
	path := nic_path(dcid, srvid, nicid)
	return is_get(path)
}

// PatchNic partial update of nic properties passed in as jason []byte
// Returns Instance struct
func PatchNic(dcid string, srvid string, nicid string, jason []byte) Instance {
	path := nic_path(dcid, srvid, nicid)
	return is_patch(path, jason)
}

// DeleteNic deletes the nic where id=nicid and returns a Resp struct
func DeleteNic(dcid, srvid, nicid string) Resp {
	path := nic_path(dcid, srvid, nicid)
	return is_delete(path)
}
