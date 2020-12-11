package day06

import (
	"aoc-2020/pkg/utils"
	"fmt"
	"strings"
)

// Run run
func Run(data []byte) {

	dataByGroups := strings.Split(string(data), "\n\n")

	sum := 0
	for _, group := range dataByGroups {
		inLineGroup := strings.ReplaceAll(group, "\n", "")

		cleaned := utils.RemoveDuplicate(inLineGroup)

		sum += len(cleaned)
	}
	fmt.Printf("%d\n", sum)
}
