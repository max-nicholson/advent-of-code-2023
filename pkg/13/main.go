package main

import (
	"errors"
	"fmt"
	"log"
	"slices"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

type Terrain int

type Pattern [][]Terrain

const (
	Ash Terrain = iota + 1
	Rock
)

type ReflectionDirection int

const (
	Vertical ReflectionDirection = iota + 1
	Horizontal
)

type Reflection struct {
	direction ReflectionDirection
	position  int
}

var ErrNoReflection = errors.New("no reflection found")

func (p *Pattern) isRowEqual(a, b int) bool {
	return slices.Equal((*p)[a], (*p)[b])
}

func (p *Pattern) rowDiff(a, b int) int {
	pattern := *p
	var diff int
	for column := 0; column < len(pattern[0]); column++ {
		if pattern[a][column] != pattern[b][column] {
			diff++
		}
	}
	return diff
}

func (p *Pattern) columnDiff(a, b int) int {
	pattern := *p
	var diff int
	for row := 0; row < len(pattern); row++ {
		if pattern[row][a] != pattern[row][b] {
			diff++
		}
	}
	return diff
}

func (p *Pattern) isColumnEqual(a, b int) bool {
	pattern := *p
	for row := 0; row < len(pattern); row++ {
		if pattern[row][a] != pattern[row][b] {
			return false
		}
	}
	return true
}

func (p *Pattern) Reflection() (Reflection, error) {
	pattern := *p
	rows := len(pattern)
	columns := len(pattern[0])

	for row := 0; row < rows-1; row++ {
		above := row
		below := row + 1
		reflection := true
		for above >= 0 && below < rows {
			if !p.isRowEqual(above, below) {
				reflection = false
				break
			}
			above--
			below++
		}
		if reflection {
			return Reflection{direction: Horizontal, position: row}, nil
		}
	}

	for column := 0; column < columns-1; column++ {
		left := column
		right := column + 1
		reflection := true
		for left >= 0 && right < columns {
			if !p.isColumnEqual(left, right) {
				reflection = false
				break
			}
			left--
			right++
		}
		if reflection {
			return Reflection{direction: Vertical, position: column}, nil
		}
	}

	// No reflection found
	return Reflection{}, ErrNoReflection
}

func (p *Pattern) ReflectionWithSmudge() (Reflection, error) {
	pattern := *p
	rows := len(pattern)
	columns := len(pattern[0])

	for row := 0; row < rows-1; row++ {
		above := row
		below := row + 1
		diff := 0
		for above >= 0 && below < rows {
			diff += p.rowDiff(above, below)
			if diff > 1 {
				break
			}
			above--
			below++
		}
		if diff == 1 {
			return Reflection{direction: Horizontal, position: row}, nil
		}
	}

	for column := 0; column < columns-1; column++ {
		left := column
		right := column + 1
		diff := 0
		for left >= 0 && right < columns {
			diff += p.columnDiff(left, right)
			left--
			right++
		}
		if diff == 1 {
			return Reflection{direction: Vertical, position: column}, nil
		}
	}

	// No reflection found
	return Reflection{}, ErrNoReflection
}

func main() {
	lines, err := lib.ReadLines("pkg/13/input.txt")
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

func Patterns(lines []string) []Pattern {
	patterns := []Pattern{}
	pattern := Pattern{}
	l := len(lines)
	i := 0

	for i < l {
		line := lines[i]
		if lines[i] == "" {
			patterns = append(patterns, pattern)
			pattern = Pattern{}
			i++
			continue
		}

		columns := len(line)
		row := make([]Terrain, columns)
		for j := 0; j < columns; j++ {
			var terrain Terrain
			if line[j] == '#' {
				terrain = Rock
			} else {
				terrain = Ash
			}
			row[j] = terrain
		}
		pattern = append(pattern, row)
		i++
	}

	if len(pattern) > 0 {
		patterns = append(patterns, pattern)
	}

	return patterns
}

func Part1(lines []string) (int, error) {
	total := 0
	patterns := Patterns(lines)

	for i, pattern := range patterns {
		reflection, err := pattern.Reflection()
		if err != nil {
			return 0, fmt.Errorf("no reflection for pattern %d: %w", i, err)
		}

		if reflection.direction == Vertical {
			total += (reflection.position + 1)
		} else {
			total += ((reflection.position + 1) * 100)
		}
	}

	return total, nil
}

func Part2(lines []string) (int, error) {
	total := 0
	patterns := Patterns(lines)

	for i, pattern := range patterns {
		reflection, err := pattern.ReflectionWithSmudge()
		if err != nil {
			return 0, fmt.Errorf("no reflection for pattern %d: %w", i, err)
		}

		if reflection.direction == Vertical {
			total += (reflection.position + 1)
		} else {
			total += ((reflection.position + 1) * 100)
		}
	}

	return total, nil
}
