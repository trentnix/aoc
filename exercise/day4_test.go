package exercise

import (
	"testing"
)

func TestDay4Part1(t *testing.T) {
	input := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	grid := make([][]rune, len(input))

	for i, str := range input {
		// Convert each string into a slice of runes
		grid[i] = []rune(str)
	}

	expectedValue := 18

	d4 := Day4{}
	calculatedValue := d4.Part1(grid)

	if calculatedValue != expectedValue {
		t.Errorf("Day 2 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay4Part2(t *testing.T) {
}
