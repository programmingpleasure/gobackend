package main

import "fmt"

func main() {
	addSmile("hey team")
	addSmile("any smile here?")
}

func addSmile(txt string) {
	fmt.Printf("%s :)\n", txt)
}
