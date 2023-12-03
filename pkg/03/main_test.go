package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	total, err := Part1(strings.Split(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 4361 {
		t.Fatalf("expected 4361, got %d", total)
	}
}

func TestPart2(t *testing.T) {
	total, err := Part2(strings.Split(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 467835 {
		t.Fatalf("expected 467835, got %d", total)
	}
}
