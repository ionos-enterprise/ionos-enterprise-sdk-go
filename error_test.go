package profitbricks

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func Test_ErrorWrapping(t *testing.T) {

	inner := errors.New("test")
	middle := BadRequest.Wrap(inner, "some problems")
	outer := errors.Wrap(middle, "another stack")
	assert.Error(t, outer)
	assert.Equal(t, BadRequest, GetType(outer))
	assert.Equal(t, middle, errors.Cause(outer))
	assert.Equal(t, inner, errors.Cause(GetCause(middle)))
	assert.Equal(t, inner, errors.Cause(GetCause(outer)))

}

func TestIsStatusNotFound(t *testing.T) {
	inner := ApiErrorResponse{HTTPStatus: http.StatusNotFound}
	outer := ClientError.Wrap(inner, "Ups")
	assert.True(t, IsStatusNotFound(outer))
}

func TestIsRequestFailed(t *testing.T) {
	inner := RequestFailed.New("ouch")
	outer := ClientError.Wrap(inner, "test")
	assert.True(t, IsRequestFailed(outer))
	assert.True(t, IsRequestFailed(inner))
}