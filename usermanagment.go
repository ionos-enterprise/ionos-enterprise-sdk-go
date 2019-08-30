package profitbricks

import (
	"net/http"
)

// Groups object
type Groups struct {
	BaseResource `json:",inline"`
	ID           string  `json:"id,omitempty"`
	PBType       string  `json:"type,omitempty"`
	Href         string  `json:"href,omitempty"`
	Items        []Group `json:"items,omitempty"`
	Response     string  `json:"Response,omitempty"`
	StatusCode   int     `json:"statuscode,omitempty"`
}

// Group object
type Group struct {
	BaseResource `json:",inline"`
	ID           string          `json:"id,omitempty"`
	PBType       string          `json:"type,omitempty"`
	Href         string          `json:"href,omitempty"`
	Properties   GroupProperties `json:"properties,omitempty"`
	Entities     *GroupEntities  `json:"entities,omitempty"`
	Response     string          `json:"Response,omitempty"`
	StatusCode   int             `json:"statuscode,omitempty"`
}

// GroupProperties object
type GroupProperties struct {
	Name              string `json:"name,omitempty"`
	CreateDataCenter  *bool  `json:"createDataCenter,omitempty"`
	CreateSnapshot    *bool  `json:"createSnapshot,omitempty"`
	ReserveIP         *bool  `json:"reserveIp,omitempty"`
	AccessActivityLog *bool  `json:"accessActivityLog,omitempty"`
}

// GroupEntities object
type GroupEntities struct {
	Users     Users     `json:"users,omitempty"`
	Resources Resources `json:"resources,omitempty"`
}

// Users object
type Users struct {
	BaseResource `json:",inline"`
	ID           string `json:"id,omitempty"`
	PBType       string `json:"type,omitempty"`
	Href         string `json:"href,omitempty"`
	Items        []User `json:"items,omitempty"`
	Response     string `json:"Response,omitempty"`
	StatusCode   int    `json:"statuscode,omitempty"`
}

// User object
type User struct {
	BaseResource `json:",inline"`
	ID           string          `json:"id,omitempty"`
	PBType       string          `json:"type,omitempty"`
	Href         string          `json:"href,omitempty"`
	Metadata     *Metadata       `json:"metadata,omitempty"`
	Properties   *UserProperties `json:"properties,omitempty"`
	Entities     *UserEntities   `json:"entities,omitempty"`
	Response     string          `json:"Response,omitempty"`
	StatusCode   int             `json:"statuscode,omitempty"`
}

// UserProperties object
type UserProperties struct {
	Firstname     string `json:"firstname,omitempty"`
	Lastname      string `json:"lastname,omitempty"`
	Email         string `json:"email,omitempty"`
	Password      string `json:"password,omitempty"`
	Administrator bool   `json:"administrator,omitempty"`
	ForceSecAuth  bool   `json:"forceSecAuth,omitempty"`
	SecAuthActive bool   `json:"secAuthActive,omitempty"`
}

// UserEntities object
type UserEntities struct {
	Groups Groups `json:"groups,omitempty"`
	Owns   Owns   `json:"owns,omitempty"`
}

// Resources object
type Resources struct {
	BaseResource `json:",inline"`
	ID           string     `json:"id,omitempty"`
	PBType       string     `json:"type,omitempty"`
	Href         string     `json:"href,omitempty"`
	Items        []Resource `json:"items,omitempty"`
	Response     string     `json:"Response,omitempty"`
	StatusCode   int        `json:"statuscode,omitempty"`
}

// Resource object
type Resource struct {
	BaseResource `json:",inline"`
	ID           string            `json:"id,omitempty"`
	PBType       string            `json:"type,omitempty"`
	Href         string            `json:"href,omitempty"`
	Metadata     *Metadata         `json:"metadata,omitempty"`
	Entities     *ResourceEntities `json:"entities,omitempty"`
	Response     string            `json:"Response,omitempty"`
	StatusCode   int               `json:"statuscode,omitempty"`
}

// ResourceEntities object
type ResourceEntities struct {
	Groups Groups `json:"groups,omitempty"`
}

// Owns object
type Owns struct {
	BaseResource `json:",inline"`
	ID           string   `json:"id,omitempty"`
	PBType       string   `json:"type,omitempty"`
	Href         string   `json:"href,omitempty"`
	Items        []Entity `json:"items,omitempty"`
	Response     string   `json:"Response,omitempty"`
	StatusCode   int      `json:"statuscode,omitempty"`
}

// Entity object
type Entity struct {
	BaseResource `json:",inline"`
	ID           string    `json:"id,omitempty"`
	PBType       string    `json:"type,omitempty"`
	Href         string    `json:"href,omitempty"`
	Metadata     *Metadata `json:"metadata,omitempty"`
	Response     string    `json:"Response,omitempty"`
	StatusCode   int       `json:"statuscode,omitempty"`
}

// Shares object
type Shares struct {
	BaseResource `json:",inline"`
	ID           string  `json:"id,omitempty"`
	PBType       string  `json:"type,omitempty"`
	Href         string  `json:"href,omitempty"`
	Items        []Share `json:"items,omitempty"`
	Response     string  `json:"Response,omitempty"`
	StatusCode   int     `json:"statuscode,omitempty"`
}

