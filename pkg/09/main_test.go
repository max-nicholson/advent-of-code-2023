package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	location, err := Part1(strings.Split(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if location != 114 {
		t.Fatalf("expected 114, got %d", location)
	}
}

func TestPart2(t *testing.T) {
	total, err := Part2(strings.Split(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 2 {
		t.Fatalf("expected 2, got %d", total)
	}
}
