package profitbricks

import (
	"context"
	"github.com/ionos-cloud/sdk-go/v5"
	"net/http"
)

// Groups object
type Groups struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Group      `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// Group object
type Group struct {
	ID         string          `json:"id,omitempty"`
	PBType     string          `json:"type,omitempty"`
	Href       string          `json:"href,omitempty"`
	Properties GroupProperties `json:"properties,omitempty"`
	Entities   *GroupEntities  `json:"entities,omitempty"`
	Response   string          `json:"Response,omitempty"`
	Headers    *http.Header    `json:"headers,omitempty"`
	StatusCode int             `json:"statuscode,omitempty"`
}

// GroupProperties object
type GroupProperties struct {
	Name                 string `json:"name,omitempty"`
	CreateDataCenter     *bool  `json:"createDataCenter,omitempty"`
	CreateSnapshot       *bool  `json:"createSnapshot,omitempty"`
	ReserveIP            *bool  `json:"reserveIp,omitempty"`
	AccessActivityLog    *bool  `json:"accessActivityLog,omitempty"`
	CreateBackupUnit     *bool  `json:"createBackupUnit,omitempty"`
	CreateInternetAccess *bool  `json:"createInternetAccess,omitempty"`
	CreateK8sCluster     *bool  `json:"createK8sCluster,omitempty"`
	CreatePcc            *bool  `json:"createPcc,omitempty"`
	S3Privilege          *bool  `json:"s3Privilege,omitempty"`
}

// GroupEntities object
type GroupEntities struct {
	Users     Users     `json:"users,omitempty"`
	Resources Resources `json:"resources,omitempty"`
}

// Users object
type Users struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []User       `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// User object
type User struct {
	ID         string          `json:"id,omitempty"`
	PBType     string          `json:"type,omitempty"`
	Href       string          `json:"href,omitempty"`
	Metadata   *Metadata       `json:"metadata,omitempty"`
	Properties *UserProperties `json:"properties,omitempty"`
	Entities   *UserEntities   `json:"entities,omitempty"`
	Response   string          `json:"Response,omitempty"`
	Headers    *http.Header    `json:"headers,omitempty"`
	StatusCode int             `json:"statuscode,omitempty"`
}

// UserProperties object
type UserProperties struct {
	Firstname         string `json:"firstname,omitempty"`
	Lastname          string `json:"lastname,omitempty"`
	Email             string `json:"email,omitempty"`
	Password          string `json:"password,omitempty"`
	Administrator     bool   `json:"administrator,omitempty"`
	ForceSecAuth      bool   `json:"forceSecAuth,omitempty"`
	SecAuthActive     bool   `json:"secAuthActive,omitempty"`
	Active            *bool  `json:"active,omitempty"`
	S3CanonicalUserID string `json:"s3CanonicalUserId,omitempty"`
}

// UserEntities object
type UserEntities struct {
	Groups Groups `json:"groups,omitempty"`
	Owns   Owns   `json:"owns,omitempty"`
}

