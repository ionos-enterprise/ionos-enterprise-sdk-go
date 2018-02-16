package profitbricks

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientParams(t *testing.T) {
	pbc := NewClient(os.Getenv("PROFITBRICKS_API_URL"), os.Getenv("PROFITBRICKS_USERNAME"))

	pbc.SetDepth(5)
	pbc.SetUserAgent("blah")
	assert.Equal(t, pbc.client.depth, "5")
	assert.Equal(t, pbc.client.agentHeader, pbc.GetUserAgent())
}
