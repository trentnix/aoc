// day22.go is the implementation for the twenty-second day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"math"
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
	maxBananas := d.Part2(input, 2000)
	w.Write([]byte(fmt.Sprintf("Day 22 - Part 2 - The maximum number of bananas we can get from the specified buyers is %d.\n", maxBananas)))
}

// Part1 calculates the secret numbers after 2000 generations
func (d *Day22) Part1(secretNumbers []string, numGenerations int) int {
	sumSecretNumbers := 0
	for _, secretNumber := range secretNumbers {
		secret, _ := strconv.Atoi(secretNumber)
		for i := 0; i < numGenerations; i++ {
			secret = calculateSecret(secret)
		}

		sumSecretNumbers += secret
	}

	return sumSecretNumbers
}

// Part2 determines the maximum number of bananas we can retrieve from providing a sequence
// of changes to the monkey sellers
func (d *Day22) Part2(secretNumbers []string, numGenerations int) int {
	// stores the bananas that would result from the sequence of changes (the key)
	sequencesDiscovered := make(map[[4]int]int)

	for _, secretNumber := range secretNumbers {
		secret, _ := strconv.Atoi(secretNumber)

		// the changes that correspond to the current secret value
		changes := make(map[[4]int]bool)

		// the set of changes we'll use as a key to our changes and sequences map
		//  the value is initialized to a nonsensical value
		last4 := [4]int{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt}

		for i := 0; i < numGenerations; i++ {
			previousPrice := secret % 10 // last digit
			secret = calculateSecret(secret)
			currentPrice := secret % 10 // last digit

			// store a new set of changes by shifting the values
			last4[0] = last4[1]
			last4[1] = last4[2]
			last4[2] = last4[3]
			last4[3] = currentPrice - previousPrice // change in price from this generation to the last

			if !changes[last4] {
				// we've now seen this sequence for the current value, so set the changes map for
				// the set of sequences to true - the monkey sells the first time it sees the sequence,
				// so we don't care if there's another instance of the sequence that provides a
				// higher value
				changes[last4] = true
				// add the price to the
				sequencesDiscovered[last4] += currentPrice
			}
		}
	}

	// find maximum number of bananas in the sequences
	maxBananas := 0
	for _, sumBananas := range sequencesDiscovered {
		if sumBananas > maxBananas {
			maxBananas = sumBananas
		}
	}

	return maxBananas
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
