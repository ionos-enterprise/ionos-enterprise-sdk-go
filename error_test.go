package profitbricks

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/assert"
)

func newApiError(code int) error {
	return ApiError{
		HTTPStatus: code,
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

type ErrorSuite struct {
	ClientBaseSuite
}

func Test_Client(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}

func (s *ErrorSuite) Test_ApiError() {
	body := []byte(`{"httpStatus" : 401, "messages" : [ {"errorCode" : "315", "message" : "Unauthorized" } ] }`)
	rsp := makeJsonResponse(http.StatusUnauthorized, body)
	httpmock.RegisterResponder(http.MethodGet, "=~/datacenters", httpmock.ResponderFromResponse(rsp))
	_, err := s.c.ListDatacenters()
	s.Error(err)
	s.True(IsStatusUnauthorized(err))
	s.False(IsStatusAccepted(err))
	s.Equal(1, httpmock.GetTotalCallCount())
}

func (s *ErrorSuite) Test_BadGatewayError() {
	body := []byte("<html><body>Service temporarily not available</body></html>")
	mRsp := &http.Response{
		Header:     http.Header{},
		StatusCode: http.StatusBadGateway,
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
		Status:     http.StatusText(http.StatusBadGateway),
	}
	mRsp.Header.Set("Content-Type", "text/html")
	httpmock.RegisterResponder(http.MethodGet, "=~/datacenters", httpmock.ResponderFromResponse(mRsp))
	_, err := s.c.ListDatacenters()
	s.Error(err)
	s.Equal(body, err.(ApiError).Body())
	s.Equal(1, httpmock.GetTotalCallCount())
}
