package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	result, err := Part1(strings.Split(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 21 {
		t.Fatalf("expected 21, got %d", result)
	}
}

func TestArrangements(t *testing.T) {
	var testCases = []struct {
		row  string
		want int
	}{
		{
			row:  `???.### 1,1,3`,
			want: 1,
		},
		{
			row:  `.??..??...?##. 1,1,3`,
			want: 4,
		},
		{
			row:  `?#?#?#?#?#?#?#? 1,3,1,6`,
			want: 1,
		},
		{
			row:  `????.#...#... 4,1,1`,
			want: 1,
		},
		{
			row:  `????.######..#####. 1,6,5`,
			want: 4,
		},
		{
			row:  `?###???????? 3,2,1`,
			want: 10,
		},
	}
	for _, testCase := range testCases {
		cache := make(map[CacheKey]int)
		row, err := ParseRow(testCase.row)
		if err != nil {
			t.Error(err)
		}
		if got := row.Arrangements(cache); got != testCase.want {
			t.Errorf("expected %d, got %d", testCase.want, got)
		}
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2(strings.Split(`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`, "\n"))
	if err != nil {
		t.Fatal(err)
	}
	if result != 525152 {
		t.Fatalf("expected 525152, got %d", result)
	}
}
