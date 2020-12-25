package day08e02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlterInstruction(t *testing.T) {
	assert := assert.New(t)

	instruction := "jmp +4"
	altered := alterInstruction(instruction)
	assert.Equal("nop +4", altered)

	instruction = "nop +4"
	altered = alterInstruction(instruction)
	assert.Equal("jmp +4", altered)

	instruction = "acc +4"
	altered = alterInstruction(instruction)
	assert.Equal("acc +4", altered)
}

// func TestSwapInstruction(t *testing.T) {
// 	assert := assert.New(t)

// 	original := []string{
// 		"jmp +4",
// 		"nop +3",
// 	}

// 	c := make([]string, 0)

// 	copy(c, original)

// 	c[1] = "jmp +7"

// 	// instructions[0] = alterInstruction(instructions[0])

// 	assert.ElementsMatch(original, c)
// }
