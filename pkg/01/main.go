package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %v\n", part1)

	part2, err := Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %v\n", part2)
}

var part1Regexp = regexp.MustCompile(`\d`)

func Part1(lines []string) (int, error) {
	total := 0
	for _, line := range lines {
		matches := part1Regexp.FindAllString(line, -1)
		if matches == nil {
			return 0, fmt.Errorf("no digits found in line: %s", line)
		}
		i, err := strconv.Atoi(matches[0] + matches[len(matches)-1])
		if err != nil {
			return 0, err
		}
		total += i
	}
	return total, nil
}

var part2ForwardRegexp = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
var part2BackwardRegexp = regexp.MustCompile(`\d|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno`)

func Part2(lines []string) (int, error) {
	total := 0
	values := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range lines {
		value := 0

		// can't use one regexp, as Go's regex engine doesn't do overlapping matches
		// e.g. "sevenine" with a single Regexp would return 77, instead of 79, due to the shared "n"
		// hacky solution is to run a "forward" Regex, and a "backward" Regex on a reversed string
		match := part2ForwardRegexp.FindString(line)
		if match == "" {
			return 0, fmt.Errorf("no digits found in line: %s", line)
		}
		if len(match) == 1 {
			v, err := strconv.Atoi(match)
			if err != nil {
				return 0, fmt.Errorf("expected digit, got %s", match)
			}
			value += v * 10
		} else {
			v, ok := values[match]
			if !ok {
				return 0, fmt.Errorf("expected %s to map to a digit", match)
			}
			value += v * 10
		}

		match = part2BackwardRegexp.FindString(lib.Reverse(line))
		if match == "" {
			return 0, fmt.Errorf("no digits found in line: %s", line)
		}
		if len(match) == 1 {
			v, err := strconv.Atoi(match)
			if err != nil {
				return 0, fmt.Errorf("expected digit, got %s", match)
			}
			value += v
		} else {
			v, ok := values[lib.Reverse(match)]
			if !ok {
				return 0, fmt.Errorf("expected %s to map to a digit", match)
			}
			value += v
		}

		total += value
	}
	return total, nil
}
