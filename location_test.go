package profitbricks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var locid string

func TestListLocations(t *testing.T) {
	setupTestEnv()
	want := 200
	resp := ListLocations()
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	locid = resp.Items[0].Id

	assert.True(t, len(resp.Items) > 0)
}

func TestGetLocation(t *testing.T) {
	want := 200
	resp := GetLocation("us/las")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id,"us/las")
	assert.Equal(t, resp.Type_,"location")
}

func TestGetRegionalLocations(t *testing.T) {
	want := 200
	resp := GetRegionalLocations("us")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}
