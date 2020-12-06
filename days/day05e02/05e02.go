package day05e02

import (
	"aoc-2020/pkg/utils"
	"fmt"
	"strings"
)

func idFromCode(code string) int {

	rowSpace := make([]int, 128)
	for i := 0; i < len(rowSpace); i++ {
		rowSpace[i] = i
	}

	sitSpace := make([]int, 8)
	for i := 0; i < len(sitSpace); i++ {
		sitSpace[i] = i
	}

	row := utils.BinarySplit(code[:7], rowSpace)
	// fmt.Printf("%d\n", row)
	sit := utils.BinarySplit(code[7:], sitSpace)
	// fmt.Printf("%d\n", sit)
	return row*8 + sit
}

// Run run
func Run(data []byte) {
	// codes := strings.Split(string(data), "\n")

	idsFreeMap := make(map[int]bool)
	for i := 0; i < 1024; i++ {
		idsFreeMap[i] = true
	}

	codes := strings.Split(string(data), "\n")
	for _, code := range codes {
		id := idFromCode(code)
		idsFreeMap[id] = false
	}

	for id, seatIsFree := range idsFreeMap {
		if id == 0 || id == len(idsFreeMap)-1 {
			continue
		}

		// fmt.Printf("%d\t%t\n", id, seatIsFree)
		if seatIsFree && idsFreeMap[id-1] == false && idsFreeMap[id+1] == false {
			fmt.Printf("My seat is: %d\n", id)
		}

	}
}
