package main

import (
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"
	red   = "\033[31m"
)

type (
	formatter func(keyWord, line string) string
)

func colorFormat(keyWord, line string) string {
	if line == "" {
		return ""
	}

	replaced := strings.ReplaceAll(line, keyWord, fmt.Sprintf("%s%s%s", red, keyWord, reset))
	return fmt.Sprintf("%s\n", replaced)
}
