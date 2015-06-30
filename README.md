# profitbricks-sdk-go
	Example
	profitbricks.SetAuth("user@name.com", "password")
	profitbricks.SetDepth("5")

	obj := profitbricks.CreateDatacenterRequest{
		Properties: profitbricks.Properties{
			Name:        "test",
			Description: "description",
			Location:    "us/lasdev",
		},
	}
	
	dc := profitbricks.CreateDatacenter(obj)
	
	sm := map[string]string{"name": "Renamed DC"}
	jason_patch := []byte(profitbricks.MkJson(sm))

	resp := profitbricks.PatchDatacenter(dc.Id,jason_patch)