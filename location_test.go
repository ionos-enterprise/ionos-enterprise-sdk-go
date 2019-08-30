package profitbricks

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	// ...
	"testing"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	fixtureLocationsString = `{
	"id":"de/fkb",
	"type":"location",
	"href":"https://api.profitbricks.com/cloudapi/v4/locations/de/fkb",
	"properties":{
		"name":"karlsruhe",
		"features":["SSD"],
		"imageAliases":[
			"opensuse:latest",
			"ubuntu:18.10_iso",
			"mssql:2017_trial_iso",
			"fedora:28_iso",
			"zenloadbalancer:latest_iso"
		]}}`
)
var client *Client
var _ = BeforeSuite(func() {
	client = RestyClient("test", "test", "")
	// block all HTTP requests
	client.SetDoNotParseResponse(false)
	client.SetDebug(false)
	httpmock.ActivateNonDefault(client.Client.GetClient())
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})

var _ = Describe("Locations", func() {

	BeforeEach(func() {
		httpmock.Reset()
	})
	It("returns a location", func() {
		l := &Location{}
		err := json.Unmarshal([]byte(fixtureLocationsString), l)
		Expect(err).NotTo(HaveOccurred())
		x := &http.Response{
			Status:        "200",
			StatusCode:    200,
			Body:          ioutil.NopCloser(bytes.NewReader([]byte(fixtureLocationsString))),
			Header:        http.Header{},
			ContentLength: -1,
		}
		x.Header.Set("Content-Type", "application/json")

		responder := httpmock.ResponderFromResponse(x)
		fakeUrl := DefaultApiUrl + "/locations/de/fkb"
		httpmock.RegisterResponder("GET", fakeUrl, responder)

		// fetch the article into struct
		loc, err := client.GetLocation("de/fkb")
		Expect(err).NotTo(HaveOccurred())
		Expect(loc).NotTo(BeNil())
		Expect(loc.ID).To(Equal("de/fkb"))

		// do stuff with the article object ...
	})
})

func Test_Locations(t *testing.T) {
	RegisterFailHandler(Fail)
	//RunSpecs(t, "Locations")
}
