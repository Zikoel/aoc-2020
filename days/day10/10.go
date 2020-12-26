package day10

import (
	"aoc-2020/pkg/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func selectCompaible(source int, chargers []int) int {
	compatibles := []int{}

	for _, val := range chargers {
		if val >= source && val <= source+3 {
			compatibles = append(compatibles, val)
		}
	}

	return utils.ListMin(compatibles)
}

func makeChargersChain(start int, chargers []int) []int {
	if len(chargers) == 0 {
		return []int{}
	}

	nextCharger := selectCompaible(start, chargers)
	remainsChargers := utils.ListRemoveVal(chargers, nextCharger)
	chain := []int{nextCharger}
	chain = append(chain, makeChargersChain(nextCharger, remainsChargers)...)
	return chain
}

func listElementsDiff(list []int) []int {
	result := []int{}
	for idx, val := range list {
		if idx == 0 {
			continue
		}

		diff := val - list[idx-1]
		result = append(result, diff)
	}

	return result
}

// Run run
func Run(data []byte) {
	values := strings.Split(string(data), "\n")

	chargers := []int{}

	for _, val := range values {
		v, err := strconv.ParseInt(val, 10, 64)

		if err != nil {
			log.Fatalf("%v", err)
		}

		chargers = append(chargers, int(v))
	}

	source := 0

	chain := makeChargersChain(source, chargers)

	pc := utils.ListMax(chargers) + 3

	finalList := append([]int{source}, chain...)
	finalList = append(finalList, pc)

	diffs := listElementsDiff(finalList)

	occurenceOf1 := utils.ListCount(diffs, 1)
	occurenceOf2 := utils.ListCount(diffs, 2)
	occurenceOf3 := utils.ListCount(diffs, 3)

	fmt.Printf("Occurence of 1: %d\n", occurenceOf1)
	fmt.Printf("Occurence of 2: %d\n", occurenceOf2)
	fmt.Printf("Occurence of 3: %d\n", occurenceOf3)
	fmt.Printf("Occurence of 1 multiplied fo occurence of 3: %d\n", occurenceOf1*occurenceOf3)
}
