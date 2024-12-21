package exercise

import (
	"testing"
)

func TestDay20Part1(t *testing.T) {
	input := []string{
		"###############",
		"#...#...#.....#",
		"#.#.#.#.#.###.#",
		"#S#...#.#.#...#",
		"#######.#.#.###",
		"#######.#.#...#",
		"#######.#.###.#",
		"###..E#...#...#",
		"###.#######.###",
		"#...###...#...#",
		"#.#####.#.###.#",
		"#.#...#.#.#...#",
		"#.#.#.#.#.#.###",
		"#...#...#...###",
		"###############",
	}

	d20 := Day20{}

	raceTrack := d20.parseInput(input)

	numCheats := d20.Part1(raceTrack, 20)
	expectedNumcheats := 5

	if numCheats != expectedNumcheats {
		t.Errorf("Day 20 - Part 1 (num cheats with minimum 20) Test:\nwant %v\ngot %v\n", expectedNumcheats, numCheats)
	}
}

func TestDay20Part2(t *testing.T) {
	input := []string{
		"###############",
		"#...#...#.....#",
		"#.#.#.#.#.###.#",
		"#S#...#.#.#...#",
		"#######.#.#.###",
		"#######.#.#...#",
		"#######.#.###.#",
		"###..E#...#...#",
		"###.#######.###",
		"#...###...#...#",
		"#.#####.#.###.#",
		"#.#...#.#.#...#",
		"#.#.#.#.#.#.###",
		"#...#...#...###",
		"###############",
	}

	d20 := Day20{}

	raceTrack := d20.parseInput(input)

	numCheats := d20.Part2Manhattan(raceTrack, 70)
	expectedNumcheats := 41

	if numCheats != expectedNumcheats {
		t.Errorf("Day 20 - Part 2 (num cheats with minimum 70) Test:\nwant %v\ngot %v\n", expectedNumcheats, numCheats)
	}
}
