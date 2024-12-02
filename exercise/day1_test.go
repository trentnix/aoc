package exercise

import (
	"testing"
)

func TestSumSmallestDistances(t *testing.T) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	d1 := Day1{}
	left, right, err := d1.parseIntoLists(input)
	if err != nil {
		t.Errorf("There was an error parsing the input")
		return
	}

	expectedValue := 11
	calculatedValue := d1.Part1(left, right)

	if calculatedValue != expectedValue {
		t.Errorf("Day One - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}

func TestSumSimilarityScores(t *testing.T) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	d1 := Day1{}
	left, right, err := d1.parseIntoLists(input)
	if err != nil {
		t.Errorf("There was an error parsing the input")
		return
	}

	expectedValue := 31
	calculatedValue := d1.Part2(left, right)

	if calculatedValue != expectedValue {
		t.Errorf("Day One - Part 2 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}