// Resources object
type Resources struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Resource   `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// Resource object
type Resource struct {
	ID         string            `json:"id,omitempty"`
	PBType     string            `json:"type,omitempty"`
	Href       string            `json:"href,omitempty"`
	Metadata   *Metadata         `json:"metadata,omitempty"`
	Entities   *ResourceEntities `json:"entities,omitempty"`
	Response   string            `json:"Response,omitempty"`
	Headers    *http.Header      `json:"headers,omitempty"`
	StatusCode int               `json:"statuscode,omitempty"`
}

// ResourceEntities object
type ResourceEntities struct {
	Groups Groups `json:"groups,omitempty"`
}

// Owns object
type Owns struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Entity     `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// Entity object
type Entity struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Metadata   *Metadata    `json:"metadata,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// Shares object
type Shares struct {
	ID         string       `json:"id,omitempty"`
	PBType     string       `json:"type,omitempty"`
	Href       string       `json:"href,omitempty"`
	Items      []Share      `json:"items,omitempty"`
	Response   string       `json:"Response,omitempty"`
	Headers    *http.Header `json:"headers,omitempty"`
	StatusCode int          `json:"statuscode,omitempty"`
}

// Share object
type Share struct {
	ID         string          `json:"id,omitempty"`
	PBType     string          `json:"type,omitempty"`
	Href       string          `json:"href,omitempty"`
	Properties ShareProperties `json:"properties,omitempty"`
	Response   string          `json:"Response,omitempty"`
	Headers    *http.Header    `json:"headers,omitempty"`
	StatusCode int             `json:"statuscode,omitempty"`
}

// ShareProperties object
type ShareProperties struct {
	EditPrivilege  *bool `json:"editPrivilege,omitempty"`
	SharePrivilege *bool `json:"sharePrivilege,omitempty"`
}

//ListGroups lists all groups
func (c *Client) ListGroups() (*Groups, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsGet(ctx).Execute()
	ret := Groups{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := groupsPath()
	ret := &Groups{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//GetGroup gets a group
func (c *Client) GetGroup(groupid string) (*Group, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsFindById(ctx, groupid).Execute()
	ret := Group{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := groupPath(groupid)
	ret := &Group{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//CreateGroup creates a group
func (c *Client) CreateGroup(grp Group) (*Group, error) {

	input := ionoscloud.Group{}
	if errConvert := convertToCore(&grp, &input); errConvert != nil {
		return nil, errConvert
	}

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsPost(ctx).Group(input).Execute()
	ret := Group{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := groupsPath()
	ret := &Group{}
	err := c.Post(url, grp, ret, http.StatusAccepted)
	return ret, err
	 */
}

//UpdateGroup updates a group
func (c *Client) UpdateGroup(groupid string, obj Group) (*Group, error) {

	input := ionoscloud.Group{}
	if errConvert := convertToCore(&obj, &input); errConvert != nil {
		return nil, errConvert
	}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsPut(ctx, groupid).Group(input).Execute()
	ret := Group{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := groupPath(groupid)
	ret := &Group{}
	err := c.Put(url, obj, ret, http.StatusAccepted)
	return ret, err
	 */
}

//DeleteGroup deletes a group
func (c *Client) DeleteGroup(groupid string) (*http.Header, error) {
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsDelete(ctx, groupid).Execute()
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}

	/*
	url := groupPath(groupid)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}

//ListShares lists all shares
func (c *Client) ListShares(grpid string) (*Shares, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsSharesGet(ctx, grpid).Execute()
	ret := Shares{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := sharesPath(grpid)
	ret := &Shares{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// GetShare gets a share
func (c *Client) GetShare(groupid string, resourceid string) (*Share, error) {

	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsSharesFindByResourceId(
		context.TODO(), groupid, resourceid).Execute()
	ret := Share{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := sharePath(groupid, resourceid)
	ret := &Share{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// AddShare adds a share
func (c *Client) AddShare(groupid string, resourceid string, share Share) (*Share, error) {

	input := ionoscloud.GroupShare{}
	if errConvert := convertToCore(&share, &input); errConvert != nil {
		return nil, errConvert
	}

	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsSharesPost(
		context.TODO(), groupid, resourceid).Resource(input).Execute()
	ret := Share{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := sharePath(groupid, resourceid)
	ret := &Share{}
	err := c.Post(url, share, ret, http.StatusAccepted)
	return ret, err
	 */
}

// UpdateShare updates a share
func (c *Client) UpdateShare(groupid string, resourceid string, obj Share) (*Share, error) {

	input := ionoscloud.GroupShare{}
	if errConvert := convertToCore(&obj, &input); errConvert != nil {
		return nil, errConvert
	}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsSharesPut(ctx, groupid, resourceid).Resource(input).Execute()
	ret := Share{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := sharePath(groupid, resourceid)
	ret := &Share{}
	err := c.Put(url, obj, ret, http.StatusAccepted)
	return ret, err
	 */
}

// DeleteShare deletes a share
func (c *Client) DeleteShare(groupid string, resourceid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsSharesDelete(ctx, groupid, resourceid).Execute()
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
	/*
	url := sharePath(groupid, resourceid)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}

//ListGroupUsers lists Users in a group
func (c *Client) ListGroupUsers(groupid string) (*Users, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsUsersGet(ctx, groupid).Execute()
	ret := Users{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := groupUsersPath(groupid)
	ret := &Users{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// AddUserToGroup adds a user to a group
func (c *Client) AddUserToGroup(groupid string, userid string) (*User, error) {

	input := ionoscloud.User{
		Id: &userid,
	}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsUsersPost(ctx, groupid).User(input).Execute()
	ret := User{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	var usr User
	usr.ID = userid
	url := groupUsersPath(groupid)
	ret := &User{}
	err := c.Post(url, usr, ret, http.StatusAccepted)
	return ret, err
	 */
}

// DeleteUserFromGroup removes a user from a group
func (c *Client) DeleteUserFromGroup(groupid string, userid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.UserManagementApi.UmGroupsUsersDelete(ctx, groupid, userid).Execute()
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}

	/*
	url := groupUserPath(groupid, userid)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}

//ListUsers lists all users
func (c *Client) ListUsers() (*Users, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmUsersGet(ctx).Execute()
	ret := Users{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := usersPath()
	ret := &Users{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

// GetUser gets a user
func (c *Client) GetUser(usrid string) (*User, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmUsersFindById(ctx, usrid).Execute()
	ret := User{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := userPath(usrid)
	ret := &User{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//CreateUser creates a user
func (c *Client) CreateUser(usr User) (*User, error) {

	input := ionoscloud.User{}
	if errConvert := convertToCore(&usr, &input); errConvert != nil {
		return nil, errConvert
	}
	/* setting this to nil to avoid marshalling it, otherwise we get a
	 * [(root).properties.secAuthActive] Attribute is not allowed in create requests
	 */
	input.Properties.SecAuthActive = nil
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmUsersPost(ctx).User(input).Execute()
	ret := User{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := usersPath()
	ret := &User{}
	err := c.Post(url, usr, ret, http.StatusAccepted)
	return ret, err
	 */
}

//UpdateUser updates user information
func (c *Client) UpdateUser(userid string, obj User) (*User, error) {

	input := ionoscloud.User{}
	if errConvert := convertToCore(&obj, &input); errConvert != nil {
		return nil, errConvert
	}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmUsersPut(ctx, userid).User(input).Execute()
	ret := User{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err
	/*
	url := userPath(userid)
	ret := &User{}
	err := c.Put(url, obj, ret, http.StatusAccepted)
	return ret, err
	 */
}

//DeleteUser deletes the specified user
func (c *Client) DeleteUser(userid string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.UserManagementApi.UmUsersDelete(ctx, userid).Execute()
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
	/*
	url := userPath(userid)
	ret := &http.Header{}
	err := c.Delete(url, ret, http.StatusAccepted)
	return ret, err
	 */
}

//ListResources lists all resources
func (c *Client) ListResources() (*Resources, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmResourcesGet(ctx).Execute()
	ret := Resources{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := resourcesPath()
	ret := &Resources{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
 	*/
}

//GetResourceByType gets a resource by type
func (c *Client) GetResourceByType(resourcetype string, resourceid string) (*Resource, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmResourcesFindByTypeAndId(ctx, resourcetype, resourceid).Execute()
	ret := Resource{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := resourcePath(resourcetype, resourceid)
	ret := &Resource{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}

//ListResourcesByType list resources by type
func (c *Client) ListResourcesByType(resourcetype string) (*Resources, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, apiResponse, err := c.CoreSdk.UserManagementApi.UmResourcesFindByType(ctx, resourcetype).Execute()
	ret := Resources{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	fillInResponse(&ret, apiResponse)
	return &ret, err

	/*
	url := resourcesTypePath(resourcetype)
	ret := &Resources{}
	err := c.Get(url, ret, http.StatusOK)
	return ret, err
	 */
}
