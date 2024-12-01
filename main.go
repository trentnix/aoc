// main.go is the entry point for each of the solutions in the 2025 Advent of Code! I've
// done a few other Advent of Code exercises over the years but have never completed an
// entire year. We'll see if this year's the year that I finally break through.
//
// Each day will be displayed in a command-line menu that allows a user to specify the
// day to run and the data file input.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/exercise"
	"github.com/trentnix/aoc2024/fileprocessing"
)

// main runs the exercise for the specified day. If no day is specified, a
// menu is displayed that allows the user to select a day to run. If the user
// selects a day to run, the default input file is used.
//
// If the exercise is specified, an optional input file can also be specified. Otherwise,
// the default input file is used.
func main() {
	choice := -1
	inputFile := ""

	var writer io.Writer
	writer = os.Stdout

	exercises := exercise.GetExercises()
	numExercises := len(exercises)

	if numExercises <= 0 {
		log.Fatalf("There are no exercises available.")
	}

	// check for a command-line argument with a preselection
	argCount := len(os.Args)
	if argCount > 1 {
		selectionNum, err := strconv.Atoi(os.Args[1])
		if err != nil || numExercises > selectionNum {
			log.Fatalf("Invalid exercise")
			return
		} else {
			choice = selectionNum
		}

		if argCount > 2 {
			inputFile = os.Args[2]
		}

		if inputFile == "" {
			exercises[selectionNum].Run(writer)
		} else {
			input, err := fileprocessing.ReadFile(inputFile)
			if err != nil {
				log.Fatalf("Could not process the specified file %s: %v", inputFile, err)
				return
			}

			exercises[selectionNum].RunFromInput(writer, input)
		}

	} else {
		// there is no argument so show the full menu
		fmt.Print("\n")
		fmt.Println("Welcome to solutions for the Advent of Code 2023!")
	}

	for {
		if choice < 0 {
			// no command-line menu choice was selected or the menu needs to be re-displayed
			// 'selection' captures the user's selection for processing
			selection := menu(exercises)
			selectionNum, err := strconv.Atoi(selection)
			if err != nil {
				fmt.Print("invalid choice:", err)
				continue
			} else {
				choice = selectionNum
			}
		}

		fmt.Print("\n")

		if 0 < choice && choice <= numExercises { // the choice is valid
			exercises[choice-1].Run(writer)
		} else if choice == 0 {
			// the user has specified the 'Exit' choice
			fmt.Println("Exiting...")
			fmt.Print("\n")
			return
		} else {
			// the choice isn't valid
			fmt.Println("Invalid choice. Please try again.")
		}

		// We are looping through, so any command-line choice should be rendered moot.
		// Resetting choice does just that.
		choice = -1
	}
}

// menu takes an array of exercises and builds a command-line menu to present to the
// user. The function returns (as a string value) the selection made by the user.
func menu(exercises []exercise.Exercise) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n\n")
	fmt.Println("Pick an option below:")
	for index, ex := range exercises {
		fmt.Println((index + 1), ":", ex.GetName())
	}
	if len(exercises) <= 0 {
		fmt.Println("No exercises available")
	}
	fmt.Println("0 : Exit")
	fmt.Print("\nChoose wisely: ")

	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(input)

	return choice
}
