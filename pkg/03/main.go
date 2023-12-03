package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/03/input.txt")
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

type Point struct {
	r int
	c int
}

var neighbours = []Point{
	{r: -1, c: -1},
	{r: -1, c: 0},
	{r: -1, c: 1},
	{r: 0, c: -1},
	{r: 0, c: 0},
	{r: 0, c: 1},
	{r: 1, c: -1},
	{r: 1, c: 0},
	{r: 1, c: 1},
}

func Part1(lines []string) (int, error) {
	total := 0
	seen := map[Point]int{}
	maxR := len(lines) - 1
	maxC := len(lines[0]) - 1
	for r, line := range lines {
		chars := []rune(line)
		for c, char := range chars {
			point := Point{r, c}
			if char == '.' {
				seen[point] = 1
			} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
				for _, direction := range neighbours {
					candidate := Point{r: direction.r + point.r, c: direction.c + point.c}
					if candidate.r < 0 || candidate.c < 0 || candidate.r > maxR || candidate.c > maxC {
						continue
					}
					_, ok := seen[candidate]
					if ok {
						continue
					}
					b := lines[candidate.r][candidate.c]
					if !unicode.IsDigit(rune(b)) {
						continue
					}

					seen[candidate] = 1
					start := candidate.c
					end := candidate.c
					for i := candidate.c - 1; i >= 0; i-- {
						char := lines[candidate.r][i]
						if !unicode.IsDigit(rune(char)) {
							break
						}
						start = i
						seen[Point{r: candidate.r, c: i}] = 1
					}
					for i := candidate.c + 1; i <= maxC; i++ {
						char := lines[candidate.r][i]
						if !unicode.IsDigit(rune(char)) {
							break
						}
						end = i
						seen[Point{r: candidate.r, c: i}] = 1
					}
					partNumber, err := strconv.Atoi(lines[candidate.r][start : end+1])
					if err != nil {
						return 0, fmt.Errorf("expected %s to be an integer: %w", lines[r][start:end], err)
					}
					total += partNumber
				}
			} else if unicode.IsDigit(char) {
				// pass
			} else {
				return 0, fmt.Errorf("unexpected char %c at Point(%d, %d)", char, c, r)
			}

		}
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	total := 0
	seen := map[Point]int{}
	maxR := len(lines) - 1
	maxC := len(lines[0]) - 1
	for r, line := range lines {
		chars := []rune(line)
		for c, char := range chars {
			point := Point{r, c}
			if char == '.' {
				seen[point] = 1
			} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
				partNumbers := []int{}
				for _, direction := range neighbours {
					candidate := Point{r: direction.r + point.r, c: direction.c + point.c}
					if candidate.r < 0 || candidate.c < 0 || candidate.r > maxR || candidate.c > maxC {
						continue
					}
					_, ok := seen[candidate]
					if ok {
						continue
					}
					b := lines[candidate.r][candidate.c]
					if !unicode.IsDigit(rune(b)) {
						continue
					}

					seen[candidate] = 1

					if len(partNumbers) >= 2 {
						break
					}

					start := candidate.c
					end := candidate.c
					for i := candidate.c - 1; i >= 0; i-- {
						char := lines[candidate.r][i]
						if !unicode.IsDigit(rune(char)) {
							break
						}
						start = i
						seen[Point{r: candidate.r, c: i}] = 1
					}
					for i := candidate.c + 1; i <= maxC; i++ {
						char := lines[candidate.r][i]
						if !unicode.IsDigit(rune(char)) {
							break
						}
						end = i
						seen[Point{r: candidate.r, c: i}] = 1
					}
					partNumber, err := strconv.Atoi(lines[candidate.r][start : end+1])
					if err != nil {
						return 0, fmt.Errorf("expected %s to be an integer: %w", lines[r][start:end], err)
					}
					partNumbers = append(partNumbers, partNumber)
				}

				if len(partNumbers) == 2 {
					gearRatio := partNumbers[0] * partNumbers[1]
					total += gearRatio
				}
			} else if unicode.IsDigit(char) {
				// pass
			} else {
				return 0, fmt.Errorf("unexpected char %c at Point(%d, %d)", char, c, r)
			}

		}
	}

	return total, nil
}
