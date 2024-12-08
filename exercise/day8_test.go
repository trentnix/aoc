package exercise

import (
	"testing"
)

func TestDay8Part1(t *testing.T) {
	input := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	d8 := Day8{}

	antennaMap := d8.parseInput(input)

	calculatedValue := d8.Part1(antennaMap)
	expectedValue := 14

	if calculatedValue != expectedValue {
		t.Errorf("Day 8 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay8Part2(t *testing.T) {
	input := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	d8 := Day8{}

	antennaMap := d8.parseInput(input)

	calculatedValue := d8.Part2(antennaMap)
	expectedValue := 34

	if calculatedValue != expectedValue {
		t.Errorf("Day 8 - Part 2 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}
