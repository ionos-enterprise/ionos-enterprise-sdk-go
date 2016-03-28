package profitbricks

import (
	"testing"
)

func TestGetRequestStatus(t *testing.T) {
	want :=200
	SetAuth("username", "password")
	resp := GetRequestStatus("https://api.profitbricks.com/rest/v2/requests/75237fe8-5d52-40dc-951a-4f8854e8014a/status")
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}

}
