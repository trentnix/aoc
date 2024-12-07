package exercise

import (
	"testing"
)

func TestDay7Part1(t *testing.T) {
	input := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	d7 := Day7{}

	equations, err := d7.parseInput(input)
	if err != nil {
		t.Fatalf("Day 7 - Part 1 - Unable to parse the input")
	}

	calculatedValue := d7.Part1(equations)
	expectedValue := uint64(3749)

	if calculatedValue != expectedValue {
		t.Errorf("Day 7 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay7Part2(t *testing.T) {
}
