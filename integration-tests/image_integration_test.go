package integration_tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListImages(t *testing.T) {
	fmt.Println("Image tests")
	c := setupTestEnv()
	resp, err := c.ListImages()

	if err != nil {
		t.Error(err)
	}

	image = &resp.Items[0]
	assert.True(t, len(resp.Items) > 0)
}

func TestGetImage(t *testing.T) {
	c := setupTestEnv()
	resp, err := c.GetImage(image.ID)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.ID, image.ID)
	assert.Equal(t, resp.PBType, "image")
}

func TestGetImageFailure(t *testing.T) {
	c := setupTestEnv()
	_, err := c.GetImage("00000000-0000-0000-0000-000000000000")

	if err == nil {
		t.Fail()
	}
}
