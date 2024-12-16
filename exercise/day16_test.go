package exercise

import (
	"testing"
)

func TestDay16Part1Simple(t *testing.T) {
	input := []string{
		"###############",
		"#S...........E#",
		"###############",
	}

	d16 := Day16{}

	maze := d16.parseInput(input)

	bestPath := d16.Part1(maze)
	expectedBestPath := 12

	if bestPath != expectedBestPath {
		t.Errorf("Day 16 - Part 1 (best path) Test:\nwant %v\ngot %v\n", expectedBestPath, bestPath)
	}
}

func TestDay16Part1(t *testing.T) {
	input := []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}

	d16 := Day16{}

	maze := d16.parseInput(input)

	bestPath := d16.Part1(maze)
	expectedBestPath := 7036

	if bestPath != expectedBestPath {
		t.Errorf("Day 16 - Part 1 (best path) Test:\nwant %v\ngot %v\n", expectedBestPath, bestPath)
	}
}

func TestDay16Part1Ex2(t *testing.T) {
	input := []string{
		"#################",
		"#...#...#...#..E#",
		"#.#.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#",
		"#.#.#.#.###.#.#.#",
		"#...#.#.#.....#.#",
		"#.#.#.#.#.#####.#",
		"#.#...#.#.#.....#",
		"#.#.#####.#.###.#",
		"#.#.#.......#...#",
		"#.#.###.#####.###",
		"#.#.#...#.....#.#",
		"#.#.#.#####.###.#",
		"#.#.#.........#.#",
		"#.#.#.#########.#",
		"#S#.............#",
		"#################",
	}

	d16 := Day16{}

	maze := d16.parseInput(input)

	bestPath := d16.Part1(maze)
	expectedBestPath := 11048

	if bestPath != expectedBestPath {
		t.Errorf("Day 16 - Part 1 (best path) Test:\nwant %v\ngot %v\n", expectedBestPath, bestPath)
	}
}

func TestDay16Part2(t *testing.T) {
}
