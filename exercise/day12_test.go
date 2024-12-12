package exercise

import (
	"testing"
)

func TestDay12Part1Simple(t *testing.T) {
	input := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}

	d12 := Day12{}

	garden := d12.parseInput(input)

	calculatedValue := d12.Part1(garden)
	expectedValue := 140

	if calculatedValue != expectedValue {
		t.Errorf("Day 12 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay12Part1Simple1(t *testing.T) {
	input := []string{
		"OOOOO",
		"OXOXO",
		"OOOOO",
		"OXOXO",
		"OOOOO",
	}

	d12 := Day12{}

	garden := d12.parseInput(input)

	calculatedValue := d12.Part1(garden)
	expectedValue := 772

	if calculatedValue != expectedValue {
		t.Errorf("Day 12 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay12Part1(t *testing.T) {
	input := []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}

	d12 := Day12{}

	garden := d12.parseInput(input)

	calculatedValue := d12.Part1(garden)
	expectedValue := 1930

	if calculatedValue != expectedValue {
		t.Errorf("Day 12 - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestDay12Part2(t *testing.T) {
}
