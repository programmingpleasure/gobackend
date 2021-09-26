package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"
)

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

type downloadDuration struct {
	duration uint64
	size     uint64
}

func main() {
	var maxDepth int
	var targetUrl string
	var verbose bool
	flag.IntVar(&maxDepth, "maxDepth", -1, "maximum depth of levels to branch out to (default unlimited)")
	flag.StringVar(&targetUrl, "url", "http://www.bbc.co.uk", "url to download the images from (defaults to http://www.bbc.co.uk)")
	flag.BoolVar(&verbose, "v", true, "print out extra info at the end of the operations")
	flag.Parse()

	var stopOperations bool = false
	imgChan := make(chan string, downloadPagesQueueSize)
	statChan := make(chan downloadDuration, 10)
	reportchan := make(chan downloadDuration)
	go keepStats(statChan, reportchan, &stopOperations)
	go downloadAndStoreTheImages(imgChan, statChan, &stopOperations)

	urlsP := makeLinkWorker(downloadPagesQueueSize, targetUrl, imgChan, maxDepth)
	if verbose {
		fmt.Println("\n\nURI visited:\n---------------\n")
		for _, urlVisited := range urlsP {
			fmt.Println(urlVisited)
		}
	}
	stopOperations = true

	avg := <-reportchan
	if verbose {
		fmt.Printf("avg image download speed: %d bytes taking %s\n", avg.size, time.Duration(avg.duration).String())
	}

	// start a queue
	//linksChan <- links
}

func keepStats(statChan <-chan downloadDuration, reportChan chan<- downloadDuration, stopped *bool) {
	var counter uint64 = 0
	var timeElapsed uint64 = 0
	var totalDownloaded uint64 = 0

	for {
		select {
		case dd := <-statChan:
			if counter != 0 && (dd.size+totalDownloaded < totalDownloaded || dd.duration+timeElapsed < timeElapsed) {
				timeElapsed = timeElapsed / counter
				totalDownloaded = totalDownloaded / counter
				counter = 1
			}
			counter++
			timeElapsed += dd.duration
			totalDownloaded += dd.size
		default:
			if *stopped {
				if counter == 0 {
					timeElapsed = 0
					totalDownloaded = 0
					counter = 1
				}
				reportChan <- downloadDuration{timeElapsed / counter, totalDownloaded / counter}
				return
			}

		}
	}

}

func downloadAndStoreTheImages(imgUrls <-chan string, statChan chan<- downloadDuration, stopped *bool) {
	scrapper := newScrapper(newHTTPGetter())
	imgSet := make(map[string]bool)

	for !*stopped {
		select {
		case imgUrl := <-imgUrls:
			_, exists := imgSet[imgUrl]
			if !exists {
				imgSet[imgUrl] = true
				written, duration, err := scrapper.downloadImage(imgUrl, imageSavePath)
				statChan <- downloadDuration{duration, written}
				if err != nil {
					log.Println(err)
				}
			}
		default:
		}
	}

}

type urlDownloadContext struct {
	url   string
	level int
}

func appendUrlsWithContext(urlsToScrape []urlDownloadContext, level int, startURL string, urlsToAdd []string) []urlDownloadContext {
	var newUrlsWithContext []urlDownloadContext = make([]urlDownloadContext, len(urlsToAdd))
	startURLData, _ := url.Parse(startURL)
	for i, urlToAdd := range urlsToAdd {
		adjustedLevel := level
		urlToAddData, _ := url.Parse(urlToAdd)
		if startURLData.Host != urlToAddData.Host {
			adjustedLevel++
		}
		newUrlsWithContext[i] = urlDownloadContext{urlToAdd, adjustedLevel}
	}

	return append(urlsToScrape, newUrlsWithContext...)
}

func makeLinkWorker(size int, startURL string, imgChan chan<- string, maxStepsRemoved int) []string {
	scrapper := newScrapper(newHTTPGetter())
	var urlsFound map[string]bool = make(map[string]bool)
	urlsFound[startURL] = false
	var urlsToScrape []urlDownloadContext = []urlDownloadContext{urlDownloadContext{startURL, 0}}
	base, _ := url.Parse(startURL)

	for {
		if len(urlsToScrape) == 0 {
			break
		}

		link := urlsToScrape[0].url
		linkData, _ := url.Parse(link)
		currentLevel := urlsToScrape[0].level
		if base.Host == linkData.Host {
			currentLevel = 0
		}
		urlsToScrape = urlsToScrape[1:]

		if downloaded, exists := urlsFound[link]; downloaded && exists {
			continue
		}

		if maxStepsRemoved == -1 || currentLevel <= maxStepsRemoved {
			data, err := scrapper.scrapPage(link)
			if err != nil {
				log.Println(err)
				continue
			}
			urlsToScrape = appendUrlsWithContext(urlsToScrape, currentLevel, startURL, data.links)
			urlsFound[link] = true

			for _, img := range data.images {
				imgChan <- img
			}
		}

	}

	var urlsProcessed []string = make([]string, len(urlsFound))
	i := 0
	for key, _ := range urlsFound {
		urlsProcessed[i] = key
		i++
	}

	return urlsProcessed
}
