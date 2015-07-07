# goprofitbricks

#### ```Install go```
      https://golang.org/doc/install

#### ```Set your Environment```
```
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

#### ``` Fetch goprofitbricks```
```go
go get "github.com/StackPointCloud/profitbricks-sdk-go"
```
####  ```Test stuff```
```go
cd $GOPATH/src/github.com/StackPointCloud/profitbricks-sdk-go
```
```go
vi config_test.go
```
* Set Username, Password, and Endpoint for testing

```go
go test -v 
```
* runs all the tests and reports pass/fail 

#### ```Use```
```
cd ~/
vi testrest.go
```
```go
package  main
import 	"github.com/StackPointCloud/profitbricks-sdk-go"

import "fmt"	

func main() {

	goprofitbricks.SetAuth("your_username","your_password")
	/**
	 List Datacenter returns a collection of (Datacenter) Instances
	See file resp.go for Instance struct
	**/
	
	resp:=goprofitbricks.ListDatacenters()

	//get the Id of the first Item in the collection
	dc := goprofitbricks.GetDatacenter(resp.Items[0].Id)

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

	}
	
```

###### ```Run```
```go 
	go run  /root/testrest.go
```



## PACKAGE DOCUMENTATION

```go
package goprofitbricks
    import "github.com/StackPointCloud/profitbricks-sdk-go"
```

#### ```Variables```
```go
var Depth string
```
Depth controls the amount of data returned from the rest server ( range 1-5 )
```go
var Endpoint string
```
* Endpoint is the base url for REST requests .
```go 
var Passwd string
```
* Password for authentication .
```go
var Username string
```
* Username for authentication .

#### ```Functions```
```go
func MkJson(i interface{}) string
```
*  Turn just about anything into Json

```go
func SetAuth(u, p string)
```
```go
func SetDepth(newdepth string) string
```

```go
func SetEndpoint(newendpoint string) string
```
*  SetEndpoint is used to set the REST Endpoint. 

### ```Resp struct```
* 	Resp is the struct returned by all Rest request functions

```go
type Resp struct {
Req        *http.Request
StatusCode int
Headers    http.Header
Body       []byte
}
```
###### Resp methods
```go
	func (r *Resp) PrintHeaders()
```
* 	PrintHeaders prints the http headers as k,v pairs

### ```Id_Type_Href struct```

* 	The Id_Type_Href struct is embedded in Instance structs and Collection structs
```go 
type Id_Type_Href struct {
Id   string `json:"id"`
Type string `json:"type"`
Href string `json:"href"`
}
```

### ```Instance struct```
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


### ```Collection struct``` 
* 	Collection Structs contain Instance arrays. 
* 	List functions return Collections

```go
type Collection struct {
Id_Type_Href
Items []Instance `json:"items,omitempty"`
Resp  Resp       `json:"-"`
}
```


## ```Functions by target```

### ```Datacenter```
```go
 func ListDatacenters() Collection  
```
```go
 func CreateDatacenter(jason []byte) Instance  
```
```go
 func GetDatacenter(dcid string) Instance  
```
```go
 func PatchDatacenter(dcid string, jason []byte) Instance  
```
```go
 func DeleteDatacenter(dcid string) Resp  
```

###  ```Server```
```go
 func ListServers(dcid string) Collection  
```
```go
 func CreateServer(dcid string, jason []byte) Instance  
```
```go
 func GetServer(dcid, srvid string) Instance  
```
```go
 func PatchServer(dcid string, srvid string, jason []byte) Instance  
```
```go
 func DeleteServer(dcid, srvid string) Resp  
```
##### ``` Server Attached Cdroms```
```go
 func ListAttachedCdroms(dcid, srvid string) Collection  
```
```go
 func AttachCdrom(dcid string, srvid string, cdid string) Instance  
```
```go
 func GetAttachedCdrom(dcid, srvid, cdid string) Instance  
```
```go
 func DetachCdrom(dcid, srvid, cdid string) Resp  
```
##### ```Server Attached Volumes```
```go
 func ListAttachedVolumes(dcid, srvid string) Collection  
```
```go
 func AttachVolume(dcid string, srvid string, volid string) Instance  
```
```go
 func GetAttachedVolume(dcid, srvid, volid string) Instance  
```
```go
 func DetachVolume(dcid, srvid, volid string) Resp  
```
```go
 func StartServer(dcid, srvid string) Resp  
```
```go
 func StopServer(dcid, srvid string) Resp  
```
```go
 func RebootServer(dcid, srvid string) Resp  
```
### ```Nics```

```go
 func ListNics(dcid, srvid string) Collection  
```
```go
 func CreateNic(dcid string, srvid string, jason []byte) Instance  
```
```go
 func GetNic(dcid, srvid, nicid string) Instance  
```
```go
 func PatchNic(dcid string, srvid string, nicid string, jason []byte) Instance  
```
```go
 func DeleteNic(dcid, srvid, nicid string) Resp  
```

### ```Firewall Rules```
```go
 func ListFwRules(dcid, srvid, nicid string) Collection  
```
```go
 func CreateFwRule(dcid string, srvid string, nicid string, jason []byte) Instance  
```
```go
 func GetFwRule(dcid, srvid, nicid, fwruleid string) Instance  
```
```go
 func PatchFWRule(dcid string, srvid string, nicid string, fwruleid string, jason []byte) Instance  
```
```go
 func DeleteFWRule(dcid, srvid, nicid, fwruleid string) Resp  
```

### ```Images```

```go
 func ListImages() Collection  
```
```go
 func GetImage(imageid string) Instance  
```
```go
 func PatchImage(imageid string, jason []byte) Instance  
```
```go
 func DeleteImage(imageid string) Resp  
```

### ```Volumes```

```go
 func ListVolumes(dcid string) Collection   
```
```go
func GetVolume(dcid string, volumeId string) Instance 
```
```go
func PatchVolume(dcid string, volid string, request VolumeProperties) Instance 
```
```go
func CreateVolume(dcid string, request CreateVolumeRequest) Instance 
```
```go
func DeleteVolume(dcid, volid string) Resp 
```
```go
func CreateSnapshot(dcid string, volid string, jason []byte) Resp
```
### ```Load Balancers```

```go
 func ListLoadbalancers(dcid string) Collection    
```
```go
 func GetLoadbalancer(dcid, lbalid string) Instance     
```
```go
func CreateLoadbalancer(dcid string, request LoablanacerCreateRequest) Instance
```
```go
func PatchLoadbalancer(dcid string, lbalid string, obj map[string]string) Instance 
```
```go
func DeleteLoadbalancer(dcid, lbalid string) Resp 
```
```go
func ListBalancedNics(dcid, lbalid string) Collection  
```
```go
func GetBalancedNic(dcid, lbalid, balnicid string) Instance 
```
```go
func AssociateNic(dcid string, lbalid string, nicid string) Instance
```
```go
func DeleteBalancedNic(dcid, lbalid, balnicid string) Resp 
```
### ```IP Blocks```

```go
func ListIpBlocks() Collection 
```
```go
func GetIpBlock(ipblockid string) Instance 
```
```go
func ReserveIpBlock(request IPBlockReserveRequest) Instance
```
```go
func ReleaseIpBlock(ipblockid string) Resp
```

### ```Note```

Details about object propreties are located [here](https://devops.profitbricks.com/api/rest/)