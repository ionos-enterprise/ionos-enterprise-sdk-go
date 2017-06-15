package profitbricks

import (
	"testing"
)

func TestGetContractResources(t *testing.T) {
	want := 200
	resp := GetContractResources()
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
