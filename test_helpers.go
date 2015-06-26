package profitbricks

func mkdcid() string {
	dc := CreateDatacenter([]byte(`{
    "properties": {
        "name": "GOSDK",
        "description": "datacenter-description",
        "location": "us/lasdev"
    }
	}`))

	return dc.Id
}

func mklocid() string {
	resp := ListLocations()

	locid := resp.Items[0].Id
	return locid
}
