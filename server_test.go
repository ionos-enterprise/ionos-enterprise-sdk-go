package profitbricks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type ApiSuite struct {
	suite.Suite
	client *Client
}

func (s *ApiSuite) SetupTest() {
	s.client = RestyClient("", "", "lame-token")
	s.client.SetDoNotParseResponse(false)
	s.client.SetDebug(os.Getenv("RESTY_DEBUG") == "true")
	httpmock.ActivateNonDefault(s.client.Client.GetClient())
}

func (s *ApiSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

type SuiteServer struct {
	ApiSuite
}

func Test_Server(t *testing.T) {
	suite.Run(t, new(SuiteServer))
}

func pstr(s string) *string { return &s }
func pint(i int) *int       { return &i }

var SSHRule = FirewallRule{
	Properties: FirewallruleProperties{
		Name:           "SSH",
		Protocol:       "TCP",
		SourceMac:      pstr("00:11:22:33:44:55"),
		PortRangeStart: pint(22),
		PortRangeEnd:   pint(22),
	},
}
var DefaultFWRules = &FirewallRules{
	Items: []FirewallRule{
		SSHRule,
	},
}

var DefaultDatacenterID = "f23f140c-21fc-40c8-a263-43f4fe14a1a9"
var DefaultServerID = "87356098-fb0e-4c7c-bb3b-b8b535202dd3"

var DefaultCompServerRespBody = `{
   "id": "87356098-fb0e-4c7c-bb3b-b8b535202dd3",
   "type": "server",
   "href": "https://api.ionos.com/cloudapi/v5/datacenters/f23f140c-21fc-40c8-a263-43f4fe14a1a9/servers/87356098-fb0e-4c7c-bb3b-b8b535202dd3",
   "metadata": {
      "etag": "77f5fdac710dddd31c5fe27657c63cb0",
      "createdDate": "2019-06-11T19:12:41Z",
      "createdBy": "schulze-wiehenbrauk@strato-rz.de",
      "createdByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
      "lastModifiedDate": "2019-06-11T19:12:41Z",
      "lastModifiedBy": "schulze-wiehenbrauk@strato-rz.de",
      "lastModifiedByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
      "state": "BUSY"
   },
   "properties": {
      "name": "GO SDK Test",
      "cores": 1,
      "ram": 1024,
      "availabilityZone": "ZONE_1",
      "vmState": null,
      "bootCdrom": null,
      "bootVolume": null,
      "cpuFamily": "INTEL_XEON"
   },
   "entities": {
      "volumes": {
         "id": "87356098-fb0e-4c7c-bb3b-b8b535202dd3/volumes",
         "type": "collection",
         "href": "https://api.ionos.com/cloudapi/v5/datacenters/f23f140c-21fc-40c8-a263-43f4fe14a1a9/servers/87356098-fb0e-4c7c-bb3b-b8b535202dd3/volumes",
         "items": [
            {
               "id": "7dc59b42-32ca-4bd9-b584-4381c49a7947",
               "type": "volume",
               "href": "https://api.ionos.com/cloudapi/v5/datacenters/87356098-fb0e-4c7c-bb3b-b8b535202dd3/volumes/7dc59b42-32ca-4bd9-b584-4381c49a7947",
               "metadata": {
                  "etag": "604bf48b8572948be3da2e1fa78bf5cd",
                  "createdDate": "2019-06-11T19:12:41Z",
                  "createdBy": "schulze-wiehenbrauk@strato-rz.de",
                  "createdByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
                  "lastModifiedDate": "2019-06-11T19:12:41Z",
                  "lastModifiedBy": "schulze-wiehenbrauk@strato-rz.de",
                  "lastModifiedByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
                  "state": "BUSY"
               },
               "properties": {
                  "name": "volume1",
                  "type": "HDD",
                  "size": 5,
                  "availabilityZone": null,
                  "image": "d322efa7-8421-11e9-84a0-525400f64d8d",
                  "imagePassword": "JWXuXR9CMghXAc6v",
                  "sshKeys": null,
                  "bus": null,
                  "licenceType": null,
                  "cpuHotPlug": false,
                  "ramHotPlug": false,
                  "nicHotPlug": false,
                  "nicHotUnplug": false,
                  "discVirtioHotPlug": false,
                  "discVirtioHotUnplug": false,
                  "deviceNumber": null
               }
            }
         ]
      },
      "nics": {
         "id": "87356098-fb0e-4c7c-bb3b-b8b535202dd3/nics",
         "type": "collection",
         "href": "https://api.ionos.com/cloudapi/v5/datacenters/f23f140c-21fc-40c8-a263-43f4fe14a1a9/servers/87356098-fb0e-4c7c-bb3b-b8b535202dd3/nics",
         "items": [
            {
               "id": "d6c95e5c-50b6-4ceb-bebf-7f9bfdda3e9f",
               "type": "nic",
               "href": "https://api.ionos.com/cloudapi/v5/datacenters/f23f140c-21fc-40c8-a263-43f4fe14a1a9/servers/87356098-fb0e-4c7c-bb3b-b8b535202dd3/nics/d6c95e5c-50b6-4ceb-bebf-7f9bfdda3e9f",
               "metadata": {
                  "etag": "109eca8c019d76648e9f9fa05f8c5381",
                  "createdDate": "2019-06-11T19:12:41Z",
                  "createdBy": "schulze-wiehenbrauk@strato-rz.de",
                  "createdByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
                  "lastModifiedDate": "2019-06-11T19:12:41Z",
                  "lastModifiedBy": "schulze-wiehenbrauk@strato-rz.de",
                  "lastModifiedByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
                  "state": "BUSY"
               },
               "properties": {
                  "name": "nic",
                  "mac": null,
                  "ips": [],
                  "dhcp": null,
                  "lan": 1,
                  "firewallActive": null,
                  "nat": null
               },
               "entities": {
                  "firewallrules": {
                     "id": "d6c95e5c-50b6-4ceb-bebf-7f9bfdda3e9f/firewallrules",
                     "type": "collection",
                     "href": "https://api.ionos.com/cloudapi/v5/datacenters/f23f140c-21fc-40c8-a263-43f4fe14a1a9/servers/87356098-fb0e-4c7c-bb3b-b8b535202dd3/nics/d6c95e5c-50b6-4ceb-bebf-7f9bfdda3e9f/firewallrules",
                     "items": [
                        {
                           "id": "3626478b-4b0c-457b-968d-ded3234cc781",
                           "type": "firewall-rule",
                           "href": "https://api.ionos.com/cloudapi/v5/datacenters/f23f140c-21fc-40c8-a263-43f4fe14a1a9/servers/87356098-fb0e-4c7c-bb3b-b8b535202dd3/nics/d6c95e5c-50b6-4ceb-bebf-7f9bfdda3e9f/firewallrules/3626478b-4b0c-457b-968d-ded3234cc781",
                           "metadata": {
                              "etag": "54451d8daf6aab67154fbc2184e2fdb3",
                              "createdDate": "2019-06-11T19:12:41Z",
                              "createdBy": "schulze-wiehenbrauk@strato-rz.de",
                              "createdByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
                              "lastModifiedDate": "2019-06-11T19:12:41Z",
                              "lastModifiedBy": "schulze-wiehenbrauk@strato-rz.de",
                              "lastModifiedByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
                              "state": "BUSY"
                           },
                           "properties": {
                              "name": "SSH",
                              "protocol": "TCP",
                              "sourceMac": "01:23:45:67:89:00",
                              "sourceIp": null,
                              "targetIp": null,
                              "icmpCode": null,
                              "icmpType": null,
                              "portRangeStart": 22,
                              "portRangeEnd": 22
                           }
                        }
                     ]
                  }
               }
            }
         ]
      }
   }
}`

func CreateResponder(method, url, body string, header http.Header, status int, expected interface{}) error {
	bodyReader := ioutil.NopCloser(bytes.NewBufferString(body))
	r := &http.Response{
		StatusCode: status,
		Body:       bodyReader,
		Header:     header,
	}
	if len(body) > 0 {
		r.Header.Set("Content-Type", "application/json")
		err := json.Unmarshal([]byte(body), expected)
		if err != nil {
			return err
		}
	}
	responder := httpmock.ResponderFromResponse(r)
	httpmock.RegisterResponder(method, url, responder)
	return nil

}

/*
func (s *ApiSuite) verifyBaseResource(exp BaseResource, got BaseResource) {
	for k := range *exp.Headers {
		s.Equal(exp.Headers.Get(k), got.Headers.Get(k))

	}
}

*/

func (s *SuiteServer) SetupTest() {
	s.ApiSuite.SetupTest()

}
func (s *SuiteServer) verifyServer(exp *Server, got *Server) {
	//	s.verifyBaseResource(exp.BaseResource, got.BaseResource)
	s.Equal(exp.Properties, got.Properties)
	s.Equal(exp.Metadata, got.Metadata)
	s.Equal(exp.Entities, got.Entities)
	s.Equal(exp.ID, got.ID)
	s.Equal(exp.PBType, got.PBType)
	s.Equal("status", got.Headers.Get("location"))
}
func (s *SuiteServer) Test_CreateServer() {
	obj := Server{
		Properties: ServerProperties{
			Name:             "GO SDK Test",
			RAM:              1024,
			Cores:            1,
			AvailabilityZone: "ZONE_1",
			CPUFamily:        "INTEL_XEON",
		},
		Entities: &ServerEntities{
			Volumes: &Volumes{
				Items: []Volume{
					{
						Properties: VolumeProperties{
							Type:          "HDD",
							Size:          5,
							Name:          "volume1",
							ImageAlias:    "ubuntu:latest",
							ImagePassword: "JWXuXR9CMghXAc6v",
						},
					},
				},
			},
			Nics: &Nics{
				Items: []Nic{
					{
						Properties: &NicProperties{
							Name: "nic",
							Lan:  1,
						},
						Entities: &NicEntities{
							FirewallRules: DefaultFWRules,
						},
					},
				},
			},
		},
	}
	exp := &Server{}
	url := DefaultApiUrl + "/datacenters/f23f140c-21fc-40c8-a263-43f4fe14a1a9/servers"
	err := CreateResponder("POST", url, DefaultCompServerRespBody, http.Header{"Location": []string{"status"}}, 202, exp)
	s.NoError(err)
	got, err := s.client.CreateServer(DefaultDatacenterID, obj)

	if s.NoError(err) && s.NotNil(got) {
		s.verifyServer(exp, got)
		s.Equal("status", got.Headers.Get("location"))
	}
}

func (s *SuiteServer) Test_CreateServer_500() {
	obj := Server{
		Properties: ServerProperties{
			Name:             "GO SDK Test",
			RAM:              1024,
			Cores:            1,
			AvailabilityZone: "ZONE_1",
			CPUFamily:        "INTEL_XEON",
		},
		Entities: &ServerEntities{
			Volumes: &Volumes{
				Items: []Volume{
					{
						Properties: VolumeProperties{
							Type:          "HDD",
							Size:          5,
							Name:          "volume1",
							ImageAlias:    "ubuntu:latest",
							ImagePassword: "JWXuXR9CMghXAc6v",
						},
					},
				},
			},
			Nics: &Nics{
				Items: []Nic{
					{
						Properties: &NicProperties{
							Name: "nic",
							Lan:  1,
						},
						Entities: &NicEntities{
							FirewallRules: DefaultFWRules,
						},
					},
				},
			},
		},
	}
	exp := &Server{}
	url := DefaultApiUrl + "/datacenters/f23f140c-21fc-40c8-a263-43f4fe14a1a9/servers"
	body := `{"httpStatus": 500, "messages": [{"errorCode": "301", "message": "Oops! Something went very wrong."}]}`
	err := CreateResponder("POST", url, body, http.Header{"Location": []string{"status"}}, 500, exp)
	s.NoError(err)
	got, err := s.client.CreateServer(DefaultDatacenterID, obj)

	s.Error(err)
	s.Empty(got.ID)
	/*	if s.NoError(err) && s.NotNil(got) {
			s.verifyServer(exp, got)
			s.Equal("status", got.Headers.Get("location"))
		}
	*/
}

func (s *SuiteServer) Test_StopServer() {
	url := fmt.Sprintf("%s/datacenters/%s/servers/%s/stop", DefaultApiUrl, DefaultDatacenterID, DefaultServerID)
	err := CreateResponder("POST", url, "", http.Header{"Location": []string{"status"}}, 202, nil)
	s.NoError(err)
	rsp, err := s.client.StopServer(DefaultDatacenterID, DefaultServerID)
	s.NoError(err)
	s.Equal("status", rsp.Get("Location"))
}
