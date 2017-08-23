package profitbricks

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

var groupid string
var resourceId string
var userid string
var email string
var ipblockId string
var snapshotResourceId string
var imageResourceId string
var TRUE bool = true
var FALSE bool = false

func setupTest() {
	setupTestEnv()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	email = "test" + strconv.Itoa(r1.Intn(1000)) + "@go.com"
	resourceId = mkdcid("GO SDK TEST")
	snapshotResourceId=mksnapshotId("GO SDK TEST",resourceId)
	ipblockId = mkipid()
	imageResourceId=getImageId(location,"ubuntu","HDD")
}

func TestCreateUser(t *testing.T) {
	setupTest()
	want := 202
	var obj = User{
		Properties: &UserProperties{
			Firstname:     "John",
			Lastname:      "Doe",
			Email:         email,
			Password:      "abc123-321CBA",
			Administrator: true,
			ForceSecAuth:  false,
			SecAuthActive: false,
		},
	}
	resp := CreateUser(obj)
	userid = resp.Id

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Properties.Firstname, "John")
	assert.Equal(t, resp.Properties.Lastname, "Doe")
	assert.Equal(t, resp.Properties.Email, email)
	assert.Equal(t, resp.Properties.Administrator, true)
}

func TestCreateUserFailure(t *testing.T) {
	want := 422
	var obj = User{
		Properties: &UserProperties{
			Firstname:     "John",
			Lastname:      "Doe",
			Password:      "abc123-321CBA",
			Administrator: true,
			ForceSecAuth:  false,
			SecAuthActive: false,
		},
	}
	resp := CreateUser(obj)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, strings.Contains(resp.Response, "Attribute 'email' is required"))
}

func TestListUsers(t *testing.T) {
	SetDepth("5")
	want := 200
	resp := ListUsers()

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.True(t, len(resp.Items) > 0)
}

func TestGetUser(t *testing.T) {
	want := 200
	resp := GetUser(userid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, userid)
	assert.Equal(t, resp.Properties.Firstname, "John")
	assert.Equal(t, resp.Properties.Lastname, "Doe")
	assert.Equal(t, resp.Properties.Email, email)
	assert.Equal(t, resp.Properties.Administrator, true)
	assert.Equal(t, resp.Type_, "user")
}

func TestUpdateUser(t *testing.T) {
	want := 202
	newName := "user updated"
	obj := User{
		Properties: &UserProperties{
			Firstname:     "John",
			Lastname:      newName,
			Email:         email,
			Administrator: false,
			ForceSecAuth:  false,
			SecAuthActive: false,
		}}

	resp := UpdateUser(userid, obj)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	if resp.Properties.Lastname != newName {
		t.Errorf("Not updated")
	}
}

func TestCreateGroup(t *testing.T) {
	want := 202
	var obj = Group{
		Properties: GroupProperties{
			Name:              "GO SDK Test",
			CreateDataCenter:  &TRUE,
			CreateSnapshot:    &TRUE,
			ReserveIp:         &TRUE,
			AccessActivityLog: &TRUE,
		},
	}
	resp := CreateGroup(obj)
	groupid = resp.Id

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, *resp.Properties.CreateDataCenter, true)
	assert.Equal(t, *resp.Properties.CreateSnapshot, true)
	assert.Equal(t, *resp.Properties.AccessActivityLog, true)
	assert.Equal(t, *resp.Properties.ReserveIp, true)
}

func TestCreateGroupFaliure(t *testing.T) {
	want := 422
	var obj = Group{
		Properties: GroupProperties{
			CreateDataCenter:  &TRUE,
			CreateSnapshot:    &TRUE,
			ReserveIp:         &TRUE,
			AccessActivityLog: &TRUE,
		},
	}
	resp := CreateGroup(obj)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, strings.Contains(resp.Response, "Attribute 'name' is required"))
}

func TestListGroups(t *testing.T) {
	SetDepth("5")
	want := 200
	resp := ListGroups()

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetGroup(t *testing.T) {
	want := 200
	resp := GetGroup(groupid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, groupid)
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, *resp.Properties.CreateDataCenter, true)
	assert.Equal(t, *resp.Properties.CreateSnapshot, true)
	assert.Equal(t, *resp.Properties.AccessActivityLog, true)
	assert.Equal(t, *resp.Properties.ReserveIp, true)
}

func TestGetGroupFailure(t *testing.T) {
	want := 404
	resp := GetGroup("00000000-0000-0000-0000-000000000000")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, strings.Contains(resp.Response, "Resource does not exist"))
}

func TestUpdateGroup(t *testing.T) {
	want := 202
	newName := "GO SDK Test - RENAME"
	obj := Group{
		Properties: GroupProperties{
			Name:              newName,
			CreateDataCenter:  &FALSE,
			CreateSnapshot:    &TRUE,
			ReserveIp:         &TRUE,
			AccessActivityLog: &TRUE,
		},
	}

	resp := UpdateGroup(groupid, obj)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	if resp.Properties.Name != newName {
		t.Errorf("Not updated")
	}

	assert.Equal(t, resp.Id, groupid)
	assert.Equal(t, resp.Properties.Name, newName)
	assert.Equal(t, *resp.Properties.CreateDataCenter, false)
	assert.Equal(t, resp.Type_, "group")
}

