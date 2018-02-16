package profitbricks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var reqID Request

func TestListRequests(t *testing.T) {
	fmt.Println("Request tests")
	c := setupTestEnv()
	resp, err := c.ListRequests()
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
	reqID = resp.Items[0]
}

func TestGetRequestStatus(t *testing.T) {
	c := setupTestEnv()
	path := reqID.Href + "/status"
	resp, err := c.GetRequestStatus(path)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.PBType, "request-status")
	assert.Equal(t, resp.Href, path)
}

func TestGetRequestFailure(t *testing.T) {
	c := setupTestEnv()
	_, err := c.GetRequest("00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, err)
}
