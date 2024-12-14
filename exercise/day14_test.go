package exercise

import (
	"testing"
)

func TestDay14Part1(t *testing.T) {
	input := []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}

	d14 := Day14{}

	var seconds, gridX, gridY int
	seconds = 100
	gridX = 11
	gridY = 7

	robots := d14.parseInput(input)

	safetyFactor := d14.Part1(robots, seconds, gridX, gridY)
	expectedSafetyFactor := 12

	if safetyFactor != expectedSafetyFactor {
		t.Errorf("Day 14 - Part 1 (safety factor) Test:\nwant %v\ngot %v\n", expectedSafetyFactor, safetyFactor)
	}
}

func TestDay14Part2(t *testing.T) {
}
