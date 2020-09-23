package profitbricks

import (
	"context"
	ionossdk "github.com/ionos-cloud/ionos-cloud-sdk-go/v5"
	"net/http"
)

// Server object
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
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Server     `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// ResourceReference object
type ResourceReference struct {
	ID     string `json:"id,omitempty"`
	PBType string `json:"type,omitempty"`
	Href   string `json:"href,omitempty"`
}

// ListServers returns a server struct collection
func (c *Client) ListServers(dcid string) (*Servers, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersGet(ctx, dcid, nil)
	ret := Servers{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := serversPath(dcid)
	ret := &Servers{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// CreateServer creates a server in given datacenter
func (c *Client) CreateServer(dcid string, server Server) (*Server, error) {

	input := ionossdk.Server{}
	if errConvert := convertToCore(&server, &input); errConvert != nil {
		return nil, errConvert
	}

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersPost(ctx, dcid, input, nil)
	ret := Server{}

	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := serversPath(dcid)
	ret := &Server{}
	err := c.Post(url, server, ret, http.StatusAccepted)
	return ret, err
	 */
}

// CreateServerAndWait creates a server, waits for the request to finish and returns a refreshed resource
// Note that an error does not necessarily means that the resource has not been created.
// If err & res are not nil, a resource with res.ID exists, but an error occurred either while waiting for
// the request or when refreshing the resource.
func (c *Client) CreateServerAndWait(ctx context.Context, dcid string, srvid Server) (res *Server, err error) {
	res, err = c.CreateServer(dcid, srvid)
	if err != nil {
		return
	}
	if err = c.WaitTillProvisionedOrCanceled(ctx, res.Headers.Get("location")); err != nil {
		return
	}
	var srv *Server
	if srv, err = c.GetServer(dcid, res.ID); err != nil {
		return
	} else {
		return srv, nil
	}
}

// GetServer pulls data for the server where id = srvid returns a Instance struct
func (c *Client) GetServer(dcid, srvid string) (*Server, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersFindById(ctx, dcid, srvid, nil)
	ret := Server{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := serverPath(dcid, srvid)
	ret := &Server{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// UpdateServer updates server with given properties and returns instance
func (c *Client) UpdateServer(dcid string, srvid string, props ServerProperties) (*Server, error) {

	input := ionossdk.ServerProperties{}
	if errConvert := convertToCore(&props, &input); errConvert != nil {
		return nil, errConvert
	}

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersPatch(ctx, dcid, srvid, input, nil)
	ret := Server{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := serverPath(dcid, srvid)
	ret := &Server{}
	err := c.Patch(url, props, ret, http.StatusAccepted)
	return ret, err

	 */
}

// UpdateServerAndWait updates a server, waits for the request to finish and
// returns a refreshed instance.
// Note that an error does not necessarily means that the resource has not been updated.
// If err & res are not nil, a resource with res.ID exists, but an error occurred either while waiting for
// the request or when refreshing the resource.
func (c *Client) UpdateServerAndWait(
	ctx context.Context, dcid, srvid string, props ServerProperties) (res *Server, err error) {
	res, err = c.UpdateServer(dcid, srvid, props)
	if err != nil {
		return
	}
	if err = c.WaitTillProvisionedOrCanceled(ctx, res.Headers.Get("location")); err != nil {

		return
	}
	var srv *Server
	if srv, err = c.GetServer(dcid, res.ID); err != nil {
		return
	} else {
		return srv, nil
	}
}

// DeleteServer deletes the server where id=srvid and returns Resp struct
func (c *Client) DeleteServer(dcid, srvid string) (*http.Header, error) {
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersDelete(ctx, dcid, srvid, nil)
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
	/*
	ret := &http.Header{}
	err := c.Delete(serverPath(dcid, srvid), ret, http.StatusAccepted)
	return ret, err
	 */
}

// DeleteServerAndWait deletes a server and waits for the request to finish
func (c *Client) DeleteServerAndWait(ctx context.Context, dcid, srvid string) error {
	rsp, err := c.DeleteServer(dcid, srvid)
	if err != nil {
		return err
	}
	return c.WaitTillProvisionedOrCanceled(ctx, rsp.Get("location"))
}

// ListAttachedCdroms returns list of attached cd roms
func (c *Client) ListAttachedCdroms(dcid, srvid string) (*Images, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersCdromsGet(ctx, dcid, srvid, nil)
	ret := Images{}

	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := cdromsPath(dcid, srvid)
	ret := &Images{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// AttachCdrom attaches a CD rom
func (c *Client) AttachCdrom(dcid string, srvid string, cdid string) (*Image, error) {

	image := ionossdk.Image{Id: &cdid}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersCdromsPost(ctx, dcid, srvid, image, nil)

	ret := Image{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	data := struct {
		ID string `json:"id,omitempty"`
	}{
		cdid,
	}
	url := cdromsPath(dcid, srvid)
	ret := &Image{}
	err := c.Post(url, data, ret, http.StatusAccepted)
	return ret, err
	 */
}

// GetAttachedCdrom gets attached cd roms
func (c *Client) GetAttachedCdrom(dcid, srvid, cdid string) (*Image, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersCdromsFindById(ctx, dcid, srvid, cdid, nil)
	ret := Image{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, err
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := cdromPath(dcid, srvid, cdid)
	ret := &Image{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// DetachCdrom detaches a CD rom
func (c *Client) DetachCdrom(dcid, srvid, cdid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersCdromsDelete(ctx, dcid, srvid, cdid, nil)
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
	/*
	url := cdromPath(dcid, srvid, cdid)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}

// ListAttachedVolumes lists attached volumes
func (c *Client) ListAttachedVolumes(dcid, srvid string) (*Volumes, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersVolumesGet(ctx, dcid, srvid, nil)
	ret := Volumes{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := attachedVolumesPath(dcid, srvid)
	ret := &Volumes{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// AttachVolume attaches a volume
func (c *Client) AttachVolume(dcid string, srvid string, volid string) (*Volume, error) {

	input := ionossdk.Volume{Id: &volid}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersVolumesPost(ctx, dcid, srvid, input, nil)
	ret := Volume{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	data := struct {
		ID string `json:"id,omitempty"`
	}{
		volid,
	}
	url := attachedVolumesPath(dcid, srvid)
	ret := &Volume{}
	err := c.Post(url, data, ret, http.StatusAccepted)

	return ret, err
	 */
}

// GetAttachedVolume gets an attached volume
func (c *Client) GetAttachedVolume(dcid, srvid, volid string) (*Volume, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersVolumesFindById(ctx, dcid, srvid, volid, nil)
	ret := Volume{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := attachedVolumePath(dcid, srvid, volid)
	ret := &Volume{}
	err := c.Get(url, ret, http.StatusOK)

	return ret, err
	 */
}

// DetachVolume detaches a volume
func (c *Client) DetachVolume(dcid, srvid, volid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersVolumesDelete(ctx, dcid, srvid, volid, nil)
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
	/*
	url := attachedVolumePath(dcid, srvid, volid)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}

// StartServer starts a server
func (c *Client) StartServer(dcid, srvid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersStartPost(ctx, dcid, srvid, nil)
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
	/*
	url := serverStartPath(dcid, srvid)
	ret := &Header{}
	err := c.Post(url, nil, ret, http.StatusAccepted)
	return ret.GetHeader(), err
	 */
}

// StopServer stops a server
func (c *Client) StopServer(dcid, srvid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersStopPost(ctx, dcid, srvid, nil)
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
	/*
	url := serverStopPath(dcid, srvid)
	ret := &Header{}
	err := c.Post(url, nil, ret, http.StatusAccepted)
	return ret.GetHeader(), err
	 */
}

// RebootServer reboots a server
func (c *Client) RebootServer(dcid, srvid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.ServerApi.DatacentersServersRebootPost(ctx, dcid, srvid, nil)
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
	/*
	url := serverRebootPath(dcid, srvid)
	ret := &Header{}
	err := c.Post(url, nil, ret, http.StatusAccepted)
	return ret.GetHeader(), err
	 */
}
