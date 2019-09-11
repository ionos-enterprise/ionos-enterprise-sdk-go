package profitbricks

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/stretchr/testify/suite"
)

type TestClientServer struct {
	ClientBaseSuite
}

func TestClient_Server(t *testing.T) {
	suite.Run(t, new(TestClientServer))
}
func (s *TestClientServer) TestClient_GetServer() {
	rsp := loadTestData(s.T(), "server_get.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/datacenters/1/servers/2\?`,
		httpmock.ResponderFromResponse(mResp))
	srv, err := s.c.GetServer("1", "2")
	s.NoError(err)
	s.Equal("Server001", srv.Properties.Name)
}
