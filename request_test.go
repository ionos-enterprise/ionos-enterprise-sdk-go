package profitbricks

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const requestTestDataPath = "testdata"

type SuiteRequest struct {
	suite.Suite
	client *Client
	apiUrl string
}

func (s *SuiteRequest) SetupTest() {
	s.client = NewClient("", "")
	s.apiUrl = s.client.client.apiURL
	httpmock.Activate()
}

func (s *SuiteRequest) TearDownTest() {
	httpmock.DeactivateAndReset()
}

type SuiteWaitTillRequests struct {
	SuiteRequest
}

func Test_WaitTillRequests(t *testing.T) {
	suite.Run(t, new(SuiteWaitTillRequests))
}

func loadTestData(t *testing.T, filename string) []byte {
	data, err := ioutil.ReadFile(filepath.Join(requestTestDataPath, filename))
	assert.NoError(t, err)
	return data
}

func (s *SuiteWaitTillRequests) Test_OK_NoSelector() {
	listResponses := [][]byte{
		loadTestData(s.T(), "request_request_till_no_request_matches_01.json"),
		loadTestData(s.T(), "request_request_till_no_request_matches_02.json"),
	}
	statusResponses := [][]byte{
		loadTestData(s.T(), "request_queued.json"),
		loadTestData(s.T(), "request_done.json"),
	}
	query := url.Values{
		"filter.url": []string{"volumes"},
		"depth":      []string{"10"},
	}
	listResponded := 0
	httpmock.RegisterResponderWithQuery(http.MethodGet, s.apiUrl+"/requests", query,
		func(r *http.Request) (*http.Response, error) {
			if listResponded >= len(listResponses) {
				return httpmock.NewStringResponse(404, "{}"), nil
			}
			rsp := httpmock.NewBytesResponse(200, listResponses[listResponded])
			listResponded++
			return rsp, nil
		})

	statusResponded := 0
	httpmock.RegisterResponder(http.MethodGet, "=~/requests/.*/status.*", func(request *http.Request) (response *http.Response, e error) {
		rsp := httpmock.NewBytesResponse(200, statusResponses[statusResponded])
		statusResponded++
		return rsp, nil
	})

	err := s.client.WaitTillRequestsFinished(context.TODO(), NewRequestListFilter().WithUrl("volumes"))
	s.NoError(err)
}

func (s *SuiteWaitTillRequests) Test_Err_ListError() {
	httpmock.RegisterResponder(http.MethodGet, s.apiUrl+"/requests",
		func(r *http.Request) (*http.Response, error) {
			rsp := httpmock.NewStringResponse(401, "{}")
			return rsp, nil
		})
	err := s.client.WaitTillMatchingRequestsFinished(context.TODO(), nil, nil)
	s.Error(err)
	s.Equal(1, httpmock.GetTotalCallCount())
}

func (s *SuiteWaitTillRequests) Test_Err_GetStatusError() {
	rsp := loadTestData(s.T(), "request_request_till_no_request_matches_01.json")
	httpmock.RegisterResponder(http.MethodGet, s.apiUrl+"/requests",
		func(r *http.Request) (*http.Response, error) {
			rsp := httpmock.NewBytesResponse(200, rsp)
			return rsp, nil
		})
	httpmock.RegisterResponder(http.MethodGet, "=~/requests/.*/status.*", func(request *http.Request) (
		response *http.Response, e error) {
		return httpmock.NewStringResponse(404, "{}"), nil
	})
	err := s.client.WaitTillRequestsFinished(context.TODO(), nil)
	s.Error(err)
	s.Equal(2, httpmock.GetTotalCallCount())
}
