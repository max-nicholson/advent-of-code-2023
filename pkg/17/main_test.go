package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result := Part1(strings.Split(`2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`, "\n"))
	if result != 102 {
		t.Fatalf("expected 102, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(strings.Split(`2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`, "\n"))
	if result != 94 {
		t.Fatalf("expected 94, got %d", result)
	}
}
