package exercise

import (
	"testing"
)

func TestDay15Part1Smaller(t *testing.T) {
	input := []string{
		"########",
		"#..O.O.#",
		"##@.O..#",
		"#...O..#",
		"#.#.O..#",
		"#...O..#",
		"#......#",
		"########",
		"",
		"<^^>>>vv<v>>v<<",
	}

	d15 := Day15{}

	boxMap, instructions := d15.parseInput(input)

	sumBoxCoordinates := d15.Part1(boxMap, instructions)
	expectedSumBoxCoordinates := 2028

	if sumBoxCoordinates != expectedSumBoxCoordinates {
		t.Errorf("Day 15 - Part 1 (sum GPS coordinates) Test:\nwant %v\ngot %v\n", expectedSumBoxCoordinates, sumBoxCoordinates)
	}
}

func TestDay15Part1(t *testing.T) {
	input := []string{
		"##########",
		"#..O..O.O#",
		"#......O.#",
		"#.OO..O.O#",
		"#..O@..O.#",
		"#O#..O...#",
		"#O..O..O.#",
		"#.OO.O.OO#",
		"#....O...#",
		"##########",
		"",
		"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
		"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
		"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
		"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
		"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
		"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
		">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
		"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
		"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
		"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	}

	d15 := Day15{}

	boxMap, instructions := d15.parseInput(input)

	sumBoxCoordinates := d15.Part1(boxMap, instructions)
	expectedSumBoxCoordinates := 10092

	if sumBoxCoordinates != expectedSumBoxCoordinates {
		t.Errorf("Day 15 - Part 1 (sum GPS coordinates) Test:\nwant %v\ngot %v\n", expectedSumBoxCoordinates, sumBoxCoordinates)
	}
}

func TestDay15Part2Smaller(t *testing.T) {
	input := []string{
		"#######",
		"#...#.#",
		"#.....#",
		"#..OO@#",
		"#..O..#",
		"#.....#",
		"#######",
		"",
		"<vv<<^^<<^^",
	}

	d15 := Day15{}

	boxMap, instructions := d15.parseInputPart2(input)

	sumBoxCoordinates := d15.Part2(boxMap, instructions)
	expectedSumBoxCoordinates := 618

	if sumBoxCoordinates != expectedSumBoxCoordinates {
		t.Errorf("Day 15 - Part 1 (sum GPS coordinates) Test:\nwant %v\ngot %v\n", expectedSumBoxCoordinates, sumBoxCoordinates)
	}
}

func TestDay15Part2(t *testing.T) {
	input := []string{
		"##########",
		"#..O..O.O#",
		"#......O.#",
		"#.OO..O.O#",
		"#..O@..O.#",
		"#O#..O...#",
		"#O..O..O.#",
		"#.OO.O.OO#",
		"#....O...#",
		"##########",
		"",
		"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
		"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
		"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
		"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
		"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
		"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
		">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
		"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
		"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
		"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	}

	d15 := Day15{}

	boxMap, instructions := d15.parseInputPart2(input)

	sumBoxCoordinates := d15.Part2(boxMap, instructions)
	expectedSumBoxCoordinates := 9021

	if sumBoxCoordinates != expectedSumBoxCoordinates {
		t.Errorf("Day 15 - Part 1 (sum GPS coordinates) Test:\nwant %v\ngot %v\n", expectedSumBoxCoordinates, sumBoxCoordinates)
	}
}
