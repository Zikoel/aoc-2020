package day03e02

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

	var routes = []slope{
		slope{
			h: 1,
			v: 1,
		},
		slope{
			h: 3,
			v: 1,
		},
		slope{
			h: 5,
			v: 1,
		},
		slope{
			h: 7,
			v: 1,
		},
		slope{
			h: 1,
			v: 2,
		},
	}

	h := 0
	trees := 0
	result := 1

	for _, route := range routes {
		landingLines := filterLines(lines, route.v)

		for _, pattern := range landingLines {
			if checkForThree(pattern, h) {
				trees++
			}
			h += route.h
			// fmt.Printf("%s (%d) -> %d\n", pattern, v, trees)
		}

		fmt.Printf("Route with slope %v encounter %d trees\n", route, trees)
		h = 0
		result *= trees
		trees = 0
	}

	fmt.Printf("Result %d\n", result)
}
