package day08e02

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type context struct {
	ip            int
	accumulator   int
	visitedIps    map[int]bool
	infiniteLoop  bool
	programExited bool
}

func decodeIstruction(istruction string) (string, int, error) {
	parts := strings.Split(istruction, " ")

	value, err := strconv.ParseInt(parts[1], 10, 32)

	if err != nil {
		return "", 0, fmt.Errorf("Unable to decode istruction: %s", istruction)
	}

	return parts[0], int(value), nil
}

func compute(istructions []string, ctx context) (context, error) {
	result := ctx
	if ctx.ip >= len(istructions) {
		result.programExited = true
		return result, nil
	}

	istruction := istructions[ctx.ip]
	command, value, err := decodeIstruction(istruction)

	if err != nil {
		return ctx, err
	}

	if result.visitedIps[result.ip] == true {
		result.infiniteLoop = true
		return result, nil
	}

	result.visitedIps[result.ip] = true

	switch command {
	case "acc":
		result.accumulator += value
		result.ip++
	case "nop":
		result.ip++
	case "jmp":
		result.ip += value
	}

	return result, nil
}

func executeProgram(istructions []string) (context, error) {

	ctx := context{
		ip:            0,
		accumulator:   0,
		visitedIps:    make(map[int]bool, 0),
		infiniteLoop:  false,
		programExited: false,
	}

	for i := range istructions {
		ctx.visitedIps[i] = false
	}

	for ctx.infiniteLoop == false && ctx.programExited == false {
		newCtx, err := compute(istructions, ctx)
		if err != nil {
			return newCtx, err
		}
		ctx = newCtx
	}

	return ctx, nil
}

func alterInstruction(instruction string) string {

	if strings.HasPrefix(instruction, "jmp") {
		return strings.Replace(instruction, "jmp", "nop", 1)
	}

	if strings.HasPrefix(instruction, "nop") {
		return strings.Replace(instruction, "nop", "jmp", 1)
	}

	return instruction
}

func bruteFixCorrection(instructions []string) (int, error) {

	for i := 0; i < len(instructions); i++ {
		fixedProgram := make([]string, len(instructions))
		copy(fixedProgram, instructions)
		fixedProgram[i] = alterInstruction(fixedProgram[i])

		result, err := executeProgram(fixedProgram)
		if err != nil {
			return 0, err
		}

		if result.programExited == true {
			return result.accumulator, nil
		}
	}

	return 0, errors.New("Impossible to correct the program")
}

// Run run
func Run(data []byte) {
	istructions := strings.Split(string(data), "\n")

	result, err := bruteFixCorrection(istructions)

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("The correct program exit with value: %d\n", result)

}
