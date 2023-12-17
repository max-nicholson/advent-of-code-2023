package main

import (
	"container/heap"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

type Direction struct {
	row int
	column int
}

var (
	W = Direction{row: 0, column: -1}
	E = Direction{row: 0, column: 1}
	N = Direction{row: -1, column: 0}
	S = Direction{row: 1, column: 0}
)

var nextDirections = map[Direction][]Direction{
	N: {W, N, E},
	E: {N, E, S},
	S: {E, S, W},
	W: {S, W, N},
}

type state struct {
	row      int
	column   int
	direction      Direction
	acc int
}

type item struct {
	cost  int
	state state
}

type PriorityQueue []*item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func main() {
	lines, err := lib.ReadLines("pkg/17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1 := Part1(lines)
	fmt.Printf("part1: %d\n", part1)

	part2 := Part2(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part2: %d\n", part2)
}

func ParseLines(lines []string) [][]int {
	rows := len(lines)
	columns := len(lines[0])
	result := make([][]int, rows)
	for row, line := range lines {
		result[row] = make([]int, columns)
		values := strings.Split(line, "")
		for column, v := range values {
			result[row][column], _ = strconv.Atoi(v)
		}
	}

	return result
}

func Part1(lines []string) int {
	grid := ParseLines(lines)
	return dijkstra(grid, 0, 3)
}

func Part2(lines []string) int {
	grid := ParseLines(lines)
	return dijkstra(grid, 4, 10)
}

func dijkstra(grid [][]int, minStraight int, maxStraight int) int {
	rows := len(grid)
	columns := len(grid[0])

	startRight := state{row: 0, column: 0, direction: E, acc: 0}
	startDown := state{row: 0, column: 0, direction: S, acc: 0}

	pq := PriorityQueue{
		&item{cost: 0, state: startRight},
		&item{cost: 0, state: startDown},
	}

	minCost := map[state]int{startRight: 0, startDown: 0}
	heap.Init(&pq)

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(*item)

		// Already visited this tile, but at a lower cost than the current path
		// No way to go any further at lower cost
		if minCost[curr.state] < curr.cost {
			continue
		}

		if curr.state.row == rows-1 && curr.state.column == columns-1 && curr.state.acc >= minStraight {
			// End state
			return curr.cost
		}

		currentDirection := curr.state.direction

		for _, dir := range nextDirections[currentDirection] {
			if dir == currentDirection && curr.state.acc == maxStraight {
				// Hit the limit going in a straight line, MUST change direction
				continue
			}

			nextRow := curr.state.row + dir.row
			nextColumn := curr.state.column + dir.column

			if nextRow < 0 || nextColumn < 0 || nextRow >= rows || nextColumn >= columns {
				// Out of bounds
				continue
			}

			nextDirTotal := curr.state.acc

			if curr.state.acc < minStraight {
				if dir != curr.state.direction {
					// Need to keep going straight
					continue
				}
				nextDirTotal += 1
			} else {
				if dir != curr.state.direction {
					// Turning
					nextDirTotal = 1
				} else {
					// Straight
					nextDirTotal = nextDirTotal + 1
				}
			}

			nextState := state{row: nextRow, column: nextColumn, acc: nextDirTotal, direction: dir}
			nextCost := curr.cost + grid[nextRow][nextColumn]
			if _, seen := minCost[nextState]; seen && minCost[nextState] <= nextCost {
				// Only skip if we've both seen the nextState before AND at a lower cost
				// (At a lower cost, we'll want to re-explore with the new "baseline" min cost)
				continue
			}

			minCost[nextState] = nextCost
			heap.Push(&pq, &item{cost: nextCost, state: nextState})
		}
	}

	return 0
}
