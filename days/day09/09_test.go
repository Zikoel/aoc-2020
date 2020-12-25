package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlterInstruction(t *testing.T) {
	assert := assert.New(t)

	preamble := []int{1, 2, 3}
	admissibileValues := calculatePermutations(preamble)

	assert.Len(admissibileValues, 3)
	assert.ElementsMatch(admissibileValues, []int{3, 4, 5})
	assert.True(false)
}
