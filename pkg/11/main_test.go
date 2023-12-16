package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 374 {
		t.Fatalf("expected 374, got %d", result)
	}
}

func TestSumOfShortestPaths(t *testing.T) {
	lines := strings.Split(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`, "\n")
	var testCases = []struct {
		expansionFactor int
		want            int
	}{
		{expansionFactor: 10, want: 1030},
		{expansionFactor: 100, want: 8410},
	}
	for _, testCase := range testCases {
		result, err := SumOfShortestPaths(lines, testCase.expansionFactor)
		if err != nil {
			t.Error(err)
		}
		if result != testCase.want {
			t.Errorf("expected %d, got %d", testCase.want, result)
		}
	}
}
