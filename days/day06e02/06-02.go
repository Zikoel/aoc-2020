package day06e02

import (
	"aoc-2020/pkg/utils"
	"fmt"
	"strings"
)

// remove from str all the rune not container on matcher
func runeMatching(str, matcher string) string {
	result := ""
	for _, r := range matcher {
		if strings.ContainsRune(str, r) {
			result += string(r)
		}
	}
	return result
}

// Run run
func Run(data []byte) {

	dataByGroups := strings.Split(string(data), "\n\n")

	sum := 0
	for _, group := range dataByGroups {
		inLineGroup := strings.ReplaceAll(group, "\n", "")
		cleaned := utils.RemoveDuplicate(inLineGroup)

		persons := strings.Split(group, "\n")
		for _, person := range persons {
			cleaned = runeMatching(cleaned, person)
		}

		sum += len(cleaned)
	}
	fmt.Printf("%d\n", sum)
}
