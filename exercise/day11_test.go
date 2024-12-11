package exercise

import (
	"testing"
)

func TestBlinkPart1(t *testing.T) {
	input := "0 1 10 99 999"
	d11 := Day11{}
	stones := d11.blink(d11.parseInput(input))
	output := d11.parseInput("1 2024 1 0 9 9 2021976")

	if compareStringSlices(stones, output) {
		t.Errorf("Day 11 - Blink Test:\nwant %v\ngot %v\n", output, stones)
	}
}

func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestDay11Part1(t *testing.T) {
	input := "125 17"

	d11 := Day11{}

	stones := d11.parseInput(input)

	blinks := 6
	calculatedValue := d11.Part1(stones, blinks)
	expectedValue := 22

	if calculatedValue != expectedValue {
		t.Errorf("Day 11 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}

	blinks = 25
	calculatedValue = d11.Part1(stones, blinks)
	expectedValue = 55312

	if calculatedValue != expectedValue {
		t.Errorf("Day 11 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay11Part2(t *testing.T) {
}