// Share object
type Share struct {
	BaseResource `json:",inline"`
	ID           string          `json:"id,omitempty"`
	PBType       string          `json:"type,omitempty"`
	Href         string          `json:"href,omitempty"`
	Properties   ShareProperties `json:"properties,omitempty"`
	Response     string          `json:"Response,omitempty"`
	StatusCode   int             `json:"statuscode,omitempty"`
}

// ShareProperties object
type ShareProperties struct {
	EditPrivilege  *bool `json:"editPrivilege,omitempty"`
	SharePrivilege *bool `json:"sharePrivilege,omitempty"`
}

type ApiResourcePath string

// ListGroups lists all groups
func (c *Client) ListGroups() (*Groups, error) {
	ret := &Groups{}
	return ret, c.GetOK(groupsPath(), ret)
}

// GetGroup gets a group
func (c *Client) GetGroup(groupid string) (*Group, error) {
	ret := &Group{}
	return ret, c.GetOK(groupPath(groupid), ret)
}

// CreateGroup creates a group
func (c *Client) CreateGroup(grp Group) (*Group, error) {
	ret := &Group{}
	return ret, c.PostAcc(groupsPath(), grp, ret)
}

// UpdateGroup updates a group
func (c *Client) UpdateGroup(groupid string, obj Group) (*Group, error) {
	ret := &Group{}
	return ret, c.PutAcc(groupPath(groupid), obj, ret)
}

// DeleteGroup deletes a group
func (c *Client) DeleteGroup(groupid string) (*http.Header, error) {
	return c.DeleteAcc(groupPath(groupid))
}

// ListShares lists all shares
func (c *Client) ListShares(grpid string) (*Shares, error) {
	ret := &Shares{}
	return ret, c.GetOK(sharesPath(grpid), ret)
}

// GetShare gets a share
func (c *Client) GetShare(groupid string, resourceid string) (*Share, error) {
	ret := &Share{}
	return ret, c.GetOK(sharePath(groupid, resourceid), ret)
}

// AddShare adds a share
func (c *Client) AddShare(groupid string, resourceid string, share Share) (*Share, error) {
	ret := &Share{}
	return ret, c.PostAcc(sharePath(groupid, resourceid), share, ret)
}

// UpdateShare updates a share
func (c *Client) UpdateShare(groupid string, resourceid string, obj Share) (*Share, error) {
	ret := &Share{}
	return ret, c.PatchAcc(sharePath(groupid, resourceid), obj, ret)
}

// DeleteShare deletes a share
func (c *Client) DeleteShare(groupid string, resourceid string) (*http.Header, error) {
	return c.DeleteAcc(sharePath(groupid, resourceid))
}

// ListGroupUsers lists Users in a group
func (c *Client) ListGroupUsers(groupid string) (*Users, error) {
	ret := &Users{}
	return ret, c.GetOK(groupUsersPath(groupid), ret)
}

// AddUserToGroup adds a user to a group
func (c *Client) AddUserToGroup(groupid string, userid string) (*User, error) {
	usr := &User{ID: userid}
	ret := &User{}
	return ret, c.PostAcc(groupUsersPath(groupid), usr, ret)
}

// DeleteUserFromGroup removes a user from a group
func (c *Client) DeleteUserFromGroup(groupid string, userid string) (*http.Header, error) {
	return c.DeleteAcc(groupUserPath(groupid, userid))
}

// ListUsers lists all users
func (c *Client) ListUsers() (*Users, error) {
	ret := &Users{}
	return ret, c.GetOK(usersPath(), ret)
}

// GetUser gets a user
func (c *Client) GetUser(usrid string) (*User, error) {
	ret := &User{}
	return ret, c.GetOK(userPath(usrid), ret)
}

// CreateUser creates a user
func (c *Client) CreateUser(usr User) (*User, error) {
	ret := &User{}
	return ret, c.PostAcc(usersPath(), usr, ret)
}

// UpdateUser updates user information
func (c *Client) UpdateUser(userid string, obj User) (*User, error) {
	ret := &User{}
	return ret, c.PutAcc(userPath(userid), obj, ret)
}

// DeleteUser deletes the specified user
func (c *Client) DeleteUser(userid string) (*http.Header, error) {
	return c.DeleteAcc(userPath(userid))
}

// ListResources lists all resources
func (c *Client) ListResources() (*Resources, error) {
	ret := &Resources{}
	return ret, c.GetOK(resourcesPath(), ret)
}

// GetResourceByType gets a resource by type
func (c *Client) GetResourceByType(resourcetype string, resourceid string) (*Resource, error) {
	ret := &Resource{}
	return ret, c.GetOK(resourcePath(resourcetype, resourceid), ret)
}

// ListResourcesByType list resources by type
func (c *Client) ListResourcesByType(resourcetype string) (*Resources, error) {
	ret := &Resources{}
	return ret, c.GetOK(resourcesTypePath(resourcetype), ret)
}
