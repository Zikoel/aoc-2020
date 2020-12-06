package day05

import (
	"aoc-2020/pkg/utils"
	"fmt"
	"sort"
	"strings"
)

func calculateID(code string) int {

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
	codes := strings.Split(string(data), "\n")

	ids := make([]int, 0)
	for _, code := range codes {
		id := calculateID(code)
		ids = append(ids, id)
		fmt.Printf("%d\n", id)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	fmt.Printf("Result: %d", ids[0])

}
