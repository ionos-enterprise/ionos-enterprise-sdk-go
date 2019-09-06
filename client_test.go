package profitbricks

import (
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/assert"
)

func TestNewClientParams(t *testing.T) {
	pbc := NewClient(os.Getenv("PROFITBRICKS_API_URL"), os.Getenv("PROFITBRICKS_USERNAME"))

	pbc.SetDepth(5)
	pbc.SetUserAgent("blah")
	assert.Equal(t, "blah", pbc.GetUserAgent())
	assert.Equal(t, "5", pbc.QueryParam.Get("depth"))
}

type ClientBaseSuite struct {
	suite.Suite
	c *Client
}

func (s *ClientBaseSuite) SetupTest() {
	s.c = NewClient("", "")
	httpmock.ActivateNonDefault(s.c.Client.GetClient())
}

func (s *ClientBaseSuite) TearDownTest() {
	httpmock.Reset()
}