func TestAddShare(t *testing.T) {
	want := 202
	var obj = Share{
		Properties: ShareProperties{
			SharePrivilege: &TRUE,
			EditPrivilege:  &TRUE,
		},
	}
	resp := AddShare(obj, groupid, resourceId)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, *resp.Properties.EditPrivilege, true)
	assert.Equal(t, *resp.Properties.SharePrivilege, true)
}

func TestAddShareFailure(t *testing.T) {
	want := 422
	var obj = Share{
		Properties: ShareProperties{},
	}
	resp := AddShare(obj, groupid, resourceId)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestListShares(t *testing.T) {
	SetDepth("5")
	want := 200
	resp := ListShares(groupid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetShare(t *testing.T) {
	want := 200
	resp := GetShare(groupid, resourceId)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, resourceId)
	assert.Equal(t, *resp.Properties.EditPrivilege, true)
	assert.Equal(t, *resp.Properties.SharePrivilege, true)
}

func TestGetShareFailure(t *testing.T) {
	want := 404
	resp := GetShare(groupid, "00000000-0000-0000-0000-000000000000")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.True(t, strings.Contains(resp.Response, "Resource does not exist"))
}

func TestUpdateShare(t *testing.T) {
	want := 202
	obj := Share{
		Properties: ShareProperties{
			SharePrivilege: &TRUE,
			EditPrivilege:  &FALSE,
		},
	}

	resp := UpdateShare(groupid, resourceId, obj)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, resourceId)
	assert.Equal(t, *resp.Properties.EditPrivilege, false)
	assert.Equal(t, *resp.Properties.SharePrivilege, true)
}

func TestAddUserToGroup(t *testing.T) {
	want := 202

	resp := AddUserToGroup(groupid, userid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, userid)
	assert.Equal(t, resp.Type_, "user")
}

func TestListGroupUsers(t *testing.T) {
	want := 200
	resp := ListGroupUsers(groupid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestListResources(t *testing.T) {
	want := 200
	resp := ListResources()

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestListIPBlockResources(t *testing.T) {
	want := 200
	resp := ListResourcesByType("ipblock")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.True(t, len(resp.Items) > 0)
}

func TestListDatacenterResources(t *testing.T) {
	want := 200
	resp := ListResourcesByType("datacenter")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.True(t, len(resp.Items) > 0)
}

func TestListImagesResources(t *testing.T) {
	want := 200
	resp := ListResourcesByType("image")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.True(t, len(resp.Items) > 0)
}

func TestListSnapshotResources(t *testing.T) {
	want := 200
	resp := ListResourcesByType("snapshot")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.True(t, len(resp.Items) > 0)
}

func TestListResourceFailure(t *testing.T) {
	want := 404
	resp := ListResourcesByType("unknown")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.Equal(t, resp.StatusCode, want)
}

func TestGetDatacenterResource(t *testing.T) {
	want := 200
	resp := GetResourceByType("datacenter", resourceId)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.Equal(t, resp.Id, resourceId)
	assert.Equal(t, resp.Type_, "datacenter")
}

func TestGetIPBlockResource(t *testing.T) {
	want := 200
	resp := GetResourceByType("ipblock", ipblockId)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, ipblockId)
	assert.Equal(t, resp.Type_, "ipblock")
}

func TestGetImageResource(t *testing.T) {
	want := 200
	resp := GetResourceByType("image", imageResourceId)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, imageResourceId)
	assert.Equal(t, resp.Type_, "image")
}

func TestGetSnapshotResource(t *testing.T) {
	want := 200
	resp := GetResourceByType("snapshot", snapshotResourceId)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, snapshotResourceId)
	assert.Equal(t, resp.Type_, "snapshot")
}

func TestGetResourceFailure(t *testing.T) {
	want := 404
	resp := GetResourceByType("snapshot", "00000000-0000-0000-0000-000000000000")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	assert.Equal(t, resp.StatusCode, want)
}

func TestDeleteUserFromGroup(t *testing.T) {
	want := 202
	resp := DeleteUserFromGroup(groupid, userid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestDeleteShare(t *testing.T) {
	want := 202
	resp := DeleteShare(groupid, resourceId)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestDeleteGroup(t *testing.T) {
	want := 202
	resp := DeleteGroup(groupid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	//clean resources
	resp1 := DeleteDatacenter(resourceId)
	if resp1.StatusCode != want {
		t.Errorf(bad_status(want, resp1.StatusCode))
	}
}

func TestDeleteUser(t *testing.T) {
	want := 202
	resp := DeleteUser(userid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	CleanUpResources()
}

func CleanUpResources() {
	dcDeleted:=DeleteDatacenter(resourceId)
	waitTillProvisioned(dcDeleted.Headers.Get("Location"))

	snapshotDeleted:=DeleteSnapshot(snapshotResourceId)
	waitTillProvisioned(snapshotDeleted.Headers.Get("Location"))

	ReleaseIpBlock(ipblockId)
}
