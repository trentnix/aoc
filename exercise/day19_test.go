package exercise

import (
	"testing"
)

func TestDay19Part1(t *testing.T) {
	input := []string{
		"r, wr, b, g, bwu, rb, gb, br",
		"",
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}

	d19 := Day19{}

	towels, desiredDesigns := d19.parseInput(input)

	possibleDesigns := d19.Part1(towels, desiredDesigns)
	expectedPossibleDesigns := 6

	if possibleDesigns != expectedPossibleDesigns {
		t.Errorf("Day 19 - Part 1 (possible designs) Test:\nwant %v\ngot %v\n", expectedPossibleDesigns, possibleDesigns)
	}
}

func TestDay19Part2(t *testing.T) {
}
