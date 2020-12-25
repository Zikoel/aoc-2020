package day04e02

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	fields []string
}

func parseRawPassportData(raw string) passport {
	lines := strings.Split(raw, "\n")
	elements := make([]string, 0)
	for _, l := range lines {
		elements = append(elements, strings.Split(l, " ")...)
	}

	return passport{
		fields: elements,
	}
}

func rangeCheck(value string, expectedDigit, min, max int) bool {
	if len(value) != expectedDigit {
		return false
	}

	numericValue, err := strconv.ParseInt(value, 10, 32)

	if err != nil {
		log.Fatalf("%v", err)
	}

	return int(numericValue) >= min && int(numericValue) <= max
}

func isHgtValid(value string) bool {
	if strings.HasSuffix(value, "cm") {
		return rangeCheck(strings.TrimSuffix(value, "cm"), 3, 150, 193)
	}

	if strings.HasSuffix(value, "in") {
		return rangeCheck(strings.TrimSuffix(value, "in"), 2, 59, 76)
	}

	return false
}

func isHclValid(value string) bool {
	re := regexp.MustCompile(`#[a-f0-9]{6}`)
	return re.MatchString(value)
}

func isEclValid(value string) bool {
	switch value {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

func isPidValid(value string) bool {
	re := regexp.MustCompile(`^\d{9}$`)
	return re.MatchString(value)
}

func isFieldValid(key, value string) bool {
	switch key {
	case "byr":
		return rangeCheck(value, 4, 1920, 2002)
	case "iyr":
		return rangeCheck(value, 4, 2010, 2020)
	case "eyr":
		return rangeCheck(value, 4, 2020, 2030)
	case "hgt":
		return isHgtValid(value)
	case "hcl":
		return isHclValid(value)
	case "ecl":
		return isEclValid(value)
	case "pid":
		return isPidValid(value)
	case "cid":
		return true
	}
	return false
}

func isPassportValid(passport passport) bool {
	// fmt.Printf("%v", passport.fields)
	if len(passport.fields) < 7 {
		return false
	}

	passportValid := true
	isCidPresent := false
	for _, v := range passport.fields {
		keyValue := strings.Split(v, ":")
		if !isFieldValid(keyValue[0], keyValue[1]) {
			// fmt.Printf("%s\n", keyValue)
			passportValid = false
		}
		if keyValue[0] == "cid" {
			isCidPresent = true
		}
	}

	if isCidPresent {
		return passportValid && len(passport.fields) == 8
	}
	return passportValid && len(passport.fields) == 7

}

// Run run
func Run(data []byte) {
	sparsePassports := strings.Split(string(data), "\n\n")

	passportValidCount := 0
	for _, passportRaw := range sparsePassports {
		passport := parseRawPassportData(passportRaw)
		if isPassportValid(passport) {
			fmt.Printf("YYY -> %v\n", passport)
			passportValidCount++
		} else {
			fmt.Printf("XXX -> %v\n", passport)
		}
	}

	fmt.Printf("We have %d valid passport", passportValidCount)
}
