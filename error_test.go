package profitbricks

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newApiError(code int) error {
	return ApiError{
		response: errorResponse{
			HTTPStatus: code,
		},
	}
}

func TestIsStatusOk(t *testing.T) {
	assert.True(t, IsStatusOK(newApiError(http.StatusOK)))
	assert.False(t, IsStatusOK(newApiError(http.StatusTeapot)))
}

func TestIsStatusAccepted(t *testing.T) {
	assert.True(t, IsStatusAccepted(newApiError(http.StatusAccepted)))
	assert.False(t, IsStatusAccepted(newApiError(http.StatusTeapot)))
}

func TestIsStatusNotModified(t *testing.T) {
	assert.True(t, IsStatusNotModified(newApiError(http.StatusNotModified)))
	assert.False(t, IsStatusNotModified(newApiError(http.StatusTeapot)))
}

func TestIsStatusBadRequest(t *testing.T) {
	assert.True(t, IsStatusBadRequest(newApiError(http.StatusBadRequest)))
	assert.False(t, IsStatusBadRequest(newApiError(http.StatusTeapot)))
}

func TestIsStatusUnauthorized(t *testing.T) {
	assert.True(t, IsStatusUnauthorized(newApiError(http.StatusUnauthorized)))
	assert.False(t, IsStatusUnauthorized(newApiError(http.StatusTeapot)))
}

func TestIsStatusForbidden(t *testing.T) {
	assert.True(t, IsStatusForbidden(newApiError(http.StatusForbidden)))
	assert.False(t, IsStatusForbidden(newApiError(http.StatusTeapot)))
}

func TestIsStatusNotFoundError(t *testing.T) {
	assert.True(t, IsStatusNotFound(newApiError(http.StatusNotFound)))
	assert.False(t, IsStatusNotFound(newApiError(http.StatusTeapot)))
}

func TestIsStatusMethodNotAllowed(t *testing.T) {
	assert.True(t, IsStatusMethodNotAllowed(newApiError(http.StatusMethodNotAllowed)))
	assert.False(t, IsStatusMethodNotAllowed(newApiError(http.StatusTeapot)))
}

func TestIsStatusUnsupportedMediaType(t *testing.T) {
	assert.True(t, IsStatusUnsupportedMediaType(newApiError(http.StatusUnsupportedMediaType)))
	assert.False(t, IsStatusUnsupportedMediaType(newApiError(http.StatusTeapot)))
}

func TestIsStatusUnprocessableEntity(t *testing.T) {
	assert.True(t, IsStatusUnprocessableEntity(newApiError(http.StatusUnprocessableEntity)))
	assert.False(t, IsStatusUnprocessableEntity(newApiError(http.StatusTeapot)))
}

func TestIsStatusTooManyRequests(t *testing.T) {
	assert.True(t, IsStatusTooManyRequests(newApiError(http.StatusTooManyRequests)))
	assert.False(t, IsStatusTooManyRequests(newApiError(http.StatusTeapot)))
}

func TestIsRequestFailed(t *testing.T) {
	assert.True(t, IsRequestFailed(ClientError{errType: RequestFailed, msg: "fail"}))
	assert.False(t, IsRequestFailed(ClientError{errType: ClientErrorType(-1), msg: "fail"}))
}
