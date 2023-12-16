package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

func main() {
	lines, err := lib.ReadLines("pkg/08/input.txt")
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

type Node struct {
	Left  string
	Right string
}

func ParseInstructions(line string) []string {
	return strings.Split(line, "")
}

func ParseNodes(lines []string) (map[string]Node, error) {
	var nodes = make(map[string]Node, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("expected line %d to split by ' = ' into 2", i)
		}
		directions := strings.Split(parts[1], ", ")
		nodes[parts[0]] = Node{Left: directions[0][1:4], Right: directions[1][:3]}
	}
	return nodes, nil
}

func Part1(lines []string) (int, error) {
	step := 0

	instructions := ParseInstructions(lines[0])
	nodes, err := ParseNodes(lines[2:])
	if err != nil {
		return 0, fmt.Errorf("failed to parse nodes: %w", err)
	}
	node := "AAA"
	for {
		if node == "ZZZ" {
			return step, nil
		}
		instruction := instructions[step%len(instructions)]
		if instruction == "L" {
			node = nodes[node].Left
		} else {
			node = nodes[node].Right
		}
		step += 1
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Part2(lines []string) (int, error) {
	instructions := ParseInstructions(lines[0])
	allNodes, err := ParseNodes(lines[2:])
	if err != nil {
		return 0, fmt.Errorf("failed to parse nodes: %w", err)
	}
	nodes := []string{}
	for node := range allNodes {
		if strings.HasSuffix(node, "A") {
			nodes = append(nodes, node)
		}
	}

	ends := make([]int, len(nodes))
	for i, start := range nodes {
		node := start
		step := 0

		for {
			if strings.HasSuffix(node, "Z") {
				// Turns out the answer only needs the first "end" node, since subsequent moves
				// are a cycle
				// This would need to be more complicated if we expected a slice of "end"s
				break
			}

			instruction := instructions[step%len(instructions)]
			if instruction == "L" {
				node = allNodes[node].Left
			} else {
				node = allNodes[node].Right
			}
			step += 1
		}

		ends[i] = step
	}

	return LCM(ends[0], ends[1], ends[2:]...), nil
}
