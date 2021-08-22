package main

import (
	"fmt"
	"strings"
)

type (
	checkFunc func(keyWord, line string) (string, error)
)

func checkFull(content []string, keyWord string, checkFunc checkFunc, formatter formatter) ([]string, error) {
	// TODO: add test
	var detected []string
	for _, line := range content {
		res, err := checkFunc(keyWord, line)
		if err != nil {
			return nil, fmt.Errorf("error while check the line:%w: %s", err, line)
		}

		detected = append(detected, formatter(keyWord, res))
	}

	return detected, nil
}

func containsCheck(keyWord, line string) (string, error) {
	// TODO: add test
	if strings.Contains(line, keyWord) {
		return line, nil
	}

	// usually "else" keyword of what we don't really need
	return "", nil
}
