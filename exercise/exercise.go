// exercise.go defines the Exercise interface and initializes the exercises array
package exercise

import "io"

// the exercises array contains the implementations of each Advent of Code day's exercise
var exercises []Exercise

// init initializes the exercises array
func init() {
	day1 := &Day1{
		name:  "2024: Day 1",
		file:  "data/day1/input.txt",
		order: 202401,
	}

	RegisterExercise(day1)
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
