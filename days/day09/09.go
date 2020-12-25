package day09

import (
	"aoc-2020/pkg/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculatePermutations(elements []int) []int {
	result := []int{}

	for i, iv := range elements {
		for _, jv := range elements[i+1:] {
			val := iv + jv

			if utils.ListContains(result, val) {
				continue
			}

			result = append(result, val)
		}
	}

	return result
}

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

	preambleLen := 25

	for idx, val := range numbers {

		if idx < preambleLen {
			continue
		}

		preceedingNumbers := numbers[idx-preambleLen : idx]
		admissibileValues := calculatePermutations(preceedingNumbers)

		if !utils.ListContains(admissibileValues, val) {
			fmt.Printf("First not admissible: %d\n", val)
			os.Exit(0)
		}

		// fmt.Printf("For val %d we have preamble of %v, admisible value are: %v\n", val, preceedingNumbers, admissibileValues)
	}

	fmt.Printf("All values are admissibile\n")
}
