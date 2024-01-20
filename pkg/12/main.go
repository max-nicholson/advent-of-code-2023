package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/max-nicholson/advent-of-code-2023/lib"
)

type Condition int

const (
	Unknown = iota
	Operational
	Damaged
)

type Spring struct {
	Condition Condition
}

type Row struct {
	Springs []Spring
	Groups  []int
}

type CacheKey uint64

func (r *Row) Key() CacheKey {
	hash := fnv.New64a()
	springs := len(r.Springs)
	groups := len(r.Groups)
	l := springs + groups + 1
	bytes := make([]byte, l)
	for i, s := range r.Springs {
		bytes[i] = byte(s.Condition)
	}
	// Sentinel value between springs and groups
	bytes[springs] = byte(255)
	for i, g := range r.Groups {
		bytes[springs+i+1] = byte(g)
	}
	hash.Write(bytes)
	return CacheKey(hash.Sum64())
}

func (r *Row) Arrangements(cache map[CacheKey]int) int {
	key := r.Key()
	value, ok := cache[key]
	if ok {
		return value
	}

	if len(r.Groups) == 0 {
		// Can only complete the row with Operational springs
		// i.e Unknown MUST be Operational
		// If any Damaged springs left in the row, not a valid permutation
		var v int
		if slices.ContainsFunc(r.Springs, func(s Spring) bool {
			return s.Condition == Damaged
		}) {
			v = 0
		} else {
			v = 1
		}
		cache[key] = v
		return v
	}

	if len(r.Springs) < (lib.Sum(r.Groups) + len(r.Groups) - 1) {
		// We need another X Damaged (and at least 1 Operational separating them)
		// May be impossible if we have fewer springs than needed
		cache[key] = 0
		return 0
	}

	if r.Springs[0].Condition == Operational {
		arrangements := (&Row{Springs: r.Springs[1:], Groups: slices.Clone(r.Groups)}).Arrangements(cache)
		cache[key] = arrangements
		return arrangements
	}

	var arrangements int
	// Can we consume a damaged group from current position (using a known Damaged or Unknown spring)
	group := r.Groups[0]
	allNonOperational := !slices.ContainsFunc(r.Springs[:group], func(s Spring) bool {
		return s.Condition == Operational
	})
	end := lib.Min(group+1, len(r.Springs))
	if allNonOperational && (
	// Does Damaged block consume all remaining springs
	// Is the first spring AFTER this Damaged block Operational OR Unknown
	(len(r.Springs) > group && r.Springs[group].Condition != Damaged) || len(r.Springs) <= group) {
		arrangements = (&Row{Springs: r.Springs[end:], Groups: r.Groups[1:]}).Arrangements(cache)
	}

	// Use Unknown as an Operational
	if r.Springs[0].Condition == Unknown {
		arrangements += (&Row{Springs: r.Springs[1:], Groups: slices.Clone(r.Groups)}).Arrangements(cache)
	}

	cache[key] = arrangements
	return arrangements
}

func (r *Row) Unfold() {
	n := 5
	newGroups := make([]int, 0, len(r.Groups)*n)
	newSprings := make([]Spring, 0, len(r.Springs)+n)

	for i := 0; i < n; i++ {
		newGroups = append(newGroups, r.Groups...)
		newSprings = append(newSprings, r.Springs...)
		if i != n-1 {
			newSprings = append(newSprings, Spring{Condition: Unknown})
		}
	}
	r.Groups = newGroups
	r.Springs = newSprings
}

func ParseSpring(condition rune) (Spring, error) {
	switch condition {
	case '#':
		return Spring{Condition: Damaged}, nil
	case '.':
		return Spring{Condition: Operational}, nil
	case '?':
		return Spring{Condition: Unknown}, nil
	default:
		return Spring{}, fmt.Errorf("unexpected condition %v", condition)
	}
}

func ParseRow(row string) (*Row, error) {
	parts := strings.Split(row, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("expected condition record and groups of damaged springs space separated")
	}
	conditionRecord := parts[0]

	springs := make([]Spring, len(conditionRecord))
	for i, r := range conditionRecord {
		spring, err := ParseSpring(r)
		if err != nil {
			return nil, fmt.Errorf("invalid spring at index %d: %w", i, err)
		}
		springs[i] = spring
	}

	groupsOfDamagedSprings := strings.Split(parts[1], ",")
	groups := make([]int, len(groupsOfDamagedSprings))
	for i, g := range groupsOfDamagedSprings {
		n, err := strconv.Atoi(g)
		if err != nil {
			return nil, fmt.Errorf("invalid group at index %d, %w", i, err)
		}
		groups[i] = n
	}

	return &Row{Springs: springs, Groups: groups}, nil
}

func main() {
	lines, err := lib.ReadLines("pkg/12/input.txt")
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
	arrangements := 0
	cache := make(map[CacheKey]int)

	for i, line := range lines {
		row, err := ParseRow(line)
		if err != nil {
			return 0, fmt.Errorf("line %d: %w", i, err)
		}
		arrangements += row.Arrangements(cache)
	}

	return arrangements, nil
}

func Part2(lines []string) (int, error) {
	var arrangements int
	cache := make(map[CacheKey]int)

	for i, line := range lines {
		row, err := ParseRow(line)
		if err != nil {
			return 0, fmt.Errorf("line %d: %w", i, err)
		}
		row.Unfold()
		arrangements += row.Arrangements(cache)
	}

	return arrangements, nil
}
