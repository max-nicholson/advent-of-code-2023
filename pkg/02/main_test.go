package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	total, err := Part1(strings.Split(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 8 {
		t.Fatalf("expected 8, got %d", total)
	}
}

func TestPart2(t *testing.T) {
	total, err := Part2(strings.Split(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 2286 {
		t.Fatalf("expected 2286, got %d", total)
	}
}
