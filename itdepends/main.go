package main

import "log"

// THE AGENDA:
// 1. Why do we need external dependencies
// 2. How to use them
// 3. Best practices
//		3.1 Use vendoring or go proxy
//		3.2 Check latest updates and support
//		3.3 Less dependencies is better
//		3.4 Check stars and issues on github
//		3.5 Check examples, check a tests inside the repo
//		3.6 Update on a regular basis
//		3.7 use go mod tidy
//		3.8 Use abstractions to isolate current dependencies
// 4. Defer calls
// 5. Interfaces
// 6. Simple web scrapper explanation

func main() {
	httpGetter := newHTTPGetter()
	linksChan, scrapper := makeLinkWorker(downloadPagesQueueSize, httpGetter)

	// download the first web page
	links, err := scrapper.downloadPage("https://bbc.com")
	if err != nil {
		log.Fatal()
	}

	// start a queue
	linksChan <- links
	wait()
}

func makeLinkWorker(size int, httpGetter httpGetter) (chan []string, scrapper) {
	scrapper := newScrapper(httpGetter)
	linksChan := make(chan []string, size)

	// go func, go! (^_^)
	go func() {
		for links := range linksChan {
			for _, link := range links {
				links, err := scrapper.downloadPage(link)
				if err != nil {
					log.Println(err)
					continue
				}

				// here is kinda "loop" when we write to channel we read from
				linksChan <- links
			}
		}
	}()

	return linksChan, scrapper
}
