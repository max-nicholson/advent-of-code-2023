package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

func ParseLine(line string) ([]int, error) {
	s := strings.Split(line, " ")
	values := make([]int, len(s))
	for i := range s {
		v, err := strconv.Atoi(s[i])
		if err != nil {
			return nil, fmt.Errorf("unable to parse line %s: %w", line, err)
		}
		values[i] = v
	}
	return values, nil
}

func allSameValue(values []int) bool {
	target := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] != target {
			return false
		}
	}
	return true
}

func buildSequences(values []int) [][]int {
	sequences := [][]int{values}
	for {
		values = sequences[len(sequences)-1]
		newValues := make([]int, len(values)-1)
		for i := 0; i < len(values)-1; i++ {
			newValues[i] = values[i+1] - values[i]
		}

		sequences = append(sequences, newValues)

		if allSameValue(newValues) {
			break
		}
	}

	return sequences
}

func Part1(lines []string) (int, error) {
	total := 0

	for _, line := range lines {
		values, err := ParseLine(line)
		if err != nil {
			return 0, err
		}
		sequences := buildSequences(values)

		for _, seq := range sequences {
			total += seq[len(seq)-1]
		}
	}
	return total, nil
}

func Part2(lines []string) (int, error) {
	total := 0

	for _, line := range lines {
		values, err := ParseLine(line)
		if err != nil {
			return 0, err
		}
		sequences := buildSequences(values)

		previousValue := sequences[0][0]
		for i := 1; i < len(sequences); i++ {
			seq := sequences[i]
			delta := seq[0]
			if i%2 == 0 {
				previousValue += delta
			} else {
				previousValue -= delta
			}
		}
		total += previousValue
	}
	return total, nil
}
