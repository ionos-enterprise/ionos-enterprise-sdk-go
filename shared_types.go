package profitbricks

import (
	"github.com/go-openapi/strfmt"
)

// NoStateMetaData no state meta data
type NoStateMetaData struct {

	// The user who has created the resource.
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// The user id of the user who has created the resource.
	// Read Only: true
	CreatedByUserID string `json:"createdByUserId,omitempty"`

	// The time the Resource was created
	// Read Only: true
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// Resource's Entity Tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11 . Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter.
	// Read Only: true
	Etag string `json:"etag,omitempty"`

	// The user who last modified the resource.
	// Read Only: true
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// The user id of the user who has last modified the resource.
	// Read Only: true
	LastModifiedByUserID string `json:"lastModifiedByUserId,omitempty"`

	// The last time the resource has been modified
	// Read Only: true
	// Format: date-time
	LastModifiedDate strfmt.DateTime `json:"lastModifiedDate,omitempty"`
}
