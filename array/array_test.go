package array

import (
	"testing"

	"github.com/testify/assert"
)

func TestIsInArray(t *testing.T) {
	assert.Equal(t, false, IsInArray(3, []int{2, 4}))
	assert.Equal(t, true, IsInArray(2, []int{2, 4}))
	assert.Equal(t, false, IsInArray("2", []int{2, 4}))
	assert.Equal(t, true, IsInArray(2, []int{2, 4}))
	assert.Equal(t, true, IsInArray(2, []int{2, 4}))
}
