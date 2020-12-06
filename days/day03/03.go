package day03

import (
	"fmt"
	"strings"
)

func checkForThree(geologyLine string, position int) bool {
	lineElements := strings.Split(geologyLine, "")
	geologyPatternLen := len(lineElements)
	positionOnPattern := position % geologyPatternLen
	return lineElements[positionOnPattern] == "#"
}

type slope struct {
	h int
	v int
}

func filterLines(lines []string, step int) []string {
	result := make([]string, 0)
	for i, e := range lines {
		if i%step == 0 {
			result = append(result, e)
		}
	}
	return result
}

// Run run
func Run(data []byte) {
	lines := strings.Split(string(data), "\n")

	fmt.Printf("%v\n", filterLines(lines, 1))

	h := 0
	trees := 0
	const slopeH = 3
	for v, pattern := range lines {
		if checkForThree(pattern, h) {
			trees++
		}
		h += slopeH
		fmt.Printf("%s (%d) -> %d\n", pattern, v, trees)
	}
	fmt.Printf("Encountered trees: %d\n", trees)
}
