package exercise

import (
	"testing"
)

func TestDay11(t *testing.T) {
	input := "125 17"

	d11 := Day11{}

	stones := d11.parseInput(input)

	blinks := 6
	calculatedValue := d11.ProcessStones(stones, blinks)
	expectedValue := uint64(22)

	if calculatedValue != expectedValue {
		t.Errorf("Day 11 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}

	blinks = 25
	calculatedValue = d11.ProcessStones(stones, blinks)
	expectedValue = uint64(55312)

	if calculatedValue != expectedValue {
		t.Errorf("Day 11 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}
