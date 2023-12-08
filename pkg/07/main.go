package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

var strengthByCard = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

type GameOptions struct {
	JokerRule bool
}

type Hand struct {
	Cards    []int
	handType int
}

func NewHand(cards []string, options *GameOptions) (*Hand, error) {
	if len(cards) != 5 {
		return nil, fmt.Errorf("expected a hand to contain 5 cards")
	}

	var JOKER_RULE_STRENGTH = 0

	hand := make([]int, 5)
	set := make(map[int]int)
	for i, card := range cards {
		strength, ok := strengthByCard[cards[i]]
		if !ok {
			return nil, fmt.Errorf("expected card %s to be one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2", cards[i])
		}
		if card == "J" && options.JokerRule {
			strength = JOKER_RULE_STRENGTH
		}
		hand[i] = strength
		set[strength] += 1
	}

	var handType int

	jokers := set[JOKER_RULE_STRENGTH]
	uniqueCards := len(set)
	if options.JokerRule && jokers > 0 {
		// since the jokers are in the `set`, they also need subtracting
		uniqueCards -= 1
	}
	if uniqueCards <= 1 {
		// 5 of a kind
		handType = 7
	} else if uniqueCards == 2 {
		// 4 of a kind OR full house
		handType = 5 // default to full house
		for _, v := range set {
			if v == 4-jokers {
				// we have a 4 of a kind
				handType = 6
				break
			}
		}
	} else if uniqueCards == 3 {
		// 3 of a kind OR 2 pair
		handType = 3 // default to 2 pair
		for _, v := range set {
			if v == 3-jokers {
				// we have a 3 of a kind
				handType = 4
				break
			}
		}
	} else if uniqueCards == 4 {
		// 1 pair
		handType = 2
	} else {
		// high card
		handType = 1
	}

	return &Hand{
		Cards:    hand,
		handType: handType,
	}, nil
}

type Round struct {
	Hand *Hand
	Bid  int
}

func ParseRound(round string, options *GameOptions) (*Round, error) {
	parts := strings.Split(round, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("expected a hand and bid separated by a space")
	}
	cards := strings.Split(parts[0], "")

	hand, err := NewHand(cards, options)
	if err != nil {
		return nil, fmt.Errorf("failed to parse hand %v: %w", cards, err)
	}

	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse bid %s: %w", parts[1], err)
	}

	return &Round{
		Hand: hand,
		Bid:  bid,
	}, nil
}

func main() {
	lines, err := lib.ReadLines("pkg/07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

func SortRounds(rounds []*Round) {
	sort.Slice(rounds, func(i, j int) bool {
		handTypeDiff := rounds[i].Hand.handType - rounds[j].Hand.handType
		if handTypeDiff > 0 {
			return false
		} else if handTypeDiff < 0 {
			return true
		}

		for card := 0; card < 5; card++ {
			cardDiff := rounds[i].Hand.Cards[card] - rounds[j].Hand.Cards[card]
			if cardDiff > 0 {
				return false
			} else if cardDiff < 0 {
				return true
			}
		}
		return true
	})
}

func Part1(lines []string) (int, error) {
	rounds := []*Round{}
	gameOptions := &GameOptions{JokerRule: false}
	for i, line := range lines {
		round, err := ParseRound(line, gameOptions)
		if err != nil {
			return 0, fmt.Errorf("failed to parse round %d: %w", i+1, err)
		}
		// fmt.Printf("round: %d\nhand: %v\nbind: %d\n\n", i+1, round.Hand, round.Bid)
		rounds = append(rounds, round)
	}
	SortRounds(rounds)
	winnings := 0
	for i, round := range rounds {
		rank := i + 1
		winnings += rank * round.Bid
	}
	return winnings, nil
}

func Part2(lines []string) (int, error) {
	rounds := []*Round{}
	gameOptions := &GameOptions{JokerRule: true}
	for i, line := range lines {
		round, err := ParseRound(line, gameOptions)
		if err != nil {
			return 0, fmt.Errorf("failed to parse round %d: %w", i+1, err)
		}
		// fmt.Printf("round: %d\nhand: %v\nbind: %d\n\n", i+1, round.Hand, round.Bid)
		rounds = append(rounds, round)
	}
	SortRounds(rounds)
	winnings := 0
	for i, round := range rounds {
		rank := i + 1
		winnings += rank * round.Bid
	}
	return winnings, nil
}
