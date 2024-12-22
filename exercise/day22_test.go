package exercise

import (
	"testing"
)

func TestDay22Part1(t *testing.T) {
	input := []string{
		"1",
		"10",
		"100",
		"2024",
	}

	d22 := Day22{}

	sumSecretNumbers := d22.Part1(input, 2000)

	expectedSumSecretNumbers := 37327623

	if sumSecretNumbers != expectedSumSecretNumbers {
		t.Errorf("Day 22 - Part 1 (sum secret numbers aftr 2000 generations) Test:\nwant %v\ngot %v\n", expectedSumSecretNumbers, sumSecretNumbers)
	}
}

func TestDay22Part2(t *testing.T) {
}
