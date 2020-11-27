// +build integration_tests integration_tests_backupunit

package integration_tests

import (
	sdk "github.com/ionos-cloud/ionos-enterprise-sdk-go/v5"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	BuName = "GO SDK Test Backup Unit"
	BuPass = "letmein"
	BuEmail = "gosdktest@mailinator.com"
	BuNameUpdated = BuName + " UPDATED"
	BuPassUpdated = BuPass + "UPDATED"
	BuEmailUpdated = "updated_" + BuEmail
)

func TestCreateBackupUnit(t *testing.T) {
	c := setupTestEnv()
	input := sdk.BackupUnit{
		Properties: &sdk.BackupUnitProperties{
			Name: BuName,
			Password: BuPass,
			Email: BuEmail,
		},
	}
	var err error
	backupUnit, err = c.CreateBackupUnit(input)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, backupUnit)
	assert.NotNil(t, backupUnit.Properties)
	assert.Equal(t, BuName, backupUnit.Properties.Name)
	assert.Equal(t, BuEmail, backupUnit.Properties.Email)
}

func TestListBackupUnits(t *testing.T) {
	c := setupTestEnv()
	ret, err := c.ListBackupUnits()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, ret)
	assert.True(t, len(ret.Items) > 0)
	assert.Equal(t, backupUnit.Properties.Name, ret.Items[0].Properties.Name)
	assert.Equal(t, backupUnit.Properties.Email, ret.Items[0].Properties.Email)
	assert.Equal(t, backupUnit.Properties.Password, ret.Items[0].Properties.Password)
}

func TestGetBackupUnit(t *testing.T) {
	c := setupTestEnv()
	ret, err := c.GetBackupUnit(backupUnit.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, ret)
	assert.Equal(t, backupUnit.Properties.Name, ret.Properties.Name)
	assert.Equal(t, backupUnit.Properties.Email, ret.Properties.Email)
	assert.Equal(t, backupUnit.Properties.Password, ret.Properties.Password)
}

func TestUpdateBackupUnit(t *testing.T) {
	c := setupTestEnv()
	ret, err := c.UpdateBackupUnit(backupUnit.ID, sdk.BackupUnit{
		Properties: &sdk.BackupUnitProperties{
			/* we're not allowed to update the name, only email and password and the api returns name and email */
			Email: BuEmailUpdated,
			Password: BuPassUpdated,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, ret)
	assert.NotNil(t, ret.Properties)
	assert.Equal(t, BuName, ret.Properties.Name)
	assert.Equal(t, BuEmailUpdated, ret.Properties.Email)
}

func TestDeleteBackupUnit(t *testing.T) {
	c := setupTestEnv()
	_, err := c.DeleteBackupUnit(backupUnit.ID)
	if err != nil {
		t.Fatal(err)
	}
}
