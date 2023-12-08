package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	location, err := Part1(strings.Split(`Time:      7  15   30
Distance:  9  40  200`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if location != 288 {
		t.Fatalf("expected 288, got %d", location)
	}
}

func TestPart2(t *testing.T) {
	total, err := Part2(strings.Split(`Time:      7  15   30
Distance:  9  40  200`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 71503 {
		t.Fatalf("expected 71503, got %d", total)
	}
}
