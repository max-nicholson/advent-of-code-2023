package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	total, err := Part1(strings.Split(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 142 {
		t.Fatalf("expected 142, got %d", total)
	}
}

func TestPart2(t *testing.T) {
	total, err := Part2(strings.Split(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 281 {
		t.Fatalf("expected 281, got %d", total)
	}
}

func TestPart2PartialNumber(t *testing.T) {
	total, err := Part2([]string{"315twonehz"})
	if err != nil {
		t.Fatal(err)
	}
	if total != 31 {
		t.Fatalf("want 31, got %d", total)
	}
}
