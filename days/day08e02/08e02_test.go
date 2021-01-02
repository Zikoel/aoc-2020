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
