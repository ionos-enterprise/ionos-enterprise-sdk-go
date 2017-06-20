package profitbricks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var reqId Request

func TestListRequests(t *testing.T) {
	setupTestEnv()
	want := 200
	resp := ListRequests()

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	reqId = resp.Items[0]
	req := GetRequest(reqId.ID)
	if req.StatusCode != want {
		t.Errorf(bad_status(want, req.StatusCode))
	}
	assert.True(t, len(resp.Items) > 0)
}

func TestGetRequestStatus(t *testing.T) {
	want := 200
	id := reqId.Href + "/status"
	resp := GetRequestStatus(id)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.Equal(t, resp.Type_, "request-status")
	assert.Equal(t, resp.Href, id)
}
