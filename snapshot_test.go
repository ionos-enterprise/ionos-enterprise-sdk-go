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

type ListSnapshotSuite struct {
	ClientBaseSuite
}

func TestListSnapshots(t *testing.T) {
	suite.Run(t, new(ListSnapshotSuite))
}

func (s *ListSnapshotSuite) SetupTest() {
	s.ClientBaseSuite.SetupTest()

	rsp := loadTestData(s.T(), "list_snapshots.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/snapshots`, httpmock.ResponderFromResponse(mResp))
}

func (s *ListSnapshotSuite) Test_WithSelector_NoSelector() {
	snapshots, err := s.c.ListSnapshotsWithSelector(nil)
	s.Error(err)
	s.Nil(snapshots)
}

func (s *ListSnapshotSuite) Test_WithSelector_Exact_Match() {
	snapshots, err := s.c.ListSnapshotsWithSelector(
		SelectExactSnapshot(
			SnapshotByState(StateAvailable),
			SnapshotByName("test-snapshot-02"),
		),
	)
	s.NoError(err)
	s.Len(snapshots, 1)
}

func (s *ListSnapshotSuite) Test_WithSelector_Exact_NoMatch() {
	snapshots, err := s.c.ListSnapshotsWithSelector(
		SelectExactSnapshot(
			SnapshotByState(StateAvailable),
			SnapshotByName("test-snapshot-03"),
		),
	)
	s.NoError(err)
	s.Len(snapshots, 0)
}

func (s *ListSnapshotSuite) Test_WithSelector_Any_Match() {
	snapshots, err := s.c.ListSnapshotsWithSelector(
		SelectAnySnapshot(
			SnapshotByDescription("some description"),
			SnapshotByName("test-snapshot-02"),
		),
	)
	s.NoError(err)
	s.Len(snapshots, 1)
}

func (s *ListSnapshotSuite) Test_WithSelector_Any_NoMatch() {
	snapshots, err := s.c.ListSnapshotsWithSelector(
		SelectAnySnapshot(
			SnapshotByDescription("my custom description"),
			SnapshotByName("test-snapshot-04"),
		),
	)
	s.NoError(err)
	s.Len(snapshots, 0)
}
