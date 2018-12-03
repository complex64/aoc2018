package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadClaim(t *testing.T) {
	s := "#123 @ 3,2: 5x4"
	c, err := readClaim(s)
	assert.NoError(t, err)
	assert.Equal(t, 123, c.id)
	assert.Equal(t, 3, c.x)
	assert.Equal(t, 2, c.y)
	assert.Equal(t, 5, c.width)
	assert.Equal(t, 4, c.height)
}
