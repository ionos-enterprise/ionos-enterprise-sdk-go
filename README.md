# Go SDK

The ProfitBricks Client Library for [Go](https://www.golang.org/) provides you with access to the ProfitBricks REST API. It is designed for developers who are building applications in Go.

This guide will walk you through getting setup with the library and performing various actions against the API.

# Table of Contents
* [Concepts](#concepts)
* [Getting Started](#getting-started)
* [Installation](#installation)
* [How to: Create Data Center](#how-to-create-data-center)
* [How to: Delete Data Center](#how-to-delete-data-center)
* [How to: Create Server](#how-to-create-server)
* [How to: List Available Images](#how-to-list-available-images)
* [How to: Create Storage Volume](#how-to-create-storage-volume)
* [How to: Update Cores and Memory](#how-to-update-cores-and-memory)
* [How to: Attach or Detach Storage Volume](#how-to-attach-or-detach-storage-volume)
* [How to: List Servers, Volumes, and Data Centers](#how-to-list-servers-volumes-and-data-centers)
* [Reference](#reference)
	* [Objects](#objects) 
		* [Datacenter](#datacenter) 
		    * [DatacenterProperties](#datacenterproperties) 
		    * [DatacenterEntities](#datacenterentities) 
		* [Datacenters](#datacenters)
		* [Server](#server)
            * [ServerProperties](#serverproperties) 
            * [ServerEntities](#serverentities)
        * [Servers](#servers)
        * [Volume](#volume)
            * [VolumeProperties](#volumeproperties) 
        * [Volumes](#volumes)
        * [Nic](#nic)
            * [NicProperties](#niceproperties) 
            * [NicEntities](#nicentities)
        * [Nics](#nics)
        * [FirewallRule](#firewallrule)
            * [FirewallruleProperties](#firewallruleproperties)
        * [FirewallRules](#firewallrules)
        * [Lan](#lan)
            * [LanProperties](#lanproperties)
        * [Lans](#lans) 
        * [Image](#image)
            * [ImageProperties](#imageproperties)
        * [Images](#images)     
        * [Loadbalancer](#loadbalancer)
            * [LoadbalancerProperties](#loadbalancerproperties) 
            * [LoadbalancerEntities](#loadbalancerentities)
        * [Loadbalancers](#loadbalancers)
        * [Location](#location)
            * [LocationProperties](#locationproperties) 
        * [Locations](#locations)
        * [IpBlock](#ipblock)
            * [IpBlockProperties](#ipblockproperties) 
        * [IpBlocks](#ipblocks)
        * [Snapshot](#snapshot)
            * [SnapshotProperties](#snapshotproperties) 
        * [Snapshots](#snapshots)
        * [Request](#request)
            * [RequestProperties](#requestproperties)
            * [RequestStatusMetadata](#requeststatusmetadata)
            * [RequestStatusMetadata](#requeststatusmetadata)  
        * [Requests](#requests)
                    
	* [Functions](#functions) 
	    * [Virtual Data Centers](#virtual-data-centers)
	    * [Servers](#servers)
	    * [Volumes](#volumes)
	    * [NICs](#nics)
	    * [Firewall Rules](#firewall-rules)
	    * [LANs](#lans)
	    * [Images](#images)
	    * [Load Balancers](#load-balancers)
	    * [IP Blocks](#ip-blocks)
	    * [Snapshot](#snapshot-functions)
	    * [Requests](#request-functions)
        * [Locations](#location-functions)    	
    
* [Example](#example)
* [Support](#support)


# Concepts

The Go SDK wraps the latest version of the ProfitBricks REST API. All API operations are performed over SSL and authenticated using your ProfitBricks portal credentials. The API can be accessed within an instance running in ProfitBricks or directly over the Internet from any application that can send an HTTPS request and receive an HTTPS response.

# Getting Started

Before you begin you will need to have [signed-up](https://www.profitbricks.com/signup) for a ProfitBricks account. The credentials you setup during sign-up will be used to authenticate against the API.

Install the Go language from: [Go Installation](https://golang.org/doc/install)

The `GOPATH` environment variable specifies the location of your Go workspace. It is likely the only environment variable you'll need to set when developing Go code. This is an example of pointing to a workspace configured underneath your home directory:

```
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

# Installation

The following go command will download `profitbricks-sdk-go` to your configured `GOPATH`:

```go
go get "github.com/profitbricks/profitbricks-sdk-go"
```

The source code of the package will be located at:

	$GOBIN\src\profitbricks-sdk-go

Create main package file *example.go*:

```go
package main

import (
	"fmt"
)

func main() {
}
```

Import GO SDK:

```go
import(
	"github.com/profitbricks/profitbricks-sdk-go"
)
```

Add your credentials for connecting to ProfitBricks:

```go
profitbricks.SetAuth("username", "password")
```

Set depth:

```go
profitbricks.SetDepth("5")
```

Depth controls the amount of data returned from the REST server ( range 1-5 ). The larger the number the more information is returned from the server. This is especially useful if you are looking for the information in the nested objects.

**Caution**: You will want to ensure you follow security best practices when using credentials within your code or stored in a file.

# How To's

## How To: Create Data Center

ProfitBricks introduces the concept of Data Centers. These are logically separated from one another and allow you to have a self-contained environment for all servers, volumes, networking, snapshots, and so forth. The goal is to give you the same experience as you would have if you were running your own physical data center.

The following code example shows you how to programmatically create a data center:

```go
dcrequest := profitbricks.Datacenter{
		Properties: profitbricks.DatacenterProperties{
			Name:        "example.go3",
			Description: "description",
			Location:    "us/lasdev",
		},
	}

datacenter := profitbricks.CreateDatacenter(dcrequest)
```

## How To: Create Data Center with Multiple Resources

To create a complex Data Center you would do this. As you can see, you can create quite a few of the objects you will need later all in one request.:

```go
datacenter := model.Datacenter{
		Properties: model.DatacenterProperties{
			Name: "composite test",
			Location:location,
		},
		Entities:model.DatacenterEntities{
			Servers: &model.Servers{
				Items:[]model.Server{
					model.Server{
						Properties: model.ServerProperties{
							Name : "server1",
							Ram: 2048,
							Cores: 1,
						},
						Entities:model.ServerEntities{
							Volumes: &model.AttachedVolumes{
								Items:[]model.Volume{
									model.Volume{
										Properties: model.VolumeProperties{
											Type_:"HDD",
											Size:10,
											Name:"volume1",
											Image:"1f46a4a3-3f47-11e6-91c6-52540005ab80",
											Bus:"VIRTIO",
											ImagePassword:"test1234",
											SshKeys: []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoLVLHON4BSK3D8L4H79aFo..."},
										},
									},
								},
							},
							Nics: &model.Nics{
								Items: []model.Nic{
									model.Nic{
										Properties: model.NicProperties{
											Name : "nic",
											Lan : "1",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	
dc := CompositeCreateDatacenter(datacenter)

```


## How To: Delete Data Center

You will want to exercise a bit of caution here. Removing a data center will destroy all objects contained within that data center -- servers, volumes, snapshots, and so on.

The code to remove a data center is as follows. This example assumes you want to remove previously data center:

```go
profitbricks.DeleteDatacenter(response.Id)
```

## How To: Create Server

The server create method has a list of required parameters followed by a hash of optional parameters. The optional parameters are specified within the "options" hash and the variable names match the [REST API](https://devops.profitbricks.com/api/rest/) parameters.

The following example shows you how to create a new server in the data center created above:

```go
req := profitbricks.Server{
 		Properties: profitbricks.ServerProperties{
 			Name:  "go01",
 			Ram:   1024,
 			Cores: 2,
 		},
}
server := CreateServer(datacenter.Id, req)
```

## How To: List Available Images

A list of disk and ISO images are available from ProfitBricks for immediate use. These can be easily viewed and selected. The following shows you how to get a list of images. This list represents both CDROM images and HDD images.

```go
images := profitbricks.ListImages()
```

This will return a [collection](#Collection) object

## How To: Create Storage Volume

ProfitBricks allows for the creation of multiple storage volumes that can be attached and detached as needed. It is useful to attach an image when creating a storage volume. The storage size is in gigabytes.

```go
volumerequest := profitbricks.Volume{
		Properties: profitbricks.VolumeProperties{
			Size:        1,
			Name:        "Volume Test",
			LicenceType: "LINUX",
			Type:        "HDD",
		},
}

storage := CreateVolume(datacenter.Id, volumerequest)
```

## How To: Update Cores and Memory

ProfitBricks allows users to dynamically update cores, memory, and disk independently of each other. This removes the restriction of needing to upgrade to the next size available size to receive an increase in memory. You can now simply increase the instances memory keeping your costs in-line with your resource needs.

Note: The memory parameter value must be a multiple of 256, e.g. 256, 512, 768, 1024, and so forth.

The following code illustrates how you can update cores and memory:

```go
serverupdaterequest := profitbricks.ServerProperties{
	Cores: 1,
	Ram:   256,
}

resp := PatchServer(datacenter.Id, server.Id, serverupdaterequest)
```

## How To: Attach or Detach Storage Volume

ProfitBricks allows for the creation of multiple storage volumes. You can detach and reattach these on the fly. This allows for various scenarios such as re-attaching a failed OS disk to another server for possible recovery or moving a volume to another location and spinning it up.

The following illustrates how you would attach and detach a volume and CDROM to/from a server:

```go
profitbricks.AttachVolume(datacenter.Id, server.Id, volume.Id)
profitbricks.AttachCdrom(datacenter.Id, server.Id, images.Items[0].Id)

profitbricks.DetachVolume(datacenter.Id, server.Id, volume.Id)
profitbricks.DetachCdrom(datacenter.Id, server.Id, images.Items[0].Id)
```

## How To: List Servers, Volumes, and Data Centers

Go SDK provides standard functions for retrieving a list of volumes, servers, and datacenters.

The following code illustrates how to pull these three list types:

```go
volumes := profitbricks.ListVolumes(datacenter.Id)

servers := profitbricks.ListServers(datacenter.Id)

datacenters := profitbricks.ListDatacenters()
```

## Reference  

-----

## Objects

### Common object properties

| Property Name |  Type | Description|
|---|-----|-----|
| Id | String | Unique identifier of the object|
| Type_ | String | Type of the object as returned from the Cloud API|
| Href | String | URL to the object’s representation|
| Metadata | *Metadata | See [Metadata](#metadata) |
| Headers | *http.Header | Response headers|
| Response | string | Raw JSON response|
| StatusCode | int | Http response status code |

### Metadata

| Property Name |  Type | Description|
|---|-----|-----|
|CreatedDate     |time.Time|The date when the resource was created.|
|CreatedBy       |string|The user who created the resource.|
|Etag            |string|The etag for the request.|
|LastModifiedDate|time.Time|The last time the resource has been modified.|
|LastModifiedBy  |string|The user who last modified the resource.|
|State            |string|*AVAILABLE* There are no pending modification requests for this item; *BUSY* There is at least one modification request pending and all following requests will be queued; *INACTIVE* Resource has been de-provisioned.|


### Resp

| Property Name |  Type | Description|
|---|-----|-----|
| Req        | *http.Request | A Request represents an HTTP request received by a server or to be sent by a client. |
| StatusCode | int           | Request status code |
| Headers    | http.Header   | A Header represents the key-value pairs in an HTTP header. |
| Body       | []byte        | Byte encoded response body |

### Datacenter

| Property Name |  Type | Description|
|---|-----|-----|
| Properties| DatacenterProperties  | See [DatacenterProperties](#datacenterproperties)|
| Entities  | DatacenterEntities    |See [DatacenterEntities](#datacenterentities)|

### DatacenterProperties

| Property Name |  Type | Description|
|---|-----|-----|
|Name        |string| Name of the data center|
|Description |string|Description of the data center|
|Location    |string|Location of the data center|
|Version     |int32 |The version of the data center|

### DatacenterEntities

| Property Name |  Type | Description|
|---|-----|-----|
|Servers       |*Servers      |See [Servers](#servers)|
|Volumes       |*Volumes      |See [Volumes](#volumes)|
|Loadbalancers |*Loadbalancers|See [LoadBalancers](#loadbalancers)|
|Lans          |*Lans         |See [Lans](#lans)|

### Datacenters

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Datacenter | Array of [Datacenters](#datacenter)|

---

### Server

| Property Name |  Type | Description|
|---|-----|-----|
| Id         |string|      |           
| Type_      |string| |
| Properties | ServerProperties |           See [ServerProperties](#serverproperties)|
| Entities   | *ServerEntities | See [ServerEntities](#serverentities)|
	
### ServerProperties

| Property Name |  Type | Description|
|---|-----|-----|
| Name | string | The hostname of the server.|
| Cores | int | The total number of cores for the server.|
| Ram | int | The amount of memory for the server in MB, e.g. 2048. | 
| AvailabilityZone | string |The availability zone in which the server should exist.|
| VmState | string | Status of the virtual Machine.|
| BootCdrom | *ResourceReference | 	Reference to a CD-ROM used for booting. |
| BootVolume | *ResourceReference | Reference to a Volume used for booting.|
| CpuFamily | string | Type of CPU assigned. |

### ServerEntities

| Property Name |  Type | Description|
|---|-----|-----|
|Cdroms  |*Cdroms | See [Cdrom](#cdroms) |
|Volumes |*Volumes | See [Volumes](#volumes)|
|Nics    |*Nics | See [Nics](#nics)|


### Servers

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Server | Array of [Servers](#server)|

---

### Volume

| Property Name |  Type | Description|
|---|-----|-----|
| Properties | VolumeProperties | See [VolumeProperties](#volueproperties) |

### VolumeProperties

| Property Name |  Type | Description|
|---|-----|-----|
| Name                | string  |	The name of the volume. | 
| Type                | string  | The volume type, HDD or SSD. | 
| Size                | int     | The size of the volume in GB. | 
| AvailabilityZone    | string  | The storage availability zone assigned to the volume. | 
| Image               | string  | The image or snapshot ID.| 
| ImagePassword       | string  | Always returns "null".| 
| SshKeys             | []string| Always returns "null".| 
| Bus                 | string  | The bus type: VIRTIO or IDE. Returns "null" if not connected to a server.| 
| LicenceType         | string  | Licence type. | 
| CpuHotPlug          | bool    | This volume is capable of CPU hot plug | 
| CpuHotUnplug        | bool    | This volume is capable of CPU hot unplug | 
| RamHotPlug          | bool    | This volume is capable of memory hot plug | 
| RamHotUnplug        | bool    | This volume is capable of memory hot unplug | 
| NicHotPlug          | bool    | This volume is capable of nic hot plug | 
| NicHotUnplug        | bool    | This volume is capable of nic hot unplug | 
| DiscVirtioHotPlug   | bool    | This volume is capable of VirtIO drive hot plug | 
| DiscVirtioHotUnplug | bool    | This volume is capable of VirtIO drive hot unplug | 
| DiscScsiHotPlug     | bool    | This volume is capable of SCSI drive hot plug | 
| DiscScsiHotUnplug   | bool    | This volume is capable of SCSI drive hot unplug | 
| DeviceNumber        | int64   | The LUN ID of the storage volume. | 

### Volumes

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Volume | Array of [Volumes](#server) |

---

### Nic

| Property Name |  Type | Description|
|---|-----|-----|
|Properties |NicProperties | See [NicProperties](#nicproperties) |
|Entities |*NicEntities | See [NicEntities](#nicentities)|


### NicProperties 

| Property Name |  Type | Description|
|---|-----|-----|
| Name           | string   |The name of the NIC.|
| Mac            | string   |The MAC address of the NIC.|
| Ips            | []string |IPs assigned to the NIC represented as a collection of strings|
| Dhcp           | bool     |Boolean value that indicates if the NIC is using DHCP or not.|
| Lan            | int      |The LAN ID the NIC sits on.|
| FirewallActive | bool     |A true value indicates the firewall is enabled. A false value indicates the firewall is disabled.|
| Nat            | bool     |Boolean value indicating if the private IP address has outbound access to the public internet.|

### NicEntities

| Property Name |  Type | Description|
|---|-----|-----|
|Firewallrules    |*Firewallrules | See [Firewallrules](#firewallrules)|

### Nics 

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Nic | Array of [Nics](#nic) |

---

### FirewallRule

| Property Name |  Type | Description|
|---|-----|-----|
|Properties |FirewallruleProperties | See [FirewallruleProperties](#firewallruleproperties) |


### FirewallruleProperties 

| Property Name |  Type | Description|
|---|-----|-----|
| Name           | string   |The name of the firewall rule.|
| Protocol           | string   |The protocol for the rule: TCP, UDP, ICMP, ANY.|
| SourceMac            | string   |Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address.|
| SourceIp            | string |Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs.|
| TargetIp           | string     |In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs.|
| IcmpCode            | int      |Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen. Value null allows all types.|
| IcmpType | int     |Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes.|
| PortRangeStart            | int      |Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd value null to allow all ports.|
| PortRangeEnd            | int      |Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.|


### FirewallRules 

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []FirewallRule | Array of [FirewallRule](#firewallrule) |

---

### Lan

| Property Name |  Type | Description|
|---|-----|-----|
|Properties |LanProperties | See [LanProperties](#lanproperties) |


### LanProperties 

| Property Name |  Type | Description|
|---|-----|-----|
| Name           | string   |The name of the firewall rule.|
| Public           | interface{} |Boolean indicating if the LAN faces the public Internet or not.| 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.|


### Lans 

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Lan | Array of [Lans](#lan) |

---

### Image

| Property Name |  Type | Description|
|---|-----|-----|
|Properties |ImageProperties | See [ImageProperties](#imageproperties) |

### ImageProperties 

| Property Name |  Type | Description|
|---|-----|-----|
|Name                | string |The name of the image.|
|Description         | string |The description of the image.|
|Location            | string |The image's location.|
|Size                | int    |The size of the image in GB.|
|CpuHotPlug          | bool   |This volume is capable of CPU hot plug (no reboot required)|
|CpuHotUnplug        | bool   |This volume is capable of CPU hot unplug (no reboot required)|
|RamHotPlug          | bool   |This volume is capable of memory hot plug (no reboot required)|
|RamHotUnplug        | bool   |This volume is capable of memory hot unplug (no reboot required)|
|NicHotPlug          | bool   |This volume is capable of nic hot plug (no reboot required)|
|NicHotUnplug        | bool   |This volume is capable of nic hot unplug (no reboot required)|
|DiscVirtioHotPlug   | bool   |This volume is capable of Virt-IO drive hot plug (no reboot required)|
|DiscVirtioHotUnplug | bool   |This volume is capable of Virt-IO drive hot unplug (no reboot required)|
|DiscScsiHotPlug     | bool   |This volume is capable of Scsi drive hot plug (no reboot required)|
|DiscScsiHotUnplug   | bool   |This volume is capable of Scsi drive hot unplug (no reboot required)|
|LicenceType         | string |The image's licence type: LINUX, WINDOWS, WINDOWS2016, OTHER or UNKNOWN.|
|ImageType           | string |The type of image: HDD, CDROM.|
|Public              | bool   |	Indicates if the image is part of the public repository or not.|
	
	
### Images 

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Image | Array of [Images](#image) |

---

### Loadbalancer

| Property Name |  Type | Description|
|---|-----|-----|
|Properties |LoadbalancerProperties | See [LoadbalancerProperties](#loadbalancerproperties) |
|Entities |LoadbalancerEntities | See [LoadbalancerEntities](#loadbalancerentities) |

### LoadbalancerProperties 

| Property Name |  Type | Description|
|---|-----|-----|
|Name               | string |The name of the loadbalancer.|
|Ip                 | string |IPv4 address of the loadbalancer. All attached NICs will inherit this IP.|
|Dhcp               | bool |Indicates if the loadbalancer will reserve an IP using DHCP.|
	
### LoadbalancerEntities 

| Property Name |  Type | Description|
|---|-----|-----|
|Balancednics               | BalancedNics |See [BalancedNics](#nics)|
	
### Loadbalancers 

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Loadbalancer | Array of [Loadbalancers](#loadbalancer) |

---

### Location

| Property Name |  Type | Description|
|---|-----|-----|
|Properties |LocationProperties | See [LocationProperties](#locationproperties) |

### LocationProperties 

| Property Name |  Type | Description|
|---|-----|-----|
|Name               | string |The name of the location.|
|Features           | []string |Features available at this location.|

### Locations

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Location | Array of [Locations](#location) |

---

---

### IpBlock

| Property Name |  Type | Description|
|---|-----|-----|
|Properties |IpBlockProperties | See [IpBlockProperties](#ipblockproperties) |

### IpBlockProperties 

| Property Name |  Type | Description|
|---|-----|-----|
|Ips               | []string |A collection of IPs associated with the IP Block.|
|Location           | string |	Location the IP block resides in.|
|Size           | int | Number of IP addresses in the block.|

### IpBlocks

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []IpBlock | Array of [IpBlock](#ipblock) |

---

### Snapshot

| Property Name |  Type | Description|
|---|-----|-----|
|Properties |SnapshotProperties | See [SnapshotProperties](#snapshotproperties) |

### SnapshotProperties 

| Property Name |  Type | Description|
|---|-----|-----|
| name | string | The name of the snapshot. | 
| description | string | The description of the snapshot. | No |
| cpuHotPlug | bool | This volume is capable of CPU hot plug (no reboot required) | 
| cpuHotUnplug | bool | This volume is capable of CPU hot unplug (no reboot required) | 
| ramHotPlug | bool | This volume is capable of memory hot plug (no reboot required) | 
| ramHotUnplug | bool | This volume is capable of memory hot unplug (no reboot required) | 
| nicHotPlug | bool | This volume is capable of NIC hot plug (no reboot required) | 
| nicHotUnplug | bool | This volume is capable of NIC hot unplug (no reboot required) | 
| discVirtioHotPlug | bool | This volume is capable of Virt-IO drive hot plug (no reboot required) | 
| discVirtioHotUnplug | bool | This volume is capable of Virt-IO drive hot unplug (no reboot required) | 
| discScsiHotPlug | bool | This volume is capable of SCSI drive hot plug (no reboot required) | 
| discScsiHotUnplug | bool | This volume is capable of SCSI drive hot unplug (no reboot required) | 
| licenceType | string | The snapshot's licence type: LINUX, WINDOWS, or UNKNOWN. | 

### Snapshots

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Snapshot | Array of [Snapshot](#snapshot) |

---

### Request

| Property Name |  Type | Description|
|---|-----|-----|
| Properties |RequestProperties | See [RequestProperties](#requestproperties) |
| Metadata | Metadata | See [Metadata](#requestmetadata) | 

### RequestProperties 

| Property Name |  Type | Description|
|---|-----|-----|
| Method  | string | | 
| Headers  | interface{} | | 
| Body     |interface{} | | 
| URL      | string | | 

### Requests

| Property Name |  Type | Description|
|---|-----|-----|
| Items | []Requests | Array of [Requests](#request) |

### RequestStatusMetadata

| Property Name |  Type | Description|
|---|-----|-----|
| Metadata | Metadata | See [RequestStatusMetadata](#requeststatusmetadata) |
 
### RequestStatusMetadata

| Property Name |  Type | Description|
|---|-----|-----|
| Status  |string||
| Message|string ||
| Etag   |string||
| Targets|[]RequestTarget||



---


## Functions

### Virtual Data Centers

Virtual Data Centers are the foundation of the ProfitBricks platform. Virtual Data Centers act as logical containers for all other objects you will be creating, e.g., servers. You can provision as many data centers as you want. Data centers have their own private network and are logically segmented from each other to create isolation.


#### List Data Centers

#### Return Type

[Datacenters](#datacenters)

```
ListDatacenters()
```
---

#### Retrieve a Data Center

The following table describes the request arguments:

| Name | Required | Type | Description |
|---|---|---|---|
| dcid | Yes | string | The ID of the data center.  |


#### Return Type

[Datacenter](#datacenter)

```
GetDatacenter(dcid string)
```
---

#### Create a Data Center

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dc | Datacenter |See [Datacenter](#datacetner) | Yes |

The following table outlines the locations currently supported:

| ID | Country | City |
|---|---|---|
| us/las | United States | Las Vegas |
| de/fra | Germany | Frankfurt |
| de/fkb | Germany | Karlsruhe |

```
CreateDatacenter(dc Datacenter)
```

#### Return Type

[Datacenter](#datacenter)

*NOTES*:
- The value for `name` cannot contain the following characters: (@, /, , |, ‘’, ‘).
- You cannot change a data center's `location` once it has been provisioned.

---

#### Update a Data Center

After retrieving a data center, you can change it's properties:

```
PatchDatacenter(dcid string, obj DatacenterProperties)
```

The following table describes the request arguments:

| Name | Type | Description | Required |
| --- | --- | --- | --- |
| dcid | string | ID of the datacenter | Yes |
| obj | DatacenterProperties | See [DatacenterProperties](#datacenterproperties | Yes |


#### Return Type

[Datacenter](#datacenter)

---

#### Delete a Data Center

This will remove all objects within the data center and remove the data center object itself.

**NOTE**: This is a highly destructive operation which should be used with extreme caution.

```
DeleteDatacenter(dcid string)
```

#### Return Type

[Resp](#resp)

### Server Functions

#### List Servers

You can retrieve a list of all servers within a data center.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |

```
ListServers(dcid string)
```

#### Return Type

[Server](#server)

---

#### Retrieve a Server

Returns information about a server such as its configuration, provisioning status, etc.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |

```
GetServer(dcid, srvid string)
```

#### Return Type

[Servers](#servers)

---

#### Create a Server

Creates a server within an existing data center. You can configure additional properties such as specifying a boot volume and connecting the server to an existing LAN.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| server | Server | See [Server](#server) | Yes |

The following table outlines the various licence types you can define:

| Licence Type | Description |
|---|---|
| WINDOWS | You must specify this if you are using your own, custom Windows image due to Microsoft's licensing terms. |
| LINUX | |
| UNKNOWN | If you are using an image uploaded to your account your OS Type will inherit as UNKNOWN. |

The following table outlines the availability zones currently supported:

| Availability Zone | Description |
|---|---|
| AUTO | Automatically selected zone |
| ZONE_1 | Zone 1 |
| ZONE_2 | Zone 2 |
| ZONE_3 | Zone 3 |

```
CreateServer(dcid string, server Server)
```


#### Return Type

[Server](#server)

---

#### Update a Server

Perform updates to attributes of a server.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvId | string | The ID of the server. | Yes |
| props | ServerProperties | See [ServerProperties](#serverproperties) | No |


```
PatchServer(dcid string, srvid string, props ServerProperties)
```

#### Return Type

[Server](#server)

---

#### Delete a Server

This will remove a server from a data center. NOTE: This will not automatically remove the storage volume(s) attached to a server. A separate API call is required to perform that action.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |


```
DeleteServer(dcid, srvid string)
```

#### Return Type

[Resp](#resp)

---

#### Reboot a Server

This will force a hard reboot of the server. Do not use this method if you want to gracefully reboot the machine. This is the equivalent of powering off the machine and turning it back on.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the data center. | Yes |
| serverId | string | The ID of the server. | Yes |


```
RebootServer(dcid, srvid string)
```

#### Return Type

[Resp](#resp)

---

#### Start a Server

This will start a server. If the server's public IP was deallocated then a new IP will be assigned.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |


```
StartServer(dcid, srvid string)
```

#### Return Type

[Resp](#resp)

---

#### Stop a Server

This will stop a server. The machine will be forcefully powered off, billing will cease, and the public IP, if one is allocated, will be deallocated.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid  | string | The ID of the data center. | Yes |
|  srvid | string | The ID of the server. | Yes |


```
StopServer(dcid, srvid string)
```

#### Return Type

[Resp](#resp)

---

#### Attach a CDROM

This will attach a CDROM to the server.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |
| cdid | string | The ID of a ProfitBricks image of type CDROM. | Yes |


```
AttachCdrom(dcid string, srvid string, cdid string)
```

#### Return Type

[Image](#image)

---

#### Detach a CDROM

This will detach the CDROM from the server. Depending on the volume "hot_unplug" settings, this may result in the server being rebooted.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvId | string | The ID of the server. | Yes |
| cdid | string | The ID of the attached CDROM. | Yes |


```
DetachCdrom(dcid, srvid, cdid string)
```

#### Return Type

[Resp](#resp)

---

#### List attached CDROMs

This will list CDROMs that are attached to the server

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |


```
ListAttachedCdroms(dcid, srvid string)
```

#### Return Type

[Images](#images)

---

#### Get attached CDROM

This will retrieve a CDROM that is attached to the server

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |
| cdid | string | The ID of the attached CDROM. | Yes |


```
GetAttachedCdrom(dcid, srvid, cdid string)
```

#### Return Type

[Image](#image)


#### Delete Volume

This will delete a volume:

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| volid | string | The ID of the volume. | Yes |

```
DeleteVolume(dcid, volid string)
```

#### Return Type

[Resp](#resp)

---

### NIC Functions

#### List NICs

Retrieve a list of LANs within the data center.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |

```
ListNics(dcid, srvid string)
```

#### Return Type

[Nics](#nics)

---

#### Get a NIC

Retrieves the attributes of a given NIC.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvId | string | The ID of the server. | Yes |
| nicid | string | The ID of the NIC. | Yes |

```
GetNic(dcid, srvid, nicid string)
```

#### Return Type

[Nic](#nic)

---

#### Create a NIC

Adds a NIC to the target server.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvId | string | The ID of the server. | Yes |
| nic | Nic | See [Nic](#nic)| Yes |

```
CreateNic(dcid string, srvid string, request Nic)
```

#### Return Type

[Nic](#nic)

---

#### Update a NIC

You can update -- in full or partially -- various attributes on the NIC; however, some restrictions are in place:

The primary address of a NIC connected to a load balancer can only be changed by changing the IP of the load balancer. You can also add additional reserved, public IPs to the NIC.

The user can specify and assign private IPs manually. Valid IP addresses for private networks are 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |
| nicid | string | The ID of the NIC. | Yes |
| obj | NicProperties | See [NicProperties](#nicproperties) | Yes |


```
PatchNic(dcid string, srvid string, nicid string, obj NicProperties)
```

#### Return Type

[Nic](#nic)

---

#### Delete a NIC

This will delete a volume:

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |
| nicid | string | The ID of the NIC. | Yes |

```
DeleteNic(dcid, srvid, nicid string)
```

#### Return Type
 
[Resp](#resp) 

---

### Firewall Rule Functions

#### List Firewall Rules

Retrieves a list of firewall rules associated with a particular NIC.

| Name | Type | Description | Required |
|---|---|---|---|
| dcId | string | The ID of the data center. | Yes |
| serverId | string | The ID of the server. | Yes |
| nicId | string | The ID of the NIC. | Yes |

```
ListFirewallRules(dcId string, serverid string, nicId string)
```

#### Return Type
 
[FirewallRules](#firewallrules) 

---

#### Get a Firewall Rule

Retrieves the attributes of a given firewall rule.

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvid | string | The ID of the server. | Yes |
| nicId | string | The ID of the NIC. | Yes |
| fwId | string | The ID of the firewall rule. | Yes |

```
GetFirewallRule(dcid string, srvid string, nicId string, fwId string)
```

#### Return Type
 
[FirewallRule](#firewallrule) 

---

#### Create a Firewall Rule

This will add a firewall rule to the NIC.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvId | string | The ID of the server. | Yes |
| nicId | string | The ID of the NIC. | Yes |
| fw | FirewallRule | See [FirewallRule)(#firewallrule) | Yes |

```
CreateFirewallRule(dcid string, srvid string, nicId string, fw FirewallRule)
```

#### Return Type
 
[FirewallRule](#firewallrule) 

---

#### Update a Firewall Rule

Perform updates to attributes of a firewall rule.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvId | string | The ID of the server. | Yes |
| nicId | string | The ID of the NIC. | Yes |
| fwId | string | The ID of the firewall rule. | Yes |
| obj | FirewallruleProperties | See [FirewallruleProperties](#firewallruleproperties) | Yes |

```
PatchFirewallRule(dcid string, srvid string, nicId string, fwId string, obj FirewallruleProperties)
```

#### Return Type
 
[FirewallRule](#firewallrule) 

---

#### Delete a Firewall Rule

Removes the specific firewall rule.

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| srvId | string | The ID of the server. | Yes |
| nicId | string | The ID of the NIC. | Yes |
| fwId | string | The ID of the firewall rule. | Yes |

```
DeleteFirewallRule(dcid string, srvid string, nicId string, fwId string)
```

#### Return Type
 
[Resp](#resp) 


### LAN functions

#### List LANs

Retrieve a list of LANs within the data center.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |

```
ListLans(dcid string)
```

#### Return Type
 
[Lans](#lans)
 
---

#### Create a LAN

Creates a LAN within a data center.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the data center. | Yes |
| reques | Lan | See [Lan](#lan) | Yes |

```
CreateLan(dcid string, request Lan)
```

#### Return Type
 
[Lan](#lan)

---

#### Get a LAN

Retrieves the attributes of a given LAN.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| lanId | string | The ID of the LAN. | Yes |

```
GetLan(dcid, lanid string)
```

#### Return Type
 
[Lan](#lan)

---

#### Update a LAN

Perform updates to attributes of a LAN.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the data center. | Yes |
| lanId | string | The ID of the LAN. | Yes |
| obj | LanProperties | See [LanProperties](#lanproperties) | Yes |

```
PatchLan(dcid string, lanid string, obj LanProperties)
```

#### Return Type
 
[Lan](#lan)

---

#### Delete a LAN

Deletes the specified LAN.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| lanId | string | The ID of the LAN. | Yes |


```
DeleteLan(dcid, lanid string)
```

#### Return Type
 
[Resp](#resp)

---

### Images

#### List Images

Retrieve a list of images.

```
ListImages()
```

#### Return Type
 
[Images](#images)

---

#### Get an Image

Retrieves the attributes of a specific image.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| imageId | string | The ID of the image. | Yes |

```
GetImage(imageid string)
```

#### Return Type
 
[Image](#image)

---

### Load Balancers

#### List Load Balancers

Retrieve a list of load balancers within the data center.

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |

```
ListLoadbalancers(dcid string)
```

---

#### Get a Load Balancer

Retrieves the attributes of a given load balancer.

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the data center. | Yes |
| loadBalancerId | string | The ID of the load balancer. | Yes |

```
GetLoadbalancer(dcid, lbalid string)
```

#### Return Type
 
[Loadbalancers](#loadbalancers)

---

#### Create a Load Balancer

Creates a load balancer within the data center. Load balancers can be used for public or private IP traffic.

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the data center. | Yes |
| request | Loadbalancer | See [Loadbalancer](#loadbalancer) | Yes |

```
CreateLoadbalancer(dcid string, request Loadbalancer)
```

#### Return Type
 
[Loadbalancer](#loadbalancer)

---

#### Update a Load Balancer

Perform updates to attributes of a load balancer.

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the data center. | Yes |
| lbalid | string | The ID of the load balancer. | Yes |
| obj | LoadbalancerProperties | See [LoadbalancerProperties](#loadbalancerproperties) | Yes |

```
PatchLoadbalancer(dcid string, lbalid string, obj LoadbalancerProperties)
```

#### Return Type
 
[Loadbalancer](#loadbalancer)

---

#### Delete a Load Balancer

Deletes the specified load balancer.

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the data center. | Yes |
| lbalid | string | The ID of the load balancer. | Yes |


```
DeleteLoadbalancer(dcid, lbalid string)
```

#### Return Type
 
[Resp](#resp)


#### List Load Balanced NICs

This will retrieve a list of NICs associated with the load balancer.

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| lbalid | string | The ID of the load balancer. | Yes |


```
ListBalancedNics(dcid, lbalid string)
```

#### Return Type
 
[Nics](#nics)

---

#### Get a Load Balanced NIC

Retrieves the attributes of a given load balanced NIC.

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| lbalid | string | The ID of the load balancer. | Yes |
| balnicid | string | The ID of the Nic | Yes |

```
GetBalancedNic(dcid, lbalid, balnicid string)
```

#### Return Type
 
[Nic](#nic)

---

#### Associate NIC to a Load Balancer

This will associate a NIC to a Load Balancer, enabling the NIC to participate in load-balancing.

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| lbalid | string | The ID of the load balancer. | Yes |
| nicId | string | The ID of the load balancer. | Yes |


```
AssociateNic(dcid string, lbalid string, nicid string)
```

#### Return Type
 
[Nic](#nic)

---

#### Remove a NIC Association

Removes the association of a NIC with a load balancer.

| Name | Type | Description | Required |
|---|---|---|---|
| dcid | string | The ID of the data center. | Yes |
| lbalid | string | The ID of the load balancer. | Yes |
| balnicid | string | The ID of the load balancer. | Yes |


```
DeleteBalancedNic(dcid, lbalid, balnicid string)
```

#### Return Type
 
[Resp](#resp)

---

### Location Functions

#### List Locations

Locations represent regions where you can provision your Virtual Data Centers.

```
ListLocations()
```

#### Return Type 

[Locations](#locations)

---

#### Get a Location

Retrieves the attributes of a given location.

The following table describes the request arguments:

| Name | Type | Description | Required |
| --- | --- | --- | --- |
| locid | string | The unique identifier consisting of country/city. | Yes |

```
GetLocation(locid string)
```

#### Return Type 

[Location](#location)

---

### IP Blocks

#### List IP Blocks

Retrieve a list of IP Blocks.

```
ListIpBlocks()
```

#### Return Type 

[IpBlocks](#ipblocks)

---

#### Get an IP Block

Retrieves the attributes of a specific IP Block.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| ipblockid | string | The ID of the IP block. | Yes |

```
GetIpBlock(ipblockid string)
```

#### Return Type 

[IpBlock](#ipblock)

---

#### Create an IP Block

Creates an IP block.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| request | IpBlock | See [IpBlock](#ipblock) | Yes |


```
ReserveIpBlock(request IpBlock)
```

#### Return Type 

[IpBlock](#ipblock)

---

#### Delete an IP Block

Deletes the specified IP Block.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| ipblockid | string | The ID of the IP block. | Yes |

```
ReleaseIpBlock(ipblockid string)
```

#### Return Type 

[Resp](#resp)

---


### Snapshot Functions

#### List Snapshots

You can retrieve a list of all snapshots.

```
ListSnapshots()
```

#### Return Type 

[Snapshots](#snapshots)

---

#### Get a Snapshot

Retrieves the attributes of a specific snapshot.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| snapshotId | string | The ID of the snapshot. | Yes |

```
GetSnapshot(snapshotId string)
```

#### Return Type 

[Snapshot](#snapshot)

---

#### Update a Snapshot

Perform updates to attributes of a snapshot.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the snapshot. | Yes |
| request | SnapshotProperties | See [SnapshotProperties](#snapshotproperties) | Yes |

```
UpdateSnapshot(snapshotId string, request SnapshotProperties)
```

#### Return Type 

[Snapshot](#snapshot)

---

#### Create a Volume Snapshot

Creates a snapshot of a volume within the data center. You can use a snapshot to create a new storage volume or to restore a storage volume.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the datacenter. | Yes |
| volumeId | string | The ID of the volume. | Yes |
| name | string | The name of the snapshot. | No |

```
CreateSnapshot(dcid string, volid string, name string)
```

#### Return Type 

[Snapshot](#snapshot)

---


#### Restore a Volume Snapshot

This will restore a snapshot onto a volume. A snapshot is created as just another image that can be used to create new volumes or to restore an existing volume.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| dataCenterId | string | The ID of the datacenter. | Yes |
| volumeId | string | The ID of the volume. | Yes |
| snapshotId | string | The ID of the snapshot. | Yes |


```
RestoreSnapshot(dcid string, volid string, snapshotId string)
```

#### Return Type 

[Resp](#resp)

---

#### Delete a Snapshot

Deletes the specified snapshot.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| snapshotId | string | The ID of the snapshot. | Yes |


```
DeleteSnapshot(snapshotId string)
```

#### Return Type 

[Resp](#resp)

---

### Request Functions

#### Get a Request status

Retrieves the status of a specific request.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| path | string | The ID of the request. Retrieved from response header `location` | Yes |

```
GetRequestStatus(path string)
```

#### Return Type

[RequestStatus](#requeststatus)

#### Get a Request

Retrieves the attributes of a specific request.

The following table describes the request arguments:

| Name | Type | Description | Required |
|---|---|---|---|
| requestId | string | The ID of the request. | Yes |

```
GetRequest(req_id string)
```

#### Return Type

[Request](#request)

#### List Requests

Retrieves list of requests.

The following table describes the request arguments:

```
ListRequests()
```

#### Return Type

[Requests](#requests)

---

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/profitbricks/profitbricks-sdk-go"
)

func main() {

	//Sets username and password
	profitbricks.SetAuth("username", "password")
	//Sets depth.
	profitbricks.SetDepth("5")

	dcrequest := profitbricks.Datacenter{
		Properties: profitbricks.DatacenterProperties{
			Name:        "example.go3",
			Description: "description",
			Location:    "us/lasdev",
		},
	}

	datacenter := profitbricks.CreateDatacenter(dcrequest)

	serverrequest := profitbricks.Server{
		Properties: profitbricks.ServerProperties{
			Name:  "go01",
			Ram:   1024,
			Cores: 2,
		},
	}
	server := profitbricks.CreateServer(datacenter.Id, serverrequest)

	volumerequest := profitbricks.Volume{
		Properties: profitbricks.VolumeProperties{
			Size:        1,
			Name:        "Volume Test",
			LicenceType: "LINUX",
			Type:        "HDD",
		},
	}

	storage := profitbricks.CreateVolume(datacenter.Id, volumerequest)

	serverupdaterequest := profitbricks.ServerProperties{
		Name:  "go01renamed",
		Cores: 1,
		Ram:   256,
	}

	profitbricks.PatchServer(datacenter.Id, server.Id, serverupdaterequest)
	//It takes a moment for a volume to be provisioned so we wait.
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
```

# Support
You are welcome to contact us with questions or comments at [ProfitBricks DevOps Central](https://devops.profitbricks.com/). Please report any issues via [GitHub's issue tracker](https://github.com/profitbricks/profitbricks-sdk-go/issues).