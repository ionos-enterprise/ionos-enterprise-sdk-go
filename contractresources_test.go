package profitbricks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetContractResources(t *testing.T) {
	want := 200
	resp := GetContractResources()
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Type_, "contract")
}
