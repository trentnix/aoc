package exercise

import (
	"testing"
)

func TestDay21Part1(t *testing.T) {
	input := []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}

	d21 := Day21{}

	sumCodeComplexity := d21.CalculateComplexity(input, 2)
	expectedSumCodeComplexity := 126384

	if sumCodeComplexity != expectedSumCodeComplexity {
		t.Errorf("Day 21 - Part 1 (sum code complexity) Test:\nwant %v\ngot %v\n", expectedSumCodeComplexity, sumCodeComplexity)
	}
}

func TestDay21Part1Solution(t *testing.T) {
	input := []string{
		"279A",
		"286A",
		"508A",
		"463A",
		"246A",
	}

	d21 := Day21{}

	sumCodeComplexity := d21.CalculateComplexity(input, 2)
	expectedSumCodeComplexity := 125742

	if sumCodeComplexity != expectedSumCodeComplexity {
		t.Errorf("Day 21 - Part 1 (sum code complexity) Test:\nwant %v\ngot %v\n", expectedSumCodeComplexity, sumCodeComplexity)
	}
}

func TestDay21Part2Solution(t *testing.T) {
	input := []string{
		"279A",
		"286A",
		"508A",
		"463A",
		"246A",
	}

	d21 := Day21{}

	sumCodeComplexity := d21.CalculateComplexity(input, 25)
	expectedSumCodeComplexity := 157055032722640

	if sumCodeComplexity != expectedSumCodeComplexity {
		t.Errorf("Day 21 - Part 1 (sum code complexity) Test:\nwant %v\ngot %v\n", expectedSumCodeComplexity, sumCodeComplexity)
	}
}
