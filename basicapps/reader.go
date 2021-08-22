package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: why this function is very bad?
func readFile(filePath string) ([]string, error) {
	// TODO: add test
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}
