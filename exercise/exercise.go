// exercise.go defines the Exercise interface and initializes the exercises array
package exercise

import "io"

// the exercises array contains the implementations of each Advent of Code day's exercise
var exercises []Exercise

// init initializes the exercises array
func init() {
	day1 := &Day1{
		name: "2024: Day 1",
		file: "data/day1/input.txt",
	}

	RegisterExercise(day1)

	day2 := &Day2{
		name: "2024: Day 2",
		file: "data/day2/input.txt",
	}

	RegisterExercise(day2)

	day3 := &Day3{
		name: "2024: Day 3",
		file: "data/day3/input.txt",
	}

	RegisterExercise(day3)

	day4 := &Day4{
		name: "2024: Day 4",
		file: "data/day4/input.txt",
	}

	RegisterExercise(day4)

	day5 := &Day5{
		name: "2024: Day 5",
		file: "data/day5/input.txt",
	}

	RegisterExercise(day5)

	day6 := &Day6{
		name: "2024: Day 6",
		file: "data/day6/input.txt",
	}

	RegisterExercise(day6)

	day7 := &Day7{
		name: "2024: Day 7",
		file: "data/day7/input.txt",
	}

	RegisterExercise(day7)

	day8 := &Day8{
		name: "2024: Day 8",
		file: "data/day8/input.txt",
	}

	RegisterExercise(day8)

	day9 := &Day9{
		name: "2024: Day 9",
		file: "data/day9/input.txt",
	}

	RegisterExercise(day9)

	day10 := &Day10{
		name: "2024: Day 10",
		file: "data/day10/input.txt",
	}

	RegisterExercise(day10)

	day11 := &Day11{
		name: "2024: Day 11",
		file: "data/day11/input.txt",
	}

	RegisterExercise(day11)

	day12 := &Day12{
		name: "2024: Day 12",
		file: "data/day12/input.txt",
	}

	RegisterExercise(day12)

	day13 := &Day13{
		name: "2024: Day 13",
		file: "data/day13/input.txt",
	}

	RegisterExercise(day13)
}

// RegisterExercise provides a way for an Exercise to register itself
func RegisterExercise(e Exercise) {
	exercises = append(exercises, e)
}

// GetExercises returns the exercises slice
func GetExercises() []Exercise {
	return exercises
}

// defines the Exercise interface
type Exercise interface {
	GetName() string
	RunFromInput(w io.Writer, input []string)
	Run(w io.Writer)
}
