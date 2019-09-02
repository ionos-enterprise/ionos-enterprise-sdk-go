package profitbricks

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestClientServer struct {
	suite.Suite
}

func TestClient_Server(t *testing.T) {
	suite.Run(t, new(TestClientServer))
}
func (s * TestClientServer) TestClient_GetServer() {

}
