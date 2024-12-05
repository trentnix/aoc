// day5.go is the implementation for the fifth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day5 represents the data necessary to process the Exercise
	Day5 struct {
		name string
		file string
	}

	orderingRule struct {
		before int
		after  int
	}

	pageNumbers []int
)

// GetName returns the name of the Day 5 exercise
func (d *Day5) GetName() string {
	return d.name
}

// Run executes the solution for Day 5 by retrieving the default file contents and uses that data
func (d *Day5) Run(w io.Writer) {
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

// RunFromInput executs the Day 5 solution using the provided input data
func (d *Day5) RunFromInput(w io.Writer, input []string) {
	rules, pages, err := d.parseInput(input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("There was an error trying to parse the input: %v.", err)))
		return
	}

	sumOfMiddleValues := d.Part1(rules, pages)

	// part 1
	w.Write([]byte(fmt.Sprintf("Day 5 - Part 1 - The sum of the middle page numbers is %d.\n", sumOfMiddleValues)))
}

// Part1 determines whether a list of pages is ordered correctly and, if so, it will
// sum and return the value of the middle elements of the lists that are ordered correctly
func (d *Day5) Part1(rules []orderingRule, pagesLists []pageNumbers) int {
	sumValidMiddleValues := 0

	rulesMap := d.getRulesMap(rules)
	for _, pages := range pagesLists {
		if pagesAreInOrder := d.pagesAreInOrder(rulesMap, pages); pagesAreInOrder {
			middleValue := pages[len(pages)/2]
			sumValidMiddleValues += middleValue
		}
	}

	return sumValidMiddleValues
}

// pagesAreInOrder determines whether, using the given rules, the pages provided are
// listed in correct order
func (d *Day5) pagesAreInOrder(rulesMap map[int][]int, pages []int) bool {
	for pageIndex := 1; pageIndex < len(pages); pageIndex++ {
		page := pages[pageIndex]
		if d.hasOverlap(rulesMap[page], pages[:pageIndex]) {
			return false
		}
	}

	return true
}

// hasOverlap determines whether a value in the first integer array overlaps with a value in
// the second array. That way, we can determine whether a current page can occur after the
// preceding specified pages.
func (d *Day5) hasOverlap(a1, a2 []int) bool {
	// Use a map to store elements of the first array
	elements := make(map[int]struct{})
	for _, num := range a1 {
		elements[num] = struct{}{}
	}

	// Check if any element in the second array exists in the map
	for _, num := range a2 {
		if _, exists := elements[num]; exists {
			return true
		}
	}

	return false
}

// Part2
func (d *Day5) Part2() int {
	return 0
}

// parseInput takes the input and parses it into a slice of orderingRule and a
// slice of pageNumbers. If an error is encountered, an error is returned.
func (d *Day5) parseInput(input []string) ([]orderingRule, []pageNumbers, error) {
	var index int

	var rules []orderingRule

	// parse page ordering rules
	for _, s := range input {
		index++

		if s == "" {
			break
		}

		parts := strings.Split(s, "|")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("error parsing input - line %d: %s", index-1, s)
		}

		before, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing input (converting parts[0]) - line %d: %s - %w", index-1, s, err)
		}

		after, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing input (converting parts[1]) - line %d: %s - %w", index-1, s, err)
		}

		rules = append(rules, orderingRule{before: before, after: after})
	}

	if index >= len(input) {
		return nil, nil, fmt.Errorf("input malformed")
	}

	var pages []pageNumbers

	// parse page numbers
	for _, s := range input[index:] {
		index++

		numbers := strings.Split(s, ",")
		var pageNum pageNumbers

		for _, num := range numbers {
			parsedNum, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing input - line %d: %s - %w", index-1, s, err)
			}
			pageNum = append(pageNum, parsedNum)
		}

		if len(pageNum) == 0 {
			return nil, nil, fmt.Errorf("error parsing input - pages list empty - line %d: %s", index-1, s)
		}

		pages = append(pages, pageNum)
	}

	return rules, pages, nil
}

// getRulesMap takes the rules specified and creates a map of 'before' elements that points to
// an integer array of all of its 'after' elements
func (d *Day5) getRulesMap(rules []orderingRule) map[int][]int {
	result := make(map[int][]int)

	for _, rule := range rules {
		result[rule.before] = append(result[rule.before], rule.after)
	}

	return result
}
