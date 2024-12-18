package exercise

import (
	"testing"
)

func TestDay18Part1(t *testing.T) {
	input := []string{
		"5,4",
		"4,2",
		"4,5",
		"3,0",
		"2,1",
		"6,3",
		"2,4",
		"1,5",
		"0,6",
		"3,3",
		"2,6",
		"5,1",
		"1,2",
		"5,5",
		"2,5",
		"6,5",
		"1,4",
		"0,4",
		"6,4",
		"1,1",
		"6,1",
		"1,0",
		"0,5",
		"1,6",
		"2,0",
	}

	d18 := Day18{}

	fallingBlocks := d18.parseInput(input)

	gridSize := 7
	startTick := 12
	steps := d18.Part1(fallingBlocks, gridSize, startTick)
	expectedSteps := 22

	if steps != expectedSteps {
		t.Errorf("Day 18 - Part 1 (shortest path) Test:\nwant %v\ngot %v\n", expectedSteps, steps)
	}
}

func TestDay18Part2(t *testing.T) {
	input := []string{
		"5,4",
		"4,2",
		"4,5",
		"3,0",
		"2,1",
		"6,3",
		"2,4",
		"1,5",
		"0,6",
		"3,3",
		"2,6",
		"5,1",
		"1,2",
		"5,5",
		"2,5",
		"6,5",
		"1,4",
		"0,4",
		"6,4",
		"1,1",
		"6,1",
		"1,0",
		"0,5",
		"1,6",
		"2,0",
	}

	d18 := Day18{}

	fallingBlocks := d18.parseInput(input)

	gridSize := 7
	startTick := 12
	y, x := d18.Part2(fallingBlocks, gridSize, startTick)
	expectedY, expectedX := 6, 1

	if y != expectedY || x != expectedX {
		t.Errorf("Day 18 - Part 2 (block that breaks the graph) Test:\nwant %d, %d\ngot %d, %d\n", y, x, expectedY, expectedX)
	}
}
