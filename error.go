package profitbricks

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	Error = ErrorType(iota)
	ClientError
	BadRequest
	NotFound     = ErrorType(http.StatusNotFound)
	UnAuthorized = ErrorType(http.StatusUnauthorized)
	RequestFailed
)

type ErrorType uint

type sdkError struct {
	errorType     ErrorType
	originalError error
}

func (e sdkError) Error() string {
	return e.originalError.Error()
}

func (etype ErrorType) New(msg string) error {
	return sdkError{
		errorType:     etype,
		originalError: errors.New(msg),
	}
}
func (etype ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, args...)

	return sdkError{errorType: etype, originalError: err}
}

// Wrap creates a new wrapped error
func (etype ErrorType) Wrap(err error, msg string) error {
	return etype.Wrapf(err, msg)
}

// Wrap creates a new wrapped error with formatted message
func (etype ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args...)

	return sdkError{
		errorType: etype, originalError: newErr}
}

func GetType(e error) ErrorType {
	underlying := errors.Cause(e)
	if et, ok := underlying.(sdkError); ok {
		return et.errorType
	}
	return Error
}

func GetCause(e error) error {
	underlying := errors.Cause(e)
	if et, ok := underlying.(sdkError); ok {
		return et.originalError
	}
	return underlying

}

type ApiErrorResponse struct {
	HTTPStatus int `json:"httpStatus"`
	Messages   []struct {
		ErrorCode string `json:"errorCode"`
		Message   string `json:"message"`
	} `json:"messages"`
}

func (e ApiErrorResponse) Error() string {
	return e.String()
}

func (e ApiErrorResponse) String() string {
	toReturn := fmt.Sprintf("HTTP Status: %s \n%s", fmt.Sprint(e.HTTPStatus), "Error Messages:")
	for _, m := range e.Messages {
		toReturn = toReturn + fmt.Sprintf("Error Code: %s Message: %s\n", m.ErrorCode, m.Message)
	}
	return toReturn
}

func (e ApiErrorResponse) HttpStatusCode() int {
	return e.HTTPStatus
}

func IsClientErrorType(err error, errType ErrorType) bool {
	if err == nil {
		return false
	}
	return GetType(err) == errType || IsClientErrorType(GetCause(err), errType)
}

func IsHttpStatus(err error, status int) bool {
	if err == nil {
		return false
	}
	if err, ok := err.(ApiErrorResponse); ok {
		return err.HttpStatusCode() == status
	}
	if int(GetType(err)) == status {
		return true
	}
	return IsHttpStatus(GetCause(err), status)
}

// IsStatusOK - (200)
func IsStatusOK(err error) bool {
	return IsHttpStatus(err, http.StatusOK)
}

// IsStatusAccepted - (202) Used for asynchronous requests using PUT, DELETE, POST and PATCH methods.
// The response will also include a Location header pointing to a resource. This can be used for polling.
func IsStatusAccepted(err error) bool {
	return IsHttpStatus(err, http.StatusAccepted)
}

// IsStatusNotModified - (304) Response for GETs on resources that have not been changed. (based on ETag values).
func IsStatusNotModified(err error) bool {
	return IsHttpStatus(err, http.StatusNotModified)
}

// IsStatusBadRequest - (400) Response to malformed requests or general client errors.
func IsStatusBadRequest(err error) bool {
	return IsHttpStatus(err, http.StatusBadRequest)
}

// IsStatusUnauthorized - (401) Response to an unauthenticated connection.
// You will need to use your API username and password to be authenticated.
func IsStatusUnauthorized(err error) bool {
	return IsHttpStatus(err, http.StatusUnauthorized)
}

// IsStatusForbidden - (403) Forbidden
func IsStatusForbidden(err error) bool {
	return IsHttpStatus(err, http.StatusForbidden)
}

// IsStatusNotFound - (404) if resource does not exist
func IsStatusNotFound(err error) bool {
	return IsHttpStatus(err, http.StatusNotFound)
}

// IsStatusMethodNotAllowed - (405) Use for any POST, PUT, PATCH, or DELETE performed
// on read-only resources. This is also the response to PATCH requests
// on resources that do not support partial updates.
func IsStatusMethodNotAllowed(err error) bool {
	return IsHttpStatus(err, http.StatusMethodNotAllowed)
}

// IsStatusUnsupportedMediaType - (415) The content-type is incorrect for the payload.
func IsStatusUnsupportedMediaType(err error) bool {
	return IsHttpStatus(err, http.StatusUnsupportedMediaType)
}

// IsStatusUnprocessableEntity - (422) Validation errors.
func IsStatusUnprocessableEntity(err error) bool {
	return IsHttpStatus(err, http.StatusUnprocessableEntity)
}

// IsStatusTooManyRequests - (429) The number of requests exceeds the rate limit.
func IsStatusTooManyRequests(err error) bool {
	return IsHttpStatus(err, http.StatusTooManyRequests)
}

// IsRequestFailed - returns true if the error reason was that the request status was failed
func IsRequestFailed(err error) bool {
	return IsClientErrorType(err, RequestFailed)
}
