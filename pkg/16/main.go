package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/16/input.txt")
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

type Point struct {
	row    int
	column int
}

var DELTA = map[byte]*Point{
	'N': {row: -1},
	'S': {row: 1},
	'W': {column: -1},
	'E': {column: 1},
}

type Visit struct {
	Point
	direction byte
}

func PrintEnergized(energized map[Point]struct{}, rows int, columns int) {
	printLines := make([]string, rows)
	for row := 0; row < rows; row++ {
		var sb strings.Builder
		for column := 0; column < columns; column++ {
			_, ok := energized[Point{row, column}]
			if ok {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		printLines[row] = sb.String()
	}
	fmt.Println(strings.Join(printLines, "\n"))
}

func CalculateEnergized(lines []string, start *Visit) (map[Point]struct{}, error) {
	energized := map[Point]struct{}{}
	visited := map[Visit]struct{}{}

	stack := []Visit{
		*start,
	}

	rows := len(lines)
	columns := len(lines[0])

	for len(stack) > 0 {
		visit := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		point := visit.Point
		if point.column < 0 || point.row < 0 || point.row >= rows || point.column >= columns {
			continue
		}

		_, ok := visited[visit]
		if ok {
			continue
		}

		value := lines[point.row][point.column]
		energized[point] = struct{}{}
		visited[visit] = struct{}{}

		if value == '.' {
			delta := DELTA[visit.direction]
			stack = append(
				stack,
				Visit{Point{row: point.row + delta.row, column: point.column + delta.column}, visit.direction},
			)
		} else if value == '/' {
			if visit.direction == 'S' {
				stack = append(
					stack,
					Visit{Point{row: point.row, column: point.column - 1}, 'W'},
				)
			} else if visit.direction == 'E' {
				stack = append(
					stack,
					Visit{Point{row: point.row - 1, column: point.column}, 'N'},
				)
			} else if visit.direction == 'N' {
				stack = append(
					stack,
					Visit{Point{row: point.row, column: point.column + 1}, 'E'},
				)
			} else if visit.direction == 'W' {
				stack = append(
					stack,
					Visit{Point{row: point.row + 1, column: point.column}, 'S'},
				)
			}
		} else if value == '\\' {
			if visit.direction == 'S' {
				stack = append(
					stack,
					Visit{Point{row: point.row, column: point.column + 1}, 'E'},
				)
			} else if visit.direction == 'W' {
				stack = append(
					stack,
					Visit{Point{row: point.row - 1, column: point.column}, 'N'},
				)
			} else if visit.direction == 'N' {
				stack = append(
					stack,
					Visit{Point{row: point.row, column: point.column - 1}, 'W'},
				)
			} else if visit.direction == 'E' {
				stack = append(
					stack,
					Visit{Point{row: point.row + 1, column: point.column}, 'S'},
				)
			}
		} else if value == '|' {
			if visit.direction == 'N' || visit.direction == 'S' {
				delta := DELTA[visit.direction]
				stack = append(
					stack,
					Visit{Point{row: point.row + delta.row, column: point.column + delta.column}, visit.direction},
				)
			} else {
				stack = append(
					stack,
					Visit{Point{row: point.row - 1, column: point.column}, 'N'},
					Visit{Point{row: point.row + 1, column: point.column}, 'S'},
				)
			}
		} else if value == '-' {
			if visit.direction == 'W' || visit.direction == 'E' {
				delta := DELTA[visit.direction]
				stack = append(
					stack,
					Visit{Point{row: point.row + delta.row, column: point.column + delta.column}, visit.direction},
				)
			} else {
				stack = append(
					stack,
					Visit{Point{row: point.row, column: point.column - 1}, 'W'},
					Visit{Point{row: point.row, column: point.column + 1}, 'E'},
				)
			}
		} else {
			return energized, fmt.Errorf("unexpected value %s at (%d, %d)", string(value), point.row, point.column)
		}
	}
	return energized, nil
}

func Part1(lines []string) (int, error) {
	energized, err := CalculateEnergized(lines, &Visit{Point{0, 0}, 'E'})
	if err != nil {
		return 0, fmt.Errorf("failed to calculate energized tiles: %w", err)
	}

	return len(energized), nil
}

func Part2(lines []string) (int, error) {
	maxEnergized := 0
	columns := len(lines[0])
	rows := len(lines)

	for column := 0; column < columns; column++ {
		var edges = []struct {
			row       int
			direction byte
		}{
			{row: 0, direction: 'S'},
			{row: rows - 1, direction: 'N'},
		}
		for _, edge := range edges {
			energized, err := CalculateEnergized(lines, &Visit{Point{edge.row, column}, edge.direction})
			if err != nil {
				return 0, fmt.Errorf("failed to calculate energized tiles for (%d, %d): %w", edge.row, column, err)
			}
			maxEnergized = max(maxEnergized, len(energized))
		}
	}

	for row := 0; row < rows; row++ {
		var edges = []struct {
			column    int
			direction byte
		}{
			{column: 0, direction: 'E'},
			{column: columns - 1, direction: 'W'},
		}
		for _, edge := range edges {
			energized, err := CalculateEnergized(lines, &Visit{Point{row, edge.column}, edge.direction})
			if err != nil {
				return 0, fmt.Errorf("failed to calculate energized tiles for (%d, %d): %w", row, edge.column, err)
			}
			maxEnergized = max(maxEnergized, len(energized))
		}
	}

	return maxEnergized, nil
}
