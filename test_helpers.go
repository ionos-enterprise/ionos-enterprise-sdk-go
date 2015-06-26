package profitbricks

import "fmt"

func mkdcid() string {
	fmt.Println("Setting up dc")
	dc := CreateDatacenter([]byte(`{
    "properties": {
        "name": "VOLUME",
        "description": "datacenter-description",
        "location": "us/lasdev"
    }
	}`))

	fmt.Println(dc.Id)
	return dc.Id
}
