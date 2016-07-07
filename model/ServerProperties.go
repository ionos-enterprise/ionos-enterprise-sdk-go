package model

type ServerProperties struct {
	Name             string  `json:"name,omitempty"`
	Cores            int32  `json:"cores,omitempty"`
	Ram              int32  `json:"ram,omitempty"`
	AvailabilityZone string  `json:"availabilityZone,omitempty"`
	VmState          string  `json:"vmState,omitempty"`
}
