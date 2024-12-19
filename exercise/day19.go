// day19.go is the implementation for the nineteenth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day19 represents the data necessary to process the Exercise
	Day19 struct {
		name string
		file string
	}

	Towels       []string
	TowelDesigns []string
)

// GetName returns the name of the Day 19 exercise
func (d *Day19) GetName() string {
	return d.name
}

// Run executes the solution for Day 19 by retrieving the default file contents and uses that data
func (d *Day19) Run(w io.Writer) {
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

// RunFromInput executs the Day 19 solution using the provided input data
func (d *Day19) RunFromInput(w io.Writer, input []string) {
	towels, towelDesigns := d.parseInput(input)

	possibleTowelDesigns := d.Part1(towels, towelDesigns)
	w.Write([]byte(fmt.Sprintf("Day 19 - Part 1 - The number of possible towel designs: %d\n", possibleTowelDesigns)))

	sumTowelCombinationsThatSolve := d.Part2(towels, towelDesigns)
	w.Write([]byte(fmt.Sprintf("Day 19 - Part 2 - The sum of the towel combinations tha that solve the designs is: %d\n", sumTowelCombinationsThatSolve)))
}

// Part1 iterates through the various designs and determines if the specified
// towels can be used to build the design
func (d *Day19) Part1(towels Towels, towelDesigns TowelDesigns) int {
	numPossible := 0
	for _, design := range towelDesigns {
		if d.canBuildDesign(design, towels) {
			numPossible++
		}
	}

	return numPossible
}

// canBuildTowel determines whether the specified towelDesign can be built
// from the set of towels.
func (d *Day19) canBuildDesign(towelDesign string, towels []string) bool {
	dict := make(map[string]bool)
	for _, d := range towels {
		dict[d] = true
	}

	n := len(towelDesign)
	// dp[i] will be true if we can form towel[:i] from the given designs
	dp := make([]bool, n+1)
	dp[0] = true // Empty string can always be formed

	for i := 1; i <= n; i++ {
		// Check all possible substrings that end at i
		for j := 0; j < i; j++ {
			// if dp[j] is true (we can form towel[:j]) and
			// towel[j:i] is in the dictionary, then dp[i] = true.
			if dp[j] && dict[towelDesign[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

// Part2 determines the sum of the possible ways to solve for each towel design
func (d *Day19) Part2(towels Towels, towelDesigns TowelDesigns) int {
	sumTowelCombinationsThatSolve := 0
	for _, design := range towelDesigns {
		sumTowelCombinationsThatSolve += d.countWaysToBuildDesign(design, towels)
	}

	return sumTowelCombinationsThatSolve
}

// countWaysToBuildTowel counts the number of sets of towels that can be
// used to solve the specified towelDesign
func (d *Day19) countWaysToBuildDesign(towelDesign string, towels []string) int {
	// Convert towelDesigns into a set for quick lookups
	dict := make(map[string]bool)
	for _, d := range towels {
		dict[d] = true
	}

	n := len(towelDesign)
	dp := make([]int, n+1)
	dp[0] = 1 // There's one way to form the empty substring: do nothing

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// If towel[j:i] is a valid segment
			if dict[towelDesign[j:i]] {
				// Add all ways to form towel[:j] to dp[i]
				dp[i] += dp[j]
			}
		}
	}

	return dp[n]
}

// parseInput takes the assignment's specified input and parses it into Towels and
// TowelDesigns structures
func (d *Day19) parseInput(input []string) (Towels, TowelDesigns) {
	if len(input) == 0 {
		return nil, nil
	}

	// first line contains the towels, separated by commas.
	towelsLine := input[0]
	towels := strings.Split(towelsLine, ",")
	for i, t := range towels {
		towels[i] = strings.TrimSpace(t)
	}

	// subsequent lines (skipping the blank line that separates towels from the various
	// designs) are towel designs
	towelDesigns := input[2:]

	return Towels(towels), TowelDesigns(towelDesigns)
}
