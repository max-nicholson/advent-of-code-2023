package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/06/input.txt")
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

type Race struct {
	Distance int
	Time     int
}

func (r *Race) Length() int {
	time := r.Time
	distance := r.Distance
	start := 0
	for press := 1; press < time-1; press++ {
		speed := press
		travelTime := time - press

		d := travelTime * speed
		if d > distance {
			start = press
			break
		}
	}
	end := 0
	for press := time - 1; press > 0; press-- {
		speed := press
		travelTime := time - press

		d := travelTime * speed
		if d > distance {
			end = press
			break
		}
	}

	return end - start + 1
}

func Part1(lines []string) (int, error) {
	var ParseLine = func(line string) ([]int, error) {
		values := []int{}

		for _, s := range strings.Split(line[9:], " ") {
			s = strings.Trim(s, " ")
			if s == "" {
				continue
			}
			v, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("unable to convert %s to int", s)
			}
			values = append(values, v)
		}
		return values, nil
	}
	times, err := ParseLine(lines[0])
	if err != nil {
		return 0, fmt.Errorf("unable to parse times: %w", err)
	}
	distances, err := ParseLine(lines[1])
	if err != nil {
		return 0, fmt.Errorf("unable to parse distances: %w", err)
	}
	if len(times) != len(distances) {
		return 0, fmt.Errorf("mismatch in length between times (%d) and distances (%d)", len(times), len(distances))
	}

	races := len(times)
	ways := 0
	for r := 0; r < races; r++ {
		race := Race{Time: times[r], Distance: distances[r]}
		length := race.Length()
		if r == 0 {
			ways = length
		} else {
			ways *= length
		}
	}

	return ways, nil
}

func Part2(lines []string) (int, error) {
	var ParseLine = func(line string) (int, error) {
		data := strings.Split(line, ": ")[1]
		s := strings.ReplaceAll(data, " ", "")
		v, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("unable to convert %s to int: %w", s, err)
		}
		return v, nil
	}

	time, err := ParseLine(lines[0])
	if err != nil {
		return 0, fmt.Errorf("unable to parse times: %w", err)
	}
	distance, err := ParseLine(lines[1])
	if err != nil {
		return 0, fmt.Errorf("unable to parse distances: %w", err)
	}

	race := Race{Time: time, Distance: distance}

	length := race.Length()

	return length, nil
}
