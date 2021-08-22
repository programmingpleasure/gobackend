package main

import (
	"log"
)

// The agenda:
// 1. App structure
//		1.1 Go don't go with a packages in most cases (it's not Java or C++)
//		1.2 Different files working very well
//		1.3 Security and immutability?
// 2. Code style
//		2.1 Functions over all (Max 5 input args, understandable name, simplicity, single responsibility (it's very hard)) etc
//		2.2 Responsibility
// 		2.3	Create files, not packages
//		2.4 Don't use "else" and "must" functions
//		2.5 Always to check error and think about the future (but not too much)
//		2.6
// 3. Dependencies
//		3.1 Inject them all
//		3.2 Use constructors (what is it?)
//		3.3 99% of your code must have an input and output - for testing
//		3.4 Dependencies must be substitutable
// 4. Good practices
//		4.1 Don't use pointers if you are not 100% sure app really need that (or by performance reasons)
//		4.2 Write tests
//		4.3 Write comments where it is necessary
//		4.4 No nested ifelse trash
// 5. Make own grep

func main() {
	config := mustGetConfig()

	content, err := readFile(config.filePath)
	if err != nil {
		log.Fatal(err)
	}

	detected, err := checkFull(content, config.keyString, containsCheck, colorFormat)
	if err != nil {
		log.Fatal(err)
	}

	print(detected)
}
