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
	s.apiUrl = s.client.client.cloudApiUrl
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
	httpmock.RegisterResponderWithQuery(http.MethodGet, s.apiUrl+"/requests", query,
		httpmock.NewBytesResponder(200, listResponses[0]).Once())
	httpmock.RegisterResponderWithQuery(http.MethodGet, s.apiUrl+"/requests", query,
		httpmock.NewBytesResponder(200, listResponses[1]).Once())

	httpmock.RegisterResponder(http.MethodGet, "=~/requests/.*/status.*",
		httpmock.NewBytesResponder(200, statusResponses[0]).Once())
	httpmock.RegisterResponder(http.MethodGet, "=~/requests/.*/status.*",
		httpmock.NewBytesResponder(200, statusResponses[1]).Once())

	err := s.client.WaitTillRequestsFinished(context.Background(), NewRequestListFilter().WithUrl("volumes"))
	s.NoError(err)
}

func (s *SuiteWaitTillRequests) Test_Err_ListError() {
	httpmock.RegisterResponder(http.MethodGet, s.apiUrl+"/requests",
		httpmock.NewStringResponder(401, "{}"))
	err := s.client.WaitTillMatchingRequestsFinished(context.Background(), nil, nil)
	s.Error(err)
	s.Equal(1, httpmock.GetTotalCallCount())
}

func (s *SuiteWaitTillRequests) Test_Err_GetStatusError() {
	rsp := loadTestData(s.T(), "request_request_till_no_request_matches_01.json")
	httpmock.RegisterResponder(http.MethodGet, s.apiUrl+"/requests",
		httpmock.NewBytesResponder(200, rsp))

	httpmock.RegisterResponder(http.MethodGet, "=~/requests/.*/status.*",
		httpmock.NewStringResponder(401, "{}"))
	err := s.client.WaitTillRequestsFinished(context.Background(), nil)
	s.Error(err)
	s.Equal(2, httpmock.GetTotalCallCount())
}
