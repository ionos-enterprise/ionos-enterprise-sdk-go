module github.com/ionos-enterprise/ionos-enterprise-sdk-go

require (
	github.com/antihax/optional v1.0.0
	github.com/go-resty/resty/v2 v2.3.0
	github.com/ionos-cloud/ionos-cloud-sdk-go/v5 v5.0.0-00010101000000-000000000000
	github.com/ionos-enterprise/ionos-enterprise-sdk-go/v6 v6.0.0-00010101000000-000000000000
	github.com/jarcoal/httpmock v1.0.5
	github.com/stretchr/testify v1.6.1
)

replace github.com/ionos-enterprise/ionos-enterprise-sdk-go/v6 => /Users/florin/work/code/cloud-sdk/sdk/compat/go

replace github.com/ionos-cloud/ionos-cloud-sdk-go/v5 => /Users/florin/work/code/cloud-sdk/sdk/core/go

go 1.13
