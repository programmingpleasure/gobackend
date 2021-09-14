package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// TODO: why this function is very bad?
// BAD, BAD EXAMPLE of how to kill the app
// with OutOfMemory error
func readFile(filePath string) ([]string, error) {
	// TODO: add test
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	flag.Parse()
	var result []string
	for scanner.Scan() {
		// TODO: handle the input here
		// to not collect the file content in memory
		// (it's a bad way, heh)
		result = append(result, scanner.Text())
	}

	return result, nil
}
