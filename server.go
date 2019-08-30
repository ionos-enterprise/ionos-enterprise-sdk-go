package profitbricks

import (
	"context"
	"net/http"
)

// Server object
type Server struct {
	BaseResource `json:",inline"`
	ID           string           `json:"id,omitempty"`
	PBType       string           `json:"type,omitempty"`
	Href         string           `json:"href,omitempty"`
	Metadata     *Metadata        `json:"metadata,omitempty"`
	Properties   ServerProperties `json:"properties,omitempty"`
	Entities     *ServerEntities  `json:"entities,omitempty"`
	Response     string           `json:"Response,omitempty"`
	StatusCode   int              `json:"statuscode,omitempty"`
}

// ServerProperties object
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

// ServerEntities object
type ServerEntities struct {
	Cdroms  *Cdroms  `json:"cdroms,omitempty"`
	Volumes *Volumes `json:"volumes,omitempty"`
	Nics    *Nics    `json:"nics,omitempty"`
}

// Servers collection
type Servers struct {
	BaseResource `json:",inline"`
	ID           string   `json:"id,omitempty"`
	PBType       string   `json:"type,omitempty"`
	Href         string   `json:"href,omitempty"`
	Items        []Server `json:"items,omitempty"`
	Response     string   `json:"Response,omitempty"`
	StatusCode   int      `json:"statuscode,omitempty"`
}

// ResourceReference object
type ResourceReference struct {
	ID     string `json:"id,omitempty"`
	PBType string `json:"type,omitempty"`
	Href   string `json:"href,omitempty"`
}

// ListServers returns a server struct collection
func (c *Client) ListServers(dcid string) (*Servers, error) {
	ret := &Servers{}
	return ret, c.GetOK(serversPath(dcid), ret)
}

// CreateServer creates a server from a jason []byte and returns a Instance struct
func (c *Client) CreateServer(dcid string, server Server) (*Server, error) {
	ret := &Server{}
	return ret, c.PostAcc(serversPath(dcid), server, ret)
}

// GetServer pulls data for the server where id = srvid returns a Instance struct
func (c *Client) GetServer(dcid, srvid string) (*Server, error) {
	ret := &Server{}
	return ret, c.GetOK(serverPath(dcid, srvid), ret)
}

// UpdateServer partial update of server properties passed in as jason []byte
// Returns Instance struct
func (c *Client) UpdateServer(dcid string, srvid string, props ServerProperties) (*Server, error) {
	ret := &Server{}
	return ret, c.PatchAcc(serverPath(dcid, srvid), props, ret)
}

// DeleteServer deletes the server where id=srvid and returns Resp struct
func (c *Client) DeleteServer(dcid, srvid string) (*http.Header, error) {
	return c.DeleteAcc(serverPath(dcid, srvid))
}

// ListAttachedCdroms returns list of attached cd roms
func (c *Client) ListAttachedCdroms(dcid, srvid string) (*Images, error) {

	ret := &Images{}
	return ret, c.GetOK(cdromsPath(dcid, srvid), ret)
}

// AttachCdrom attaches a CD rom
func (c *Client) AttachCdrom(dcid string, srvid string, cdid string) (*Image, error) {
	data := map[string]string{"id": cdid}
	ret := &Image{}
	return ret, c.PostAcc(cdromsPath(dcid, srvid), data, ret)
}

// GetAttachedCdrom gets attached cd roms
func (c *Client) GetAttachedCdrom(dcid, srvid, cdid string) (*Image, error) {
	ret := &Image{}
	return ret, c.GetOK(cdromPath(dcid, srvid, cdid), ret)
}

// DetachCdrom detaches a CD rom
func (c *Client) DetachCdrom(dcid, srvid, cdid string) (*http.Header, error) {
	return c.DeleteAcc(cdromPath(dcid, srvid, cdid))
}

// ListAttachedVolumes lists attached volumes
func (c *Client) ListAttachedVolumes(dcid, srvid string) (*Volumes, error) {
	ret := &Volumes{}
	return ret, c.GetOK(attachedVolumesPath(dcid, srvid), ret)
}

// AttachVolume attaches a volume
func (c *Client) AttachVolume(dcid string, srvid string, volid string) (*Volume, error) {
	data := map[string]string{"id": volid}
	ret := &Volume{}
	return ret, c.PostAcc(attachedVolumesPath(dcid, srvid), data, ret)
}

// GetAttachedVolume gets an attached volume
func (c *Client) GetAttachedVolume(dcid, srvid, volid string) (*Volume, error) {
	ret := &Volume{}
	return ret, c.GetOK(attachedVolumePath(dcid, srvid, volid), ret)
}

// DetachVolume detaches a volume
func (c *Client) DetachVolume(dcid, srvid, volid string) (*http.Header, error) {
	return c.DeleteAcc(attachedVolumePath(dcid, srvid, volid))
}

func (c *Client) SyncDetachVolume(ctx context.Context, dcid, srvid, volid string) error {
	rsp, err := c.DetachVolume(dcid, srvid, volid)
	if err != nil {
		return err
	}
	return c.WaitTillProvisionedOrCanceled(ctx, rsp.Get("location"))
}

// StartServer starts a server
func (c *Client) StartServer(dcid, srvid string) (*http.Header, error) {
	ret := &BaseResource{}
	err := c.PostAcc(serverStartPath(dcid, srvid), nil, ret)
	return ret.GetHeaders(), err
}

// StopServer stops a server
func (c *Client) StopServer(dcid, srvid string) (*http.Header, error) {
	ret := &BaseResource{}
	err := c.PostAcc(serverStopPath(dcid, srvid), nil, ret)
	return ret.GetHeaders(), err
}

// RebootServer reboots a server
func (c *Client) RebootServer(dcid, srvid string) (*http.Header, error) {
	ret := &BaseResource{}
	err := c.PostAcc(serverRebootPath(dcid, srvid), nil, ret)
	return ret.GetHeaders(), err
}
