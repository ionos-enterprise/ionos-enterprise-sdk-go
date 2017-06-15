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

	assert.Equal(t, a, b, "The two words should be the same.")
}

func TestGetLocation(t *testing.T) {
	//t.Parallel()
	want := 200
	resp := GetLocation("us/las")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestGetRegionalLocations(t *testing.T) {
	//t.Parallel()
	want := 200
	resp := GetRegionalLocations("us")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
