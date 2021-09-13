package main

// ScrapData represent a result of scrapping: list of image links and links to other pages
type scrapData struct {
	images []string
	links  []string
}
