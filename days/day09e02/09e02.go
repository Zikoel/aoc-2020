package day09e02

import (
	"aoc-2020/pkg/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Run run
func Run(data []byte) {
	values := strings.Split(string(data), "\n")

	numbers := []int{}

	for _, val := range values {
		v, err := strconv.ParseInt(val, 10, 64)

		if err != nil {
			log.Fatalf("%v", err)
		}

		numbers = append(numbers, int(v))
	}

	invalidNumber := 127 // Put here your 09 day result

	sum := 0
	contiguousLen := 2
	subList := make([]int, 0)
	for sum != invalidNumber {
		fmt.Printf("Try to find a soulution with len %d\n", contiguousLen)
		for idx := 0; idx < len(numbers) && sum != invalidNumber; idx++ {

			if idx < contiguousLen {
				continue
			}

			endIdx := idx
			startIdx := endIdx - contiguousLen
			subList = numbers[startIdx:endIdx]
			sum = utils.ListElementsSum(subList)

			fmt.Printf("Try with: %v -> %d\n", subList, sum)

		}
		contiguousLen++
	}

	result := utils.ListMax(subList) + utils.ListMin(subList)

	fmt.Printf("Solution found with len %d (%d) from %v\n", contiguousLen, result, subList)
}
