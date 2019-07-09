package integration_tests

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	sdk "github.com/profitbricks/profitbricks-sdk-go"
	"github.com/stretchr/testify/assert"
)

var reqID sdk.Request

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

func TestWaitTillProvisionedOrCanceled(t *testing.T) {
	c := setupTestEnv()
	t.Run("cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		err := c.WaitTillProvisionedOrCanceled(ctx, "a/path")
		if assert.Error(t, err) {
			assert.Equal(t, context.Canceled, err)
		}
	})
	t.Run("error getting request status", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := c.WaitTillProvisionedOrCanceled(ctx, "no/such/path")
		if assert.Error(t, err) {
			assert.IsType(t, &url.Error{}, err)
		}
	})
}
