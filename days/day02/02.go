package day02

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type policy struct {
	password      string
	character     string
	minOccurence  int64
	maxOccurrence int64
	isValid       bool
}

func stringToPolicy(str string) (policy, error) {
	re := regexp.MustCompile(`(\d+)-(\d+) (.): (.+)`)
	matches := re.FindStringSubmatch(str)

	if len(matches) <= 1 {
		return policy{}, fmt.Errorf("unrecognized pattern of %s", str)
	}

	minOccurence, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return policy{}, fmt.Errorf("unrecognized pattern of %s", str)
	}

	maxOccurence, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return policy{}, fmt.Errorf("unrecognized pattern of %s", str)
	}

	character := matches[3]
	password := matches[4]

	var count int64 = 0
	for _, c := range strings.Split(password, "") {
		if c == character {
			count++
		}
	}

	isPasswordValid := true
	if count < minOccurence || count > maxOccurence {
		isPasswordValid = false
	}

	return policy{
		minOccurence:  minOccurence,
		maxOccurrence: maxOccurence,
		character:     character,
		password:      password,
		isValid:       isPasswordValid,
	}, nil
}

// Run run
func Run(data []byte) {
	valuesStr := strings.Split(string(data), "\n")

	validPasswords := 0
	for _, policyStr := range valuesStr {
		policy, err := stringToPolicy(policyStr)
		if err != nil {
			log.Fatal(err)
		}

		if policy.isValid {
			validPasswords++
		}

	}

	fmt.Printf("We have %d valid passwords", validPasswords)
}
