package profitbricks

import (
	"github.com/ionos-cloud/sdk-go/v5"
	"reflect"
)

func fillInResponse(obj interface{}, response *ionoscloud.APIResponse) {
	if response == nil || response.Response == nil {
		return
	}

	objValue := reflect.ValueOf(obj)
	if objValue.Type().Kind() != reflect.Ptr {
		panic("fillInResponse() requires a pointer")
	}

	headersVal := reflect.ValueOf(obj).Elem().FieldByName("Headers")
	if headersVal.IsValid() {
		h := response.Header
		headersVal.Set(reflect.ValueOf(&h))
	}

	responseVal := reflect.ValueOf(obj).Elem().FieldByName("Response")
	if responseVal.IsValid() {
		body := string(response.Payload)
		responseVal.Set(reflect.ValueOf(body))
	}

	statusVal := reflect.ValueOf(obj).Elem().FieldByName("StatusCode")
	if statusVal.IsValid() && statusVal.Type().Kind() == reflect.Int {
		statusVal.Set(reflect.ValueOf(response.StatusCode))
	}
}
