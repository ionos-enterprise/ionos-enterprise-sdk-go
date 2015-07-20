## GO SDK

The ProfitBricks Client Library for [GO](https://www.golang.org/) provides you with access to the ProfitBricks REST API. It is designed for developers who are building applications in GO.

This guide will walk you through getting setup with the library and performing various actions against the API.

## Concepts

The GO SDK wraps the latest version of the ProfitBricks REST API. All API operations are performed over SSL and authenticated using your ProfitBricks portal credentials. The API can be accessed within an instance running in ProfitBricks or directly over the Internet from any application that can send an HTTPS request and receive an HTTPS response.

## Getting Started

Before you begin you will need to have [signed-up](https://www.profitbricks.com/signup) for a ProfitBricks account. The credentials you setup during sign-up will be used to authenticate against the API. 

## Pre-Requisites

GO SDK has some pre-requisities before you're able to use it. You will need to:
 
- Install GO language environment from

```	
	https://golang.org/doc/install 
```
#### Set your Environment
The GOPATH environment variable specifies the location of your workspace. It is likely the only environment variable you'll need to set when developing Go code.

```
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

#### Fetch profitbricks-sdk-go

This will download profitbricks-sdk-go to the [GOPATH](#setenvironment)

```go
go get "github.com/profitbricks/profitbricks-sdk-go" 
```

The source code of the package will be located at 

	$GOBIN\src\profitbricks-sdk-go

Create main package file *example.go*:

	package main
	
	import (
		"fmt"
	)
	
	func main() {
	}

Import GO SDK

	import(
		"profitbricks-sdk-go"
	)

	
Set Username, Password, and Endpoint for testing

	profitbricks.SetAuth("username", "password")

Set depth:

	profitbricks.SetDepth("5")


Depth controls the amount of data returned from the rest server ( range 1-5 ). Higher the number more information is returned from the server. This is especially useful if you are looking for the information in the nested objects. 


**Caution**: You will want to ensure you follow security best practices when using credentials within your code or stored in a file.


### HOW TO'S

###HOW TO: CREATE A DATA CENTER

ProfitBricks introduces the concept of Data Centers. These are logically separated from one another and allow you to have a self-contained environment for all servers, volumes, networking, snapshots, and so forth. The goal is to give you the same experience as you would have if you were running your own physical data center.

The following code example shows you how to programmatically create a data center:

		request := profitbricks.CreateDatacenterRequest{
		DCProperties: profitbricks.DCProperties{
			Name:        "test",
			Description: "description",
			Location:    "us/lasdev",
		},
	}

	response := profitbricks.CreateDatacenter(request)

###HOW TO: Delete a Data Center

You will want to exercise a bit of caution here. Removing a data center will destroy all objects contained within that data center -- servers, volumes, snapshots, and so on.

The code to remove a data center is as follows. This example assumes you want to remove previously data center:

	profitbricks.DeleteDatacenter(response.Id)

###HOW TO: CREATE A SERVER

The server create method has a list of required parameters followed by a hash of optional parameters. The optional parameters are specified within the "options" hash and the variable names match the [REST API](https://devops.profitbricks.com/api/rest/) parameters.

The following example shows you how to create a new server in the data center created above:

	request = CreateServerRequest{
		ServerProperties: ServerProperties{
			Name:  "go01",
			Ram:   1024,
			Cores: 2,
		},
	}
	server := CreateServer(datacenter.Id, req)

### HOW TO: LIST AVAILABLE DISK AND ISO IMAGES

A list of disk and ISO images are available from ProfitBricks for immediate use. These can be easily viewed and selected. The following shows you how to get a list of images. This list represents both CDROM images and HDD images.

	images := profitbricks.ListImages()

This will return [collection](#Collection) object

### HOW TO: CREATE A STORAGE VOLUME
ProfitBricks allows for the creation of multiple storage volumes that can be attached and detached as needed. It is useful to attach an image when creating a storage volume. The storage size is in gigabytes.

	volumerequest := CreateVolumeRequest{
		VolumeProperties: VolumeProperties{
			Size:        1,
			Name:        "Volume Test",
			LicenceType: "LINUX",
		},
	}

	storage := CreateVolume(datacenter.Id, volumerequest)

 
###HOW TO: UPDATE SERVER CORES, AND MEMORY
ProfitBricks allows users to dynamically update cores, memory, and disk independently of each other. This removes the restriction of needing to upgrade to the next size available size to receive an increase in memory. You can now simply increase the instances memory keeping your costs in-line with your resource needs.

Note: The memory parameter value must be a multiple of 256, e.g. 256, 512, 768, 1024, and so forth.

The following code illustrates how you can update cores and memory:

	serverupdaterequest := profitbricks.ServerProperties{
		Cores: 1,
		Ram:   256,
	}
	resp := PatchServer(datacenter.Id, server.Id, serverupdaterequest)

###HOW TO: ATTACH AND DETACH A STORAGE VOLUME
ProfitBricks allows for the creation of multiple storage volumes. You can detach and reattach these on the fly. This allows for various scenarios such as re-attaching a failed OS disk to another server for possible recovery or moving a volume to another location and spinning it up.

The following illustrates how you would attach and detach a volume and CDROM to/from a server:

	profitbricks.AttachVolume(datacenter.Id, server.Id, volume.Id)
	profitbricks.AttachCdrom(datacenter.Id, server.Id, images.Items[0].Id)

	profitbricks.DetachVolume(datacenter.Id, server.Id, volume.Id)
	profitbricks.DetachCdrom(datacenter.Id, server.Id, images.Items[0].Id)

###HOW TO: LIST SERVERS, VOLUMES, AND DATA CENTERS

GO SDK provides standard functions for retrieving a list of volumes, servers, and datacenters.

The following code illustrates how to pull these three list types:

	volumes := profitbricks.ListVolumes(datacenter.Id)
	servers := profitbricks.ListServers(datacenter.Id)
	datacenters := profitbricks.ListDatacenters()


###Example 

	
	package main
	
	import (
		"fmt"
		"profitbricks-sdk-go"
		"time"
	)
	
	func main() {
	
		//Sets username and password
		profitbricks.SetAuth("username", "password")
		//Sets depth.
		profitbricks.SetDepth("5")
	
		dcrequest := profitbricks.CreateDatacenterRequest{
			DCProperties: profitbricks.DCProperties{
				Name:        "example.go3",
				Description: "description",
				Location:    "us/lasdev",
			},
		}
	
		datacenter := profitbricks.CreateDatacenter(dcrequest)
	
		serverrequest := profitbricks.CreateServerRequest{
			ServerProperties: profitbricks.ServerProperties{
				Name:  "go01",
				Ram:   1024,
				Cores: 2,
			},
		}
		server := profitbricks.CreateServer(datacenter.Id, serverrequest)
	
		images := profitbricks.ListImages()
	
		fmt.Println(images.Items)
		
		volumerequest := profitbricks.CreateVolumeRequest{
			VolumeProperties: profitbricks.VolumeProperties{
				Size:        1,
				Name:        "Volume Test",
				LicenceType: "LINUX",
			},
		}
	
		storage := profitbricks.CreateVolume(datacenter.Id, volumerequest)
	
		serverupdaterequest := profitbricks.ServerProperties{
			Name:  "go01renamed",
			Cores: 1,
			Ram:   256,
		}
	
		profitbricks.PatchServer(datacenter.Id, server.Id, serverupdaterequest)

		//It takes a moment for a volume to be provisioned so we wait before we attach it to a server
		time.Sleep(60 * time.Second)
		
		profitbricks.AttachVolume(datacenter.Id, server.Id, storage.Id)
	
		volumes := profitbricks.ListVolumes(datacenter.Id)
		fmt.Println(volumes.Items)

		servers := profitbricks.ListServers(datacenter.Id)
		fmt.Println(servers.Items)

		datacenters := profitbricks.ListDatacenters()
		fmt.Println(datacenters.Items)
	
		profitbricks.DeleteServer(datacenter.Id, server.Id)
		profitbricks.DeleteDatacenter(datacenter.Id)
	}


## Return Types  ##

## Resp struct
* 	Resp is the struct returned by all Rest request functions

```go
type Resp struct {
Req        *http.Request
StatusCode int
Headers    http.Header
Body       []byte
}
```

## ```Instance struct```
* 	"Get", "Create", and "Patch" functions all return an Instance struct.
*	A Resp struct is embedded in the Instance struct,
*	the raw server response is available as Instance.Resp.Body
		
```go
type Instance struct {
Id_Type_Href
MetaData   StringMap           `json:"metaData"`
Properties StringIfaceMap      `json:"properties"`
Entities   StringCollectionMap `json:"entities"`
Resp       Resp                `json:"-"`
}
```

## Collection struct 
* 	Collection Structs contain Instance arrays. 
* 	List functions return Collections

```go
type Collection struct {
Id_Type_Href
Items []Instance `json:"items,omitempty"`
Resp  Resp       `json:"-"`
}
```
