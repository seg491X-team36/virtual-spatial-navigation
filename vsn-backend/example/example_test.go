package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, Add(1, 1), 2)
	assert.Equal(t, Add(1, 2), 3)
	assert.Equal(t, Add(1, -1), 0)
}
