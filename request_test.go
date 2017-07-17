package profitbricks

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
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

func TestGetRequestFailure(t *testing.T) {
	want := 404
	req := GetRequest("00000000-0000-0000-0000-000000000000")
	if req.StatusCode != want {
		t.Errorf(bad_status(want, req.StatusCode))
	}

	assert.True(t, strings.Contains(req.Response, "Resource does not exist"))
}
