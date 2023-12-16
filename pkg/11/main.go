package main

import (
	"fmt"
	"log"
	"math"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/11/input.txt")
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

func SumOfShortestPaths(lines []string, expansionFactor int) (int, error) {
	rows := len(lines)
	columns := len(lines[0])
	// If we want to "double", we only need to multiply each row/column by 1
	expansionFactor -= 1

	rowsWithoutGalaxies := make(map[int]struct{}, rows)
	columnsWithoutGalaxies := make(map[int]struct{}, columns)
	for row := 0; row < rows; row++ {
		rowsWithoutGalaxies[row] = struct{}{}
	}
	for column := 0; column < columns; column++ {
		columnsWithoutGalaxies[column] = struct{}{}
	}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if lines[row][column] == '#' {
				delete(rowsWithoutGalaxies, row)
				delete(columnsWithoutGalaxies, column)
			}
		}
	}

	galaxies := []Point{}
	var rowOffset int
	for row := 0; row < rows; row++ {
		_, noGalaxies := rowsWithoutGalaxies[row]
		if noGalaxies {
			rowOffset += expansionFactor
			continue
		}

		var columnOffset int
		for column := 0; column < columns; column++ {
			_, noGalaxies := columnsWithoutGalaxies[column]
			if noGalaxies {
				columnOffset += expansionFactor
				continue
			}

			if lines[row][column] == '#' {
				galaxies = append(galaxies, Point{
					row: row + rowOffset, column: column + columnOffset,
				})
			}
		}
	}

	pairs := lib.Pairs(galaxies)
	steps := 0

	for _, pair := range pairs {
		start := pair[0]
		end := pair[1]

		steps += int(math.Abs(float64(end.column-start.column)) + math.Abs(float64(end.row-start.row)))
	}

	return steps, nil
}

func Part1(lines []string) (int, error) {
	return SumOfShortestPaths(lines, 2)
}

func Part2(lines []string) (int, error) {
	return SumOfShortestPaths(lines, 1_000_000)
}
