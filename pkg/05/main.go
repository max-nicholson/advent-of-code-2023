package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

type CategoryRange struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

type Category struct {
	Ranges []CategoryRange
}

func (c *Category) Map(number int) int {
	for _, r := range c.Ranges {
		if number < r.SourceStart {
			continue
		}

		diff := number - r.SourceStart
		if diff < r.Length {
			return r.DestinationStart + diff
		}
	}

	return number
}

func main() {
	lines, err := lib.ReadLines("pkg/05/input.txt")
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

func Part1(lines []string) (int, error) {
	seeds := strings.Split(lines[0][7:], " ")
	seedIds := make([]int, len(seeds))
	for i, seed := range seeds {
		v, err := strconv.Atoi(seed)
		if err != nil {
			return 0, err
		}
		seedIds[i] = v
	}

	categories := []Category{}
	category := Category{Ranges: []CategoryRange{}}
	for i := 3; i < len(lines); {
		line := lines[i]
		if line == "" {
			i += 2
			categories = append(categories, category)
			category = Category{Ranges: []CategoryRange{}}
			continue
		}

		values := strings.Split(line, " ")
		ints := make([]int, 3)
		for j, value := range values {
			ints[j], _ = strconv.Atoi(value)
		}

		category.Ranges = append(category.Ranges, CategoryRange{DestinationStart: ints[0], SourceStart: ints[1], Length: ints[2]})

		i += 1
	}

	min := math.MaxUint >> 1
	for _, seed := range seedIds {
		for _, category := range categories {
			seed = category.Map(seed)
		}
		if seed < min {
			min = seed
		}
	}

	return min, nil
}

type Range struct {
	Start int
	End   int
}

type MapRange struct {
	Destination Range
	Source      Range
}

type Map struct {
	Ranges []MapRange
}

func Part2(lines []string) (int, error) {
	seeds := strings.Split(lines[0][7:], " ")
	if len(seeds)%2 != 0 {
		return 0, fmt.Errorf("expected pairs of seeds")
	}

	ranges := []Range{}
	for i := 0; i < len(seeds); i += 2 {
		start, _ := strconv.Atoi(seeds[i])
		length, _ := strconv.Atoi(seeds[i+1])

		ranges = append(ranges, Range{Start: start, End: start + length - 1})
	}

	maps := []Map{}
	m := Map{}
	for i := 3; i < len(lines); {
		line := lines[i]
		if line == "" {
			i += 2
			maps = append(maps, m)
			m = Map{}
			continue
		}

		values := strings.Split(line, " ")
		ints := make([]int, 3)
		for j, value := range values {
			ints[j], _ = strconv.Atoi(value)
		}

		m.Ranges = append(m.Ranges, MapRange{
			Destination: Range{
				Start: ints[0],
				End:   ints[0] + ints[2] - 1,
			},
			Source: Range{
				Start: ints[1],
				End:   ints[1] + ints[2] - 1,
			},
		})

		i += 1
	}

	// 3 scenarios:
	// 1. No overlap -> range continues to the next layer (1:1)
	// 2. Complete ovelap -> map to Destination (1:1)
	// 3. Partial overlap -> map to Destination + Unchanged (1: 2)
	nextRanges := []Range{}
	for _, m := range maps {
		nextRanges = []Range{}
		for len(ranges) > 0 {
			r := ranges[len(ranges)-1]
			ranges = ranges[:len(ranges)-1]
			noMapping := true
			for _, mapRange := range m.Ranges {
				if r.Start > mapRange.Source.End || r.End < mapRange.Source.Start {
					// No overlap
					// fmt.Printf("no overlap for %v and %v\n", mapRange, r)
					continue
				}

				noMapping = false

				startOverlap := r.Start >= mapRange.Source.Start
				endOverlap := r.End <= mapRange.Source.End
				if startOverlap && endOverlap {
					// Complete overlap
					// fmt.Printf("complete overlap for %v and %v\n", mapRange, r)
					nextRanges = append(nextRanges, Range{
						Start: r.Start - mapRange.Source.Start + mapRange.Destination.Start,
						End:   r.End - mapRange.Source.Start + mapRange.Destination.Start,
					})
					break
				}

				// Partial overlap
				if startOverlap {
					// end overhangs
					// fmt.Printf("end overhang for %v and %v\n", mapRange, r)
					nextRanges = append(nextRanges, Range{
						Start: r.Start - mapRange.Source.Start + mapRange.Destination.Start,
						End:   mapRange.Destination.End,
					})
					ranges = append(ranges, Range{
						Start: mapRange.Source.End + 1,
						End:   r.End,
					})
				} else {
					// start overhangs
					// fmt.Printf("start overhang for %v and %v\n", mapRange, r)
					nextRanges = append(nextRanges, Range{
						Start: mapRange.Destination.Start,
						End:   r.End - mapRange.Source.Start + mapRange.Destination.Start,
					})
					ranges = append(ranges, Range{
						Start: r.Start,
						End:   mapRange.Source.Start - 1,
					})
				}
				break
			}

			// Any source numbers that aren't mapped correspond to the same destination number
			if noMapping {
				nextRanges = append(nextRanges, Range{Start: r.Start, End: r.End})
			}
		}
		ranges = nextRanges
	}

	min := math.MaxUint >> 1
	for _, r := range nextRanges {
		if r.Start < min {
			min = r.Start
		}
	}

	return min, nil
}
