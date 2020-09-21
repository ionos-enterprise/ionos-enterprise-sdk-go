package integration_tests

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	sdk "github.com/ionos-cloud/ionos-enterprise-sdk-go/v5"
	"github.com/stretchr/testify/assert"
)

var (
	groupid     string
	user        *sdk.User
	group       *sdk.Group
	email       string
	TRUE        = true
	FALSE       = false
	onceUmDC    sync.Once
	onceUmUser  sync.Once
	onceUmGroup sync.Once
	onceUmIP    sync.Once
	onceUmShare sync.Once
	s3Key		*sdk.S3Key
	onceUmS3Key	sync.Once
)

const s3KeySecret = "testsecret"

func createUser() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	email = "test" + strconv.Itoa(r1.Intn(1000)) + "@go.com"
	c := setupTestEnv()
	var obj = sdk.User{
		Properties: &sdk.UserProperties{
			Firstname:     "John",
			Lastname:      "Doe",
			Email:         email,
			Password:      "abc123-321CBA",
			Administrator: false,
			ForceSecAuth:  false,
		},
	}
	resp, err := c.CreateUser(obj)
	if err != nil {
		fmt.Println("[error] error creating user: ", err)
		fmt.Println(resp.Response)
		os.Exit(1)
	}
	user = resp
}

func createS3Key() {

	onceUmUser.Do(createUser)

	c := setupTestEnv()
	key, err := c.CreateS3Key(user.ID)

	/* TODO: remove hack when fixed upstream: we're ignoring time parsing errors until the cloud api fixes
	 * the createdTime in the s3KeyMetadata field returned by post - it should have
	 * a trailing 'Z' */
	if err != nil && !strings.Contains(err.Error(), "parsing time") {
		fmt.Println("[error] error creating S3 key: ", err)
		os.Exit(1)
	}

	s3Key = key
}

func createGroup() {
	c := setupTestEnv()
	var obj = sdk.Group{
		Properties: sdk.GroupProperties{
			Name:              "GO SDK Test",
			CreateDataCenter:  &TRUE,
			CreateSnapshot:    &TRUE,
			ReserveIP:         &TRUE,
			AccessActivityLog: &TRUE,
		},
	}
	resp, err := c.CreateGroup(obj)
	if err != nil {
		fmt.Println("[error] error creating group: ", err)
		fmt.Println(resp.Response)
		os.Exit(1)
	}
	group = resp
}

func addShare() {
	c := setupTestEnv()
	var obj = sdk.Share{
		Properties: sdk.ShareProperties{
			SharePrivilege: &TRUE,
			EditPrivilege:  &TRUE,
		},
	}
	c.AddShare(group.ID, dataCenter.ID, obj)
}
func TestCreateUser(t *testing.T) {
	fmt.Println("User management tests")
	onceUmUser.Do(createUser)

	assert.Equal(t, user.Properties.Firstname, "John")
	assert.Equal(t, user.Properties.Lastname, "Doe")
	assert.Equal(t, user.Properties.Email, email)
	assert.Equal(t, user.Properties.Administrator, false)
}

func TestCreateUserFailure(t *testing.T) {
	c := setupTestEnv()
	var obj = sdk.User{
		Properties: &sdk.UserProperties{
			Firstname:     "John",
			Lastname:      "Doe",
			Password:      "abc123-321CBA",
			Administrator: true,
			ForceSecAuth:  false,
			SecAuthActive: false,
		},
	}
	_, err := c.CreateUser(obj)
	assert.NotNil(t, err)
}

func TestListUsers(t *testing.T) {
	c := setupTestEnv()
	onceUmUser.Do(createUser)
	resp, err := c.ListUsers()
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetUser(t *testing.T) {
	c := setupTestEnv()
	onceUmUser.Do(createUser)

	resp, err := c.GetUser(user.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, user.ID)
	assert.Equal(t, resp.Properties.Firstname, "John")
	assert.Equal(t, resp.Properties.Lastname, "Doe")
	assert.Equal(t, resp.Properties.Email, email)
	assert.Equal(t, resp.Properties.Administrator, false)
	assert.Equal(t, resp.PBType, "user")
}

func TestUpdateUser(t *testing.T) {
	c := setupTestEnv()
	newName := "user updated"
	obj := sdk.User{
		Properties: &sdk.UserProperties{
			Firstname:     "John",
			Lastname:      newName,
			Email:         email,
			Administrator: false,
			ForceSecAuth:  false,
			SecAuthActive: false,
		}}

	resp, err := c.UpdateUser(user.ID, obj)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.Properties.Lastname, newName)
}

