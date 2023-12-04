package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

type Scratchcard struct {
	WinningNumbers map[int]int
	Numbers        []int
}

func NewScratchcard(card string) (*Scratchcard, error) {
	cardInfo := strings.Split(card, ": ")
	if len(cardInfo) != 2 {
		return nil, fmt.Errorf("unable to parse %s", card)
	}
	data := strings.Split(cardInfo[1], " | ")
	if len(data) != 2 {
		return nil, fmt.Errorf("unable to parse %s", cardInfo)
	}
	numbers := make([]int, 8)
	for i := 0; i < len(data[1]); i += 3 {
		s := data[1][i : i+2]
		v, err := strconv.Atoi(strings.TrimLeft(s, " "))
		if err != nil {
			return nil, fmt.Errorf("expected integer, got %s: %w", s, err)
		}
		numbers = append(numbers, v)
	}
	winningNumbers := make(map[int]int, 25)
	for i := 0; i < len(data[0]); i += 3 {
		s := data[0][i : i+2]
		v, err := strconv.Atoi(strings.TrimLeft(s, " "))
		if err != nil {
			return nil, fmt.Errorf("expected integer, got %s: %w", s, err)
		}
		winningNumbers[v] = 1
	}
	s := Scratchcard{WinningNumbers: winningNumbers, Numbers: numbers}
	return &s, nil
}

func (s *Scratchcard) IsWinningNumber(number int) bool {
	_, ok := s.WinningNumbers[number]
	return ok
}

func (s *Scratchcard) Matches() int {
	matches := 0
	for _, number := range s.Numbers {
		if s.IsWinningNumber(number) {
			matches += 1
		}
	}
	return matches
}

func (s *Scratchcard) Points() int {
	matches := s.Matches()
	if matches == 0 {
		return 0
	}
	value := 1
	for i := 1; i < matches; i++ {
		value *= 2
	}
	return value
}

func main() {
	lines, err := lib.ReadLines("pkg/04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %v\n", part1)

	part2, err := Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %v\n", part2)
}

func Part1(lines []string) (int, error) {
	total := 0

	for i, line := range lines {
		card, err := NewScratchcard(line)
		if err != nil {
			return 0, fmt.Errorf("unable to parse scratchcard %d: %w", i+1, err)
		}
		total += card.Points()
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	winsByCard := make(map[int]int, len(lines))

	for i, line := range lines {
		card, err := NewScratchcard(line)
		if err != nil {
			return 0, fmt.Errorf("unable to parse scratchcard %d: %w", i+1, err)
		}
		winsByCard[i+1] = card.Matches()
	}

	numberOfCards := make(map[int]int, len(lines))

	for i := 0; i < len(lines); i++ {
		cardId := i + 1
		wins, ok := winsByCard[cardId]
		if !ok {
			return 0, fmt.Errorf("card %d not found in wins map", cardId)
		}
		copies := numberOfCards[cardId]
		var multiplier = 1
		if copies != 0 {
			multiplier += copies
		}
		// original
		numberOfCards[cardId] += 1
		// copies
		for win := 1; win <= wins; win++ {
			numberOfCards[cardId+win] += multiplier
		}
	}

	total := 0
	for _, t := range numberOfCards {
		total += t
	}

	return total, nil
}
