package image_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 3, Max(1, 2, 3))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2, 3))
}
