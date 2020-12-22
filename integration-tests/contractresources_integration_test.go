// +build integration_tests integration_tests_contractresources

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
		t.Fatal(err)
	}
	assert.Equal(t, resp.PBType, "contract")
}
