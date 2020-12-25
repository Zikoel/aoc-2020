package main

import (
	"aoc-2020/days/day01"
	"aoc-2020/days/day01e02"
	"aoc-2020/days/day02"
	"aoc-2020/days/day02e02"
	"aoc-2020/days/day03"
	"aoc-2020/days/day03e02"
	"aoc-2020/days/day04e02"
	"aoc-2020/days/day05"
	"aoc-2020/days/day05e02"
	"aoc-2020/days/day06"
	"aoc-2020/days/day06e02"
	"aoc-2020/days/day07"
	"aoc-2020/days/day07e02"
	"aoc-2020/days/day08"
	"aoc-2020/days/day08e02"
	"aoc-2020/days/day09"
	"aoc-2020/days/day09e02"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var day string = ""

	flag.StringVar(&day, "day", "", "day to run")
	flag.StringVar(&day, "d", "", "day to run (shorthand)")

	flag.Parse()

	if day == "" {
		flag.Usage()
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal("Impossible read input")
	}

	switch day {
	case "01":
		day01.Run(data)
	case "01e02":
		day01e02.Run(data)
	case "02":
		day02.Run(data)
	case "02e02":
		day02e02.Run(data)
	case "03":
		day03.Run(data)
	case "03e02":
		day03e02.Run(data)
	case "04":
		fmt.Println("To be restored")
	case "04e02":
		day04e02.Run(data)
	case "05":
		day05.Run(data)
	case "05e02":
		day05e02.Run(data)
	case "06":
		day06.Run(data)
	case "06e02":
		day06e02.Run(data)
	case "07":
		day07.Run(data)
	case "07e02":
		day07e02.Run(data)
	case "08":
		day08.Run(data)
	case "08e02":
		day08e02.Run(data)
	case "09":
		day09.Run(data)
	case "09e02":
		day09e02.Run(data)
	default:
		fmt.Println("Day not present")
	}
}
