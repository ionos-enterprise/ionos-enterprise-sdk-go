package profitbricks

import (
	"net/http"
)

// Snapshot object
type Snapshot struct {
	BaseResource `json:",inline"`
	ID           string             `json:"id,omitempty"`
	PBType       string             `json:"type,omitempty"`
	Href         string             `json:"href,omitempty"`
	Metadata     Metadata           `json:"metadata,omitempty"`
	Properties   SnapshotProperties `json:"properties,omitempty"`
	Response     string             `json:"Response,omitempty"`
	StatusCode   int                `json:"statuscode,omitempty"`
}

// SnapshotProperties properties
type SnapshotProperties struct {
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Location            string `json:"location,omitempty"`
	Size                int    `json:"size,omitempty"`
	CPUHotPlug          bool   `json:"cpuHotPlug,omitempty"`
	CPUHotUnplug        bool   `json:"cpuHotUnplug,omitempty"`
	RAMHotPlug          bool   `json:"ramHotPlug,omitempty"`
	RAMHotUnplug        bool   `json:"ramHotUnplug,omitempty"`
	NicHotPlug          bool   `json:"nicHotPlug,omitempty"`
	NicHotUnplug        bool   `json:"nicHotUnplug,omitempty"`
	DiscVirtioHotPlug   bool   `json:"discVirtioHotPlug,omitempty"`
	DiscVirtioHotUnplug bool   `json:"discVirtioHotUnplug,omitempty"`
	DiscScsiHotPlug     bool   `json:"discScsiHotPlug,omitempty"`
	DiscScsiHotUnplug   bool   `json:"discScsiHotUnplug,omitempty"`
	LicenceType         string `json:"licenceType,omitempty"`
}

// Snapshots object
type Snapshots struct {
	BaseResource `json:",inline"`
	ID           string     `json:"id,omitempty"`
	PBType       string     `json:"type,omitempty"`
	Href         string     `json:"href,omitempty"`
	Items        []Snapshot `json:"items,omitempty"`
	Response     string     `json:"Response,omitempty"`
	StatusCode   int        `json:"statuscode,omitempty"`
}

// ListSnapshots lists all snapshots
func (c *Client) ListSnapshots() (*Snapshots, error) {
	ret := &Snapshots{}
	return ret, c.GetOK(snapshotsPath(), ret)
}

// GetSnapshot gets a specific snapshot
func (c *Client) GetSnapshot(snapshotID string) (*Snapshot, error) {
	ret := &Snapshot{}
	return ret, c.GetOK(snapshotPath(snapshotID), ret)
}

// DeleteSnapshot deletes a specified snapshot
func (c *Client) DeleteSnapshot(snapshotID string) (*http.Header, error) {
	return c.DeleteAcc(snapshotPath(snapshotID))
}

// UpdateSnapshot updates a snapshot
func (c *Client) UpdateSnapshot(snapshotID string, request SnapshotProperties) (*Snapshot, error) {
	ret := &Snapshot{}
	return ret, c.PatchAcc(snapshotPath(snapshotID), request, ret)
}
