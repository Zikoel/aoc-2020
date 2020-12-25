package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySplit(t *testing.T) {
	assert := assert.New(t)

	val := BinarySplit("F", []int{4, 5})
	assert.Equal(4, val)

	val = BinarySplit("L", []int{4, 5})
	assert.Equal(4, val)

	val = BinarySplit("B", []int{4, 5})
	assert.Equal(5, val)

	val = BinarySplit("R", []int{4, 5})
	assert.Equal(5, val)

	val = BinarySplit("FF", []int{44, 45, 46, 47})
	assert.Equal(44, val)

	val = BinarySplit("BFF", []int{40, 41, 42, 43, 44, 45, 46, 47})
	assert.Equal(44, val)
}

func TestListContains(t *testing.T) {
	assert := assert.New(t)

	result := ListContains([]int{4, 5}, 4)
	assert.True(result)

	result = ListContains([]int{4, 5}, 10)
	assert.False(result)
}
