package profitbricks

import (
	"fmt"
)

type errorResponse struct {
	HTTPStatus int `json:"httpStatus"`
	Messages   []struct {
		ErrorCode string `json:"errorCode"`
		Message   string `json:"message"`
	} `json:"messages"`
}

type ApiError struct {
	response errorResponse
}

func (e ApiError) Error() string {
	return e.response.String()
}

func (e errorResponse) Error() string {
	return e.String()
}

func (e errorResponse) String() string {
	toReturn := fmt.Sprintf("HTTP Status: %s \n%s", fmt.Sprint(e.HTTPStatus), "Error Messages:")
	for _, m := range e.Messages {
		toReturn = toReturn + fmt.Sprintf("Error Code: %s Message: %s\n", m.ErrorCode, m.Message)
	}
	return toReturn
}

func (e ApiError) HttpStatusCode() int {
	return e.response.HTTPStatus
}
