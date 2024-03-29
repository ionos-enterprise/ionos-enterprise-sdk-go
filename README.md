[![No Maintenance Intended](http://unmaintained.tech/badge.svg)](http://unmaintained.tech/)
# DEPRECATED

This is no longer supported, please consider using [IONOS Cloud SDK Go](https://github.com/ionos-cloud/sdk-go) instead.

# Go SDK

Version: profitbricks-sdk-go **5.0**

The ProfitBricks Client Library for [Go](https://www.golang.org/) provides you with access to the ProfitBricks Cloud API. It is designed for developers who are building applications in Go.

This guide will walk you through getting setup with the library and performing various actions against the API.

## Table of Contents

- [Description](#description)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Authentication](#authentication)
  - [Error Handling](#error-handling)
- [Reference](#reference)
  - [Client Configuration](#client-configuration)
    - [SetDepth](#setdepth)
    - [SetURL](#seturl)
    - [SetUserAgent](#setuseragent)
  - [Data Centers](#data-centers)
    - [List Data Centers](#list-data-centers)
    - [Retrieve a Data Center](#retrieve-a-data-center)
    - [Create a Data Center](#create-a-data-center)
    - [Update a Data Center](#update-a-data-center)
    - [Delete a Data Center](#delete-a-data-center)
  - [Locations](#locations)
    - [List Locations](#list-locations)
    - [Get a Location](#get-a-location)
  - [Servers](#servers)
    - [List Servers](#list-servers)
    - [Retrieve a Server](#retrieve-a-server)
    - [Create a Server](#create-a-server)
    - [Update a Server](#update-a-server)
    - [Delete a Server](#delete-a-server)
    - [List Attached Volumes](#list-attached-volumes)
    - [Attach a Volume](#attach-a-volume)
    - [Retrieve an Attached Volume](#retrieve-an-attached-volume)
    - [Detach a Volume](#detach-a-volume)
    - [List Attached CD-ROMs](#list-attached-cd-roms)
    - [Attach a CD-ROM](#attach-a-cd-rom)
    - [Retrieve an Attached CD-ROM](#retrieve-an-attached-cd-rom)
    - [Detach a CD-ROM](#detach-a-cd-rom)
    - [Reboot a Server](#reboot-a-server)
    - [Start a Server](#start-a-server)
    - [Stop a Server](#stop-a-server)
  - [Images](#images)
    - [List Images](#list-images)
    - [Get an Image](#get-an-image)
    - [Update an Image](#update-an-image)
    - [Delete an Image](#delete-an-image)
  - [Volumes](#volumes)
    - [List Volumes](#list-volumes)
    - [Get a Volume](#get-a-volume)
    - [Create a Volume](#create-a-volume)
    - [Update a Volume](#update-a-volume)
    - [Delete a Volume](#delete-a-volume)
    - [Create a Volume Snapshot](#create-a-volume-snapshot)
    - [Restore a Volume Snapshot](#restore-a-volume-snapshot)
  - [Snapshots](#snapshots)
    - [List Snapshots](#list-snapshots)
    - [Get a Snapshot](#get-a-snapshot)
    - [Update a Snapshot](#update-a-snapshot)
    - [Delete a Snapshot](#delete-a-snapshot)
  - [IP Blocks](#ip-blocks)
    - [List IP Blocks](#list-ip-blocks)
    - [Get an IP Block](#get-an-ip-block)
    - [Create an IP Block](#create-an-ip-block)
    - [Delete an IP Block](#delete-an-ip-block)
  - [LANs](#lans)
    - [List LANs](#list-lans)
    - [Create a LAN](#create-a-lan)
    - [Get a LAN](#get-a-lan)
    - [Update a LAN](#update-a-lan)
    - [Delete a LAN](#delete-a-lan)
  - [Network Interfaces (NICs)](#network-interfaces-nics)
    - [List NICs](#list-nics)
    - [Get a NIC](#get-a-nic)
    - [Create a NIC](#create-a-nic)
    - [Update a NIC](#update-a-nic)
    - [Delete a NIC](#delete-a-nic)
  - [Firewall Rules](#firewall-rules)
    - [List Firewall Rules](#list-firewall-rules)
    - [Get a Firewall Rule](#get-a-firewall-rule)
    - [Create a Firewall Rule](#create-a-firewall-rule)
    - [Update a Firewall Rule](#update-a-firewall-rule)
    - [Delete a Firewall Rule](#delete-a-firewall-rule)
  - [Load Balancers](#load-balancers)
    - [List Load Balancers](#list-load-balancers)
    - [Get a Load Balancer](#get-a-load-balancer)
    - [Create a Load Balancer](#create-a-load-balancer)
    - [Update a Load Balancer](#update-a-load-balancer)
    - [List Load Balanced NICs](#list-load-balanced-nics)
    - [Get a Load Balanced NIC](#get-a-load-balanced-nic)
    - [Associate NIC to a Load Balancer](#associate-nic-to-a-load-balancer)
    - [Remove a NIC Association](#remove-a-nic-association)
  - [Requests](#requests)
    - [List Requests](#list-requests)
    - [Get a Request](#get-a-request)
    - [Get a Request Status](#get-a-request-status)
  - [Contract Resources](#contract-resources)
  - [Users Management](#users-management)
    - [List Groups](#list-groups)
    - [Retrieve a Group](#retrieve-a-group)
    - [Create a Group](#create-a-group)
    - [Update a Group](#update-a-group)
    - [Delete a Group](#delete-a-group)
    - [List Shares](#list-shares)
    - [Retrieve a Share](#retrieve-a-share)
    - [Add a Share](#add-a-share)
    - [Update a Share](#update-a-share)
    - [Delete a Share](#delete-a-share)
    - [List Users in a Group](#list-users-in-a-group)
    - [Add User to Group](#add-user-to-group)
    - [Remove User from a Group](#remove-user-from-a-group)
    - [List Users](#list-users)
    - [Retrieve a User](#retrieve-a-user)
    - [Create a User](#create-a-user)
    - [Update a User](#update-a-user)
    - [Delete a User](#delete-a-user)
    - [List Resources](#list-resources)
    - [List All Resources of a Type](#list-all-resources-of-a-type)
    - [List a specific Resource Type](#list-a-specific-resource-type)
  - [Managed Kubernetes](#managed-kubernetes)
    - [Create Cluster](#create-cluster)
    - [Read Cluster](#read-cluster)
    - [List Clusters](#list-clusters)
    - [Update Cluster](#update-cluster)
    - [Delete Cluster](#delete-cluster)
    - [Read Kubeconfig](#read-kubeconfig)
    - [Create Node Pool](#create-node-pool)
    - [Read Node Pool](#read-node-pool)
    - [List Node Pools](#list-node-pools)
    - [Update Node Pool](#update-node-pool)
    - [Delete Node Pool](#delete-node-pool)
  - [Backup Units](#backup-units)
    - [Create Backup Unit](#create-backup-unit)
    - [Read Backup Unit](#read-backup-unit)
    - [List Backup Units](#list-backup-units)
    - [Update Backup Unit](#update-backup-unit)
    - [Delete Backup Unit](#delete-backup-unit)
    - [Get Backup Unit SSO URL](#get-backup-unit-sso-url)
  - [User S3 Keys](#user-s3-keys)
    - [Create S3 Key](#create-s3-key)
    - [Read S3 Key](#create-s3-key)
    - [List S3 Keys](#list-s3-keys)
    - [Update S3 Key](#update-s3-key)
    - [Delete S3 Key](#delete-s3-key)
  - [Private Cross Connects](#private-cross-connects)
    - [Create Private Cross Connect](#create-private-cross-connect)
    - [Read Private Cross Connect](#read-private-cross-connect)
    - [List Private Cross Connects](#list-private-cross-connects)
    - [Update Private Cross Connect](#update-private-cross-connect)
    - [Delete Private Cross Connect](#delete-private-cross-connect)
- [Example](#example)
- [Support](#support)
- [Testing](#testing)
- [Contributing](#contributing)

# Description

The Go SDK wraps the latest version of the ProfitBricks Cloud API. All API operations are performed over SSL and authenticated using your ProfitBricks portal credentials. The API can be accessed within an instance running in ProfitBricks or directly over the Internet from any application that can send an HTTPS request and receive an HTTPS response.

## Getting Started

Before you begin you will need to have [signed-up](https://www.profitbricks.com/signup) for a ProfitBricks account. The credentials you setup during sign-up will be used to authenticate against the API.

#### Installation

Install the Go language from from the official [Go installation](https://golang.org/doc/install) guide or using your Linux distribution package management system.

The `GOPATH` environment variable specifies the location of your Go workspace. It is likely the only environment variable you will need to set when developing Go code. This is an example of pointing to a workspace configured under your home directory:

```
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

The following `go` command will download `profitbricks-sdk-go` to your configured `GOPATH`:

```go
go get "github.com/profitbricks/profitbricks-sdk-go/v5"
```

The source code of the package will be located here:

    $GOPATH/src/github.com/profitbricks/profitbricks-sdk-go

Create main package file _example.go_:

```go
package main

import (
	"fmt"
)

func main() {
}
```

Include the ProfitBricks SDK for Go under the list of imports.

```go
import(
	"fmt"

	"github.com/profitbricks/profitbricks-sdk-go/v5"
)
```

#### Authentication

Add your ProfitBricks credentials to the client connection.

```go
client := profitbricks.NewClient("username", "password")
```

It might be necessary to accept credentials through environment variables in a containerized environment.

```go
import (
	"fmt"
	"os"

	"github.com/profitbricks/profitbricks-sdk-go/v5"
)

func main() {
	client := profitbricks.NewClient(
		os.Getenv("PROFITBRICKS_USERNAME"),
		os.Getenv("PROFITBRICKS_PASSWORD"),
	)
```

**Caution**: You will want to ensure you follow security best practices when using credentials within your code or stored in a file.

#### Error Handling

The SDK will raise custom exceptions when the Cloud API returns an error. There are four response types:

| HTTP Code | Description                                         |
| --------- | --------------------------------------------------- |
| 400       | Bad request.                                        |
| 401       | The supplied user credentials are invalid.          |
| 404       | The requested resource cannot be found.             |
| 422       | Invalid parameters have been passed onto Cloud API. |
| 429       | The Cloud API rate limit has been exceeded.         |

Therefore each client function returns two parameters expected response and an error. For example:

```go
datacenters, err := ListDatacenters()
```

In the example the first parameter is a collection of virtual data centers and an error.

## Reference

This section provides details on all the available operations and the arguments they accept. Brief code snippets demonstrating usage are also included.

## Client Configuration

Clients can be created using either username / password or a bearer token for client authentification. A bearer token can be obtained from the auth api, e.g. by using a curl request

```bash
curl --user '<username:password>' https://api.profitbricks.com/auth/v1/tokens/generate
```

```go
client := profitbricks.NewClient("username", "password")
token_client := profitbricks.NewClientbyToken("token")
```

#### SetDepth

Many of the _List_ or _Get_ operations will accept an optional _depth_ argument. Setting this to a value between 0 and 5 affects the amount of data that is returned. The details returned vary depending on the resource being queried, however, it generally follows this pattern.

| Depth | Description                                                                         |
| :---: | ----------------------------------------------------------------------------------- |
|   0   | Only direct properties are included. Children are not included.                     |
|   1   | Direct properties and children's references are returned.                           |
|   2   | Direct properties and children's properties are returned.                           |
|   3   | Direct properties, children's properties, and descendant's references are returned. |
|   4   | Direct properties, children's properties, and descendant's properties are returned. |
|   5   | Returns all available properties.                                                   |

This SDK sets the _Depth=5_ by default as that works well in the majority of cases. You may find that setting _Depth_ to a lower or higher value could simplify a later operation by reducing or increasing the data available in the response object.

```go
client := profitbricks.NewClient("username", "password")
client.SetDepth(3)
```

#### SetURL

Set Cloud API URL:

```go
client.SetURL("https://api.profitbricks.com/cloudapi/v5")
```

#### SetUserAgent

Set "User-Agent" HTTP header for all Cloud API calls:

```go
client.SetUserAgent("myproject/1.0.0")
```

## Data Centers

Virtual Data Centers (VDCs) are the foundation of the ProfitBricks platform. VDCs act as logical containers for all other objects you will be creating, e.g., servers. You can provision as many VDCs as you want. VDCs have their own private network and are logically segmented from each other to create isolation.

#### List Data Centers

This operation will list all currently provisioned VDCs that your account credentials provide access to.

There are no request arguments that need to be supplied.

Call `ListDatacenters`:

    ListDatacenters()

---

#### Retrieve a Data Center

Use this to retrieve details about a specific VDC.

The following table describes the request arguments:

| Name | Required | Type   | Description                |
| ---- | -------- | ------ | -------------------------- |
| dcid | Yes      | string | The ID of the data center. |

Pass the arguments to `GetDatacenter`:

    GetDatacenter(dcid string)

---

#### Create a Data Center

Use this operation to create a new VDC. You can create a "simple" VDC by supplying just the required _Name_ and _Location_ arguments. This operation also has the capability of provisioning a "complex" VDC by supplying additional arguments for servers, volumes, LANs, and/or load balancers.

The following table describes the request arguments:

| Name       | Required | Type   | Description                                                                          |
| ---------- | :------: | ------ | ------------------------------------------------------------------------------------ |
| datacenter | **yes**  | object | A [Datacenter object](#datacenter-resource-object) describing the VDC being created. |

Build the `Datacenter` resource object:

    var obj = Datacenter{
    	Properties: DatacenterProperties{
    	Name:        "GO SDK Test",
    	Description: "GO SDK test datacenter",
    	Location:    location,
    	},
    }

Pass the object to `CreateDatacenter`:

    CreateDatacenter(obj)

##### Datacenter Resource Object

| Name          | Required | Type   | Description                                                                                                    |
| ------------- | :------: | ------ | -------------------------------------------------------------------------------------------------------------- |
| Name          | **yes**  | string | The name of the VDC.                                                                                           |
| Location      | **yes**  | string | The physical ProfitBricks location where the VDC will be created (Please see [Locations](#locations) section). |
| Description   |    no    | string | A description for the VDC, e.g. staging, production.                                                           |
| Servers       |    no    | list   | A list of one or more [Server objects](#server-resource-object) to be created.                                 |
| Volumes       |    no    | list   | A list of one or more [Volume objects](#volume-resource-object) to be created.                                 |
| Lans          |    no    | list   | A list of one or more [LAN objects](#lan-resource-object) to be created.                                       |
| Loadbalancers |    no    | list   | A list of one or more [LoadBalancer objects](#load-balancer-resource-object) to be created.                    |

**NOTES**:

- The value for `Name` cannot contain the following characters: (@, /, , |, ‘’, ‘).
- You cannot change the VDC `Location` once it has been provisioned.

---

#### Update a Data Center

After retrieving a VDC, either by ID or as a create response object, you can change its properties by calling the `update_datacenter` method. Some arguments may not be changed using `update_datacenter`.

The following table describes the available request arguments:

| Name        | Required | Type   | Description                     |
| ----------- | :------: | ------ | ------------------------------- |
| dcid        | **yes**  | string | The ID of the VDC.              |
| Name        |    no    | string | The new name of the VDC.        |
| Description |    no    | string | The new description of the VDC. |

Build the `DatacenterProperties` resource object:

    var obj = DatacenterProperties{Name: "new Name",Description: "new desc"}

Pass the arguments to `UpdateDatacenter`:

UpdateDatacenter(dcid string, obj DatacenterProperties)

---

#### Delete a Data Center

This will remove all objects within the VDC and remove the VDC object itself.

**NOTE**: This is a highly destructive operation which should be used with extreme caution!

The following table describes the available request arguments:

| Name | Required | Type   | Description                                |
| ---- | :------: | ------ | ------------------------------------------ |
| dcid | **yes**  | string | The ID of the VDC that you want to delete. |

Pass the argument to `DeleteDatacenter`:

    DeleteDatacenter(dcid)

---

## Locations

Locations are the physical ProfitBricks data centers where you can provision your VDCs.

The following table outlines the currently supported locations:

| Value  | Country        | City      | Supported **CPUFamily** values |
| ------ | -------------- | --------- | ------------------------------ |
| us/las | United States  | Las Vegas | AMD_OPTERON, INTEL_XEON        |
| us/ewr | United States  | Newark    | AMD_OPTERON, INTEL_XEON        |
| de/fra | Germany        | Frankfurt | AMD_OPTERON, INTEL_XEON        |
| de/fkb | Germany        | Karlsruhe | AMD_OPTERON, INTEL_XEON        |
| de/txl | Germany        | Berlin    | INTEL_SKYLAKE                  |
| gb/lhr | United Kingdom | London    | INTEL_SKYLAKE                  |

#### List Locations

The `ListLocations` operation will return the list of currently available locations.

There are no request arguments to supply.

    ListLocations()

---

#### Get a Location

Retrieves the attributes of a specific location.

The following table describes the request arguments:

| Name       | Required | Type   | Description                        |
| ---------- | :------: | ------ | ---------------------------------- |
| locationid | **yes**  | string | The ID consisting of country/city. |

Pass the argument to `GetLocation`:

    GetLocation("us/las")

---

#### Get a Regional Location

Retrieves the locations available in a specific region.

The following table describes the request arguments:

| Name     | Required | Type   | Description                        |
| -------- | :------: | ------ | ---------------------------------- |
| regionid | **yes**  | string | The ID consisting of country/city. |

Pass the argument to `GetRegionalLocations`:

    GetRegionalLocations("us")

---

## Servers

#### List Servers

You can retrieve a list of all the servers provisioned inside a specific VDC.

The following table describes the request arguments:

| Name | Required | Type   | Description        |
| ---- | :------: | ------ | ------------------ |
| dcid | **yes**  | string | The ID of the VDC. |

Pass the arguments to `ListServers`:

    ListServers(dcid)

---

#### Retrieve a Server

Returns information about a specific server such as its configuration, provisioning status, etc.

The following table describes the request arguments:

| Name     | Required | Type   | Description           |
| -------- | :------: | ------ | --------------------- |
| dcId     | **yes**  | string | The ID of the VDC.    |
| serverId | **yes**  | string | The ID of the server. |

Pass the arguments to `GetServer`:

    GetServer(dcId, serverId)

---

#### Create a Server

Creates a server within an existing VDC. You can configure additional properties such as specifying a boot volume and connecting the server to a LAN.

The following table describes the request arguments:

| Name         | Required | Type   | Description                                                                     |
| ------------ | :------: | ------ | ------------------------------------------------------------------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.                                                              |
| server       | **yes**  | object | A [Server object](#server-resource-object) describing the server being created. |

Build a [Server](#server-resource-object) object:

    var server = Server{
    	Properties: ServerProperties{
    		Name:             "GO SDK Test",
    		RAM:              1024,
    		Cores:            1,
    		AvailabilityZone: "ZONE_1",
    		CPUFamily:        "INTEL_XEON",
    	},
    }

Pass the object and other arguments to `CreateServer`:

    CreateServer(datacenterId, server)

#### Server Resource Object

| Name             | Required | Type   | Description                                                                                                                                                                                                      |
| ---------------- | :------: | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Name             | **yes**  | string | The name of the server.                                                                                                                                                                                          |
| Cores            | **yes**  | int    | The total number of cores for the server.                                                                                                                                                                        |
| RAM              | **yes**  | int    | The amount of memory for the server in MB, e.g. 2048. Size must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set `RamHotPlug` to _true_ then you must use a minimum of 1024 MB. |
| AvailabilityZone |    no    | string | The availability zone in which the server should exist. Possible Values are: AUTO, ZONE_1, ZONE_2                                                                                                                |
| CPUFamily        |    no    | string | Sets the CPU type. For all supported values (depending on location), please see the [Locations](#locations) section.                                                                                             |
| BootVolume       |    no    | string | A volume ID that the server will boot from. If not _nil_ then `BootCdrom` has to be _nil_.                                                                                                                       |
| BootCdrom        |    no    | string | A CD-ROM image ID used for booting. If not _nil_ then `BootVolume` has to be _nil_.                                                                                                                              |
| Cdroms           |    no    | list   | A list of existing volume IDs that you want to connect to the server.                                                                                                                                            |
| Volumes          |    no    | list   | One or more [Volume objects](#volume-resource-object) that you want to create and attach to the server.                                                                                                          |
| Nics             |    no    | list   | One or more [NIC objects](#nic-resource-object) that you wish to create at the time the server is provisioned.                                                                                                   |

---

#### Update a Server

Perform updates to the attributes of a server.

The following table describes the request arguments:

| Name             | Required | Type   | Description                                                                                                            |
| ---------------- | :------: | ------ | ---------------------------------------------------------------------------------------------------------------------- |
| datacenterId     | **yes**  | string | The ID of the VDC.                                                                                                     |
| serverId         | **yes**  | string | The ID of the server.                                                                                                  |
| Name             |    no    | string | The name of the server.                                                                                                |
| Cores            |    no    | int    | The number of cores for the server.                                                                                    |
| RAM              |    no    | int    | The amount of memory in the server.                                                                                    |
| AvailabilityZone |    no    | string | The new availability zone for the server. Possible Values are: AUTO, ZONE_1, ZONE_2                                    |
| CPUFamily        |    no    | string | Sets the CPU type. For all supported values (depending on location), please check the [Locations](#locations) section. |
| BootVolume       |    no    | string | A volume ID used for booting. If not _nil_ then `BootCdrom` has to be _nil_.                                           |
| BootCdrom        |    no    | string | A CD-ROM image ID used for booting. If not _nil_ then `BootVolume` has to be _nil_.                                    |

Build a [ServerProperties](#serverproperties) object:

    var server = ServerProperties{
    	Name: "GO SDK Test RENAME",
    }

Pass the arguments to `update_server`:

    UpdateServer(datacenterId, serverId, server)

---

#### Delete a Server

This will remove a server from a VDC. **NOTE**: This will not automatically remove the storage volume(s) attached to a server. A separate operation is required to delete a storage volume.

The following table describes the request arguments:

| Name         | Required | Type   | Description                                |
| ------------ | :------: | ------ | ------------------------------------------ |
| datacenterId | **yes**  | string | The ID of the VDC.                         |
| serverId     | **yes**  | string | The ID of the server that will be deleted. |

Pass the arguments to `delete_server`:

     DeleteServer(datacenterId, serverId)

---

#### List Attached Volumes

Retrieves a list of volumes attached to the server.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |

Pass the arguments to `ListAttachedVolumes`:

    ListAttachedVolumes(datacenterId, serverId)

---

#### Attach a Volume

This will attach a pre-existing storage volume to the server.

The following table describes the request arguments:

| Name         | Required | Type   | Description                 |
| ------------ | :------: | ------ | --------------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.          |
| serverId     | **yes**  | string | The ID of the server.       |
| volumeId     | **yes**  | string | The ID of a storage volume. |

Pass the arguments to `AttachVolume`:

AttachVolume(datacenterId, serverId, volumeId)

---

#### Retrieve an Attached Volume

This will retrieve the properties of an attached volume.

The following table describes the request arguments:

| Name         | Required | Type   | Description                    |
| ------------ | :------: | ------ | ------------------------------ |
| datacenterId | **yes**  | string | The ID of the VDC.             |
| serverId     | **yes**  | string | The ID of the server.          |
| volumeId     | **yes**  | string | The ID of the attached volume. |

Pass the arguments to `get_attached_volume`:

    GetAttachedVolume(srv_dc_id, srv_srvid, srv_vol)

---

#### Detach a Volume

This will detach the volume from the server. Depending on the volume `hot_unplug` settings, this may result in the server being rebooted. If `disc_virtio_hot_unplug` has been set to _true_, then a reboot should not be required.

This will **NOT** delete the volume from your VDC. You will need to make a separate request to delete a volume.

The following table describes the request arguments:

| Name         | Required | Type   | Description                    |
| ------------ | :------: | ------ | ------------------------------ |
| datacenterId | **yes**  | string | The ID of the VDC.             |
| serverId     | **yes**  | string | The ID of the server.          |
| volumeId     | **yes**  | string | The ID of the attached volume. |

Pass the arguments to `detach_volume`:

     DetachVolume(datacenterId, serverId, volumeId)

---

#### List Attached CD-ROMs

Retrieves a list of CD-ROMs attached to a server.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |

Pass the arguments to `ListAttachedCdroms`:

ListAttachedCdroms(srv_dc_id, srv_srvid)

---

#### Attach a CD-ROM

You can attach a CD-ROM to an existing server.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |
| cdromId      | **yes**  | string | The ID of a CD-ROM.   |

Pass the arguments to `attach_cdrom`:

    AttachCdrom(datacenterId, serverId, cdromId)

---

#### Retrieve an Attached CD-ROM

You can retrieve a specific CD-ROM attached to the server.

The following table describes the request arguments:

| Name         | Required | Type   | Description                    |
| ------------ | :------: | ------ | ------------------------------ |
| datacenterId | **yes**  | string | The ID of the VDC.             |
| serverId     | **yes**  | string | The ID of the server.          |
| cdromId      | **yes**  | string | The ID of the attached CD-ROM. |

Pass the arguments to `GetAttachedCdrom`:

GetAttachedCdrom(datacenterId, serverId, cdromId)

---

#### Detach a CD-ROM

This will detach a CD-ROM from the server.

The following table describes the request arguments:

| Name         | Required | Type   | Description                    |
| ------------ | :------: | ------ | ------------------------------ |
| datacenterId | **yes**  | string | The ID of the VDC.             |
| serverId     | **yes**  | string | The ID of the server.          |
| cdromId      | **yes**  | string | The ID of the attached CD-ROM. |

Pass the arguments to `DetachCdrom`:

    DetachCdrom(datacenterId, serverId, cdromId)

---

#### Reboot a Server

This will force a hard reboot of the server. Do not use this method if you want to gracefully reboot the machine. This is the equivalent of powering off the machine and turning it back on.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |

Pass the arguments to `RebootServer`:

    RebootServer(datacenterId, serverId)

---

#### Start a Server

This will start a server. If a DHCP assigned public IP was deallocated when the server was stopped, then a new IP will be assigned.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |

Pass the arguments to `StartServer`:

    StartServer(datacenterId, serverId)

---

#### Stop a Server

This will stop a server. The machine will be forcefully powered off, billing will cease, and the public IP, if one is allocated, will be deallocated.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |

Pass the arguments to `StopServer`:

    StopServer(datacenterId, serverId)

---

## Images

#### List Images

Retrieve a list of images.

Just call the `ListImages`:

    ListImages()

---

#### Get an Image

Retrieves the attributes of a specific image.

The following table describes the request arguments:

| Name  | Required | Type   | Description          |
| ----- | :------: | ------ | -------------------- |
| imgId | **yes**  | string | The ID of the image. |

Pass the arguments to `GetImage`:

    GetImage(imgid)

---

## Volumes

#### List Volumes

Retrieve a list of volumes within the VDC. If you want to retrieve a list of volumes attached to a server please see the [List Attached Volumes](#list-attached-volumes) entry in the Server section for details.

The following table describes the request arguments:

| Name         | Required | Type   | Description        |
| ------------ | :------: | ------ | ------------------ |
| datacenterId | **yes**  | string | The ID of the VDC. |

Pass the arguments to `ListVolumes`:

    ListVolumes(datacenterId)

---

#### Get a Volume

Retrieves the attributes of a given volume.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| volumeId     | **yes**  | string | The ID of the volume. |

Pass the arguments to `GetVolume`:

    GetVolume(datacenterId, volumeId)

---

#### Create a Volume

Creates a volume within the VDC. This will NOT attach the volume to a server. Please see the [Attach a Volume](#attach-a-volume) entry in the Server section for details on how to attach storage volumes.

The following table describes the request arguments:

| Name          | Required | Type   | Description                                                    |
| ------------- | :------: | ------ | -------------------------------------------------------------- |
| datacenter_id | **yes**  | string | The ID of the VDC.                                             |
| volume        | **yes**  | object | A [Volume object](#volume-resource-object) you wish to create. |

Build the `Volume` resource object:

    var request = Volume{
    	Properties: VolumeProperties{
    		Size:             2,
    		Name:             "GO SDK Test",
    		ImageAlias:       "ubuntu:latest",
    		Bus:              "VIRTIO",
    		SSHKeys:          []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCoLVLHON4BSK3D8L4H79aFo..."},
    		Type:             "HDD",
    		ImagePassword:    "test1234",
    		AvailabilityZone: "ZONE_3",
    	},
    }

Pass the object and arguments to `CreateVolume`:

    CreateVolume(dcID, request)

#### Volume Resource Object

| Name             | Required | Type   | Description                                                                                                                                                                                                                                                                                                      |
| ---------------- | :------: | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Name             |    no    | string | The name of the volume.                                                                                                                                                                                                                                                                                          |
| Size             | **yes**  | int    | The size of the volume in GB.                                                                                                                                                                                                                                                                                    |
| Bus              |    no    | string | The bus type of the volume (VIRTIO or IDE). Default: VIRTIO.                                                                                                                                                                                                                                                     |
| Image            | **yes**  | string | The image or snapshot ID. Can be left empty for a data volume, however you'll need to set the `licence_type`. Default: _null_                                                                                                                                                                                    |
| Type             | **yes**  | string | The volume type, HDD or SSD. Default: HDD                                                                                                                                                                                                                                                                        |
| LicenceType      | **yes**  | string | The licence type of the volume. Options: LINUX, WINDOWS, WINDOWS2016, UNKNOWN, OTHER. Default: UNKNOWN                                                                                                                                                                                                           |
| ImagePassword    | **yes**  | string | A password to set on the volume for the appropriate root or administrative account. This field may only be set in creation requests. When reading, it always returns _null_. The password has to contain 8-50 characters. Only these characters are allowed: [abcdefghjkmnpqrstuvxABCDEFGHJKLMNPQRSTUVX23456789] |
| ImageAlias       | **yes**  | string | An alias to a ProfitBricks public image. Use instead of "image".]                                                                                                                                                                                                                                                |
| SSHKeys          | **yes**  | string | SSH keys to allow access to the volume via SSH.                                                                                                                                                                                                                                                                  |
| AvailabilityZone |    no    | string | The storage availability zone assigned to the volume. Valid values: AUTO, ZONE_1, ZONE_2, or ZONE_3. This only applies to HDD volumes. Leave blank or set to AUTO when provisioning SSD volumes.                                                                                                                 |

The following table outlines the various licence types you can define:

| Licence Type | Comment                                                                                                                           |
| ------------ | --------------------------------------------------------------------------------------------------------------------------------- |
| WINDOWS2016  | Use this for the Microsoft Windows Server 2016 operating system.                                                                  |
| WINDOWS      | Use this for the Microsoft Windows Server 2008 and 2012 operating systems.                                                        |
| LINUX        | Use this for Linux distributions such as CentOS, Ubuntu, Debian, etc.                                                             |
| OTHER        | Use this for any volumes that do not match one of the other licence types.                                                        |
| UNKNOWN      | This value may be inherited when you've uploaded an image and haven't set the license type. Use one of the options above instead. |

The following table outlines the storage availability zones currently supported:

| Availability Zone | Comment                     |
| ----------------- | --------------------------- |
| AUTO              | Automatically Selected Zone |
| ZONE_1            | Fire Zone 1                 |
| ZONE_2            | Fire Zone 2                 |
| ZONE_3            | Fire Zone 3                 |

**Note:** You will need to provide either the `Image` or the `LicenceType` arguments when creating a volume. A `LicenceType` is required, but if `Image` is supplied, it is already set and cannot be changed. Either the `ImagePassword` or `SshKeys` arguments need to be supplied when creating a volume using one of the official ProfitBricks images. Only official ProfitBricks provided images support the `SshKeys` and `ImagePassword` arguments.

---

#### Update a Volume

You can update various attributes of an existing volume; however, some restrictions are in place:

You can increase the size of an existing storage volume. You cannot reduce the size of an existing storage volume. The volume size will be increased without requiring a reboot if the relevant hot plug settings (`disc_virtio_hot_plug`, `disc_virtio_hot_unplug`, etc.) have been set to _true_. The additional capacity is not added automatically added to any partition, therefore you will need to handle that inside the OS afterwards. Once you have increased the volume size you cannot decrease the volume size.

Since an existing volume is being modified, none of the request arguments are specifically required as long as the changes being made satisfy the requirements for creating a volume.

The following table describes the request arguments:

| Name         | Required | Type   | Description                                                                                                                                                                                                                                                                                                         |
| ------------ | :------: | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.                                                                                                                                                                                                                                                                                                  |
| volumeId     | **yes**  | string | The ID of the volume.                                                                                                                                                                                                                                                                                               |
| Name         |    no    | string | The name of the volume.                                                                                                                                                                                                                                                                                             |
| Size         |    no    | int    | The size of the volume in GB. You may only increase the `size` when updating.                                                                                                                                                                                                                                       |
| Bus          |    no    | string | The bus type of the volume (VIRTIO or IDE). Default: VIRTIO.                                                                                                                                                                                                                                                        |
| LicenceType  |    no    | string | The licence type of the volume. Options: LINUX, WINDOWS, WINDOWS2016, UNKNOWN, OTHER. You may get an error trying to update `LicenceType` depending on the `Image` that was used to create the volume. For example, you cannot update the `LicenceType` for a volume created from a ProfitBricks supplied OS image. |

**Note**: Trying to change the `Image`, `Type`, or `AvailabilityZone` in an update request will result in an error.

Pass the arguments to `UpdateVolume`:

    var obj := VolumeProperties{
    	Name: "GO SDK Test - RENAME",
    	Size: 5,
    }
    UpdateVolume(datacenterId, volumeId, obj)

---

#### Delete a Volume

Deletes the specified volume. This will result in the volume being removed from your data center. Use this with caution.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| volumeId     | **yes**  | string | The ID of the volume. |

Pass the arguments to `DeleteVolume`:

    DeleteVolume(datacenterId, volumeId)

---

#### Create a Volume Snapshot

Creates a snapshot of a volume within the VDC. You can use a snapshot to create a new storage volume or to restore a storage volume.

The following table describes the request arguments:

| Name         | Required | Type   | Description                      |
| ------------ | :------: | ------ | -------------------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.               |
| volumeId     | **yes**  | string | The ID of the volume.            |
| Name         |    no    | string | The name of the snapshot.        |
| Description  |    no    | string | The description of the snapshot. |

Pass the arguments to `CreateSnapshot`:

    CreateSnapshot(datacenterId, volumeId, Name,Description)

---

#### Restore a Volume Snapshot

This will restore a snapshot onto a volume. A snapshot is created as just another image that can be used to create new volumes or to restore an existing volume.

The following table describes the request arguments:

| Name         | Required | Type   | Description             |
| ------------ | :------: | ------ | ----------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.      |
| volumeId     | **yes**  | string | The ID of the volume.   |
| snapshotId   | **yes**  | string | The ID of the snapshot. |

Pass the arguments to `restore_snapshot`:

RestoreSnapshot(datacenterId, volumeId, snapshotId)

---

## Snapshots

#### List Snapshots

Call the `ListSnapshots`:

    ListSnapshots()

---

#### Get a Snapshot

Retrieves the attributes of a specific snapshot.

The following table describes the request arguments:

| Name       | Required | Type   | Description             |
| ---------- | :------: | ------ | ----------------------- |
| snapshotId | **yes**  | string | The ID of the snapshot. |

Pass the arguments to `GetSnapshot`:

    GetSnapshot(snapshotId)

---

#### Update a Snapshot

Perform updates to attributes of a snapshot.

The following table describes the request arguments:

| Name                | Required | Type   | Description                                                            |
| ------------------- | :------: | ------ | ---------------------------------------------------------------------- |
| snapshotId          | **yes**  | string | The ID of the snapshot.                                                |
| Name                |    no    | string | The name of the snapshot.                                              |
| Description         |    no    | string | The description of the snapshot.                                       |
| LicenceType         |    no    | string | The snapshot's licence type: LINUX, WINDOWS, WINDOWS2016, or OTHER.    |
| CPUHotPlug          |    no    | bool   | This volume is capable of CPU hot plug (no reboot required)            |
| CPUHotUnplug        |    no    | bool   | This volume is capable of CPU hot unplug (no reboot required)          |
| RAMHotPlug          |    no    | bool   | This volume is capable of memory hot plug (no reboot required)         |
| RAMHotUnplug        |    no    | bool   | This volume is capable of memory hot unplug (no reboot required)       |
| NicHotPlug          |    no    | bool   | This volume is capable of NIC hot plug (no reboot required)            |
| NicHotUnplug        |    no    | bool   | This volume is capable of NIC hot unplug (no reboot required)          |
| DiscVirtioHotPlug   |    no    | bool   | This volume is capable of VirtIO drive hot plug (no reboot required)   |
| DiscVirtioHotUnplug |    no    | bool   | This volume is capable of VirtIO drive hot unplug (no reboot required) |
| DiscScsiHotPlug     |    no    | bool   | This volume is capable of SCSI drive hot plug (no reboot required)     |
| DiscScsiHotUnplug   |    no    | bool   | This volume is capable of SCSI drive hot unplug (no reboot required)   |

Pass the arguments to `UpdateSnapshot`:

    UpdateSnapshot(snapshotId, SnapshotProperties{Name: newValue})

---

#### Delete a Snapshot

Deletes the specified snapshot.

The following table describes the request arguments:

| Name       | Required | Type   | Description             |
| ---------- | :------: | ------ | ----------------------- |
| snapshotId | **yes**  | string | The ID of the snapshot. |

Pass the arguments to `DeleteSnapshot`:

    DeleteSnapshot(snapshotId)

---

## IP Blocks

The IP block operations assist with managing reserved /static public IP addresses.

#### List IP Blocks

Retrieve a list of available IP blocks.

    ListIpBlocks()

---

#### Get an IP Block

Retrieves the attributes of a specific IP block.

The following table describes the request arguments:

| Name      | Required | Type   | Description             |
| --------- | :------: | ------ | ----------------------- |
| ipblockid | **yes**  | string | The ID of the IP block. |

Pass the arguments to `GetIPBlock`:

    response = client.GetIPBlock('UUID')

---

#### Create an IP Block

Creates an IP block. Creating an IP block is a bit different than some of the other available create operations. IP blocks are not attached to a particular VDC, but rather to a location. Therefore, you must specify a valid `location` along with a `size` argument indicating the number of IP addresses you want to reserve in the IP block. Any resources using an IP address from an IP block must be in the same `location`.

The following table describes the request arguments:

| Name    | Required | Type   | Description                                                       |
| ------- | :------: | ------ | ----------------------------------------------------------------- |
| ipblock | **yes**  | object | An [IPBlock object](#ipblock-resource-object) you wish to create. |

To create an IP block, define the `IPBlock` resource object:

    var ipblock = IpBlock{
    	Properties: IpBlockProperties{
    		Name:     "GO SDK Test",
    		Size:     2,
    		Location: location,
    	},
    }

Pass it to `ReserveIpBlock`:

    ReserveIpBlock(ipblock)

#### IPBlock Resource Object

| Name     | Required | Type   | Description                                                           |
| -------- | :------: | ------ | --------------------------------------------------------------------- |
| Location | **yes**  | string | This must be one of the available locations: [Locations](#locations). |
| Size     | **yes**  | int    | The size of the IP block you want.                                    |
| Name     |    no    | string | A descriptive name for the IP block                                   |

---

#### Update an IPBlock

Perform updates to the attributes of a ipblock.

The following table describes the request arguments:

| Name    | Required | Type   | Description                                |
| ------- | :------: | ------ | ------------------------------------------ |
| ipblkid | **yes**  | string | The ID of the IP Block you want to update. |
| Name    |    no    | string | A descriptive name for the IP Block.       |

Build an [IpBlockProperties] object:

    var ipblock = IPBlockProperties{
    	Name: "GO SDK Test RENAME",
    }

Pass the arguments to `UpdateIPBlock`:

    UpdateIPBlock(ipblkid, ipblock)

---

#### Delete an IP Block

Deletes the specified IP Block.

The following table describes the request arguments:

| Name    | Required | Type   | Description             |
| ------- | :------: | ------ | ----------------------- |
| ipblkid | **yes**  | string | The ID of the IP block. |

Pass the arguments to `ReleaseIpBlock`:

    ReleaseIpBlock(ipblkid)

---

## LANs

#### List LANs

Retrieve a list of LANs within the VDC.

The following table describes the request arguments:

| Name         | Required | Type   | Description        |
| ------------ | :------: | ------ | ------------------ |
| datacenterId | **yes**  | string | The ID of the VDC. |

Pass the arguments to `ListLans`:

    ListLans(datacenterId)

---

#### Create a LAN

Creates a LAN within a VDC.

The following table describes the request arguments:

| Name         | Required | Type   | Description                                                        |
| ------------ | :------: | ------ | ------------------------------------------------------------------ |
| datacenterId | **yes**  | string | The ID of the VDC.                                                 |
| lan          | **yes**  | object | A [LAN object](#lan-resource-object) describing the LAN to create. |

Create the `LAN` resource object:

    var request = CreateLanRequest{
    	Properties: LanProperties{
    		Public: true,
    		Name:   "GO SDK Test with failover",
    	},
    	Entities: &LanEntities{
    		Nics: lanNics,
    	},
    }

Pass the object and arguments to `create_lan`:

    CreateLan(datacenterId, request)

#### LAN Resource Object

| Name   | Required | Type   | Description                                                                                                   |
| ------ | :------: | ------ | ------------------------------------------------------------------------------------------------------------- |
| Name   |    no    | string | The name of your LAN.                                                                                         |
| Public | **Yes**  | bool   | Boolean indicating if the LAN faces the public Internet or not.                                               |
| Nics   |    no    | list   | One or more NIC IDs attached to the LAN.                                                                      |
| PCC    |    no    | string | Optionally, add this LAN as a peer to a [PrivateCrossConnect](#Private-Cross-Connects) by specifying its UUID |

#### IPFailover Resource Object

| Name    | Required | Type   | Description             |
| ------- | -------- | ------ | ----------------------- |
| IP      | **yes**  | string | The failover IP Address |
| NicUUID | **yes**  | string | The NIC UUID            |

#### Get a LAN

Retrieves the attributes of a given LAN.

The following table describes the request arguments:

| Name         | Required | Type   | Description        |
| ------------ | :------: | ------ | ------------------ |
| datacenterId | **yes**  | string | The ID of the VDC. |
| lanId        | **yes**  | int    | The ID of the LAN. |

Pass the arguments to `GetLan`:

    GetLan(datacenterId, lanId)

---

#### Update a LAN

Perform updates to attributes of a LAN.

The following table describes the request arguments:

| Name         | Required | Type                                                           | Description                                                     |
| ------------ | :------: | -------------------------------------------------------------- | --------------------------------------------------------------- |
| datacenterId | **yes**  | string                                                         | The ID of the VDC.                                              |
| lanId        | **yes**  | int                                                            | The ID of the LAN.                                              |
| Name         |    no    | string                                                         | A descriptive name for the LAN.                                 |
| Public       |    no    | bool                                                           | Boolean indicating if the LAN faces the public Internet or not. |
| IPFailover   |    no    | \*[] seeIPFailover - [IPFailover](#ipfailover-resource-object) | Only supported when updating a LAN                              |

Pass the arguments to `update_lan`:

    var obj = LanProperties{
    	Properties: LanProperties{
    		Public: true,
    		Name:   "GO SDK Test with failover",
    		IPFailover: &[]IPFailover {
    			IPFailover {
    				IP: "1.1.1.1",
    				NicUUID: ".....",
    			}
    		},
    	}
    UpdateLan(datacenterId, lanId, obj)

---

#### Delete a LAN

Deletes the specified LAN.

The following table describes the request arguments:

| Name         | Required | Type   | Description        |
| ------------ | :------: | ------ | ------------------ |
| datacenterId | **yes**  | string | The ID of the VDC. |
| lanId        | **yes**  | string | The ID of the LAN. |

Pass the arguments to `delete_lan`:

    DeleteLan(lan_dcid, lanid)

---

## Network Interfaces (NICs)

#### List NICs

Retrieve a list of LANs within the VDC.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |

Pass the arguments to `ListNics`:

    ListNics(datacenterId, serverId)

---

#### Get a NIC

Retrieves the attributes of a given NIC.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |
| nicId        | **yes**  | string | The ID of the NIC.    |

Pass the arguments to `GetNic`:

    GetNic(datacenterId, serverId, nicId)

---

#### Create a NIC

Adds a NIC to the target server.

The following table describes the request arguments:

| Name         | Required | Type   | Description                                                            |
| ------------ | :------: | ------ | ---------------------------------------------------------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.                                                     |
| serverId     | **yes**  | string | The ID of the server.                                                  |
| nic          | **yes**  | object | A [NIC object](#nic-resource-object) describing the NIC to be created. |

Create the `NIC` resource object:

    var nic = Nic{
    	Properties: &NicProperties{
    		Lan:            1,
    		Name:           "GO SDK Test",
    		Nat:            false,
    		Dhcp:           true,
    		FirewallActive: true,
    		Ips:            []string{"10.0.0.1"},
    	},
    }

Pass the object and arguments to `create_nic`:

CreateNic(datacenterId, serverId, nic)

#### NIC Resource Object

| Name           | Required | Type   | Description                                                                                  |
| -------------- | :------: | ------ | -------------------------------------------------------------------------------------------- |
| Name           |    no    | string | The name of the NIC.                                                                         |
| Ips            |    no    | list   | IP addresses assigned to the NIC.                                                            |
| Dhcp           |    no    | bool   | Set to _false_ if you wish to disable DHCP on the NIC. Default: _true_.                      |
| Lan            | **yes**  | int    | The LAN ID the NIC will sit on. If the LAN ID does not exist it will be created.             |
| Nat            |    no    | bool   | Indicates the private IP address has outbound access to the public internet.                 |
| FirewallActive |    no    | bool   | Set this to _true_ to enable the ProfitBricks firewall, _false_ to disable.                  |
| Firewallrules  |    no    | list   | A list of [FirewallRule objects](#firewall-rule-resource-object) to be created with the NIC. |

---

#### Update a NIC

You can update -- in full or partially -- various attributes on the NIC; however, some restrictions are in place:

The primary address of a NIC connected to a load balancer can only be changed by changing the IP of the load balancer. You can also add additional reserved, public IPs to the NIC.

The user can specify and assign private IPs manually. Valid IP addresses for private networks are 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16.

The following table describes the request arguments:

| Name           | Required | Type   | Description                                                                  |
| -------------- | :------: | ------ | ---------------------------------------------------------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.                                                           |
| serverId       | **yes**  | string | The ID of the server.                                                        |
| nicId          | **yes**  | string | The ID of the NIC.                                                           |
| Name           |    no    | string | The name of the NIC.                                                         |
| Ips            |    no    | list   | IPs assigned to the NIC represented as a list of strings.                    |
| Dhcp           |    no    | bool   | Boolean value that indicates if the NIC is using DHCP or not.                |
| Lan            |    no    | int    | The LAN ID the NIC sits on.                                                  |
| Nat            |    no    | bool   | Indicates the private IP address has outbound access to the public internet. |
| FirewallActive |    no    | bool   | Set this to _true_ to enable the ProfitBricks firewall, _false_ to disable.  |

Pass the arguments to `update_nic`:

    var obj = NicProperties{Name: "GO SDK Test - RENAME", Lan: 1}
    UpdateNic(nic_dcid, nic_srvid, nicid, obj)

---

#### Delete a NIC

Deletes the specified NIC.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |
| nicId        | **yes**  | string | The ID of the NIC.    |

Pass the arguments to `DeleteNic`:

    DeleteNic(nic_dcid, nic_srvid, nicid)

---

## Firewall Rules

#### List Firewall Rules

Retrieves a list of firewall rules associated with a particular NIC.

The following table describes the request arguments:

| Name         | Required | Type   | Description           |
| ------------ | :------: | ------ | --------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.    |
| serverId     | **yes**  | string | The ID of the server. |
| nicId        | **yes**  | string | The ID of the NIC.    |

Pass the arguments to `ListFirewallRules`:

    ListFirewallRules(datacenterId, serverId, nicId)

---

#### Get a Firewall Rule

Retrieves the attributes of a given firewall rule.

The following table describes the request arguments:

| Name           | Required | Type   | Description                  |
| -------------- | :------: | ------ | ---------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.           |
| serverId       | **yes**  | string | The ID of the server.        |
| nicId          | **yes**  | string | The ID of the NIC.           |
| firewallRuleId | **yes**  | string | The ID of the firewall rule. |

Pass the arguments to `get_firewall_rule`:

    GetFirewallRule(datacenterId, serverId, nicId, firewallRuleId)

---

#### Create a Firewall Rule

This will add a firewall rule to the NIC.

The following table describes the request arguments:

| Name         | Required | Type   | Description                                                                                         |
| ------------ | :------: | ------ | --------------------------------------------------------------------------------------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.                                                                                  |
| serverId     | **yes**  | string | The ID of the server.                                                                               |
| nicId        | **yes**  | string | The ID of the NIC.                                                                                  |
| firewallRule | **yes**  | object | A [FirewallRule object](#firewall-rule-resource-object) describing the firewall rule to be created. |

Create the `FirewallRule` resource object:

    var firewallRule FirewallRule{
    	Properties: FirewallruleProperties{
    		Name:           "SSH",
    		Protocol:       "TCP",
    		SourceMac:      "01:23:45:67:89:00",
    		PortRangeStart: 22,
    		PortRangeEnd:   22,
    	},
    }

Pass the object and arguments to `create_firewall_rule`:

    CreateFirewallRule(datacenterId, serverId, nicId, firewallRule)

#### Firewall Rule Resource Object

| Name           | Required | Type   | Description                                                                                                                                                                    |
| -------------- | :------: | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Name           |    no    | string | The name of the firewall rule.                                                                                                                                                 |
| Protocol       | **yes**  | string | The protocol for the rule: TCP, UDP, ICMP, ANY.                                                                                                                                |
| SourceMac      |    no    | string | Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. A _nil_ value allows all source MAC address.                             |
| SourceIP       |    no    | string | Only traffic originating from the respective IPv4 address is allowed. A _nil_ value allows all source IPs.                                                                     |
| TargetIP       |    no    | string | In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. A _nil_ value allows all target IPs.               |
| PortRangeStart |    no    | string | Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave `PortRangeStart` and `PortRangeEnd` value as _nil_ to allow all ports.   |
| PortRangeEnd   |    no    | string | Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave `PortRangeStart` and `PortRangeEnd` value as _nil_ to allow all ports. |
| IcmpType       |    no    | string | Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen. A _nil_ value allows all types.                                                                       |
| IcmpCode       |    no    | string | Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. A _nil_ value allows all codes.                                                                           |

---

#### Update a Firewall Rule

Perform updates to an existing firewall rule. You will notice that some arguments, such as `protocol` cannot be updated. If the `protocol` needs to be changed, you can [delete](#delete-a-firewall-rule) the firewall rule and then [create](#create-a-firewall-rule) new one to replace it.

The following table describes the request arguments:

| Name           | Required | Type   | Description                                                                                                                                                                    |
| -------------- | :------: | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| datacenterId   | **yes**  | string | The ID of the VDC.                                                                                                                                                             |
| serverId       | **yes**  | string | The ID of the server.                                                                                                                                                          |
| nicId          | **yes**  | string | The ID of the NIC.                                                                                                                                                             |
| firewallRuleId | **yes**  | string | The ID of the firewall rule.                                                                                                                                                   |
| Name           |    no    | string | The name of the firewall rule.                                                                                                                                                 |
| SourceMac      |    no    | string | Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. A _nil_ value allows all source MAC address.                             |
| SourceIP       |    no    | string | Only traffic originating from the respective IPv4 address is allowed. A _nil_ value allows all source IPs.                                                                     |
| TargetIP       |    no    | string | In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. A _nil_ value allows all target IPs.               |
| PortRangeStart |    no    | string | Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave `PortRangeStart` and `PortRangeEnd` value as _nil_ to allow all ports.   |
| PortRangeEnd   |    no    | string | Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave `PortRangeStart` and `PortRangeEnd` value as _nil_ to allow all ports. |
| IcmpType       |    no    | string | Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen. A _nil_ value allows all types.                                                                       |
| IcmpCode       |    no    | string | Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. A _nil_ value allows all codes.                                                                           |

Pass the arguments to `UpdateFirewallRule`:

    props := FirewallruleProperties{
    	Name: "SSH - RENAME",
    }
    UpdateFirewallRule(dcID, srv_srvid, nicid, fwId, props)

---

#### Delete a Firewall Rule

Removes a firewall rule.

| Name           | Required | Type   | Description                  |
| -------------- | :------: | ------ | ---------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.           |
| serverId       | **yes**  | string | The ID of the server.        |
| nicId          | **yes**  | string | The ID of the NIC.           |
| firewallRuleId | **yes**  | string | The ID of the firewall rule. |

Pass the arguments to `DeleteFirewallRule`:

    DeleteFirewallRule(dcID, srv_srvid, nicid, fwId)

---

## Load Balancers

#### List Load Balancers

Retrieve a list of load balancers within the data center.

The following table describes the request arguments:

| Name         | Required | Type   | Description        |
| ------------ | :------: | ------ | ------------------ |
| datacenterId | **yes**  | string | The ID of the VDC. |

Pass the arguments to `ListLoadbalancers`:

    ListLoadbalancers(datacenterId)

---

#### Get a Load Balancer

Retrieves the attributes of a given load balancer.

The following table describes the request arguments:

| Name           | Required | Type   | Description                  |
| -------------- | :------: | ------ | ---------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.           |
| loadbalancerId | **yes**  | string | The ID of the load balancer. |

Pass the arguments to `GetLoadbalancer`:

    GetLoadbalancer(datacenterId, loadbalancerId)

---

#### Create a Load Balancer

Creates a load balancer within the VDC. Load balancers can be used for public or private IP traffic.

The following table describes the request arguments:

| Name         | Required | Type   | Description                                                                                         |
| ------------ | :------: | ------ | --------------------------------------------------------------------------------------------------- |
| datacenterId | **yes**  | string | The ID of the VDC.                                                                                  |
| loadbalancer | **yes**  | object | A [LoadBalancer object](#load-balancer-resource-object) describing the load balancer to be created. |

Create the `LoadBalancer` resource object:

    var loadbalancer = Loadbalancer{
    	Properties: LoadbalancerProperties{
    		Name: "GO SDK Test",
    		Ip:   "10.0.0.1",
    		Dhcp: true,
    	}
    }

Pass the object and arguments to `CreateLoadbalancer`:

    CreateLoadbalancer(datacenterId, loadbalancer)

#### Load Balancer Resource Object

| Name         | Required | Type   | Description                                                                                           |
| ------------ | :------: | ------ | ----------------------------------------------------------------------------------------------------- |
| Name         | **yes**  | string | The name of the load balancer.                                                                        |
| Ip           |    no    | string | IPv4 address of the load balancer. All attached NICs will inherit this IP.                            |
| Dhcp         |    no    | bool   | Indicates if the load balancer will reserve an IP using DHCP.                                         |
| Balancednics |    no    | list   | List of NIC IDs taking part in load-balancing. All balanced NICs inherit the IP of the load balancer. |

---

#### Update a Load Balancer

Perform updates to attributes of a load balancer.

The following table describes the request arguments:

| Name           | Required | Type   | Description                                                   |
| -------------- | :------: | ------ | ------------------------------------------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.                                            |
| loadbalancerId | **yes**  | string | The ID of the load balancer.                                  |
| Name           |    no    | string | The name of the load balancer.                                |
| Ip             |    no    | string | The IP of the load balancer.                                  |
| Dhcp           |    no    | bool   | Indicates if the load balancer will reserve an IP using DHCP. |

Pass the arguments to `UpdateLoadbalancer`:

    var obj = LoadbalancerProperties{Name: "GO SDK Test - RENAME"}
    UpdateLoadbalancer(datacenterId, loadbalancerId, obj)

---

#### Delete a Load Balancer

Deletes the specified load balancer.

The following table describes the request arguments:

| Name           | Required | Type   | Description                  |
| -------------- | :------: | ------ | ---------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.           |
| loadbalancerId | **yes**  | string | The ID of the load balancer. |

Pass the arguments to `DeleteLoadbalancer`:

    DeleteLoadbalancer(datacenterId, loadbalancerId)

---

#### List Load Balanced NICs

This will retrieve a list of NICs associated with the load balancer.

The following table describes the request arguments:

| Name           | Required | Type   | Description                  |
| -------------- | :------: | ------ | ---------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.           |
| loadbalancerId | **yes**  | string | The ID of the load balancer. |

Pass the arguments to `ListBalancedNics`:

    ListBalancedNics(datacenterId, loadbalancerId)

---

#### Get a Load Balanced NIC

Retrieves the attributes of a given load balanced NIC.

The following table describes the request arguments:

| Name           | Required | Type   | Description                  |
| -------------- | :------: | ------ | ---------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.           |
| loadbalancerId | **yes**  | string | The ID of the load balancer. |
| nicId          | **yes**  | string | The ID of the NIC.           |

Pass the arguments to `GetBalancedNic`:

    GetBalancedNic(datacenterId, loadbalancerId, nicId)

---

#### Associate NIC to a Load Balancer

This will associate a NIC to a load balancer, enabling the NIC to participate in load-balancing.

The following table describes the request arguments:

| Name           | Required | Type   | Description                  |
| -------------- | :------: | ------ | ---------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.           |
| loadbalancerId | **yes**  | string | The ID of the load balancer. |
| nicId          | **yes**  | string | The ID of the NIC.           |

Pass the arguments to `add_loadbalanced_nics`:

    AssociateNic(datacenterId, loadbalancerId, nicId)

---

#### Remove a NIC Association

Removes the association of a NIC with a load balancer.

The following table describes the request arguments:

| Name           | Required | Type   | Description                                                |
| -------------- | :------: | ------ | ---------------------------------------------------------- |
| datacenterId   | **yes**  | string | The ID of the VDC.                                         |
| loadbalancerId | **yes**  | string | The ID of the load balancer.                               |
| nicId          | **yes**  | string | The ID of the NIC you are removing from the load balancer. |

Pass the arguments to `DeleteBalancedNic`:

    DeleteBalancedNic(datacenterId, loadbalancerId, nicId)

---

## Requests

Each call to the ProfitBricks Cloud API is assigned a request ID. These operations can be used to get information about the requests that have been submitted and their current status.

#### List Requests

    ListRequests()

---

#### Get a Request

Retrieves the attributes of a specific request. This operation shares the same `get_request` method used for getting request status, however the response it determined by the boolean value you pass for _status_. To get details about the request itself, you want to pass a _status_ of _False_.

The following table describes the request arguments:

| Name       | Required | Type   | Description                                          |
| ---------- | :------: | ------ | ---------------------------------------------------- |
| request_id | **yes**  | string | The ID of the request.                               |
| status     | **yes**  | bool   | Set to _False_ to have the request details returned. |

Pass the arguments to `get_request`:

    response = client.get_request(
        request_id='UUID',
        status=False)

---

#### Get a Request Status

Retrieves the status of a request. This operation shares the same `get_request` method used for getting the details of a request, however the response it determined by the boolean value you pass for _status_. To get the request status, you want to pass a _status_ of _True_.

The following table describes the request arguments:

| Name | Required | Type   | Description                                                    |
| ---- | :------: | ------ | -------------------------------------------------------------- |
| path | **yes**  | string | The ID of the request. Retrieved from response header location |

Pass the arguments to `get_request`:

GetRequestStatus(path)

---

## Contract Resources

#### List Contract Resources

Returns information about the resource limits for a particular contract and the current resource usage.

```
GetContractResources()
```

---

## Users Management

These operations are designed to allow you to orchestrate users and resources via the Cloud API. Previously this functionality required use of the DCD (Data Center Designer) web application.

#### List Groups

This retrieves a full list of all groups.

```
ListGroups()
```

#### Retrieve a Group

The following table describes the request arguments:

| Name    | Required | Type   | Description                               |
| ------- | -------- | ------ | ----------------------------------------- |
| groupId | Yes      | string | The ID of the specific group to retrieve. |

```
GetGroup(groupid)
```

#### Create a Group

The following table describes the request arguments:

| Name  | Type  | Description                                | Required |
| ----- | ----- | ------------------------------------------ | -------- |
| group | Group | See [Group Object](#group-resource-object) | Yes      |

Build the `Group` resource object:

    var group = Group{
    	Properties: GroupProperties{
    		Name:              "GO SDK Test",
    		CreateDataCenter:  &TRUE,
    		CreateSnapshot:    &TRUE,
    		ReserveIp:         &TRUE,
    		AccessActivityLog: &TRUE,
    	},
    }

Pass the object to `CreateGroup`:

```
CreateGroup(group Group)
```

#### Group Resource Object

| Name                 | Required | Type   | Description                                               |
| -------------------- | :------: | ------ | --------------------------------------------------------- |
| Name                 | **yes**  | string | A name that was given to the group.                       |
| CreateDataCenter     |    no    | bool   | The group has permission to create virtual data centers.  |
| CreateSnapshot       |    no    | bool   | The group has permission to create snapshots.             |
| ReserveIP            |    no    | bool   | The group has permission to reserve IP addresses.         |
| AccessActivityLog    |    no    | bool   | The group has permission to access the activity log.      |
| CreateBackupUnit     |    no    | bool   | The group has permission to create backup units.          |
| CreateInternetAccess |    no    | bool   | The group has permission to create/grant internet access. |
| CreateK8sCluster     |    no    | bool   | The group has permission to create k8s clusters.          |
| CreatePcc            |    no    | bool   | The group has permission to create PCC.                   |
| S3Privilege          |    no    | bool   | The group has permission to create S3 Block Storage.      |

#### Update a Group

Use this operation to update a group.

The following table describes the request arguments:

| Name    | Type    | Description                                | Required                                  |
| ------- | ------- | ------------------------------------------ | ----------------------------------------- |
| groupId | **yes** | string                                     | The ID of the specific group to retrieve. |
| group   | Group   | See [Group Object](#group-resource-object) | Yes                                       |

```
UpdateGroup(groupId, group Group)
```

---

#### Delete a Group

This will remove all objects within the data center and remove the data center object itself.
Use this operation to delete a single group. Resources that are assigned to the group are NOT deleted, but are no longer accessible to the group members unless the member is a Contract Owner, Admin, or Resource Owner.

The following table describes the request arguments:

| Name    | Type    | Description | Required                                  |
| ------- | ------- | ----------- | ----------------------------------------- |
| groupId | **yes** | string      | The ID of the specific group to retrieve. |

```
DeleteGroup(groupId)
```

---

#### List Shares

Retrieves a full list of all the resources that are shared through this group and lists the permissions granted to the group members for each shared resource.

```
ListShares()
```

#### Retrieve a Share

The following table describes the request arguments:

| Name       | Required | Type   | Description                                  |
| ---------- | -------- | ------ | -------------------------------------------- |
| groupid    | **yes**  | string | The ID of the specific group to retrieve.    |
| resourceId | **yes**  | string | The ID of the specific resource to retrieve. |

```
GetShare(groupid, resourceId)
```

---

#### Add a Share

The following table describes the request arguments:

| Name       | Required | Type   | Description                                         |
| ---------- | -------- | ------ | --------------------------------------------------- |
| groupid    | **yes**  | string | The ID of the specific group to add a resource too. |
| resourceId | **yes**  | string | The ID of the specific resource to add.             |
| share      | **yes**  | Share  | See [Share Object](#share-resource-object)          |

Build the `Share` resource object:

    var share = Share{
    	Properties: ShareProperties{
    		SharePrivilege: true,
    		EditPrivilege:  true,
    	},
    }

Pass the object to `AddShare`:

```
AddShare(share Share, groupid, resourceId)
```

#### Share Resource Object

| Name           | Required | Type | Description                                                   |
| -------------- | :------: | ---- | ------------------------------------------------------------- |
| EditPrivilege  |    no    | bool | The group has permission to edit privileges on this resource. |
| SharePrivilege |    no    | bool | The group has permission to share this resource.              |

---

#### Update a Share

Use this to update the permissions that a group has for a specific resource share.

The following table describes the request arguments:

| Name       | Required | Type   | Description                                         |
| ---------- | -------- | ------ | --------------------------------------------------- |
| groupid    | **yes**  | string | The ID of the specific group to add a resource too. |
| resourceId | **yes**  | string | The ID of the specific resource to add.             |
| share      | **yes**  | Share  | See [Share Object](#share-resource-object)          |

```
UpdateShare(groupid, resourceId, obj)
```

#### Delete a Share

This will remove all objects within the data center and remove the data center object itself.
Use this operation to delete a single group. Resources that are assigned to the group are NOT deleted, but are no longer accessible to the group members unless the member is a Contract Owner, Admin, or Resource Owner.

The following table describes the request arguments:

| Name       | Type    | Description | Required                                                        |
| ---------- | ------- | ----------- | --------------------------------------------------------------- |
| groupid    | **yes** | string      | The ID of the specific group containing the resource to delete. |
| resourceId | **yes** | string      | The ID of the specific resource to delete.                      |

```
DeleteShare(groupid, resourceId)
```

---

#### List Users in a Group

Retrieves a full list of all the users that are members of a particular group.

The following table describes the request arguments:

| Name    | Type    | Description | Required                                                  |
| ------- | ------- | ----------- | --------------------------------------------------------- |
| groupid | **yes** | string      | The ID of the specific group to retrieve a user list for. |

```
ListGroupUsers(groupid)
```

---

#### Add User to Group

The following table describes the request arguments:

| Name    | Required | Type   | Description                                             |
| ------- | -------- | ------ | ------------------------------------------------------- |
| groupid | **yes**  | string | The ID of the specific group you want to add a user to. |
| userid  | **yes**  | string | The ID of the specific user to add to the group.        |

```
AddUserToGroup(groupid, userid)
```

---

#### Remove User from a Group

Use this operation to remove a user from a group.

The following table describes the request arguments:

| Name    | Type    | Description | Required                                                     |
| ------- | ------- | ----------- | ------------------------------------------------------------ |
| groupid | **yes** | string      | The ID of the specific group you want to remove a user from. |
| userid  | **yes** | string      | The ID of the specific user to remove from the group.        |

```
DeleteUserFromGroup(groupid, userid)
```

---

#### List Users

Retrieve a list of all the users that have been created under a contract.

```
ListUsers()
```

---

#### Retrieve a User

Retrieve details about a specific user including what groups and resources the user is associated with.

The following table describes the request arguments:

| Name   | Required | Type   | Description                                                |
| ------ | :------: | ------ | ---------------------------------------------------------- |
| userid | **yes**  | string | The ID of the specific user to retrieve information about. |

```
GetUser(userid)
```

---

#### Create a User

Creates a new user under a particular contract.

The following table describes the request arguments:

| Name | Required | Type                          | Description     |
| ---- | :------: | ----------------------------- | --------------- |
| user | **yes**  | [User](#user-resource-object) | The user data.  |

Build the `User` resource object:

```go
user := User{
    Properties: &UserProperties{
        Firstname:     "John",
        Lastname:      "Doe",
        Email:         "test@go.com",
        Password:      "abc123-321CBA",
        Administrator: false,
        ForceSecAuth:  false,
        SecAuthActive: false,
    },
}
```

Pass the object to `CreateUser`:

```go
CreateUser(user)
```

#### User Resource Object

| Name              | Required | Type  | Description                                                               |
| ----------------- | :-----: | ------ | ------------------------------------------------------------------------- |
| Firstname         | **yes** | string | The first name of the user.                                               |
| Lastname          | **yes** | string | The last name of the user.                                                |
| Email             | **yes** | string | The e-mail address of the user.                                           |
| Password          | **yes** | string | A password for the user.                                                  |
| Administrator     |    no   | bool   | Indicates if the user has administrative rights.                          |
| ForceSecAuth      |    no   | bool   | Indicates if secure (two-factor) authentication was enabled for the user. |
| SecAuthActive     |    no   | bool   | Indicates if secure (two-factor) authentication is enabled for the user.  |
| Active            |    no   | *bool  | Indicates if the user is active (true) or disabled (false).               |
| S3CanonicalUserID |    no   | string | The user's S3 ID.                                                         |
---

#### Update a User

Update details about a specific user including its privileges.

The following table describes the request arguments:

| Name   | Required | Type                          | Description                            |
| ------ | :------: | ----------------------------- | -------------------------------------- |
| userid | **yes**  | string                        | The ID of the specific user to update. |
| user   | **yes**  | [User](#user-resource-object) | The user data.                         |

```go
t := true
user := User{
    Properties: &UserProperties{
        Firstname:     "go sdk",
        Lastname:      "newName",
        Email:         "test@go.com",
        Administrator: false,
        ForceSecAuth:  false,
        SecAuthActive: false,
        Active:        &t,
    },
}
```

Pass the object to `UpdateUser`:

```go
UpdateUser(userid, user)
```

---

#### Delete a User

Blacklists the user, disabling them. The user is not completely purged, therefore if you anticipate needing to create a user with the same name in the future, we suggest renaming the user before you delete it.

The following table describes the request arguments:

| Name   | Type    | Description | Required                               |
| ------ | :-----: | ----------- | -------------------------------------- |
| userid | **yes** | string      | The ID of the specific user to delete. |

```go
DeleteUser(userid)
```

---

#### List Resources

Retrieves a list of all resources and optionally their group associations.

_Note_: This API call can take a significant amount of time to return when there are a large number of provisioned resources. You may wish to consult the next section on how to list resources of a particular type.

```
ListResources()
```

---

#### List All Resources of a Type

Lists all shareable resources of a specific type. Optionally include their association with groups, permissions that a group has for the resource, and users that are members of the group. Because you are scoping your request to a specific resource type, this API will likely return faster than querying `/um/resources`.

The following table describes the request arguments:

| Name         | Type    | Description | Required                                                      |
| ------------ | ------- | ----------- | ------------------------------------------------------------- |
| resourcetype | **Yes** | string      | The specific type of resources to retrieve information about. |

The values available for resourcetype are listed in this table:

| Resource Type | Description                                             |
| ------------- | ------------------------------------------------------- |
| datacenter    | A virtual data center.                                  |
| image         | A private image that has been uploaded to ProfitBricks. |
| snapshot      | A snapshot of a storage volume.                         |
| ipblock       | An IP block that has been reserved.                     |

```
ListResourcesByType(resourcetype)
```

---

#### List a specific Resource Type

The following table describes the request arguments:

| Name         | Type    | Description | Required                                                       |
| ------------ | ------- | ----------- | -------------------------------------------------------------- |
| resourcetype | **Yes** | string      | The specific type of resources to retrieve information about.  |
| resourceId   | **Yes** | string      | The ID of the specific resource to retrieve information about. |

The values available for resourcetype are listed in this table:

| Resource Type | Description                                             |
| ------------- | ------------------------------------------------------- |
| datacenter    | A virtual data center.                                  |
| image         | A private image that has been uploaded to ProfitBricks. |
| snapshot      | A snapshot of a storage volume.                         |
| ipblock       | An IP block that has been reserved.                     |

```
GetResourceByType(resourcetype, resourceId)
```

---

## Managed Kubernetes

##### KubernetesCluster Resource Object

| Name       | Required | Type                                                                          | Description                               |
| ---------- | :------: | ----------------------------------------------------------------------------- | ----------------------------------------- |
| Properties | **yes**  | \*[KubernetesClusterProperties](#KubernetesClusterProperties-resource-object) | The properties of the Kubernetes Cluster. |

##### UpdatedKubernetesCluster Resource Object

| Name       | Required | Type                                                                          | Description                               |
| ---------- | :------: | ----------------------------------------------------------------------------- | ----------------------------------------- |
| Properties | **yes**  | \*[KubernetesClusterProperties](#KubernetesClusterProperties-resource-object) | The properties of the Kubernetes Cluster. |

## AutoScaling Resource Object

| Name         | Required | Type   | Description                                                                                        |
| ------------ | :------: | ------ | -------------------------------------------------------------------------------------------------- |
| MinNodeCount | **yes**  | uint32 | The minimum number of nodes this node pool should scale down to. Should be less than MaxNodeCount  |
| MaxNodeCount | **yes**  | uint32 | The maximum number of nodes this node pool should scale up to. Should be greater than MinNodeCount |

## KubernetesNodePoolLAN Resource Object

| Name | Required | Type   | Description                  |
| ---- | -------- | ------ | ---------------------------- |
| ID   | **yes**  | uint32 | The ID of the LAN in the VDC |

## MaintenanceWindow Resource Object

| Name         | Required | Type   | Description                                                                                                 |
| ------------ | :------: | ------ | ----------------------------------------------------------------------------------------------------------- |
| DayOfTheWeek | **yes**  | string | The english name of the day of the week (One of: Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday). |
| Time         | **yes**  | string | The time of the day in **HH:mm:ss** format                                                                  |

##### KubernetesClusterProperties Resource Object

| Name              | Required | Type                                                    | Description                                 |
| ----------------- | :------: | ------------------------------------------------------- | ------------------------------------------- |
| Name              | **yes**  | string                                                  | The desired name for the Kubernetes Cluster |
| K8sVersion        |  **no**  | string                                                  | The desired Kubernetes Version              |
| MaintenanceWindow |  **no**  | [MaintenanceWindow](#MaintenanceWindow-resource-object) | The desired Maintenance Window              |

##### KubernetesNodePool Resource Object

| Name       | Required | Type                                                                            | Description                                 |
| ---------- | :------: | ------------------------------------------------------------------------------- | ------------------------------------------- |
| Properties | **yes**  | \*[KubernetesNodePoolProperties](#KubernetesNodePoolProperties-resource-object) | The properties of the kubernetes node pool. |

##### KubernetesNodePoolProperties Resource Object

| Name              | Required | Type                                                                  | Description                                                                                                    |
| ----------------- | :------: | --------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| Name              | **yes**  | string                                                                | The desired name of the node pool - The name of each node in the pool will have this as prefix                 |
| DatacenterID      | **yes**  | string                                                                | The desired datacenter location - see [Locations](#locations)                                                  |
| NodeCount         | **yes**  | uint32                                                                | The desired number of nodes in the cluster                                                                     |
| CPUFamily         | **yes**  | string                                                                | CPU Family for the nodes - see [Locations](#locations)                                                         |
| CoresCount        | **yes**  | uint32                                                                | Desired number of CPU cores per node                                                                           |
| AvailabilityZone  | **yes**  | string                                                                | Desired availability zone (one of AUTO, ZONE_1, ZONE_2)                                                        |
| RAMSize           | **yes**  | uint32                                                                | Desired amount of RAM per node in MB - Size must be specified in multiples of 256 MB with a minimum of 2048 MB |
| StorageSize       | **yes**  | uint32                                                                | Desired storage size, in MB                                                                                    |
| StorageType       | **yes**  | string                                                                | The storage type for each node (one of HDD, SSD)                                                               |
| K8sVersion        |   yes    | string                                                                | The kubernetes version in which a nodepool is running                                                          |
| AutoScaling       |    no    | [\*AutoScaling](#AutoScaling-resource-object)                         | Whether this Node Pool should autoscale. Comprised of a minimum and a maximum number of nodes                  |
| LANs              |  **no**  | \*\[\][KubernetesNodePoolLAN](#KubernetesNodePoolLAN-Resource-Object) | A list of Local Area Networks the nodes in the pool should be a part of                                        |
| MaintenanceWindow |  **no**  | [MaintenanceWindow](#MaintenanceWindow-resource-object)               | An optional object with 2 keys: dayOfTheWeek and time.                                                         |
| PublicIPs         |  **no**  | \*\[\]string                                                          | An optional array of strings with public IPs associated with the node pool.                                    |

### Create Cluster

Creates a Kubernetes cluster

The following table describes the request arguments:

| Name    | Type                                                          | Description                           | Required |
| ------- | ------------------------------------------------------------- | ------------------------------------- | -------- |
| cluster | \*[KubernetesCluster](#KubernetesNodePoolLan-Resource-Object) | An Kubernetes cluster resource object | **yes**  |

```golang

cluster := profitbricks.KubernetesCluster{
    Properties: &profitbricks.KubernetesClusterProperties{
        Name: "demo",
    },
}

createdCluster, err := client.CreateKubernetesCluster(cluster)

if err != nil {
    fmt.Printf("Error while creating cluster: %+v\n", err)
} else {
    fmt.Printf("Successfully created cluster: %+v\n", createdCluster)
}

```

### Read Cluster

Retrieves an existing Kubernetes cluster

The following table describes the request arguments:

| Name      | Type   | Description                                | Required |
| --------- | ------ | ------------------------------------------ | -------- |
| clusterID | string | The uuid of an existing Kubernetes cluster | **yes**  |

```golang

cluster, err := client.GetKubernetesCluster("b924d955-4d2e-4e0d-806a-4350ac1ec92x")

if err != nil {
    fmt.Printf("Error while reading the cluster: %+v\n", err)
} else {
    fmt.Printf("Successfully read the cluster: %+v\n", cluster)
}

```

### List Clusters

Retrieves a list of all of the Kubernetes clusters the current user has access to

```golang
clusters, err := client.ListKubernetesClusters()

if err != nil {
    fmt.Printf("Error while retrieving cluster list: %+v\n", err)
} else {
    fmt.Printf("Successfully retrieved cluster list: %+v\n", clusters)
}

```

### Update Cluster

Updates an existing Kubernetes cluster

The following table describes the request arguments:

| Name      | Type                                                                    | Description                                      | Required |
| --------- | ----------------------------------------------------------------------- | ------------------------------------------------ | -------- |
| clusterID | string                                                                  | The UUID of the Kubernetes cluster to be updated | **yes**  |
| cluster   | \*[UpdatedKubernetesCluster](#UpdatedKubernetesCluster-resource-object) | An Kubernetes cluster resource object            | **yes**  |

```golang

cluster := profitbricks.UpdatedKubernetesCluster{
    Properties: &profitbricks.KubernetesClusterProperties{
        Name: "demo-renamed",
    },
}

updatedCluster, err := client.UpdateKubernetesCluster("b924d955-4d2e-4e0d-806a-4350ac1ec92x", cluster)

if err != nil {
    fmt.Printf("Error while updating the cluster: %+v\n", err)
} else {
    fmt.Printf("Successfully updated the cluster: %+v\n", updatedCluster)
}

```

### Delete Cluster

Deletes an existing Kubernetes cluster

The following table describes the request arguments:

| Name      | Type   | Description                                      | Required |
| --------- | ------ | ------------------------------------------------ | -------- |
| clusterID | string | The UUID of the Kubernetes cluster to be deleted | **yes**  |

```golang
deleteClusterResponse, err := client.DeleteKubernetesCluster("b924d955-4d2e-4e0d-806a-4350ac1ec92x")

if err != nil {
    fmt.Printf("Error deleting Kubernetes cluster: %+v\n", err)
} else {
    fmt.Printf("Successfully deleted Kubernetes cluster: %+v\n", deleteClusterResponse)
}
```

### Read Kubeconfig

Retrieves the Kubeconfig (The kubectl configuration) as a string, for an existing cluster, so that it can be used remotely

The following table describes the request arguments:

| Name      | Type   | Description                                                     | Required |
| --------- | ------ | --------------------------------------------------------------- | -------- |
| clusterID | string | The UUID of the Kubernetes which you want the configuration for | **yes**  |

```golang
kubeConfig, err := client.GetKubeconfig("b924d955-4d2e-4e0d-806a-4350ac1ec92x")

if err != nil {
    fmt.Printf("Error retrieving kubeconfig: %+v\n", err)
} else {
    fmt.Printf("Successfully retrieved kubeconfig: %+v\n", kubeConfig)
}
```

### Create Node Pool

Creates a Kubernetes node pool for an existing Kubernetes Cluster

The following table describes the request arguments:

| Name      | Type                                                        | Description                                                            | Required |
| --------- | ----------------------------------------------------------- | ---------------------------------------------------------------------- | -------- |
| clusterID | string                                                      | The UUID of the Kubernetes cluster the node pool should be assigned to | **yes**  |
| nodePool  | \*[KubernetesNodePool](#KubernetesNodePool-Resource-Object) | An Kubernetes node pool resource object                                | **yes**  |

```golang
nodePool := profitbricks.KubernetesNodePool{
    Properties: &profitbricks.KubernetesNodePoolProperties{
        Name:             "demo-nodepool",
        DatacenterID:     existingDC,
        NodeCount:        3,
        CPUFamily:        "INTEL_XEON",
        CoresCount:       4,
        AvailabilityZone: "AUTO",
        RAMSize:          16384,
        StorageSize:      300,
        StorageType:      "SSD",
    },
}

createdNodePool, err := client.CreateKubernetesNodePool("b924d955-4d2e-4e0d-806a-4350ac1ec92x", nodePool)

if err != nil {
    fmt.Printf("Error creating Kubernetes node pool: %+v\n", err)
} else {
    fmt.Printf("Successfully created a Kubernetes node pool: %+v\n", createdNodePool)
}
```

### Read Node Pool

Retrieves an existing Kubernetes node pool

The following table describes the request arguments:

| Name       | Type   | Description                                                          | Required |
| ---------- | ------ | -------------------------------------------------------------------- | -------- |
| clusterID  | string | The UUID of the Kubernetes cluster the existing node pool belongs to | **yes**  |
| nodePoolID | string | The UUID of the Kubernetes node pool to be retrieved                 | **yes**  |

```golang
kubernetesNodePool, err := client.GetKubernetesNodePool("b924d955-4d2e-4e0d-806a-4350ac1ec92x", "d0ed6161-1e56-4db3-a094-08dddba2bc51")

if err != nil {
    fmt.Printf("Error reading Kubernetes node pool: %+v\n", err)
} else {
    fmt.Printf("Successfully read a Kubernetes node pool: %+v\n", kubernetesNodePool)
}
```

### List Node Pools

Retrieves all existing Kubernetes node pools found in an existing Kubernetes Cluster

The following table describes the request arguments:

| Name      | Type   | Description                                                       | Required |
| --------- | ------ | ----------------------------------------------------------------- | -------- |
| clusterID | string | The UUID of the Kubernetes cluster to retrieve the node pools for | **yes**  |

```golang
kubernetesNodePools, err := client.GetKubernetesNodePools("b924d955-4d2e-4e0d-806a-4350ac1ec92x")

if err != nil {
    fmt.Printf("Error retrieving Kubernetes node pools for cluster: %+v\n", err)
} else {
    fmt.Printf("Successfully retrieved Kubernetes node pools for cluster: %+v\n", kubernetesNodePools)
}
```

### Update Node Pool

Updates an existing Kubernetes Node Pool

The following table describes the request arguments:

| Name       | Type                                                        | Description                                                             | Required |
| ---------- | ----------------------------------------------------------- | ----------------------------------------------------------------------- | -------- |
| clusterID  | string                                                      | The UUID of the existingKubernetes cluster the node pool is assigned to | **yes**  |
| nodePoolID | string                                                      | The UUID of the Kubernetes node pool to be updated                      | **yes**  |
| nodePool   | \*[KubernetesNodePool](#KubernetesNodePool-Resource-Object) | An Kubernetes node pool resource object                                 | **yes**  |

_Note:_ Please Note that currently, the only property that can be updated for a Kubernetes node pool is the **NodeCount**

```golang
nodePoolUpdateRequest := profitbricks.KubernetesNodePool{
    Properties: &profitbricks.KubernetesNodePoolProperties{
        NodeCount: 3,
    },
}

updatedNodePool, err := client.UpdateKubernetesNodePool("b924d955-4d2e-4e0d-806a-4350ac1ec92x", "d0ed6161-1e56-4db3-a094-08dddba2bc51", nodePoolUpdateRequest)

if err != nil {
    fmt.Printf("Error updating Kubernetes node pool: %+v\n", err)
} else {
    fmt.Printf("Successfully updated Kubernetes node pool: %+v\n", updatedNodePool)
}
```

### Delete Node Pool

Deletes an existing Kubernetes node pool

The following table describes the request arguments:

| Name       | Type   | Description                                                          | Required |
| ---------- | ------ | -------------------------------------------------------------------- | -------- |
| clusterID  | string | The UUID of the Kubernetes cluster the existing node pool belongs to | **yes**  |
| nodePoolID | string | The UUID of the Kubernetes node pool to be deleted                   | **yes**  |

```golang
deleteNodePoolResponse, err := client.DeleteKubernetesNodePool("b924d955-4d2e-4e0d-806a-4350ac1ec92x", "d0ed6161-1e56-4db3-a094-08dddba2bc51")

if err != nil {
    fmt.Printf("Error deleting Kubernetes node pool: %+v\n", err)
} else {
    fmt.Printf("Successfully deleted Kubernetes node pool: %+v\n", deleteNodePoolResponse)
}
```

---

## Backup Units

##### BackupUnit Resource Object

| Name       | Required | Type                                                            | Description                       |
| ---------- | :------: | --------------------------------------------------------------- | --------------------------------- |
| Properties | **yes**  | \*[BackupUnitProperties](#BackupUnitProperties-resource-object) | The properties of the backup unit |

##### BackupUnitProperties Resource Object

| Name     | Required | Type   | Description                                                |
| -------- | :------: | ------ | ---------------------------------------------------------- |
| Name     | **yes**  | string | Alphanumeric name you want assigned to the backup unit     |
| Email    | **yes**  | string | Alphanumeric password you want assigned to the backup unit |
| Passowrd | **yes**  | string | The e-mail address you want assigned to the backup unit    |

### Create Backup Unit

Creates a new Backup Unit

The following table describes the request arguments:

| Name       | Type                                        | Description                                     | Required |
| ---------- | ------------------------------------------- | ----------------------------------------------- | -------- |
| backupUnit | \*[BackupUnit](#BackupUnit-resource-object) | a BackupUnit object with the desired properties | **yes**  |

```golang
backupUnit := profitbricks.BackupUnit{
			Properties: &profitbricks.BackupUnitProperties{
				Name:     "demobackup",
				Email:    "demo@demo.com",
				Password: "verystrongpassword",
			},
		}

createdBackupUnit, err := client.CreateBackupUnit(backupUnit)

if err != nil {
    fmt.Printf("Error creating backup unit: \n%+v\n", err)
} else {
    fmt.Printf("Successfully created backup unit: \n%+v\n", createdBackupUnit)
}
```

### Read Backup Unit

Retrieves an existing backup unit

The following table describes the request arguments:

| Name         | Type   | Description                                      | Required |
| ------------ | ------ | ------------------------------------------------ | -------- |
| backupUnitID | string | The UUID of the backup unit you want to retrieve | **yes**  |

```golang
backupUnit, err := client.GetBackupUnit("6e7cd5e2-d2ad-4ebc-ab35-985ded43ee3x")

if err != nil {
    fmt.Printf("Error reading backup unit: \n%+v\n", err)
} else {
    fmt.Printf("Successfully read backup unit: \n%+v\n", backupUnit)
}
```

### List Backup Units

Retrieves a list of all existing backup units associated with the current user account

```golang
backupUnits, err := client.ListBackupUnits()

if err != nil {
    fmt.Printf("Error listing backup units: \n%+v\n", err)
} else {
    fmt.Printf("Existing backup units: \n%+v\n", backupUnits)
}
```

### Update Backup Unit

Updates the email / password associated with an existing backup unit

The following table describes the request arguments:

| Name         | Type                                        | Description                                    | Required |
| ------------ | ------------------------------------------- | ---------------------------------------------- | -------- |
| backupUnitID | string                                      | The UUID of the backup unit you want to update | **yes**  |
| BackupUnit   | \*[BackupUnit](#BackupUnit-resource-object) | The updated BackupUnit object                  | **yes**  |

```golang
backupUnit := profitbricks.BackupUnit{
			Properties: &profitbricks.BackupUnitProperties{
				Name:     "demobackup",
				Email:    "updateddemo@demo.com",
				Password: "verystrongUPDATEDpassword",
			},
		}

updatedBackupUnit, err := client.UpdateBackupUnit("6e7cd5e2-d2ad-4ebc-ab35-985ded43ee3x", backupUnit)

if err != nil {
    fmt.Printf("Error updating backup unit: \n%+v\n", err)
} else {
    fmt.Printf("Successfully updated backup unit: \n%+v\n", updatedBackupUnit)
}
```

### Delete Backup Unit

Deletes an existing backup unit

The following table describes the request arguments:

| Name         | Type   | Description                                    | Required |
| ------------ | ------ | ---------------------------------------------- | -------- |
| backupUnitID | string | The UUID of the backup unit you want to delete | **yes**  |

```golang
backupUnitDeletionResult, err := client.DeleteBackupUnit("6e7cd5e2-d2ad-4ebc-ab35-985ded43ee3x")

if err != nil {
    fmt.Printf("Error deleting backup unit: \n%+v\n", err)
} else {
    fmt.Printf("Successfully deleted backup unit: \n%+v\n", backupUnitDeletionResult)
}
```

### Get Backup Unit SSO URL

Retrieves the SSO URL for an existing backup unit

The following table describes the request arguments:

| Name         | Type   | Description                                                            | Required |
| ------------ | ------ | ---------------------------------------------------------------------- | -------- |
| backupUnitID | string | The UUID of the backup unit for which you want to retrieve the SSO URL | **yes**  |

```golang
backupUnitSSOURL, err := client.DeleteBackupUnit("6e7cd5e2-d2ad-4ebc-ab35-985ded43ee3x")

if err != nil {
    fmt.Printf("Error retrieving backup unit SSO URL: \n%+v\n", err)
} else {
    fmt.Printf("Successfully retrieved the backup unit's SSO URL: \n%+v\n", backupUnitDeletionResult)
}
```

---

## User S3 Keys

##### S3Key Resource Object

| Name       | Required | Type                                                  | Description                  |
| ---------- | :------: | ----------------------------------------------------- | ---------------------------- |
| Properties | **yes**  | \*[S3KeyProperties](#S3KeyProperties-resource-object) | The properties of the S3 key |

##### S3KeyProperties Resource Object

| Name   | Required | Type   | Description                                   |
| ------ | :------: | ------ | --------------------------------------------- |
| S3Key  |  **no**  | string | The S3 secret key                             |
| Active | **yes**  | bool   | Indicates whether the S3 key is active or not |

### Create S3 Key

Creates a new S3 Key

The following table describes the request arguments:

| Name   | Type   | Description                                        | Required |
| ------ | ------ | -------------------------------------------------- | -------- |
| userID | string | The UUID of the user the key should be assigned to | **yes**  |

```golang

createdS3Key, err := client.CreateS3Key("d4246339-9c0f-41fb-a96d-1c640ae4501x")

if err != nil {
    fmt.Printf("Error creating S3 Key: \n%+v", err)
} else {
    fmt.Printf("Successfully created S3 Key: \n%+v", createdS3Key)
}

```

### Read S3 Key

Retrieves an existing S3 Key from the API

The following table describes the request arguments:

| Name    | Type   | Description                             | Required |
| ------- | ------ | --------------------------------------- | -------- |
| userID  | string | The UUID of the user the key belongs to | **yes**  |
| s3KeyID | string | The UUID of the existing S3 key         | **yes**  |

```golang
s3Key, err := client.GetS3Key("d4246339-9c0f-41fb-a96d-1c640ae4501x", "00ddbdff682631d4c0f8")

if err != nil {
    fmt.Printf("Error retrieving S3 Key: \n%+v", err)
} else {
    fmt.Printf("Existing S3 Key: \n%+v", s3Key)
}
```

### List S3 Keys

Retrieves the existing S3 Keys for an user

The following table describes the request arguments:

| Name   | Type   | Description                                  | Required |
| ------ | ------ | -------------------------------------------- | -------- |
| userID | string | The UUID of the user to list the S3 keys for | **yes**  |

```golang

s3Keys, err := client.ListS3Keys("d4246339-9c0f-41fb-a96d-1c640ae4501x")

if err != nil {
    fmt.Printf("Error retrieving S3 Keys: \n%+v", err)
} else {
    fmt.Printf("Successfully retrieved S3 Keys: \n%+v", s3Keys)
}

```

### Update S3 Key

Updates the state of an existing S3 Key

The following table describes the request arguments:

| Name    | Type                              | Description                                | Required |
| ------- | --------------------------------- | ------------------------------------------ | -------- |
| userID  | string                            | The UUID of the user the S3 key belongs to | **yes**  |
| s3KeyID | string                            | The UUID of the existing S3 key            | **yes**  |
| s3Key   | \*[S3Key](#S3Key-resource-object) | The updated S3 Key Object                  | **yes**  |

```golang

s3Key := profitbricks.S3Key{
    Properties: &profitbricks.S3KeyProperties{
        Active: false,
    },
}

updatedS3Key, err := client.UpdateS3Key("d4246339-9c0f-41fb-a96d-1c640ae4501x", "00ddbdff682631d4c0f8", s3Key)

if err != nil {
    fmt.Printf("Error updating S3 Key: \n%+v\n", err)
} else {
    fmt.Printf("Successfully updated S3 Key: \n%+v\n", updatedS3Key)
}

```

### Delete S3 Key

Deletes an existing S3 Key

The following table describes the request arguments:

| Name    | Type   | Description                                | Required |
| ------- | ------ | ------------------------------------------ | -------- |
| userID  | string | The UUID of the user the S3 key belongs to | **yes**  |
| s3KeyID | string | The UUID of the existing S3 key            | **yes**  |

```golang

deletedS3KeyResponse, err := client.DeleteS3Key("d4246339-9c0f-41fb-a96d-1c640ae4501a", "00ddbdff682631d4c0f8")

if err != nil {
    fmt.Printf("Error deleting S3 Key: \n%+v\n", err)
} else {
    fmt.Printf("Successfully deleted S3 Key: \n%+v\n", deletedS3KeyResponse)
}

```

---

## Private Cross Connects

**NOTE:** Please note that currently, only VDC's created in the same physical location can be cross-connected

##### PrivateCrossConnect Resource Object

| Name       | Required | Type                                                                              | Description                             |
| ---------- | :------: | --------------------------------------------------------------------------------- | --------------------------------------- |
| Properties | **yes**  | \*[PrivateCrossConnectProperties](#PrivateCrossConnectProperties-resource-object) | The properties of private cross connect |

##### PrivateCrossConnectProperties Resource Object

| Name                   | Required | Type                                                                      | Description                                                       |
| ---------------------- | :------: | ------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| Name                   | **yes**  | string                                                                    | A humanly readable name for the private cross-connect             |
| Description            |  **no**  | string                                                                    | A short description for the private cross-connect                 |
| Peers                  |  **no**  | \*[][pccpeer](#PCCPeer-resource-object)                                   | A read-only property containing all the currentl peers of the PCC |
| ConnectableDatacenters |  **no**  | \*[][pccconnectabledatacenter](#PCCConnectableDataCenter-resource-object) | A read-only property containing all the connectable VDC's         |

##### PCCPeer Resource Object

| Name           | Required | Type   | Description                                                                                          |
| -------------- | :------: | ------ | ---------------------------------------------------------------------------------------------------- |
| LANId          |  **no**  | string | The id of the cross-connected LAN                                                                    |
| LANName        |  **no**  | string | The name of the cross-connected LAN                                                                  |
| DataCenterID   |  **no**  | string | The id of the cross-connected datacenter                                                             |
| DataCenterName |  **no**  | string | THe name of the cross-connected datacenter                                                           |
| Location       |  **no**  | string | The physical location of the cross-connected datacenter (Please see [Locations](#locations) section) |

##### PCCConnectableDataCenter Resource Object

| Name     | Required | Type   | Description                                                                                      |
| -------- | :------: | ------ | ------------------------------------------------------------------------------------------------ |
| ID       |  **no**  | string | The UUID of the connectable datacenter                                                           |
| Name     |  **no**  | string | The name of the connectable datacenter                                                           |
| Location |  **no**  | string | The physical location of the connectable datacenter (Please see [Locations](#locations) section) |

### Create Private Cross Connect

Creates a new, empty private cross-connect

The following table describes the request arguments:

| Name | Type                                                          | Description                                    | Required |
| ---- | ------------------------------------------------------------- | ---------------------------------------------- | -------- |
| pcc  | \*[PrivateCrossConnect](#PrivateCrossConnect-resource-object) | The desired name for the private cross-connect | **yes**  |

```golang

pcc := profitbricks.PrivateCrossConnect{
    Properties: &profitbricks.PrivateCrossConnectProperties{
        Name: "demo",
        Description: "description for demo",
    },
}

createdPCC, err := client.CreatePrivateCrossConnect(pcc)

if err != nil {
    fmt.Printf("Error creating private cross-connect: \n%+v", err)
} else {
    fmt.Printf("Successfully created private cross-connect: \n%+v", createdPCC)
}

```

### Read Private Cross Connect

Retrieves an existing private cross-connect from the API

The following table describes the request arguments:

| Name  | Type   | Description                       | Required |
| ----- | ------ | --------------------------------- | -------- |
| pccID | string | The UUID of private cross-connect | **yes**  |

```golang
pccID := "d4246339-9c0f-41fb-a96d-1c640ae4501x"

pcc, err := client.GetPrivateCrossConnect(pccID)

if err != nil {
    fmt.Printf("Error retrieving private cross-connect: \n%+v", err)
} else {
    fmt.Printf("Existing private cross-connect: \n%+v", pcc)
}
```

### List Private Cross Connects

Retrieves the existing private cross-connects

```golang

pccs, err := client.ListPrivateCrossConnects()

if err != nil {
    fmt.Printf("Error retrieving private cross-connects: \n%+v", err)
} else {
    fmt.Printf("Successfully retrieved private cross-connects: \n%+v", pccs)
}

```

### Update Private Cross Connect

Updates the state of an existing Private cross-connect

The following table describes the request arguments:

| Name  | Type                                                          | Description                                              | Required |
| ----- | ------------------------------------------------------------- | -------------------------------------------------------- | -------- |
| pccID | string                                                        | The UUID of the private cross-connect you wish to update | **yes**  |
| pcc   | \*[PrivateCrossConnect](#PrivateCrossConnect-resource-object) | The desired name for the private cross-connect           | **yes**  |

```golang

pccID := "d4246339-9c0f-41fb-a96d-1c640ae4501x"

pcc := profitbricks.PrivateCrossConnect{
    Properties: &profitbricks.PrivateCrossConnectProperties{
        Name: "demo-renamed",
        Description: "updated description for demo-renamed",
    },
}

updatedPCC, err := client.UpdatePrivateCrossConnect(pccID, pcc)

if err != nil {
    fmt.Printf("Error updating private cross-connect: \n%+v", err)
} else {
    fmt.Printf("Successfully updated private cross-connect: \n%+v", updatedPCC)
}

```

### Delete Private Cross Connect

Deletes an existing private cross-connect

The following table describes the request arguments:

| Name  | Type   | Description                                              | Required |
| ----- | ------ | -------------------------------------------------------- | -------- |
| pccID | string | The UUID of the private cross-connect you wish to delete | **yes**  |

```golang

pccID := "d4246339-9c0f-41fb-a96d-1c640ae4501x"

deletedPCCResponse, err := client.DeletePrivateCrossConnect(pccID)

if err != nil {
    fmt.Printf("Error deleting private cross-connect: \n%+v\n", err)
} else {
    fmt.Printf("Successfully deleted private cross-connect: \n%+v\n", deletedPCCResponse)
}

```

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/profitbricks/profitbricks-sdk-go/v5"
)

func main() {

	//Sets username and password
	client := profitbricks.NewClient(
        os.Getenv("PROFITBRICKS_USERNAME"),
        os.Getenv("PROFITBRICKS_PASSWORD"),
    )

	dcrequest := profitbricks.Datacenter{
		Properties: profitbricks.DatacenterProperties{
			Name:        "Example",
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
```

## Support

Please report any issues via [GitHub's issue tracker](https://github.com/profitbricks/profitbricks-sdk-go/issues).

## Testing

You can run all test by using the command `go test -timeout=120m` or run a single test by specifying the name of the test file `go test servers_test.go`

## Contributing

1. Fork it ( https://github.com/profitbricks/profitbricks-sdk-go/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
