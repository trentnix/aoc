package exercise

import (
	"testing"
)

func TestDay9Part1(t *testing.T) {
	input := "12345"

	d9 := Day9{}

	diskMap := d9.parseInput(input)

	calculatedValue := d9.Part1(diskMap)
	expectedValue := int64(60)

	if calculatedValue != expectedValue {
		t.Errorf("Day 9 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}

	input2 := "2333133121414131402"

	diskMap2 := d9.parseInput(input2)

	calculatedValue = d9.Part1(diskMap2)
	expectedValue = int64(1928)

	if calculatedValue != expectedValue {
		t.Errorf("Day 9 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay9Part2(t *testing.T) {
}
