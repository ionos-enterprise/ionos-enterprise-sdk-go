package profitbricks

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var fixt = `{
  "entities": {
    "lans": {
      "href": "https://api.ionos.com/cloudapi/v5/datacenters/00cf9d60-8c53-4722-9486-fc7ed0dc35be/lans",
      "id": "00cf9d60-8c53-4722-9486-fc7ed0dc35be/lans",
      "items": [],
      "type": "collection"
    },
    "loadbalancers": {
      "href": "https://api.ionos.com/cloudapi/v5/datacenters/00cf9d60-8c53-4722-9486-fc7ed0dc35be/loadbalancers",
      "id": "00cf9d60-8c53-4722-9486-fc7ed0dc35be/loadbalancers",
      "items": [],
      "type": "collection"
    },
    "servers": {
      "href": "https://api.ionos.com/cloudapi/v5/datacenters/00cf9d60-8c53-4722-9486-fc7ed0dc35be/servers",
      "id": "00cf9d60-8c53-4722-9486-fc7ed0dc35be/servers",
      "items": [],
      "type": "collection"
    },
    "volumes": {
      "href": "https://api.ionos.com/cloudapi/v5/datacenters/00cf9d60-8c53-4722-9486-fc7ed0dc35be/volumes",
      "id": "00cf9d60-8c53-4722-9486-fc7ed0dc35be/volumes",
      "items": [],
      "type": "collection"
    }
  },
  "href": "https://api.ionos.com/cloudapi/v5/datacenters/00cf9d60-8c53-4722-9486-fc7ed0dc35be",
  "id": "00cf9d60-8c53-4722-9486-fc7ed0dc35be",
  "metadata": {
    "createdBy": "schulze-wiehenbrauk@strato-rz.de",
    "createdByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
    "createdDate": "2019-06-11T16:32:37.000Z",
    "etag": "981feaddd39a1677c150d296f9a79ea3",
    "lastModifiedBy": "schulze-wiehenbrauk@strato-rz.de",
    "lastModifiedByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
    "lastModifiedDate": "2019-06-11T16:32:37.000Z",
    "state": "AVAILABLE"
  },
  "properties": {
    "description": "GO SDK test datacenter",
    "features": [
      "SSD"
    ],
    "location": "us/las",
    "name": "GO SDK Test",
    "version": 1
  },
  "type": "datacenter"
}`

var responseBody = `
{
   "id": "687d86c1-0063-4327-8f58-ccc4afb00e92",
   "type": "datacenter",
   "href": "https://api.ionos.com/cloudapi/v5/datacenters/687d86c1-0063-4327-8f58-ccc4afb00e92",
   "metadata": {
      "etag": "1df58452e22040e628257226878ee4a5",
      "createdDate": "2019-06-11T17:39:39Z",
      "createdBy": "schulze-wiehenbrauk@strato-rz.de",
      "createdByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
      "lastModifiedDate": "2019-06-11T17:39:39Z",
      "lastModifiedBy": "schulze-wiehenbrauk@strato-rz.de",
      "lastModifiedByUserId": "8e22b57d-874f-4e13-8ed7-6851c80f2aad",
      "state": "BUSY"
   },
   "properties": {
      "name": "GO SDK Test Composite",
      "description": "GO SDK test composite datacenter",
      "location": "us/las",
      "version": null,
      "features": [],
      "secAuthProtection": null
   },
   "entities": {
      "servers": {
         "id": "687d86c1-0063-4327-8f58-ccc4afb00e92/servers",
         "type": "collection",
         "href": "https://api.ionos.com/cloudapi/v5/datacenters/687d86c1-0063-4327-8f58-ccc4afb00e92/servers"
      },
      "volumes": {
         "id": "687d86c1-0063-4327-8f58-ccc4afb00e92/volumes",
         "type": "collection",
         "href": "https://api.ionos.com/cloudapi/v5/datacenters/687d86c1-0063-4327-8f58-ccc4afb00e92/volumes"
      }
   }
}


`
var DefaultLocation = "us/las"
var _ = Describe("Datacenters", func() {
	BeforeEach(func() {
		httpmock.Reset()
	})
	It("should return a valid response on POST", func() {
		dc := &Datacenter{}
		var obj = Datacenter{
			Properties: DatacenterProperties{
				Name:        "GO SDK Test Composite",
				Description: "GO SDK test composite datacenter",
				Location:    DefaultLocation,
			},
			Entities: DatacenterEntities{
				Servers: &Servers{
					Items: []Server{
						{
							Properties: ServerProperties{
								Name:  "GO SDK Test",
								RAM:   1024,
								Cores: 1,
							},
						},
					},
				},
				Volumes: &Volumes{
					Items: []Volume{
						{
							Properties: VolumeProperties{
								Type:             "HDD",
								Size:             2,
								Name:             "GO SDK Test",
								Bus:              "VIRTIO",
								LicenceType:      "UNKNOWN",
								AvailabilityZone: "ZONE_3",
							},
						},
					},
				},
			},
		}
		err := json.Unmarshal([]byte(responseBody), dc)
		Expect(err).NotTo(HaveOccurred())
		location := "http://requests/status"
		x := &http.Response{
			Status:     "202",
			StatusCode: 202,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(responseBody))),
			Header: http.Header{
				"Location":     []string{location},
				"Content-Type": []string{"application/json"},
			},
			ContentLength: -1,
		}
		responder := httpmock.ResponderFromResponse(x)
		fakeUrl := DefaultApiUrl + "/datacenters"
		httpmock.RegisterResponder("POST", fakeUrl, responder)

		// fetch the article into struct
		got, err := client.CreateDatacenter(obj)
		Expect(err).NotTo(HaveOccurred())
		Expect(got).NotTo(BeNil())
		Expect(got.ID).To(Equal(dc.ID), "dc id should be correct")
		Expect(got.Headers.Get("Location")).To(Equal(location))

	})
	It("should be gettable", func() {
		exp := &Datacenter{}
		err := json.Unmarshal([]byte(responseBody), exp)
		x := &http.Response{
			Status:     "200",
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(responseBody))),
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			ContentLength: -1,
		}
		responder := httpmock.ResponderFromResponse(x)
		fakeUrl := DefaultApiUrl + "/datacenters/" + exp.ID
		httpmock.RegisterResponder("GET", fakeUrl, responder)

		// fetch the article into struct
		got, err := client.GetDatacenter(exp.ID)
		Expect(err).NotTo(HaveOccurred())
		Expect(got).NotTo(BeNil())
		Expect(got.ID).To(Equal(exp.ID), "dc id should be correct")

	})
})

func Test_Datacenter(t *testing.T) {
	RegisterFailHandler(Fail)
	//RunSpecs(t, "Datacenters")
}
