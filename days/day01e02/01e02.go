package day01e02

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
	for i, v := range values {
		for j, c := range values[i+1:] {
			for _, d := range values[j+1:] {
				if v+c+d == 2020 {
					fmt.Printf("The triplette is %d, %d, %d\n", v, c, d)
					result = v * c * d
				}
			}
		}
	}
	fmt.Printf("Result: %d\n", result)
}
