package profitbricks

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"testing"
	"time"

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
	s.apiUrl = s.client.CloudApiUrl
	httpmock.ActivateNonDefault(s.client.Client.GetClient())
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
	listResponses := []*http.Response{
		makeJsonResponse(200, loadTestData(s.T(), "request_request_till_no_request_matches_01.json")),
		makeJsonResponse(200, loadTestData(s.T(), "request_request_till_no_request_matches_02.json")),
	}
	statusResponses := []*http.Response{
		makeJsonResponse(200, loadTestData(s.T(), "request_queued.json")),
		makeJsonResponse(200, loadTestData(s.T(), "request_done.json")),
	}
	query := url.Values{
		"filter.url": []string{"volumes"},
		"depth":      []string{"10"},
	}
	listCalled := 0
	statusCalled := 0
	var lr httpmock.Responder = func(req *http.Request) (*http.Response, error) {
		rs := listResponses[listCalled]
		listCalled++
		return rs, nil
	}
	var sr httpmock.Responder = func(req *http.Request) (*http.Response, error) {
		rs := statusResponses[statusCalled]
		statusCalled++
		return rs, nil

	}
	httpmock.RegisterResponderWithQuery(http.MethodGet, s.apiUrl+"/requests", query, lr.Times(2))

	httpmock.RegisterResponder(http.MethodGet, "=~/requests/.*/status.*", sr.Times(2))

	err := s.client.WaitTillRequestsFinished(context.Background(), NewRequestListFilter().WithUrl("volumes"))
	s.NoError(err)
	s.Equal(4, httpmock.GetTotalCallCount())
}

func (s *SuiteWaitTillRequests) Test_Err_ListError() {
	httpmock.RegisterResponder(http.MethodGet, "=~/requests",
		httpmock.NewStringResponder(401, "{}"))
	err := s.client.WaitTillMatchingRequestsFinished(context.Background(), nil, nil)
	s.Error(err)
	s.Equal(1, httpmock.GetTotalCallCount())
}

func makeJsonResponse(status int, data []byte) *http.Response {
	body := ioutil.NopCloser(bytes.NewReader(data))
	rsp := http.Response{Body: body, Header: http.Header{}, StatusCode: status, Status: http.StatusText(status)}
	rsp.Header.Set("Content-Type", "application/json")
	return &rsp
}

func (s *SuiteWaitTillRequests) Test_Err_GetStatusError() {
	rsp := loadTestData(s.T(), "request_request_till_no_request_matches_01.json")
	listResponse := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/requests\?.*`,
		httpmock.ResponderFromResponse(listResponse))
	statusResponse := makeJsonResponse(http.StatusUnauthorized, []byte("{}"))
	httpmock.RegisterResponder(http.MethodGet, "=~/requests/.*/status.*",
		httpmock.ResponderFromResponse(statusResponse))
	err := s.client.WaitTillRequestsFinished(context.Background(), nil)
	s.Error(err)
	s.Equal(2, httpmock.GetTotalCallCount())
}

func TestRequestListFilter_AddCreatedAfter(t *testing.T) {
	r := NewRequestListFilter()
	tm, err := time.Parse(timeFormat, "2019-03-25 10:40:00")
	assert.NoError(t, err)
	r.AddCreatedAfter(tm)
	assert.Equal(t, "2019-03-25 10:40:00", r.Get("filter.createdAfter"))
}

func TestRequestListFilter_AddCreatedBefore(t *testing.T) {
	r := NewRequestListFilter()
	tm, err := time.Parse(timeFormat, "2019-03-25 10:40:00")
	assert.NoError(t, err)
	r.AddCreatedBefore(tm)
	assert.Equal(t, "2019-03-25 10:40:00", r.Get("filter.createdBefore"))
}

func TestRequestListFilter_AddRequestStatus(t *testing.T) {
	r := NewRequestListFilter()
	r.AddRequestStatus(RequestStatusDone)
	assert.Equal(t, RequestStatusDone, r.Get("filter.status"))
}

func TestRequestListFilter_Clone(t *testing.T) {
	r := NewRequestListFilter()
	r.AddMethod(http.MethodPost)
	r.AddUrl("foo/bar")
	c := r.Clone()
	assert.Equal(t, r, c)
}
