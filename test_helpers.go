package profitbricks

import "fmt"


func mkdcid(name string) string {
	request := CreateDatacenterRequest{
		DCProperties: DCProperties{
			Name:        name,
			Description: "description",
			Location:    "us/las",
		},
	}
	dc := CreateDatacenter(request)
	return dc.Id
}

func mklocid() string {
	resp := ListLocations()

	locid := resp.Items[0].Id
	return locid
}

func mksrvid(srv_dcid string) string {
	
	var req = CreateServerRequest{
		ServerProperties: ServerProperties{
			Name:        "test",
			Ram: 1024,
			Cores:    2,
		},
	}
	srv := CreateServer(srv_dcid, req)

	return srv.Id
}

func mknic(lbal_dcid, serverid string) string {
	resp := CreateNic(lbal_dcid, serverid, []byte(`{"properties": {"name":"Original Nic","lan":1}}`))
	fmt.Println("===========================")
	fmt.Println("created a nic with id " + resp.Id)
	fmt.Println("===========================")
	return resp.Id
}
