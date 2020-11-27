package profitbricks

import (
	"context"
	ionossdk "github.com/ionos-cloud/ionos-cloud-sdk-go/v5"
	"net/http"
)

// S3Keys type
type S3Keys struct {
	// Enum: [backupunits]
	// Read Only: true
	ID string `json:"id,omitempty"`
	// Enum: [collection]
	// Read Only: true
	Type string `json:"type,omitempty"`
	// Format: uri
	Href string `json:"href"`
	// Read Only: true
	Items []S3Key `json:"items"`
}

// S3Key Object
type S3Key struct {
	// URL to the object representation (absolute path)
	// Read Only: true
	// Format: uri
	Href string `json:"href,omitempty"`

	// The resource's unique identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The type of object. In this case s3key
	// Read Only: true
	Type string `json:"type,omitempty"`

	// The metadata for the S3 key
	// Read Only: true
	Metadata *Metadata `json:"metadata,omitempty"`

	// The properties of the S3 key
	// Read Only: false
	Properties *S3KeyProperties `json:"properties,omitempty"`
}

// S3KeyProperties object
type S3KeyProperties struct {
	// Read only: yes
	SecretKey string `json:"secretKey,omitempty"`
	// Required: yes
	// Read only: no
	Active bool `json:"active"`
}

// CreateS3Key creates an S3 Key for an user
func (c *Client) CreateS3Key(userID string) (*S3Key, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, _, err := c.CoreSdk.UserManagementApi.UmUsersS3keysPost(ctx, userID).Execute()
	ret := S3Key{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	return &ret, err
	/*
	rsp := &S3Key{}
	var requestBody interface{}
	err := c.Post(s3KeysPath(userID), requestBody, rsp, http.StatusCreated)
	return rsp, err
	 */
}

// ListS3Keys lists all available S3 keys for an user
func (c *Client) ListS3Keys(userID string) (*S3Keys, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	rsp, _, err := c.CoreSdk.UserManagementApi.UmUsersS3keysGet(ctx, userID).Execute()
	ret := S3Keys{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	return &ret, err
	/*
	rsp := &S3Keys{}
	return rsp, c.GetOK(s3KeysListPath(userID), rsp)
	 */
}

// UpdateS3Key updates an existing S3 key
func (c *Client) UpdateS3Key(userID string, s3KeyID string, s3Key S3Key) (*S3Key, error) {

	input := ionossdk.S3Key{}
	if errConvert := convertToCore(&s3Key, &input); errConvert != nil {
		return nil, errConvert
	}
	rsp, _, err := c.CoreSdk.UserManagementApi.UmUsersS3keysPut(
		context.TODO(), userID, s3KeyID).S3Key(input).Execute()
	ret := S3Key{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}

	return &ret, err
	/*
	rsp := &S3Key{}
	return rsp, c.PutAcc(s3KeyPath(userID, s3KeyID), s3Key, rsp)
	 */
}

// DeleteS3Key deletes an existing S3 key
func (c *Client) DeleteS3Key(userID string, s3KeyID string) (*http.Header, error) {

    ctx, cancel := c.GetContext()
    if cancel != nil { defer cancel() }
	_, apiResponse, err := c.CoreSdk.UserManagementApi.UmUsersS3keysDelete(ctx, userID, s3KeyID).Execute()
	if apiResponse != nil {
		return &apiResponse.Header, err
	} else {
		return nil, err
	}
}

// GetS3Key retrieves an existing S3 key
func (c *Client) GetS3Key(userID string, s3KeyID string) (*S3Key, error) {

	rsp, _, err := c.CoreSdk.UserManagementApi.UmUsersS3keysFindByKeyId(
		context.TODO(), userID, s3KeyID).Execute()
	ret := S3Key{}
	if errConvert := convertToCompat(&rsp, &ret); errConvert != nil {
		return nil, errConvert
	}
	return &ret, err

	/*
	rsp := &S3Key{}
	return rsp, c.GetOK(s3KeyPath(userID, s3KeyID), rsp)
	 */
}
