package profitbricks

import (
	"net/http"
	"encoding/json"
	"strconv"
)

type ContractResources struct {
	Id         string                     `json:"id,omitempty"`
	Type_      string                     `json:"type,omitempty"`
	Href       string                     `json:"href,omitempty"`
	Properties ContractResourcesProperties       `json:"properties,omitempty"`
	Response   string                     `json:"Response,omitempty"`
	Headers    *http.Header               `json:"headers,omitempty"`
	StatusCode int                        `json:"headers,omitempty"`
}

type ContractResourcesProperties struct {
	PBContractNumber string `json:"PB-Contract-Number,omitempty"`
	Owner            string `json:"owner,omitempty"`
	Status           string `json:"status,omitempty"`
	ResourceLimits   *ResourcesLimits  `json:"resourceLimits,omitempty"`
}

type ResourcesLimits struct {
	CoresPerServer        int32       `json:"coresPerServer,omitempty"`
	CoresPerContract      int32       `json:"coresPerContract,omitempty"`
	CoresProvisioned      int32       `json:"coresProvisioned,omitempty"`
	RamPerServer          int64       `json:"ramPerServer,omitempty"`
	RamPerContract        int64       `json:"ramPerContract,omitempty"`
	RamProvisioned        int64       `json:"ramProvisioned,omitempty"`
	HddLimitPerVolume     int32       `json:"hddLimitPerVolume,omitempty"`
	HddLimitPerContract   int32       `json:"hddLimitPerContract,omitempty"`
	HddVolumeProvisioned  int32       `json:"hddVolumeProvisioned,omitempty"`
	SsdLimitPerVolume     int32       `json:"ssdLimitPerVolume,omitempty"`
	SsdLimitPerContract   int32       `json:"ssdLimitPerContract,omitempty"`
	SsdVolumeProvisioned  int32       `json:"ssdVolumeProvisioned,omitempty"`
	ReservableIps         int64       `json:"reservableIps,omitempty"`
	ReservedIpsOnContract int64       `json:"reservedIpsOnContract,omitempty"`
	ReservedIpsInUse      int64       `json:"reservedIpsInUse,omitempty"`
}

func GetContractResources() ContractResources {
	path := contract_resource_path()
	url := mk_url(path) + `?depth=` + Depth + `&pretty=` + strconv.FormatBool(Pretty)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", FullHeader)
	resp := do(req)
	return toContractResources(resp)
}

func toContractResources(resp Resp) ContractResources {
	var col ContractResources
	json.Unmarshal(resp.Body, &col)
	col.Response = string(resp.Body)
	col.Headers = &resp.Headers
	col.StatusCode = resp.StatusCode
	return col
}
