// ipblock_test.go
package profitbricks

import "testing"

var ipblkid string

func TestListIpBlocks(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListIpBlocks()

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestReserveIpBlock(t *testing.T) {
	//t.Parallel()
	want := 202
	var obj = IPBlockReserveRequest{
		IPBlockProperties: IPBlockProperties{
			Size:     1,
			Location: "us/lasdev",
		},
	}

	resp := ReserveIpBlock(obj)
	ipblkid = resp.Id
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestGetIpBlock(t *testing.T) {
	shouldbe := "ipblock"
	want := 200
	resp := GetIpBlock(ipblkid)

	if resp.Type != shouldbe {
		t.Errorf(bad_type(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.Resp.StatusCode))
	}
}

func TestReleaseIpBlock(t *testing.T) {
	want := 202
	resp := ReleaseIpBlock(ipblkid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}
