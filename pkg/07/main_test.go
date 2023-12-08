package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	location, err := Part1(strings.Split(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if location != 6440 {
		t.Fatalf("expected 6440, got %d", location)
	}
}

func TestPart2(t *testing.T) {
	total, err := Part2(strings.Split(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if total != 5905 {
		t.Fatalf("expected 5905, got %d", total)
	}
}

func TestHandStrength(t *testing.T) {
	type test struct {
		cards   string
		options *GameOptions
		want    int
	}

	tests := []test{
		{cards: "KKKKK", options: &GameOptions{}, want: 7},
		{cards: "JJJJJ", options: &GameOptions{JokerRule: true}, want: 7},
		{cards: "KKKKJ", options: &GameOptions{}, want: 6},
		{cards: "KKKKJ", options: &GameOptions{JokerRule: true}, want: 7},
		{cards: "KKKJJ", options: &GameOptions{}, want: 5},
		{cards: "KKKJJ", options: &GameOptions{JokerRule: true}, want: 7},
		{cards: "KKJJJ", options: &GameOptions{}, want: 5},
		{cards: "KKJJJ", options: &GameOptions{JokerRule: true}, want: 7},
		{cards: "234JJ", options: &GameOptions{}, want: 2},
		{cards: "234JJ", options: &GameOptions{JokerRule: true}, want: 4},
		{cards: "2345J", options: &GameOptions{}, want: 1},
		{cards: "2345J", options: &GameOptions{JokerRule: true}, want: 2},
		{cards: "32T3K", options: &GameOptions{}, want: 2},
		{cards: "32T3K", options: &GameOptions{JokerRule: true}, want: 2},
		{cards: "KK677", options: &GameOptions{}, want: 3},
		{cards: "KK677", options: &GameOptions{JokerRule: true}, want: 3},
		{cards: "KTJJT", options: &GameOptions{}, want: 3},
		{cards: "KTJJT", options: &GameOptions{JokerRule: true}, want: 6},
		{cards: "T55J5", options: &GameOptions{}, want: 4},
		{cards: "T55J5", options: &GameOptions{JokerRule: true}, want: 6},
		{cards: "QQQJA", options: &GameOptions{}, want: 4},
		{cards: "QQQJA", options: &GameOptions{JokerRule: true}, want: 6},
	}

	for _, tc := range tests {
		hand, _ := NewHand(strings.Split(tc.cards, ""), tc.options)
		got := hand.handType
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
