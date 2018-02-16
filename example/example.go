package main

import (
	"fmt"
	"os"

	"github.com/profitbricks/profitbricks-sdk-go"
)

func main() {

	//Sets username and password
	client := profitbricks.NewClient(os.Getenv("PROFITBRICKS_USERNAME"), os.Getenv("PROFITBRICKS_PASSWORD"))

	dcrequest := profitbricks.Datacenter{
		Properties: profitbricks.DatacenterProperties{
			Name:        "Eexample",
			Description: "description",
			Location:    "us/las",
		},
	}

	datacenter, err := client.CreateDatacenter(dcrequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	serverrequest := profitbricks.Server{
		Properties: profitbricks.ServerProperties{
			Name:  "go01",
			RAM:   1024,
			Cores: 2,
		},
	}

	server, err := client.CreateServer(datacenter.ID, serverrequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = client.WaitTillProvisioned(server.Headers.Get("Location"))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	volumerequest := profitbricks.Volume{
		Properties: profitbricks.VolumeProperties{
			Size:        1,
			Name:        "Volume Test",
			LicenceType: "LINUX",
			Type:        "HDD",
		},
	}

	volume, err := client.CreateVolume(datacenter.ID, volumerequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = client.WaitTillProvisioned(volume.Headers.Get("Location"))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	serverupdaterequest := profitbricks.ServerProperties{
		Name:  "go01renamed",
		Cores: 1,
		RAM:   256,
	}

	server, err = client.UpdateServer(datacenter.ID, server.ID, serverupdaterequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = client.WaitTillProvisioned(server.Headers.Get("Location"))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	volume, err = client.AttachVolume(datacenter.ID, server.ID, volume.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = client.WaitTillProvisioned(volume.Headers.Get("Location"))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	volumes, err := client.ListVolumes(datacenter.ID)
	fmt.Println(volumes.Items)
	servers, err := client.ListServers(datacenter.ID)
	fmt.Println(servers.Items)
	datacenters, err := client.ListDatacenters()
	fmt.Println(datacenters.Items)

	resp, err := client.DeleteServer(datacenter.ID, server.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = client.WaitTillProvisioned(resp.Get("Location"))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	_, err = client.DeleteDatacenter(datacenter.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
