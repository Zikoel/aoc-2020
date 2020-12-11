package day07e02

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type ruleCapacity struct {
	bagColorCode string
	qt           int
}

type rule struct {
	bagColorCode string
	capacity     []ruleCapacity
}

type bagCapacity struct {
	qt  int
	bag *bag
}

type bag struct {
	color string
	inner []bagCapacity
}

func cleanBagName(rawName string) string {
	return strings.Trim(rawName, " ")
}

func bagNameFromRule(rule string) string {
	name := strings.Split(rule, "bags contain")
	return cleanBagName(name[0])
}

func ruleFromRuleStr(ruleStr string) rule {
	ruleParts := strings.Split(ruleStr, "bags contain")

	name := cleanBagName(ruleParts[0])

	innerBagsRaw := ruleParts[1]
	innerBagsStr := strings.Split(innerBagsRaw, ",")

	re := regexp.MustCompile(`(\d+) ([a-z]+ [a-z]+) bag[s]?`)

	innerBags := make([]ruleCapacity, 0)
	for _, bag := range innerBagsStr {
		matches := re.FindStringSubmatch(bag)

		if matches == nil {
			continue
		}

		qt, err := strconv.ParseInt(matches[1], 10, 32)

		if err != nil {
			log.Fatalf("%v", err)
		}

		innerBags = append(innerBags, ruleCapacity{
			qt:           int(qt),
			bagColorCode: matches[2],
		})
	}
	return rule{
		bagColorCode: name,
		capacity:     innerBags,
	}
}

func printBag(bag bag) string {
	result := fmt.Sprintf("Color: %s ->\n", bag.color)
	for _, in := range bag.inner {
		result += fmt.Sprintf("\t%d of %s\t\t", in.qt, printBag(*in.bag))
	}
	return result
}

func howMutchBagsBagCanContain(bag bag) int {
	result := 0

	for _, i := range bag.inner {
		result += i.qt * (1 + howMutchBagsBagCanContain(*i.bag))
	}
	return result
}

// Run run
func Run(data []byte) {

	rulesStr := strings.Split(string(data), "\n")

	rules := make([]rule, 0)
	for _, ruleStr := range rulesStr {
		rules = append(rules, ruleFromRuleStr(ruleStr))
	}

	// Make a list of bags
	bags := map[string]*bag{}
	for _, rule := range rules {
		bags[rule.bagColorCode] = &bag{
			color: rule.bagColorCode,
		}
	}

	// Now I can polulate the inner capacity
	for _, rule := range rules {
		inners := []bagCapacity{}
		for _, capacity := range rule.capacity {
			inners = append(inners, bagCapacity{
				bag: bags[capacity.bagColorCode],
				qt:  capacity.qt,
			})
		}
		bags[rule.bagColorCode].inner = inners
	}

	container := "shiny gold"
	fmt.Printf("%s\n", printBag(*bags[container]))

	n := howMutchBagsBagCanContain(*bags[container])

	fmt.Printf("%s bag can contain %d bags\n", container, n)
}