func TestCreateGroup(t *testing.T) {
	onceUmGroup.Do(createGroup)
	assert.Equal(t, group.Properties.Name, "GO SDK Test")
	assert.Equal(t, *group.Properties.CreateDataCenter, true)
	assert.Equal(t, *group.Properties.CreateSnapshot, true)
	assert.Equal(t, *group.Properties.AccessActivityLog, true)
	assert.Equal(t, *group.Properties.ReserveIP, true)
}

func TestCreateGroupFaliure(t *testing.T) {
	c := setupTestEnv()
	var obj = sdk.Group{
		Properties: sdk.GroupProperties{
			CreateDataCenter:  &TRUE,
			CreateSnapshot:    &TRUE,
			ReserveIP:         &TRUE,
			AccessActivityLog: &TRUE,
		},
	}
	_, err := c.CreateGroup(obj)

	assert.NotNil(t, err)
}

func TestListGroups(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	resp, err := c.ListGroups()
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetGroup(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	resp, err := c.GetGroup(group.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.ID, group.ID)
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, *resp.Properties.CreateDataCenter, true)
	assert.Equal(t, *resp.Properties.CreateSnapshot, true)
	assert.Equal(t, *resp.Properties.AccessActivityLog, true)
	assert.Equal(t, *resp.Properties.ReserveIP, true)
}

func TestGetGroupFailure(t *testing.T) {
	c := setupTestEnv()
	_, err := c.GetGroup("00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, err)
}

func TestUpdateGroup(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)

	newName := "GO SDK Test - RENAME"
	obj := sdk.Group{
		Properties: sdk.GroupProperties{
			Name:              newName,
			CreateDataCenter:  &FALSE,
			CreateSnapshot:    &TRUE,
			ReserveIP:         &TRUE,
			AccessActivityLog: &TRUE,
		},
	}

	resp, err := c.UpdateGroup(group.ID, obj)
	if err != nil {
		if resp != nil {
			t.Error(resp.Response)
		}
		t.Fatal(err)
	}

	assert.Equal(t, resp.Properties.Name, newName)
}

func TestAddShare(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	onceUmDC.Do(createDataCenter)

	var obj = sdk.Share{
		Properties: sdk.ShareProperties{
			SharePrivilege: &TRUE,
			EditPrivilege:  &TRUE,
		},
	}
	resp, err := c.AddShare(group.ID, dataCenter.ID, obj)
	if err != nil {
		t.Error(err)
	}

	share = resp
	assert.Equal(t, *resp.Properties.EditPrivilege, true)
	assert.Equal(t, *resp.Properties.SharePrivilege, true)
}

func TestAddShareFailure(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	onceUmDC.Do(createDataCenter)

	var obj = sdk.Share{
		Properties: sdk.ShareProperties{},
	}
	_, err := c.AddShare(groupid, dataCenter.ID, obj)
	assert.NotNil(t, err)
}

