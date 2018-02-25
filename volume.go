package profitbricks

import (
	"net/http"
	"net/url"
	"strconv"
)

//Volume object
type Volume struct {
	ID         string           `json:"id,omitempty"`
	PBType     string           `json:"type,omitempty"`
	Href       string           `json:"href,omitempty"`
	Metadata   *Metadata        `json:"metadata,omitempty"`
	Properties VolumeProperties `json:"properties,omitempty"`
	Response   string           `json:"Response,omitempty"`
	Headers    *http.Header     `json:"headers,omitempty"`
	StatusCode int              `json:"statuscode,omitempty"`
}

//VolumeProperties object
type VolumeProperties struct {
	Name                string   `json:"name,omitempty"`
	Type                string   `json:"type,omitempty"`
	Size                int      `json:"size,omitempty"`
	AvailabilityZone    string   `json:"availabilityZone,omitempty"`
	Image               string   `json:"image,omitempty"`
	ImageAlias          string   `json:"imageAlias,omitempty"`
	ImagePassword       string   `json:"imagePassword,omitempty"`
	SSHKeys             []string `json:"sshKeys,omitempty"`
	Bus                 string   `json:"bus,omitempty"`
	LicenceType         string   `json:"licenceType,omitempty"`
	CPUHotPlug          bool     `json:"cpuHotPlug,omitempty"`
	CPUHotUnplug        bool     `json:"cpuHotUnplug,omitempty"`
	RAMHotPlug          bool     `json:"ramHotPlug,omitempty"`
	RAMHotUnplug        bool     `json:"ramHotUnplug,omitempty"`
	NicHotPlug          bool     `json:"nicHotPlug,omitempty"`
	NicHotUnplug        bool     `json:"nicHotUnplug,omitempty"`
	DiscVirtioHotPlug   bool     `json:"discVirtioHotPlug,omitempty"`
	DiscVirtioHotUnplug bool     `json:"discVirtioHotUnplug,omitempty"`
	DiscScsiHotPlug     bool     `json:"discScsiHotPlug,omitempty"`
	DiscScsiHotUnplug   bool     `json:"discScsiHotUnplug,omitempty"`
	DeviceNumber        int64    `json:"deviceNumber,omitempty"`
}

//Volumes object
type Volumes struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Volume     `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// ListVolumes returns a Collection struct for volumes in the Datacenter
func (c *Client) ListVolumes(dcid string) (*Volumes, error) {
	url := volumeColPath(dcid) + `?depth=` + c.client.depth
	ret := &Volumes{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//GetVolume gets a volume
func (c *Client) GetVolume(dcid string, volumeID string) (*Volume, error) {
	url := volumePath(dcid, volumeID) + `?depth=` + c.client.depth
	ret := &Volume{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//UpdateVolume updates a volume
func (c *Client) UpdateVolume(dcid string, volid string, request VolumeProperties) (*Volume, error) {
	url := volumePath(dcid, volid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Volume{}
	err := c.client.Patch(url, request, ret, http.StatusAccepted)
	return ret, err
}

//CreateVolume creates a volume
func (c *Client) CreateVolume(dcid string, request Volume) (*Volume, error) {
	url := volumeColPath(dcid) + `?depth=` + c.client.depth
	ret := &Volume{}
	err := c.client.Post(url, request, ret, http.StatusAccepted)
	return ret, err
}

// DeleteVolume deletes a volume
func (c *Client) DeleteVolume(dcid, volid string) (*http.Header, error) {
	url := volumePath(dcid, volid)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}

//CreateSnapshot creates a volume snapshot
func (c *Client) CreateSnapshot(dcid string, volid string, name string, description string) (*Snapshot, error) {
	path := volumePath(dcid, volid) + "/create-snapshot"
	data := url.Values{}
	data.Set("name", name)
	data.Add("description", description)

	ret := &Snapshot{}
	err := c.client.Post(path, data, ret, http.StatusAccepted)
	return ret, err
}

// RestoreSnapshot restores a volume with provided snapshot
func (c *Client) RestoreSnapshot(dcid string, volid string, snapshotID string) (*http.Header, error) {
	path := volumePath(dcid, volid) + "/restore-snapshot"
	data := url.Values{}
	data.Set("snapshotId", snapshotID)
	ret := &http.Header{}
	err := c.client.Post(path, data, ret, http.StatusAccepted)
	return ret, err
}
