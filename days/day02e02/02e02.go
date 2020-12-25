package day02e02

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
	firstPosition int64
	lastPosition  int64
	isValid       bool
}

func stringToPolicy(str string) (policy, error) {
	re := regexp.MustCompile(`(\d+)-(\d+) (.): (.+)`)
	matches := re.FindStringSubmatch(str)

	if len(matches) <= 1 {
		return policy{}, fmt.Errorf("unrecognized pattern of %s", str)
	}

	firstPosition, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return policy{}, fmt.Errorf("unrecognized pattern of %s", str)
	}

	lastPosition, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return policy{}, fmt.Errorf("unrecognized pattern of %s", str)
	}

	character := matches[3]
	password := matches[4]

	passwordCharacters := strings.Split(password, "")

	firstPositionMatch := passwordCharacters[firstPosition-1] == character
	secondPositionMatch := passwordCharacters[lastPosition-1] == character

	fmt.Printf("first position %d is %t ", firstPosition, firstPositionMatch)
	fmt.Printf("second position %d is %t ", lastPosition, secondPositionMatch)

	isPasswordValid := false

	if firstPositionMatch != secondPositionMatch {
		isPasswordValid = true
	}

	return policy{
		firstPosition: firstPosition,
		lastPosition:  lastPosition,
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
			fmt.Printf("YYY -> %s\n", policyStr)
			validPasswords++
		} else {
			fmt.Printf("XXX -> %s\n", policyStr)
		}
	}

	fmt.Printf("We have %d valid passwords", validPasswords)
}
