package profitbricks

// Image object
type Image struct {
	BaseResource `json:",inline"`
	ID           string          `json:"id,omitempty"`
	PBType       string          `json:"type,omitempty"`
	Href         string          `json:"href,omitempty"`
	Metadata     *Metadata       `json:"metadata,omitempty"`
	Properties   ImageProperties `json:"properties,omitempty"`
	StatusCode   int             `json:"statuscode,omitempty"`
}

// ImageProperties object
type ImageProperties struct {
	Name                string   `json:"name,omitempty"`
	Description         string   `json:"description,omitempty"`
	Location            string   `json:"location,omitempty"`
	Size                float64  `json:"size,omitempty"`
	CPUHotPlug          bool     `json:"cpuHotPlug,omitempty"`
	CPUHotUnplug        bool     `json:"cpuHotUnplug,omitempty"`
	RAMHotPlug          bool     `json:"ramHotPlug,omitempty"`
	RAMHotUnplug        bool     `json:"ramHotUnplug,omitempty"`
	NicHotPlug          bool     `json:"nicHotPlug,omitempty"`
	NicHotUnplug        bool     `json:"nicHotUnplug,omitempty"`
	DiscVirtioHotPlug   bool     `json:"discVirtioHotPlug,omitempty"`
	DiscVirtioHotUnplug bool     `json:"discVirtioHotUnplug,omitempty"`
	DiscScsiHotPlug     bool     `json:"discScsiHotPlug,omitempty"`
	DiscScsiHotUnplug   bool     `json:"discScsiHotUnplug,omitempty"`
	LicenceType         string   `json:"licenceType,omitempty"`
	ImageType           string   `json:"imageType,omitempty"`
	ImageAliases        []string `json:"imageAliases,omitempty"`
	Public              bool     `json:"public,omitempty"`
}

// Images object
type Images struct {
	BaseResource `json:",inline"`
	ID           string  `json:"id,omitempty"`
	PBType       string  `json:"type,omitempty"`
	Href         string  `json:"href,omitempty"`
	Items        []Image `json:"items,omitempty"`
	StatusCode   int     `json:"statuscode,omitempty"`
}

// Cdroms object
type Cdroms struct {
	BaseResource `json:",inline"`
	ID           string  `json:"id,omitempty"`
	PBType       string  `json:"type,omitempty"`
	Href         string  `json:"href,omitempty"`
	Items        []Image `json:"items,omitempty"`
}

// ListImages returns an Collection struct
func (c *Client) ListImages() (*Images, error) {
	ret := &Images{}
	return ret, c.GetOK(imagesPath(), ret)

}

// GetImage returns an Instance struct where id ==imageid
func (c *Client) GetImage(imageId string) (*Image, error) {
	ret := &Image{}
	return ret, c.GetOK(imagePath(imageId), ret)
}
