package main

import "fmt"

func print(detected []string) {
	for _, line := range detected {
		fmt.Print(line)
	}
}
