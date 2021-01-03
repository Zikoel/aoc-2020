package day10e02

import (
	"aoc-2020/pkg/utils"
	"fmt"
	"sort"
)

func intRange(start, stop int) []int {
	r := []int{}
	for i := start; i < stop; i++ {
		r = append(r, i)
	}

	return r
}

// In this day I have tryied do apply a more fast reasoning technic so, the code
// is almost unfriendly

// Run run
func Run(data []byte) {
	c := utils.ListOfInts(data)

	c = append(c, 0)

	pc := utils.ListMax(c) + 3
	c = append(c, pc)

	sort.Ints(c)

	mem := map[int]int64{}

	var dp func(int) int64

	dp = func(i int) int64 {
		fmt.Printf("On %d\n", i)
		if i == len(c)-1 {
			return 1
		}

		if m, ok := mem[i]; ok == true {
			return m
		}

		var ans int64 = 0

		for _, j := range intRange(i+1, len(c)) {
			if c[j]-c[i] <= 3 {
				ans += dp(j)
			}
		}

		mem[i] = ans

		return ans
	}

	fmt.Printf("Solution: %d\n", dp(0))
}
