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
		xPrizeLocation int
		yPrizeLocation int
	}

	ClawGameButton struct {
		price        int
		xMove, yMove int
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
	clawGames := d.parseInput(input)

	// part 1
	numPrizesWon, numTokensSpent := d.Part1(clawGames)
	w.Write([]byte(fmt.Sprintf("Day 13 - Part 1 - prizes won: %d, tokens spent: %d\n", numPrizesWon, numTokensSpent)))
}

// Part1 determines which games are winnable given the input and, for each game,
// what the minimum cost would be to win the game. The number of winnable games and
// the sum of the minimum costs to win those games is returned
func (d *Day13) Part1(games []ClawGame) (prizes, sumTokensSpent int) {
	prizes = 0
	sumTokensSpent = 0

	for _, game := range games {
		x1 := game.A.xMove
		x2 := game.B.xMove
		x3 := game.xPrizeLocation

		var solves [][2]int

		abXSolves := d.solveEquation(x1, x2, x3)
		for _, solve := range abXSolves {
			a := solve[0]
			b := solve[1]

			yVal := a*game.A.yMove + b*game.B.yMove
			if yVal == game.yPrizeLocation {
				solves = append(solves, solve)
			}
		}

		tokensSpent := 0
		if len(solves) > 0 {
			prizes++

			for _, solve := range solves {
				gameCost := solve[0]*game.A.price + solve[1]*game.B.price
				// find cheapest option
				if tokensSpent == 0 || gameCost < tokensSpent {
					tokensSpent = gameCost
				}
			}

			sumTokensSpent += tokensSpent
		}
	}

	return prizes, sumTokensSpent
}

// Part2
func (d *Day13) Part2() int {
	return 0
}

const (
	priceA = 3
	priceB = 1
)

// parseInput parses the specified input into a slice of ClawGame instances
func (d *Day13) parseInput(input []string) []ClawGame {
	var games []ClawGame

	// Process every 3 lines, skipping the blank line
	for i := 0; i < len(input); i += 4 {
		// button A
		var aX, aY int
		fmt.Sscanf(input[i], "Button A: X+%d, Y+%d", &aX, &aY)

		// button B
		var bX, bY int
		fmt.Sscanf(input[i+1], "Button B: X+%d, Y+%d", &bX, &bY)

		// prize location
		var pX, pY int
		fmt.Sscanf(input[i+2], "Prize: X=%d, Y=%d", &pX, &pY)

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

// solveEquation determines all of the a,b pairs that solve for the specified
// input. x1 is the "A" moves, x2 is the "B" moves, and x3 is the prize location
// for the current axis
//
// x1 * a + x2 * b = x3 - solving for (a, b)
//
// b = (x3 - x1 * a) / x2
func (d *Day13) solveEquation(x1, x2, x3 int) [][2]int {
	var results [][2]int

	// Iterate over possible values of 'a'
	for a := 1; x1*a < x3; a++ {
		// Calculate b
		remainder := x3 - x1*a
		if remainder > 0 && remainder%x2 == 0 {
			b := remainder / x2
			if b > 0 {
				results = append(results, [2]int{a, b})
			}
		}
	}

	return results
}
