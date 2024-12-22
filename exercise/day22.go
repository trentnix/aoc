// day22.go is the implementation for the twenty-second day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strconv"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day22 represents the data necessary to process the Exercise
	Day22 struct {
		name string
		file string
	}
)

// GetName returns the name of the Day 22 exercise
func (d *Day22) GetName() string {
	return d.name
}

// Run executes the solution for Day 22 by retrieving the default file contents and uses that data
func (d *Day22) Run(w io.Writer) {
	if d.file == "" {
		w.Write([]byte(fmt.Sprintf("A default input file is not specified.")))
		return
	}

	input, err := fileprocessing.ReadFile(d.file)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to read the input file %s: %v.", d.file, err)))
		return
	}

	d.RunFromInput(w, input)
}

// RunFromInput executs the Day 22 solution using the provided input data
func (d *Day22) RunFromInput(w io.Writer, input []string) {
	// part 1
	sumSecretNumbers := d.Part1(input, 2000)
	w.Write([]byte(fmt.Sprintf("Day 22 - Part 1 - The sum of the secret numbers after 2000 generations is %d.\n", sumSecretNumbers)))

	// part 2
	// sumCodeComplexity = d.CalculateComplexity(input, 25)
	// w.Write([]byte(fmt.Sprintf("Day 21 - Part 2 - The sum of the complexity of the provided codes (25 robots) is %d.\n", sumCodeComplexity)))
}

// Part1 calculates the secret numbers after 2000 generations
func (d *Day22) Part1(secretNumbers []string, numGenerated int) int {
	sumSecretNumbers := 0
	for _, secretNumber := range secretNumbers {
		secret, _ := strconv.Atoi(secretNumber)
		for i := 0; i < numGenerated; i++ {
			secret = calculateSecret(secret)
		}

		sumSecretNumbers += secret
	}

	return sumSecretNumbers
}

// Part2
func (d *Day22) Part2() int {
	return 0
}

// calculateSecret calculates the secret number according to the rules specified in the assignment:
//   - Calculate the result of multiplying the secret number by 64. Then, mix this result into the
//     secret number. Finally, prune the secret number.
//   - Calculate the result of dividing the secret number by 32. Round the result down to the nearest
//     integer. Then, mix this result into the secret number. Finally, prune the secret number.
//   - Calculate the result of multiplying the secret number by 2048. Then, mix this result into the
//     secret number. Finally, prune the secret number.
func calculateSecret(secret int) int {
	secret = prune(mix(secret, secret<<6))
	secret = prune(mix(secret, secret>>5))
	secret = prune(mix(secret, secret<<11))

	return secret
}

// mix returns the bitwise XOR of a given value and the secret number
func mix(secret int, val int) int {
	return secret ^ val
}

// prune returns the modulus of the secret number and 16777216
func prune(secret int) int {
	return secret % 16777216
}
