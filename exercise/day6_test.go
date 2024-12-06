package exercise

import (
	"testing"
)

func TestDay6Part1(t *testing.T) {
	input := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	d6 := Day6{}

	g := d6.parseInput(input)
	calculatedValue := d6.Part1(g)
	expectedValue := 41

	if calculatedValue != expectedValue {
		t.Errorf("Day 6 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay6Part2(t *testing.T) {
}
