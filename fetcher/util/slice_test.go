package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	ss := Slice(0, 9)
	assert.True(t, len(ss) == 10)
}
