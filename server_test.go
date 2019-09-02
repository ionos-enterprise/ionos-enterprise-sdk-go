package profitbricks

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestClientServer struct {
	suite.Suite
}

func TestClient_Server(t *testing.T) {
	suite.Run(t, new(TestClientServer))
}
func (s *TestClientServer) TestClient_GetServer() {

}
