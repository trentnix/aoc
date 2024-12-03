package exercise

import (
	"testing"
)

func TestDay3AddMultiplicationResults(t *testing.T) {
	input := []string{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	}

	d3 := Day3{}
	data, err := d3.parseInput(input)
	if err != nil {
		t.Errorf("There was an error parsing the input")
		return
	}

	expectedValue := 161
	calculatedValue := d3.Part1(data)

	if calculatedValue != expectedValue {
		t.Errorf("Day 2 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay3P2(t *testing.T) {
	input := []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}

	d3 := Day3{}
	data, err := d3.parseInput(input)
	if err != nil {
		t.Errorf("There was an error parsing the input")
		return
	}

	expectedValue := 48
	calculatedValue := d3.Part2(data)

	if calculatedValue != expectedValue {
		t.Errorf("Day 2 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}
