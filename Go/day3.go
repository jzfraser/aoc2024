package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func Day3() {
	input := GetInput("day3input.txt")

	d3p1(input)
	d3p2(input)
}

func d3p1(input string) {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var total int
	for _, match := range matches {
		a, aErr := strconv.Atoi(match[1])
		b, bErr := strconv.Atoi(match[2])

		if aErr != nil || bErr != nil {
			fmt.Println("Day3: Error converting to integers:", aErr, bErr)
			return
		}

		total += a * b
	}

	fmt.Println("Day3 part 1 total:", total)
}

func d3p2(input string) {
	type Pattern string

	const (
		DoPattern   Pattern = "do"
		DontPattern Pattern = "dont"
		MulPattern  Pattern = "mul"
	)

	type Match struct {
		Pattern Pattern
		Match   []int
		Index   int
	}

	doPat := regexp.MustCompile(`do\(\)`)
	dontPat := regexp.MustCompile(`don't\(\)`)
	mulPat := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	doMatches := doPat.FindAllStringIndex(input, -1)
	dontMatches := dontPat.FindAllStringIndex(input, -1)
	mulMatches := mulPat.FindAllStringSubmatchIndex(input, -1)

	var allMatches []Match

	for _, match := range doMatches {
		allMatches = append(allMatches, Match{
			Pattern: DoPattern,
			Match:   match,
			Index:   match[0],
		})
	}
	for _, match := range dontMatches {
		allMatches = append(allMatches, Match{
			Pattern: DontPattern,
			Match:   match,
			Index:   match[0],
		})
	}
	for _, match := range mulMatches {
		allMatches = append(allMatches, Match{
			Pattern: MulPattern,
			Match:   match,
			Index:   match[0],
		})
	}

	sort.Slice(allMatches, func(i, j int) bool {
		return allMatches[i].Index < allMatches[j].Index
	})

	mulEnabled := true
	total := 0
	for _, match := range allMatches {
		switch match.Pattern {
		case DoPattern:
			mulEnabled = true
		case DontPattern:
			mulEnabled = false
		case MulPattern:
			if mulEnabled {
				total += getMulValue(input, match.Match)
			}
		}
	}

	fmt.Println("Day 3 Part 2 total:", total)
}

func getMulValue(input string, match []int) int {
	aStr := input[match[2]:match[3]]
	bStr := input[match[4]:match[5]]

	a, aErr := strconv.Atoi(aStr)
	b, bErr := strconv.Atoi(bStr)

	if aErr != nil || bErr != nil {
		err := "Day 3 Part 2: Error converting to integers"
		panic(err)
	}

	return a * b
}
