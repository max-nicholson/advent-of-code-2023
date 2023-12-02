package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/02/input.txt")
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

var totalCubesByColour = map[string]int{"red": 12, "green": 13, "blue": 14}

func IsPossibleGame(line string) (bool, error) {
	slices := strings.Split(line, ": ")
	if slices == nil || len(slices) != 2 {
		return false, fmt.Errorf("expected line %s to split by ':'", line)
	}
	sets := strings.Split(slices[1], "; ")
	if len(sets) == 0 {
		return false, fmt.Errorf("no game information found for line %s", slices[1])
	}
	for _, set := range sets {
		cubes := strings.Split(set, ", ")
		if len(cubes) == 0 {
			return false, fmt.Errorf("no sets found for line %s", slices[1])
		}
		for _, cube := range cubes {
			record := strings.Split(cube, " ")
			if len(record) != 2 {
				return false, fmt.Errorf("expected cube data to have format [1-9]+ blue|green|red")
			}
			count, err := strconv.Atoi(record[0])
			if err != nil {
				return false, fmt.Errorf("unable to parse number of cubes: %w", err)
			}
			colour := record[1]
			max, ok := totalCubesByColour[colour]
			if !ok {
				log.Printf("got unexpected colour %s", colour)
				continue
			}
			if count > max {
				return false, nil
			}
		}
	}
	return true, nil
}

func Part1(lines []string) (int, error) {
	total := 0
	for i, line := range lines {
		isPossible, err := IsPossibleGame(line)
		if err != nil {
			return 0, err
		}
		if isPossible {
			total += i + 1
		}
	}
	return total, nil
}

func PowerOfSet(line string) (int, error) {
	slices := strings.Split(line, ": ")
	if slices == nil || len(slices) != 2 {
		return 0, fmt.Errorf("expected line %s to split by ':'", line)
	}
	sets := strings.Split(slices[1], "; ")
	if len(sets) == 0 {
		return 0, fmt.Errorf("no game information found for line %s", slices[1])
	}
	minimumSet := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
	for _, set := range sets {
		cubes := strings.Split(set, ", ")
		if len(cubes) == 0 {
			return 0, fmt.Errorf("no sets found for line %s", slices[1])
		}
		for _, cube := range cubes {
			record := strings.Split(cube, " ")
			if len(record) != 2 {
				return 0, fmt.Errorf("expected cube data to have format [1-9]+ blue|green|red")
			}
			count, err := strconv.Atoi(record[0])
			if err != nil {
				return 0, fmt.Errorf("unable to parse number of cubes: %w", err)
			}
			colour := record[1]
			min, ok := minimumSet[colour]
			if !ok {
				log.Printf("got unexpected colour %s", colour)
				continue
			}
			if count > min {
				minimumSet[colour] = count
			}
		}
	}
	power := 1
	for _, val := range minimumSet {
		power *= val
	}
	return power, nil
}

func Part2(lines []string) (int, error) {
	total := 0
	for _, line := range lines {
		power, err := PowerOfSet(line)
		if err != nil {
			return 0, err
		}
		total += power
	}
	return total, nil
}
