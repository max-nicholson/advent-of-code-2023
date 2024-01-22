package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

type Direction int

const (
	North Direction = iota + 1
	West
	South
	East
)

var cycle []Direction = []Direction{North, West, South, East}

type Space int

const (
	RoundedRock Space = iota + 1
	CubeShapedRock
	Empty
)

func ParseSpace(c rune) (Space, error) {
	switch c {
	case 'O':
		return RoundedRock, nil
	case '#':
		return CubeShapedRock, nil
	case '.':
		return Empty, nil
	default:
		return 0, fmt.Errorf("invalid space %v", c)
	}
}

type Platform [][]Space

func (p *Platform) Tilt(direction Direction) {
	platform := *p
	columns := len(platform[0])
	rows := len(platform)

	if direction == North {
		for column := 0; column < columns; column++ {
			availableSpaces := 0
			for row := 0; row < rows; row++ {
				current := platform[row][column]
				if current == Empty {
					availableSpaces += 1
					continue
				}

				if current == CubeShapedRock {
					availableSpaces = 0
					continue
				}

				// RoundedRock
				if availableSpaces > 0 {
					platform[row][column] = Empty
					platform[row-availableSpaces][column] = RoundedRock
					// availableSpaces remains the same, as we have "swapped" the current space for
					// a space above
				} else {
					availableSpaces = 0
				}
			}
		}
	} else if direction == South {
		for column := 0; column < columns; column++ {
			availableSpaces := 0
			for row := rows - 1; row >= 0; row-- {
				current := platform[row][column]
				if current == Empty {
					availableSpaces += 1
					continue
				}

				if current == CubeShapedRock {
					availableSpaces = 0
					continue
				}

				// RoundedRock
				if availableSpaces > 0 {
					platform[row][column] = Empty
					platform[row+availableSpaces][column] = RoundedRock
					// availableSpaces remains the same, as we have "swapped" the current space for
					// a space above
				} else {
					availableSpaces = 0
				}
			}
		}
	} else if direction == West {
		for row := 0; row < rows; row++ {
			availableSpaces := 0
			for column := 0; column < columns; column++ {
				current := platform[row][column]
				if current == Empty {
					availableSpaces += 1
					continue
				}

				if current == CubeShapedRock {
					availableSpaces = 0
					continue
				}

				// RoundedRock
				if availableSpaces > 0 {
					platform[row][column] = Empty
					platform[row][column-availableSpaces] = RoundedRock
					// availableSpaces remains the same, as we have "swapped" the current space for
					// a space above
				} else {
					availableSpaces = 0
				}
			}
		}
	} else if direction == East {
		for row := 0; row < rows; row++ {
			availableSpaces := 0
			for column := columns - 1; column >= 0; column-- {
				current := platform[row][column]
				if current == Empty {
					availableSpaces += 1
					continue
				}

				if current == CubeShapedRock {
					availableSpaces = 0
					continue
				}

				// RoundedRock
				if availableSpaces > 0 {
					platform[row][column] = Empty
					platform[row][column+availableSpaces] = RoundedRock
					// availableSpaces remains the same, as we have "swapped" the current space for
					// a space above
				} else {
					availableSpaces = 0
				}
			}
		}
	}
}

func (p *Platform) Cycle() {
	for _, direction := range cycle {
		p.Tilt(direction)
	}
}

func (p *Platform) Load() int {
	platform := *p
	rows := len(platform)
	load := 0

	for i, row := range platform {
		for _, space := range row {
			if space != RoundedRock {
				continue
			}

			load += (rows - i)
		}
	}

	return load
}

func (p *Platform) Hash() uint64 {
	hash := fnv.New64a()
	platform := *p
	for _, row := range platform {
		for _, space := range row {
			hash.Write([]byte{byte(space)})
		}
	}
	return hash.Sum64()
}

func (p *Platform) Debug() string {
	platform := *p
	var sb strings.Builder
	for _, row := range platform {
		for _, space := range row {
			switch space {
			case RoundedRock:
				sb.WriteRune('O')
			case CubeShapedRock:
				sb.WriteRune('#')
			case Empty:
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func ParsePlatform(lines []string) (*Platform, error) {
	columns := len(lines[0])
	rows := len(lines)
	platform := make(Platform, rows)
	for i, line := range lines {
		spaces := make([]Space, columns)
		for j, c := range line {
			space, err := ParseSpace(c)
			if err != nil {
				return &Platform{}, fmt.Errorf("invalid space line %d position %d: %w", i, j, err)
			}
			spaces[j] = space
		}
		platform[i] = spaces
	}
	return &platform, nil
}

func main() {
	lines, err := lib.ReadLines("pkg/14/input.txt")
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

func Part1(lines []string) (int, error) {
	platform, err := ParsePlatform(lines)
	if err != nil {
		return 0, fmt.Errorf("failed to parse platform: %w", err)
	}

	platform.Tilt(North)

	return platform.Load(), nil
}

func Part2(lines []string) (int, error) {
	platform, err := ParsePlatform(lines)
	if err != nil {
		return 0, fmt.Errorf("failed to parse platform: %w", err)
	}

	cache := map[uint64]int{
		platform.Hash(): 0,
	}
	iterations := 1_000_000_000
	for i := 0; i < iterations; i++ {
		platform.Cycle()
		hash := platform.Hash()
		firstSeen, ok := cache[hash]
		if ok {
			loopLength := i + 1 - firstSeen
			left := iterations - i - 1
			remainder := left % loopLength
			for j := 0; j < remainder; j++ {
				platform.Cycle()
			}
			fmt.Println(platform.Debug())
			return platform.Load(), nil
		}
		cache[hash] = i + 1
	}

	fmt.Println(platform.Debug())
	return platform.Load(), nil
}