func TestListShares(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)

	resp, err := c.ListShares(group.ID)
	if err != nil {
		if resp != nil {
			t.Error(resp.Response)
		}
		t.Fatal(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetShare(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	onceUmDC.Do(createDataCenter)

	resp, err := c.GetShare(group.ID, share.ID)
	if err != nil {
		if resp != nil {
			t.Error(resp.Response)
		}
		t.Fatal(err)
	}

	assert.Equal(t, resp.ID, dataCenter.ID)
	assert.Equal(t, *resp.Properties.EditPrivilege, true)
	assert.Equal(t, *resp.Properties.SharePrivilege, true)
}

func TestGetShareFailure(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	onceUmDC.Do(createDataCenter)
	_, err := c.GetShare(group.ID, "00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, err)
}

func TestUpdateShare(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	onceUmDC.Do(createDataCenter)

	obj := sdk.Share{
		Properties: sdk.ShareProperties{
			SharePrivilege: &TRUE,
			EditPrivilege:  &FALSE,
		},
	}

	resp, err := c.UpdateShare(group.ID, share.ID, obj)
	if err != nil {
		if resp != nil {
			t.Error(resp.Response)
		}
		t.Fatal(err)
	}

	assert.Equal(t, resp.ID, dataCenter.ID)
	assert.Equal(t, *resp.Properties.EditPrivilege, false)
	assert.Equal(t, *resp.Properties.SharePrivilege, true)
}

func TestAddUserToGroup(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	onceUmUser.Do(createUser)

	resp, err := c.AddUserToGroup(group.ID, user.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, user.ID)
	assert.Equal(t, resp.PBType, "user")
}

func TestListGroupUsers(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	onceUmUser.Do(createUser)

	resp, err := c.ListGroupUsers(group.ID)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestListResources(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.ListResources()
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestListIPBlockResources(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.ListResourcesByType("ipblock")
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestListDatacenterResources(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.ListResourcesByType("datacenter")
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestListImagesResources(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.ListResourcesByType("image")
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestListSnapshotResources(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.ListResourcesByType("snapshot")
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestListResourceFailure(t *testing.T) {
	c := setupTestEnv()
	_, err := c.ListResourcesByType("unknown")

	assert.NotNil(t, err)
}

func TestGetDatacenterResource(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.GetResourceByType("datacenter", dataCenter.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, dataCenter.ID)
	assert.Equal(t, resp.PBType, "datacenter")
}

func TestGetImageResource(t *testing.T) {
	c := setupTestEnv()
	imageResourceID := getImageID("us/las", "ubuntu", "hdd")
	_, err := c.GetResourceByType("image", imageResourceID)
	if err != nil {
		t.Error(err)
	}
}

func TestGetResourceFailure(t *testing.T) {
	c := setupTestEnv()
	_, err := c.GetResourceByType("snapshot", "00000000-0000-0000-0000-000000000000")

	assert.NotNil(t, err)
}


func TestCreateS3Key(t *testing.T) {
	onceUmS3Key.Do(createS3Key)

	assert.NotNil(t, s3Key)
	assert.True(t, s3Key.ID != "")
}

func TestListS3Keys(t *testing.T) {
	c := setupTestEnv()
	onceUmUser.Do(createUser)
	onceUmS3Key.Do(createS3Key)

	keys, err := c.ListS3Keys(user.ID)

	/* TODO: remove hack when fixed upstream: we're ignoring time parsing errors until the cloud api fixes
	 * the createdTime in the s3KeyMetadata field returned by post - it should have
	 * a trailing 'Z' */
	if err != nil && !strings.Contains(err.Error(), "parsing time") {
		t.Fatal(err)
	}

	assert.NotNil(t, keys)
	if keys != nil {
		assert.True(t, len(keys.Items) > 0)
	} else {
		t.Fatal(errors.New("ListS3Keys returned nil"))
	}
}

func TestUpdateS3Key(t *testing.T) {

	onceUmS3Key.Do(createS3Key)

	c := setupTestEnv()

	keyUpdate := sdk.S3Key{
		Properties: &sdk.S3KeyProperties{
			Active: false,
		},
	}
	key, err := c.UpdateS3Key(user.ID, s3Key.ID, keyUpdate)

	/* TODO: remove hack when fixed upstream: we're ignoring time parsing errors until the cloud api fixes
	 * the createdTime in the s3KeyMetadata field returned by post - it should have
	 * a trailing 'Z' */
	if err != nil && !strings.Contains(err.Error(), "parsing time") {
		t.Fatal(err)
	}

	assert.NotNil(t, key)
	if key != nil {
		assert.Equal(t, s3Key.ID, key.ID)
		/* TODO: uncomment this when upstream fixes createdDate in S3KeyMetadata
		 * for now, because the parser fails at Metadata, Properties are left nil
		 */
		// assert.Equal(t, false, key.Properties.Active)
	} else {
		t.Fatal("UpdateS3Key returned nil")
	}
}

func TestGetS3Key(t *testing.T) {
	c := setupTestEnv()
	onceUmUser.Do(createUser)
	onceUmS3Key.Do(createS3Key)

	key, err := c.GetS3Key(user.ID, s3Key.ID)

	/* TODO: remove hack when fixed upstream: we're ignoring time parsing errors until the cloud api fixes
	 * the createdTime in the s3KeyMetadata field returned by post - it should have
	 * a trailing 'Z' */
	if err != nil && !strings.Contains(err.Error(), "parsing time") {
		t.Fatal(err)
	}

	assert.NotNil(t, key)
	if key != nil {
		assert.Equal(t, s3Key.ID, key.ID)
	} else {
		t.Fatal("GetS3Key returned nil")
	}
}

func TestDeleteS3Key(t *testing.T) {
	c := setupTestEnv()
	onceUmUser.Do(createUser)
	onceUmS3Key.Do(createS3Key)

	_, err := c.DeleteS3Key(user.ID, s3Key.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserFromGroup(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	onceUmUser.Do(createUser)
	_, err := c.DeleteUserFromGroup(group.ID, user.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteShare(t *testing.T) {
	c := setupTestEnv()
	onceUmGroup.Do(createGroup)
	onceUmDC.Do(createDataCenter)
	onceUmShare.Do(addShare)
	_, err := c.DeleteShare(group.ID, dataCenter.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteGroup(t *testing.T) {
	c := setupTestEnv()
	onceUmDC.Do(createDataCenter)
	onceUmGroup.Do(createGroup)
	_, err := c.DeleteGroup(group.ID)
	if err != nil {
		t.Error(err)
	}
	//clean resources
	_, err = c.DeleteDatacenter(dataCenter.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteUser(t *testing.T) {
	c := setupTestEnv()
	onceUmUser.Do(createUser)
	_, err := c.DeleteUser(user.ID)
	if err != nil {
		t.Error(err)
	}
}
