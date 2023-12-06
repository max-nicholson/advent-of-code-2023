package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	location, err := Part1(strings.Split(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if location != 35 {
		t.Fatalf("expected 35, got %d", location)
	}
}

func TestPart2(t *testing.T) {
	total, err := Part2(strings.Split(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 46 {
		t.Fatalf("expected 46, got %d", total)
	}
}
