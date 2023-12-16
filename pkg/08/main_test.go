package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	type testCase struct {
		document string
		want int
	}
	testCases := []testCase{
	{ document: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`, want: 2},
	{ document: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`, want: 6},
	}

	for _, test := range testCases {
		result, err := Part1(strings.Split(test.document, "\n"))
		if err != nil {
			t.Error(err)
		}
		if result != test.want {
			t.Errorf("expected %d, got %d", test.want, result)
		}
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 6 {
		t.Fatalf("expected 6, got %d", result)
	}
}
