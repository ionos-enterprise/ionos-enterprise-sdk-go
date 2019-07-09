package integration_tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContractResources(t *testing.T) {
	fmt.Println("ContractResources tests")
	c := setupTestEnv()
	resp, err := c.GetContractResources()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.PBType, "contract")
}
