/* This file contains helpers to check whether given error
is specific http status code or not.

*/
package profitbricks

import "net/http"

func IsHttpStatus(err error, status int) bool {
	if err, ok := err.(ApiError); ok {
		return err.HttpStatusCode() == status
	}
	return false
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
