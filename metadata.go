package profitbricks

// Metadata represents metadata recieved from Cloud API
type Metadata struct {
	CreatedDate          string `json:"createdDate,omitempty"`
	CreatedBy            string `json:"createdBy,omitempty"`
	CreatedByUserID      string `json:"createdByUserId,omitempty"`
	Etag                 string `json:"etag,omitempty"`
	LastModifiedDate     string `json:"lastModifiedDate,omitempty"`
	LastModifiedBy       string `json:"lastModifiedBy,omitempty"`
	LastModifiedByUserID string `json:"lastModifiedByUserId,omitempty"`
	State                string `json:"state,omitempty"`
}

// GetMetadata returns the metadata of the embedding object
func (m *Metadata) GetMetadata() *Metadata {
	return m
}
