// day13.go is the implementation for the thirteenth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day13 represents the data necessary to process the Exercise
	Day13 struct {
		name string
		file string
	}

	ClawGame struct {
		A              ClawGameButton
		B              ClawGameButton
		xPrizeLocation int64
		yPrizeLocation int64
	}

	ClawGameButton struct {
		price        int
		xMove, yMove int64
	}
)

// GetName returns the name of the Day 13 exercise
func (d *Day13) GetName() string {
	return d.name
}

// Run executes the solution for Day 13 by retrieving the default file contents and uses that data
func (d *Day13) Run(w io.Writer) {
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

// RunFromInput executs the Day 13 solution using the provided input data
func (d *Day13) RunFromInput(w io.Writer, input []string) {
	var numPrizesWon, sumTokensSpent int64

	// part 1
	clawGames := d.parseInput(input, false)
	numPrizesWon, sumTokensSpent = d.Solve(clawGames)
	w.Write([]byte(fmt.Sprintf("Day 13 - Part 1 - prizes won: %d, tokens spent: %d\n", numPrizesWon, sumTokensSpent)))

	// part 2
	clawGames = d.parseInput(input, true)
	numPrizesWon, sumTokensSpent = d.Solve(clawGames)
	w.Write([]byte(fmt.Sprintf("Day 13 - Part 2 - prizes won: %d, tokens spent: %d\n", numPrizesWon, sumTokensSpent)))
}

// Solve uses Cramer's rule to identify the solve for each specified ClawGame instance
func (d *Day13) Solve(games []ClawGame) (prizes, sumTokensSpent int64) {
	for _, game := range games {
		a, b := d.solveUsingCramersRule(game)
		if a != 0 && b != 0 {
			sumTokensSpent += a*int64(game.A.price) + b*int64(game.B.price)
			prizes++
		}
	}

	return prizes, sumTokensSpent
}

const (
	priceA = 3
	priceB = 1
)

// parseInput parses the specified input into a slice of ClawGame instances
func (d *Day13) parseInput(input []string, p2 bool) []ClawGame {
	var games []ClawGame

	// Process every 3 lines, skipping the blank line
	for i := 0; i < len(input); i += 4 {
		// button A
		var aX, aY int64
		fmt.Sscanf(input[i], "Button A: X+%d, Y+%d", &aX, &aY)

		// button B
		var bX, bY int64
		fmt.Sscanf(input[i+1], "Button B: X+%d, Y+%d", &bX, &bY)

		// prize location
		var pX, pY int64
		fmt.Sscanf(input[i+2], "Prize: X=%d, Y=%d", &pX, &pY)

		if p2 {
			pX += 10000000000000
			pY += 10000000000000
		}

		game := ClawGame{
			A:              ClawGameButton{xMove: aX, yMove: aY, price: priceA},
			B:              ClawGameButton{xMove: bX, yMove: bY, price: priceB},
			xPrizeLocation: pX,
			yPrizeLocation: pY,
		}
		games = append(games, game)
	}

	return games
}

// solveUsingCramersRule solves for a game using Cramer's rule:
//
// A = (p_x*b_y - prize_y*b_x) / (a_x*b_y - a_y*b_x)
// B = (a_x*p_y - a_y*p_x) / (a_x*b_y - a_y*b_x)
//
// If integer values aren't calculated (whole numbers), then (0,0) is returned
func (d *Day13) solveUsingCramersRule(g ClawGame) (int64, int64) {
	pX := g.xPrizeLocation
	pY := g.yPrizeLocation

	var denominator, numeratorA, numeratorB int64
	denominator = g.A.xMove*g.B.yMove - g.A.yMove*g.B.xMove
	if denominator == 0 {
		// If the denominator is 0, the system is either inconsistent or dependent
		return 0, 0
	}

	numeratorA = pX*g.B.yMove - pY*g.B.xMove
	numeratorB = g.A.xMove*pY - g.A.yMove*pX

	if numeratorA%denominator != 0 {
		// a is not an integer
		return 0, 0
	}
	a := numeratorA / denominator

	if numeratorB%denominator != 0 {
		// b is not an integer
		return 0, 0
	}
	b := numeratorB / denominator

	if a < 0 || b < 0 {
		return 0, 0
	}

	// (a, b) solves for the system
	return int64(a), int64(b)
}
