package profitbricks

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
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

func (s *SnapshotSuite) Test_ListSnapshotsWithSelector_Exact_Match() {
	rsp := loadTestData(s.T(), "list_snapshots.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/snapshots`, httpmock.ResponderFromResponse(mResp))

	snapshots, err := s.c.ListSnapshotsWithSelector(
		SelectExactSnapshot(
			SnapshotByState(StateAvailable),
			SnapshotByName("test-snapshot-02"),
		),
	)
	s.NoError(err)
	s.Len(snapshots, 1)
}

func (s *SnapshotSuite) Test_ListSnapshotsWithSelector_Exact_NoMatch() {
	rsp := loadTestData(s.T(), "list_snapshots.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/snapshots`, httpmock.ResponderFromResponse(mResp))

	snapshots, err := s.c.ListSnapshotsWithSelector(
		SelectExactSnapshot(
			SnapshotByState(StateAvailable),
			SnapshotByName("test-snapshot-03"),
		),
	)
	s.NoError(err)
	s.Len(snapshots, 0)
}

func (s *SnapshotSuite) Test_ListSnapshotsWithSelector_Any_Match() {
	rsp := loadTestData(s.T(), "list_snapshots.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/snapshots`, httpmock.ResponderFromResponse(mResp))

	snapshots, err := s.c.ListSnapshotsWithSelector(
		SelectAnySnapshot(
			SnapshotByDescription("some description"),
			SnapshotByName("test-snapshot-02"),
		),
	)
	s.NoError(err)
	s.Len(snapshots, 1)
}

func (s *SnapshotSuite) Test_ListSnapshotsWithSelector_Any_NoMatch() {
	rsp := loadTestData(s.T(), "list_snapshots.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/snapshots`, httpmock.ResponderFromResponse(mResp))

	snapshots, err := s.c.ListSnapshotsWithSelector(
		SelectAnySnapshot(
			SnapshotByDescription("my custom description"),
			SnapshotByName("test-snapshot-04"),
		),
	)
	s.NoError(err)
	s.Len(snapshots, 0)
}
