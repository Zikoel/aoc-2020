package day01

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Run run
func Run(data []byte) {
	valuesStr := strings.Split(string(data), "\n")

	values := []int64{}
	for _, v := range valuesStr {
		val, err := strconv.ParseInt(v, 10, 64)

		if err != nil {
			log.Fatalf("Impossibile to convert %s", v)
		}

		values = append(values, val)
	}

	var result int64 = 0
	for _, v := range values {
		for _, c := range values[1:] {
			if v+c == 2020 {
				fmt.Printf("The pair is %d, %d\n", v, c)
				result = v * c
			}
		}
	}
	fmt.Printf("Result: %d\n", result)
}
