package exercise

import (
	"testing"
)

func TestDay13Part1(t *testing.T) {
	input := []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
	}

	d13 := Day13{}

	games := d13.parseInput(input)

	prizesWon, minimumTokensSpent := d13.Part1(games)
	expectedPrizesWon := 2
	expectedTokensSpent := 480

	if prizesWon != expectedPrizesWon {
		t.Errorf("Day 13 - Part 1 (prizes) Test:\nwant %v\ngot %v\n", expectedPrizesWon, prizesWon)
	}
	if minimumTokensSpent != expectedTokensSpent {
		t.Errorf("Day 13 - Part 1 (tokens) Test:\nwant %v\ngot %v\n", expectedTokensSpent, minimumTokensSpent)
	}
}

func TestDay13Part2(t *testing.T) {
}
