// +build integration_tests integration_tests_pcc

package integration_tests

import (
	"fmt"
	sdk "github.com/ionos-cloud/ionos-enterprise-sdk-go/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	pcc *sdk.PrivateCrossConnect
)

func createSecondDc() {
	c := setupTestEnv()

	var obj = sdk.Datacenter{
		Properties: sdk.DatacenterProperties{
			Name:        "PCC VDC 2",
			Description: "GO SDK test datacenter for pcc",
			Location:    location,
		},
	}
	resp, err := c.CreateDatacenter(obj)
	if err != nil {
		panic(err)
	}
	err = c.WaitTillProvisioned(resp.Headers.Get("Location"))
	if err != nil {
		panic(err)
	}

}

func TestCreatePcc(t *testing.T) {
	fmt.Println("PCC tests")

	c := setupTestEnv()
	p, err := c.CreatePrivateCrossConnect(sdk.PrivateCrossConnect{
		Properties: &sdk.PrivateCrossConnectProperties{
			Peers:                  nil,
			ConnectableDatacenters: nil,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, p)
	pcc = p
}

func TestListPcc(t *testing.T) {
	c := setupTestEnv()

	pccs, err := c.ListPrivateCrossConnects()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, pccs)
	assert.True(t, len(pccs.Items) > 0)
}

func TestGetPcc(t *testing.T) {
	c := setupTestEnv()

	if pcc == nil {
		t.Fatal("PCC was not created")
	}
	p, err := c.GetPrivateCrossConnect(pcc.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, p)
	assert.Equal(t, pcc.ID, p.ID)
}

func TestUpdatePcc(t *testing.T) {
	c := setupTestEnv()
	if pcc == nil {
		t.Fatal("PCC was not created")
	}
	p, err := c.UpdatePrivateCrossConnect(pcc.ID, *pcc)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, p)
	assert.Equal(t, pcc.ID, p.ID)
}

func TestDeletePcc(t *testing.T) {
	c := setupTestEnv()
	if pcc == nil {
		t.Fatal("PCC was not created")
	}
	_, err := c.DeletePrivateCrossConnect(pcc.ID)
	if err != nil {
		t.Fatal(err)
	}
}
