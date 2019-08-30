package profitbricks

import (
	"os"
	"testing"
)

func TestNewClientParams(t *testing.T) {
	pbc := NewClient(os.Getenv("PROFITBRICKS_API_URL"), os.Getenv("PROFITBRICKS_USERNAME"))

	pbc.SetDepth(5)
	pbc.SetUserAgent("blah")
}
