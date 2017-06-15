package profitbricks

import (
	"testing"
	"math/rand"
	"strconv"
	"time"
)

var groupid string
var resourceId string
var userid string

func TestCreateUser(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	//create a datacenter to be used as a resource later on
	var dcObj = Datacenter{
		Properties: DatacenterProperties{
			Name:        "GO SDK",
			Description: "description",
			Location:    location,
		},
	}
	dc := CompositeCreateDatacenter(dcObj)
	resourceId = dc.Id
	want := 202
	var obj = User{
		Properties: &UserProperties{
			Firstname:     "go sdk",
			Lastname:      "user",
			Email:         "test" + strconv.Itoa(r1.Intn(100)) + "@go.com",
			Password:      "abc123-321CBA",
			Administrator: false,
			ForceSecAuth:  false,
			SecAuthActive: false,
		},
	}
	resp := CreateUser(obj)
	userid = resp.Id

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestListUsers(t *testing.T) {
	SetDepth("5")
	want := 200
	resp := ListUsers()

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestGetUser(t *testing.T) {
	want := 200
	resp := GetUser(userid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestUpdateUser(t *testing.T) {
	//want := 202
	//newName := "user updated"
	//obj := UserProperties{
	//	Firstname:     "go sdk ",
	//	Lastname:      newName,
	//	Email:         "test@go.com",
	//	Password:      "abc123-321CBA",
	//	Administrator: false,
	//	ForceSecAuth:  false,
	//	SecAuthActive: false,
	//}
	//
	//resp := UpdateUser(groupid, obj)
	//if resp.StatusCode != want {
	//	t.Errorf(bad_status(want, resp.StatusCode))
	//}
	//if resp.Properties.Lastname != newName {
	//	t.Errorf("Not updated")
	//}
}

func TestCreateGroup(t *testing.T) {
	want := 202
	var obj = Group{
		Properties: GroupProperties{
			Name:              "GO SDK Test",
			CreateDataCenter:  true,
			CreateSnapshot:    true,
			ReserveIp:         false,
			AccessActivityLog: false,

		},
	}
	resp := CreateGroup(obj)
	groupid = resp.Id

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestListGroups(t *testing.T) {
	SetDepth("5")
	want := 200
	resp := ListGroups()

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestGetGroup(t *testing.T) {
	want := 200
	resp := GetGroup(groupid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestUpdateGroup(t *testing.T) {
	want := 202
	newName := "Renamed group"
	obj := GroupProperties{
		Name:              newName,
		CreateSnapshot:    true,
		CreateDataCenter:  false,
		ReserveIp:         true,
		AccessActivityLog: false,
	}

	resp := UpdateGroup(groupid, obj)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
	if resp.Properties.Name != newName {
		t.Errorf("Not updated")
	}
}

func TestAddShare(t *testing.T) {
	want := 202
	var obj = Share{
		Properties: ShareProperties{
			SharePrivilege: true,
			EditPrivilege:  true,
		},
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
}

func TestGetShare(t *testing.T) {
	want := 200
	resp := GetShare(groupid, resourceId)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestUpdateShare(t *testing.T) {
	want := 202
	obj := ShareProperties{SharePrivilege: true, EditPrivilege: false, }

	resp := UpdateShare(groupid, resourceId, obj)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestAddUserToGroup(t *testing.T) {
	want := 202

	resp := AddUserToGroup(groupid, userid)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestListGroupUsers(t *testing.T) {
	want := 200
	resp := ListGroupUsers(groupid)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestListResources(t *testing.T) {
	want := 200
	resp := ListResources()

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestGetResourceByType(t *testing.T) {
	want := 200
	resp := GetResourceByType("datacenter", resourceId)

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestListResourcesByType(t *testing.T) {
	want := 200
	resp := ListResourcesByType("datacenter")

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
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
}
