package profitbricks

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type SnapshotSuite struct {
	ClientBaseSuite
}

func TestSnapshot(t *testing.T) {
	suite.Run(t, new(SnapshotSuite))
}

func (s *SnapshotSuite) Test_Delete() {
	mRsp := makeJsonResponse(http.StatusAccepted, nil)
	mRsp.Header.Set("location", "status")
	httpmock.RegisterResponder(http.MethodDelete, "=~/snapshots/111", httpmock.ResponderFromResponse(mRsp))
	rsp, err := s.c.DeleteSnapshot("111")
	s.NoError(err)
	s.Equal("status", rsp.Get("location"))
}
