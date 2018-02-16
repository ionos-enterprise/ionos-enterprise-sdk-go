package profitbricks

import (
	"net/http"
	"strconv"
)

//Snapshot object
type Snapshot struct {
	ID         string             `json:"id,omitempty"`
	PBType     string             `json:"type,omitempty"`
	Href       string             `json:"href,omitempty"`
	Metadata   Metadata           `json:"metadata,omitempty"`
	Properties SnapshotProperties `json:"properties,omitempty"`
	Response   string             `json:"Response,omitempty"`
	Headers    *http.Header       `json:"headers,omitempty"`
	StatusCode int                `json:"statuscode,omitempty"`
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

//Snapshots object
type Snapshots struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Snapshot   `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

//ListSnapshots lists all snapshots
func (c *Client) ListSnapshots() (*Snapshots, error) {
	url := snapshotColPath() + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Snapshots{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//GetSnapshot gets a specific snapshot
func (c *Client) GetSnapshot(snapshotID string) (*Snapshot, error) {
	url := snapshotColPath() + slash(snapshotID) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Snapshot{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// DeleteSnapshot deletes a specified snapshot
func (c *Client) DeleteSnapshot(snapshotID string) (*http.Header, error) {
	url := snapshotColPath() + slash(snapshotID) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}

// UpdateSnapshot updates a snapshot
func (c *Client) UpdateSnapshot(snapshotID string, request SnapshotProperties) (*Snapshot, error) {
	url := snapshotColPath() + slash(snapshotID)
	ret := &Snapshot{}
	err := c.client.Patch(url, request, ret, http.StatusAccepted)
	return ret, err
}
