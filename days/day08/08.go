package day08

import (
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

// Run run
func Run(data []byte) {
	istructions := strings.Split(string(data), "\n")

	ctx, err := executeProgram(istructions)

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%v\n", ctx)

	if ctx.infiniteLoop == true {
		fmt.Printf("Program exited due a loop on ip: %d (accumulator: %d)\n", ctx.ip, ctx.accumulator)
	}

	if ctx.programExited == true {
		fmt.Printf("Program exited correctly with value: %d\n", ctx.accumulator)
	}
}
