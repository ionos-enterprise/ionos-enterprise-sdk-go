package profitbricks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cache_DeepCopy(t *testing.T) {
	c := newCache()
	vol := &Volume{
		ID: "132",
		Metadata: &Metadata{
			CreatedBy: "ntr0",
		},
	}
	c.Add("vol", 2, vol)
	nVol := c.Get("vol", 2)

	assert.Equal(t, vol, nVol)

	// modify metadata
	vol.Metadata.CreatedBy = "0rtn"
	nVol = c.Get("vol", 2)
	assert.Equal(t, vol, nVol)
}
