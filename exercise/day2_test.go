package exercise

import (
	"testing"
)

func TestDay2CountSafeReports(t *testing.T) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	d2 := Day2{}
	reports, err := d2.parseIntoReports(input)
	if err != nil {
		t.Errorf("There was an error parsing the input")
		return
	}

	expectedValue := 2
	calculatedValue := d2.Part1(reports)

	if calculatedValue != expectedValue {
		t.Errorf("Day One - Part 1 Test:\nwant %v\ngot %v\n", expectedValue, calculatedValue)
	}
}
