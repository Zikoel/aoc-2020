package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlterInstruction(t *testing.T) {
	assert := assert.New(t)

	val1 := selectCompaible(0, []int{1, 5, 6})
	assert.Equal(1, val1)

	val2 := selectCompaible(4, []int{1, 5, 6})
	assert.Equal(5, val2)
}

func TestMakeCharghersChain(t *testing.T) {
	assert := assert.New(t)

	chain := makeChargersChain(0, []int{4, 2})
	assert.ElementsMatch(chain, []int{2, 4})

	chain2 := makeChargersChain(3, []int{4, 5, 9, 7})
	assert.ElementsMatch(chain2, []int{4, 5, 7, 9})
}

func TestElementsDiff(t *testing.T) {
	assert := assert.New(t)

	diffs := listElementsDiff([]int{0, 3, 5, 6})
	assert.ElementsMatch(diffs, []int{3, 2, 1})

}
