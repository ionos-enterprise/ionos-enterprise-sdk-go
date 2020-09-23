package profitbricks

import (
	"github.com/ionos-cloud/ionos-cloud-sdk-go/v5"
	"net/http"
)


// BackupUnits type
type BackupUnits struct {
	// Enum: [backupunits]
	// Read Only: true
	ID string `json:"id,omitempty"`
	// Enum: [collection]
	// Read Only: true
	Type string `json:"type,omitempty"`
	// Format: uri
	Href string `json:"href"`
	// Read Only: true
	Items []BackupUnit `json:"items"`
}

// BackupUnit Object
type BackupUnit struct {
	// URL to the object representation (absolute path)
	// Read Only: true
	// Format: uri
	Href string `json:"href,omitempty"`

	// The resource's unique identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The type of object. In this case backupunit
	// Read Only: true
	Type string `json:"type,omitempty"`

	// The metadata for the backup unit
	// Read Only: true
	Metadata *Metadata `json:"metadata,omitempty"`

	Properties *BackupUnitProperties `json:"properties,omitempty"`
}

// BackupUnitProperties object
type BackupUnitProperties struct {
	// Required: on create
	// Read only: yes
	Name string `json:"name,omitempty"`
	// Required: yes
	Password string `json:"password,omitempty"`
	// Required: yes
	Email string `json:"email,omitempty"`
}

// BackupUnitSSOURL object
type BackupUnitSSOURL struct {

	// The type of object. In this case backupunit
	// Read Only: true
	Type string `json:"type,omitempty"`
	// SSO URL

	// Read Only: true
	SSOUrl string `json:"ssoURL,omitempty"`
}

//

// CreateBackupUnit creates a Backup Unit
func (c *Client) CreateBackupUnit(backupUnit BackupUnit) (*BackupUnit, error) {
	// rsp := &BackupUnit{}
	input := ionossdk.BackupUnit{}
	err := convertToCore(&backupUnit, &input)
	
	ctx, cancel := c.GetContext()
	if cancel != nil { defer cancel() }
	
	rsp, _, err := c.CoreSdk.BackupUnitApi.BackupunitsPost(ctx, input, nil)

	ret := BackupUnit{}
	errConvert := convertToCompat(&rsp, &ret)
	if errConvert != nil {
		return nil, errConvert
	}
	return &ret, err
	// return rsp, c.PostAcc(backupUnitsPath(), backupUnit, rsp)
}

// ListBackupUnits lists all available backup units
func (c *Client) ListBackupUnits() (*BackupUnits, error) {
	
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, _, err := c.CoreSdk.BackupUnitApi.BackupunitsGet(ctx, nil)
	ret := BackupUnits{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	return &ret, err
}

// UpdateBackupUnit updates an existing backup unit
func (c *Client) UpdateBackupUnit(backupUnitID string, backupUnit BackupUnit) (*BackupUnit, error) {

	input := ionossdk.BackupUnit{}
	if errConv := convertToCore(&backupUnit, &input); errConv != nil {
		return nil, errConv
	}
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, _, err := c.CoreSdk.BackupUnitApi.BackupunitsPut(ctx, backupUnitID, input, nil)
	ret := BackupUnit{}
	if errConv := convertToCompat(&rsp, &ret); errConv != nil {
		return nil, errConv
	}
	return &ret, err
	// rsp := &BackupUnit{}
	// return rsp, c.PutAcc(backupUnitPath(backupUnitID), backupUnit, rsp)
}

// DeleteBackupUnit deletes an existing backup unit
func (c *Client) DeleteBackupUnit(backupUnitID string) (*http.Header, error) {
    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, httpResponse, err := c.CoreSdk.BackupUnitApi.BackupunitsDelete(ctx, backupUnitID, nil)
	if httpResponse == nil || err != nil {
		return nil, err
	}
	return &httpResponse.Header, err
	// return c.DeleteAcc(backupUnitPath(backupUnitID))
}

// GetBackupUnit retrieves an existing backup unit
func (c *Client) GetBackupUnit(backupUnitID string) (*BackupUnit, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, _, err := c.CoreSdk.BackupUnitApi.BackupunitsFindById(ctx, backupUnitID, nil)
	ret := BackupUnit{}
	if errConv := convertToCompat(&rsp, &ret); errConv != nil {
		return nil, errConv
	}
	return &ret, err

	// rsp := &BackupUnit{}
	// return rsp, c.GetOK(backupUnitPath(backupUnitID), rsp)
}

// GetBackupUnitSSOURL retrieves the SSO URL for an existing backup unit
func (c *Client) GetBackupUnitSSOURL(backupUnitID string) (*BackupUnitSSOURL, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, _, err := c.CoreSdk.BackupUnitApi.BackupunitsSsourlGet(ctx, backupUnitID, nil)

	if err != nil {
		return nil, err
	}

	return &BackupUnitSSOURL{
		Type: "backupunit",
		SSOUrl: *rsp.SsoUrl,
	}, err
	// rsp := &BackupUnitSSOURL{}
	// return rsp, c.GetOK(backupUnitSSOURLPath(backupUnitID), rsp)
}
