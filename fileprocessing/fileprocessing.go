// fileprocessing.go provides interaction with the input file
package fileprocessing

import (
	"bufio"
	"os"
)

// Readfile returns the contents of the specified filename if no error is encountered
func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
