package main

import (
	"fmt"
	"log"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

func main() {
	grid, err := lib.ReadLines("pkg/10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(grid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: %d\n", part1)

	part2, err := Part2(grid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

type Point struct {
	Row    int
	Column int
}

func NextDirection(pipe byte, direction byte) (byte, error) {
	if pipe == '|' {
		if direction == 'S' {
			return 'S', nil
		} else if direction == 'N' {
			return 'N', nil
		}
	} else if pipe == '-' {
		if direction == 'W' {
			return 'W', nil
		} else if direction == 'E' {
			return 'E', nil
		}
	} else if pipe == 'L' {
		if direction == 'W' {
			return 'N', nil
		} else if direction == 'S' {
			return 'E', nil
		}
	} else if pipe == 'J' {
		if direction == 'E' {
			return 'N', nil
		} else if direction == 'S' {
			return 'W', nil
		}
	} else if pipe == '7' {
		if direction == 'N' {
			return 'W', nil
		} else if direction == 'E' {
			return 'S', nil
		}
	} else if pipe == 'F' {
		if direction == 'N' {
			return 'E', nil
		} else if direction == 'W' {
			return 'S', nil
		}
	}

	return 0, fmt.Errorf("invalid pipe/direction")
}

func FindStart(grid []string) (*Point, bool) {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if grid[row][column] == 'S' {
				return &Point{Row: row, Column: column}, true
			}
		}
	}
	return &Point{}, false
}

func FindStartDirection(grid []string, start *Point) byte {
	directions := []struct {
		dr        int
		dc        int
		direction byte
	}{
		{-1, 0, 'N'},
		{0, 1, 'E'},
		{0, -1, 'W'},
		{1, 0, 'S'},
	}
	for _, dir := range directions {
		row := start.Row + dir.dr
		column := start.Column + dir.dc
		if row < 0 || column < 0 {
			continue
		}
		nextDirection, err := NextDirection(grid[row][column], dir.direction)
		if err == nil {
			return nextDirection
		}
	}
	return 0
}

var DELTA = map[byte]*Point{
	'N': {Row: -1},
	'S': {Row: 1},
	'W': {Column: -1},
	'E': {Column: 1},
}

func Part1(grid []string) (int, error) {
	start, ok := FindStart(grid)
	if !ok {
		return 0, fmt.Errorf("start not found")
	}

	column := start.Column
	row := start.Row
	direction := FindStartDirection(grid, start)

	if direction == 0 {
		return 0, fmt.Errorf("unable to find a direction from S")
	}

	delta := DELTA[direction]
	row += delta.Row
	column += delta.Column

	distance := 1

	for !(column == start.Column && row == start.Row) {
		nextDirection, err := NextDirection(grid[row][column], direction)
		//fmt.Printf("[%d][%d] %s %s\n", row, column, string(grid[row][column]), string(direction))
		if nextDirection == 0 {
			return 0, fmt.Errorf(
				"unable to find a next direction from %s at [%d][%d] going %s: %w",
				string(grid[row][column]),
				row,
				column,
				string(direction),
				err,
			)
		}
		direction = nextDirection
		delta := DELTA[direction]
		row += delta.Row
		column += delta.Column
		distance += 1
	}

	if distance%2 == 0 {
		return distance / 2, nil
	} else {
		return (distance + 1) / 2, nil
	}
}

func Part2(grid []string) (int, error) {
	start, ok := FindStart(grid)
	if !ok {
		return 0, fmt.Errorf("start not found")
	}

	column := start.Column
	row := start.Row
	direction := FindStartDirection(grid, start)

	if direction == 0 {
		return 0, fmt.Errorf("unable to find a direction from S")
	}

	delta := DELTA[direction]
	row += delta.Row
	column += delta.Column
	loop := map[Point]struct{}{
		*start: {},
	}

	for !(column == start.Column && row == start.Row) {
		loop[Point{Row: row, Column: column}] = struct{}{}
		nextDirection, err := NextDirection(grid[row][column], direction)
		//fmt.Printf("[%d][%d] %s %s\n", row, column, string(grid[row][column]), string(direction))
		if nextDirection == 0 {
			return 0, fmt.Errorf(
				"unable to find a next direction from %s at [%d][%d] going %s: %w",
				string(grid[row][column]),
				row,
				column,
				string(direction),
				err,
			)
		}
		direction = nextDirection
		delta := DELTA[direction]
		row += delta.Row
		column += delta.Column
	}

	var minRow = 99999
	var minColumn = 99999
	var maxRow = 0
	var maxColumn = 0
	for point := range loop {
		minRow = min(minRow, point.Row)
		maxRow = max(maxRow, point.Row)
		minColumn = min(minColumn, point.Column)
		maxColumn = max(maxColumn, point.Column)
	}

	tiles := 0
	for row := minRow; row <= maxRow; row++ {
		for column := minColumn; column <= maxColumn; column++ {
			point := Point{Row: row, Column: column}
			_, partOfLoop := loop[point]
			if partOfLoop {
				continue
			}

			// diagonal raycasting
			crosses := 0
			r, c := row, column

			for r <= maxRow && c <= maxColumn {
				p := grid[r][c]
				_, ok := loop[Point{Row: r, Column: c}]
				if ok && p != 'L' && p != '7' {
					crosses += 1
				}
				r += 1
				c += 1
			}

			if crosses%2 == 1 {
				tiles += 1
			}
		}
	}

	return tiles, nil
}
