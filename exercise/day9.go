// day9.go is the implementation for the ninth day of the Advent of Code 2024
package exercise

import (
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/trentnix/aoc2024/fileprocessing"
)

type (
	// Day9 represents the data necessary to process the Exercise
	Day9 struct {
		name string
		file string
	}

	DiskMap      []DiskMapBlock
	DiskMapBlock struct {
		Index           int
		FileLength      int
		FreeSpaceLength int
		StartIndex      int
	}

	DiskData  []DiskBlock
	DiskBlock struct {
		Id       int
		HasValue bool
	}
)

// GetName returns the name of the Day 9 exercise
func (d *Day9) GetName() string {
	return d.name
}

// Run executes the solution for Day 9 by retrieving the default file contents and uses that data
func (d *Day9) Run(w io.Writer) {
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

// RunFromInput executs the Day 9 solution using the provided input data
func (d *Day9) RunFromInput(w io.Writer, input []string) {
	if len(input) != 1 {
		log.Fatalf("the input was invalid")
		return
	}

	diskMap := d.parseInput(input[0])

	inputChecksum := d.Part1(diskMap)
	w.Write([]byte(fmt.Sprintf("Day 9 - Part 1 - The checksum of the input is %d.\n", inputChecksum)))

	inputChecksum = d.Part2(diskMap)
	w.Write([]byte(fmt.Sprintf("Day 9 - Part 2 - The checksum of the input is %d.\n", inputChecksum)))
}

// Part1 takes the specified DiskMap, creates a DiskData instance, compresses it, and
// calculates the checksum of the DiskData instance
func (d *Day9) Part1(m DiskMap) int64 {
	diskData := NewDiskData(m)
	diskData.Compress()
	return diskData.CalculateChecksum()
}

// Part2 takes the specified DiskMap, creates a DiskData instance, compresses it using
// whole file compression (according to the assignment rules), and calculates the checksum
// of the DiskData instance
func (d *Day9) Part2(m DiskMap) int64 {
	diskData := NewDiskData(m)
	diskData.CompressWholeFiles(m)
	return diskData.CalculateChecksum()
}

// parseInput processes an input string into a DiskMap instance
func (d *Day9) parseInput(input string) DiskMap {
	var diskMap DiskMap

	inputLength := len(input)
	var i int
	for i = 0; i+1 < inputLength; i = i + 2 {
		cFileLength := input[i]
		cSpaceLength := input[i+1]

		iFileLength, _ := strconv.Atoi(string(cFileLength))
		iSpaceLength, _ := strconv.Atoi(string(cSpaceLength))

		block := DiskMapBlock{
			Index:           i / 2,
			FileLength:      iFileLength,
			FreeSpaceLength: iSpaceLength,
		}

		diskMap = append(diskMap, block)
	}

	cFileLength := input[i]
	iFileLength, _ := strconv.Atoi(string(cFileLength))
	// need to handle the last element, which has no free block
	block := DiskMapBlock{
		Index:           i / 2,
		FileLength:      iFileLength,
		FreeSpaceLength: 0,
	}

	diskMap = append(diskMap, block)

	return diskMap
}

// NewDiskData creates a DiskData instance from the specified DiskMap
func NewDiskData(m DiskMap) DiskData {
	var diskData DiskData

	diskDataIndex := 0
	for index := 0; index < len(m); index++ {
		m[index].StartIndex = diskDataIndex
		for i := 0; i < m[index].FileLength; i++ {
			diskBlock := DiskBlock{Id: m[index].Index, HasValue: true}
			diskData = append(diskData, diskBlock)
		}

		for i := 0; i < m[index].FreeSpaceLength; i++ {
			diskBlock := DiskBlock{Id: m[index].Index, HasValue: false}
			diskData = append(diskData, diskBlock)
		}

		diskDataIndex += m[index].FileLength + m[index].FreeSpaceLength
	}

	return diskData
}

// Compress takes the specified DiskData instance and fills in the free space with the data
// at the end of the disk
func (d DiskData) Compress() {
	end := len(d) - 1

	for i := 0; i < len(d); i++ {
		// Find a block with HasValue == false
		if !d[i].HasValue {
			// Move backward from the end to find a block with HasValue == true
			for end > i && !d[end].HasValue {
				end--
			}

			// If we've found a block with HasValue == true, move it
			if end > i {
				d[i] = d[end]
				d[end] = DiskBlock{} // Reset the moved block
				end--                // Move the end pointer
			} else {
				// No more blocks with HasValue == true to move
				break
			}
		}
	}
}

// CompressWholeFiles takes the specified DiskData instance and fills in the free
// space from the beginning of the disk with entire "files" from the end of the disk,
// trying each file once from the end to the beginning
func (d DiskData) CompressWholeFiles(m DiskMap) {
	for f := len(m) - 1; f >= 0; f-- {
		file := m[f]
		fileLength := file.FileLength
		if fileLength == 0 {
			continue
		}

		originalStart := file.StartIndex
		needed := fileLength

		currentRunStart := -1
		currentRunLength := 0

		chosenRunStart := -1

		// Find a free space run that is fully to the left of originalStart
		for i := 0; i < len(d); i++ {
			if !d[i].HasValue {
				if currentRunStart == -1 {
					currentRunStart = i
				}
				currentRunLength++
				// Check if we found a sufficiently large run
				if currentRunLength >= needed {
					// Check leftward condition
					if currentRunStart+needed <= originalStart {
						chosenRunStart = currentRunStart
						break
					} else {
						// Even though large enough, not leftward.
						// Continue scanning. Reset and keep looking.
						// Move to next block after currentRunStart
						i = currentRunStart + 1
						currentRunStart = -1
						currentRunLength = 0
					}
				}
			} else {
				// Non-free block, reset the run
				currentRunStart = -1
				currentRunLength = 0
			}
		}

		if chosenRunStart == -1 {
			// No suitable run found
			continue
		}

		// Move the file
		for j := 0; j < fileLength; j++ {
			d[chosenRunStart+j] = DiskBlock{Id: file.Index, HasValue: true}
			d[originalStart+j] = DiskBlock{} // free the old location
		}

		// Update DiskMap
		file.StartIndex = chosenRunStart
		m[f] = file
	}
}

// CalculateChecksum iterates of the DiskBlock entries of the specified DiskData instance
// and calculates the checksum (according to the rules in the assignment):
//
// To calculate the checksum, add up the result of multiplying each of these blocks'
// position with the file ID number it contains. The leftmost block is in position 0.
// If a block contains free space, skip it instead.
func (d DiskData) CalculateChecksum() int64 {
	var checksum int64
	for i := 0; i < len(d); i++ {
		if !d[i].HasValue {
			continue
		}

		checksum += int64(i) * int64(d[i].Id)
	}

	return checksum
}

// Print writes the contents of specified DiskData to stdout in a similar format to what's in
// the assignemnt readme. It's helpful for debugging.
func (d DiskData) Print() {
	fmt.Printf("DiskData:\n")
	for i := 0; i < len(d); i++ {
		if !d[i].HasValue {
			fmt.Print(".")
		} else {
			if d[i].Id > 9 {
				fmt.Printf("[%d]", d[i].Id)
			} else {
				fmt.Printf("%d", d[i].Id)
			}
		}
	}

	fmt.Printf("\n")
}
