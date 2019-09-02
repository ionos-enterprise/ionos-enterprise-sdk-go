package profitbricks

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientParams(t *testing.T) {
	pbc := NewClient(os.Getenv("PROFITBRICKS_API_URL"), os.Getenv("PROFITBRICKS_USERNAME"))

	pbc.SetDepth(5)
	pbc.SetUserAgent("blah")
	assert.Equal(t, pbc.GetUserAgent(), pbc.GetUserAgent())
}

type ClientBaseSuite struct {
	suite.Suite
	c *Client
}

func (s *ClientBaseSuite) SetupTest() {
	s.c = NewClient("","")
	httpmock.ActivateNonDefault(s.c.Client.GetClient())
}

func (s *ClientBaseSuite) TearDownTest() {
	httpmock.Reset()
}


type SuiteClient struct {
	ClientBaseSuite
}

func Test_Client(t *testing.T) {
	suite.Run(t, new(SuiteClient))
}

func (s *SuiteClient)Test_ApiError() {
	body := []byte(`{"httpStatus" : 401, "messages" : [ {"errorCode" : "315", "message" : "Unauthorized" } ] }`)
	rsp := makeJsonResponse(http.StatusUnauthorized, body)
	httpmock.RegisterResponder(http.MethodGet, "=~/datacenters", httpmock.ResponderFromResponse(rsp))
	_, err := s.c.ListDatacenters()
	s.Error(err)
	s.True(IsStatusUnauthorized(err))
	s.False(IsStatusAccepted(err))
	s.Equal(1, httpmock.GetTotalCallCount())
}
