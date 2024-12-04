package exercise

import (
	"testing"
)

func TestDay4XmasCount(t *testing.T) {
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
		t.Errorf("Day 4 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay4MasXCount(t *testing.T) {
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

	expectedValue := 9

	d4 := Day4{}
	calculatedValue := d4.Part2(grid)

	if calculatedValue != expectedValue {
		t.Errorf("Day 4 - Part 2 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}
