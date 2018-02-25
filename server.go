package profitbricks

import (
	"net/http"
	"strconv"
)

//Server object
type Server struct {
	ID         string           `json:"id,omitempty"`
	PBType     string           `json:"type,omitempty"`
	Href       string           `json:"href,omitempty"`
	Metadata   *Metadata        `json:"metadata,omitempty"`
	Properties ServerProperties `json:"properties,omitempty"`
	Entities   *ServerEntities  `json:"entities,omitempty"`
	Response   string           `json:"Response,omitempty"`
	Headers    *http.Header     `json:"headers,omitempty"`
	StatusCode int              `json:"statuscode,omitempty"`
}

//ServerProperties object
type ServerProperties struct {
	Name             string             `json:"name,omitempty"`
	Cores            int                `json:"cores,omitempty"`
	RAM              int                `json:"ram,omitempty"`
	AvailabilityZone string             `json:"availabilityZone,omitempty"`
	VMState          string             `json:"vmState,omitempty"`
	BootCdrom        *ResourceReference `json:"bootCdrom,omitempty"`
	BootVolume       *ResourceReference `json:"bootVolume,omitempty"`
	CPUFamily        string             `json:"cpuFamily,omitempty"`
}

//ServerEntities object
type ServerEntities struct {
	Cdroms  *Cdroms  `json:"cdroms,omitempty"`
	Volumes *Volumes `json:"volumes,omitempty"`
	Nics    *Nics    `json:"nics,omitempty"`
}

//Servers collection
type Servers struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Server     `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

//ResourceReference object
type ResourceReference struct {
	ID     string `json:"id,omitempty"`
	PBType string `json:"type,omitempty"`
	Href   string `json:"href,omitempty"`
}

// ListServers returns a server struct collection
func (c *Client) ListServers(dcid string) (*Servers, error) {
	url := serverColPath(dcid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Servers{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// CreateServer creates a server from a jason []byte and returns a Instance struct
func (c *Client) CreateServer(dcid string, server Server) (*Server, error) {
	url := serverColPath(dcid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Server{}
	err := c.client.Post(url, server, ret, http.StatusAccepted)
	return ret, err
}

// GetServer pulls data for the server where id = srvid returns a Instance struct
func (c *Client) GetServer(dcid, srvid string) (*Server, error) {
	url := serverPath(dcid, srvid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Server{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

// UpdateServer partial update of server properties passed in as jason []byte
// Returns Instance struct
func (c *Client) UpdateServer(dcid string, srvid string, props ServerProperties) (*Server, error) {
	url := serverPath(dcid, srvid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Server{}
	err := c.client.Patch(url, props, ret, http.StatusAccepted)
	return ret, err
}

// DeleteServer deletes the server where id=srvid and returns Resp struct
func (c *Client) DeleteServer(dcid, srvid string) (*http.Header, error) {
	ret := &http.Header{}
	err := c.client.Delete(serverPath(dcid, srvid), ret, http.StatusAccepted)
	return ret, err
}

//ListAttachedCdroms returns list of attached cd roms
func (c *Client) ListAttachedCdroms(dcid, srvid string) (*Images, error) {
	url := serverCdromColPath(dcid, srvid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Images{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//AttachCdrom attaches a CD rom
func (c *Client) AttachCdrom(dcid string, srvid string, cdid string) (*Image, error) {
	data := struct {
		ID string `json:"id,omitempty"`
	}{
		cdid,
	}
	url := serverCdromColPath(dcid, srvid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Image{}
	err := c.client.Post(url, data, ret, http.StatusAccepted)
	return ret, err
}

//GetAttachedCdrom gets attached cd roms
func (c *Client) GetAttachedCdrom(dcid, srvid, cdid string) (*Image, error) {
	url := serverCdromPath(dcid, srvid, cdid) // + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Image{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//DetachCdrom detaches a CD rom
func (c *Client) DetachCdrom(dcid, srvid, cdid string) (*http.Header, error) {
	url := serverCdromPath(dcid, srvid, cdid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}

//ListAttachedVolumes lists attached volumes
func (c *Client) ListAttachedVolumes(dcid, srvid string) (*Volumes, error) {
	url := serverVolumeColPath(dcid, srvid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Volumes{}
	err := c.client.Get(url, ret, http.StatusOK)
	return ret, err
}

//AttachVolume attaches a volume
func (c *Client) AttachVolume(dcid string, srvid string, volid string) (*Volume, error) {
	data := struct {
		ID string `json:"id,omitempty"`
	}{
		volid,
	}
	url := serverVolumeColPath(dcid, srvid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Volume{}
	err := c.client.Post(url, data, ret, http.StatusAccepted)

	return ret, err
}

//GetAttachedVolume gets an attached volume
func (c *Client) GetAttachedVolume(dcid, srvid, volid string) (*Volume, error) {
	url := serverVolumePath(dcid, srvid, volid) + `?depth=` + c.client.depth + `&pretty=` + strconv.FormatBool(c.client.pretty)
	ret := &Volume{}
	err := c.client.Get(url, ret, http.StatusOK)

	return ret, err
}

//DetachVolume detaches a volume
func (c *Client) DetachVolume(dcid, srvid, volid string) (*http.Header, error) {
	url := serverVolumePath(dcid, srvid, volid)
	ret := &http.Header{}
	err := c.client.Delete(url, ret, http.StatusAccepted)
	return ret, err
}

// StartServer starts a server
func (c *Client) StartServer(dcid, srvid string) (*http.Header, error) {
	url := serverPath(dcid, srvid) + "/start"
	ret := &http.Header{}
	err := c.client.Post(url, nil, ret, http.StatusAccepted)
	return ret, err
}

// StopServer stops a server
func (c *Client) StopServer(dcid, srvid string) (*http.Header, error) {
	url := serverPath(dcid, srvid) + "/stop"
	ret := &http.Header{}
	err := c.client.Post(url, nil, ret, http.StatusAccepted)
	return ret, err
}

// RebootServer reboots a server
func (c *Client) RebootServer(dcid, srvid string) (*http.Header, error) {
	url := serverPath(dcid, srvid) + "/reboot"
	ret := &http.Header{}
	err := c.client.Post(url, nil, ret, http.StatusAccepted)
	return ret, err
}
