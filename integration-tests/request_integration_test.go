package integration_tests

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	sdk "github.com/ionos-enterprise/ionos-enterprise-sdk-go/v6"
	"github.com/stretchr/testify/assert"
)

var reqID sdk.Request

func TestListRequests(t *testing.T) {
	fmt.Println("Request tests")
	c := setupTestEnv()
	resp, err := c.ListRequests()
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, len(resp.Items) > 0)
	reqID = resp.Items[0]
}

func TestGetRequestStatus(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.GetRequestStatus(reqID.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "request-status", resp.PBType)
	assert.Equal(t, reqID.Href + "/status", resp.Href)
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
		err := c.WaitTillProvisionedOrCanceled(ctx, c.CloudApiUrl + "/a/path")
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), context.Canceled.Error())
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
