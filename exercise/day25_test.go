package exercise

import (
	"testing"
)

func TestDay25Part1(t *testing.T) {
	input := []string{
		"#####",
		".####",
		".####",
		".####",
		".#.#.",
		".#...",
		".....",
		"",
		"#####",
		"##.##",
		".#.##",
		"...##",
		"...#.",
		"...#.",
		".....",
		"",
		".....",
		"#....",
		"#....",
		"#...#",
		"#.#.#",
		"#.###",
		"#####",
		"",
		".....",
		".....",
		"#.#..",
		"###..",
		"###.#",
		"###.#",
		"#####",
		"",
		".....",
		".....",
		".....",
		"#....",
		"#.#..",
		"#.#.#",
		"#####",
	}

	d25 := Day25{}

	locks, keys := d25.parseInput(input)
	keysThatFitLocks := d25.Part1(locks, keys)

	expectedKeysThatFitLocks := 3

	if keysThatFitLocks != expectedKeysThatFitLocks {
		t.Errorf("Day 25 - Part 1 (keys that fit locks) Test:\nwant %v\ngot %v\n", expectedKeysThatFitLocks, keysThatFitLocks)
	}
}

func TestDay25Part2(t *testing.T) {
}
