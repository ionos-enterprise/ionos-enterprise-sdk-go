package profitbricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lanPath(t *testing.T) {
	assert.Equal(t, "datacenters/1/lans/3", lanPath("1", "3"))
}
