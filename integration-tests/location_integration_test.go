package integration_tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var locid string

func TestListLocations(t *testing.T) {
	fmt.Println("Location tests")
	c := setupTestEnv()

	resp, err := c.ListLocations()
	if err != nil {
		t.Errorf(err.Error())
	}

	locid = resp.Items[0].ID

	assert.True(t, len(resp.Items) > 0)
}

func TestGetLocation(t *testing.T) {
	c := setupTestEnv()

	resp, err := c.GetLocation("us/las")
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, resp.ID, "us/las")
	assert.Equal(t, resp.PBType, "location")
}

func TestGetRegionalLocations(t *testing.T) {
	c := setupTestEnv()
	_, err := c.GetRegionalLocations("us")
	if err != nil {
		t.Errorf(err.Error())
	}
}
