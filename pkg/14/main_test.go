package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 136 {
		t.Fatalf("expected 136, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 64 {
		t.Fatalf("expected 64, got %d", result)
	}
}
