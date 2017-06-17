package profitbricks

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var dcID string
var compositedcId string

func TestListDatacenters(t *testing.T) {
	setupTestEnv()
	want := 200

	resp := ListDatacenters()

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.True(t, len(resp.Items) > 0)
}

func TestGetDatacenterFail(t *testing.T) {
	want := 404
	resp := GetDatacenter("231")
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}
}

func TestCreate(t *testing.T) {
	want := 202
	var obj = Datacenter{
		Properties: DatacenterProperties{
			Name:        "GO SDK Test",
			Description: "GO SDK test datacenter",
			Location:    location,
		},
	}
	resp := CompositeCreateDatacenter(obj)
	dcID = resp.Id

	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Type_, "datacenter")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Description, "GO SDK test datacenter")
	assert.Equal(t, resp.Properties.Location, location)
}

func TestCreateFailure(t *testing.T) {
	want := 422
	var obj = Datacenter{
		Properties: DatacenterProperties{
			Name:        "GO SDK Test",
			Description: "GO SDK test datacenter",
		},
	}
	resp := CompositeCreateDatacenter(obj)

	assert.Equal(t, resp.StatusCode, want)
}

func TestCreateComposite(t *testing.T) {
	want := 202
	var obj = Datacenter{
		Properties: DatacenterProperties{
			Name:        "GO SDK Test Composite",
			Description: "GO SDK test composite datacenter",
			Location:    location,
		},
		Entities: DatacenterEntities{
			Servers: &Servers{
				Items: []Server{
					{
						Properties: ServerProperties{
							Name:  "GO SDK Test",
							Ram:   1024,
							Cores: 1,
						},
					},
				},
			},
			Volumes: &Volumes{
				Items: []Volume{
					{
						Properties: VolumeProperties{
							Type:             "HDD",
							Size:             2,
							Name:             "GO SDK Test",
							Bus:              "VIRTIO",
							LicenceType:      "UNKNOWN",
							AvailabilityZone: "ZONE_3",
						},
					},
				},
			},
		},
	}
	resp := CompositeCreateDatacenter(obj)
	compositedcId = resp.Id
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Type_, "datacenter")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test Composite")
	assert.Equal(t, resp.Properties.Description, "GO SDK test composite datacenter")
	assert.Equal(t, resp.Properties.Location, location)
	assert.True(t, len(resp.Entities.Servers.Items) > 0)
	assert.True(t, len(resp.Entities.Volumes.Items) > 0)
}

func TestGetDatacenter(t *testing.T) {
	want := 200
	resp := GetDatacenter(dcID)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, dcID)
	assert.Equal(t, resp.Type_, "datacenter")
	assert.Equal(t, resp.Properties.Name, "GO SDK Test")
	assert.Equal(t, resp.Properties.Description, "GO SDK test datacenter")
	assert.Equal(t, resp.Properties.Location, location)
}

func TestPatchDatacenter(t *testing.T) {
	want := 202
	newName := "GO SDK Test - RENAME"
	obj := DatacenterProperties{Name: newName}

	resp := PatchDatacenter(dcID, obj)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	assert.Equal(t, resp.Id, dcID)
	assert.Equal(t, resp.Properties.Name, newName)
}

func TestDeleteDatacenter(t *testing.T) {
	want := 202
	resp := DeleteDatacenter(dcID)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, resp.StatusCode))
	}

	compositeResp := DeleteDatacenter(compositedcId)
	if resp.StatusCode != want {
		t.Errorf(bad_status(want, compositeResp.StatusCode))
	}
}
