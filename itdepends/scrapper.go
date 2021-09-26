package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type (
	// Scrapper is a simple struct that able to find images and links on the HTML page
	// and donload images
	scrapper interface {
		// DownloadPage save all images on the page and return links that must be downloaded
		scrapPage(URL string) (scrapData, error)
		downloadImage(URL, savePath string) (uint64, uint64, error)
	}

	// scrapperImpl implement scrapper interface
	scrapperImpl struct {
		httpGetter httpGetter
	}
)

func newScrapper(httpGetter httpGetter) scrapper {
	return &scrapperImpl{
		httpGetter: httpGetter,
	}
}

func (si *scrapperImpl) scrapPage(seed string) (scrapData, error) {
	// Request the HTML page.
	body, err := si.httpGetter.get(seed)
	if err != nil {
		return scrapData{}, fmt.Errorf("error while makeing GET: %w", err)
	}

	// "defer" means execute after function returned
	defer body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return scrapData{}, fmt.Errorf("error while create new document reader: %w", err)
	}

	//  si.find(doc, img, src) -> must be readed like:
	// "find in the document all 'img' tags that contains 'src' attributes"
	images, err := si.find(doc, img, src, seed)
	if err != nil {
		return scrapData{}, fmt.Errorf("can't extract images: %w", err)
	}

	links, err := si.find(doc, a, href, seed)
	if err != nil {
		return scrapData{}, fmt.Errorf("can't extract links: %w", err)
	}

	return scrapData{
		images: images,
		links:  links,
	}, nil
}

// find use github.com/PuerkitoBio/goquery dependency for range through the document and
// extract elements by key
func (si *scrapperImpl) find(doc *goquery.Document, elem, key string, startURL string) ([]string, error) {
	if doc == nil {
		return nil, errors.New("the doc is nil")
	}

	var res []string
	// Find the review items
	doc.Find(elem).Each(func(i int, s *goquery.Selection) {
		for _, node := range s.Nodes {
			for _, attr := range node.Attr {
				if attr.Key == key {
					if strings.HasPrefix(attr.Val, "#") {
						continue
					}

					// We use only full links (here a point for improvement)
					if !strings.HasPrefix(attr.Val, "http") {
						continue
					}
					base, err := url.Parse(startURL)
					src, err2 := url.Parse(attr.Val)
					if err == nil && err2 == nil {
						res = append(res, base.ResolveReference(src).String())
					}
				}
			}
		}
	})

	return res, nil
}

// downloadImage used to save image on disk
func (si *scrapperImpl) downloadImage(imageLink, imagePath string) (uint64, uint64, error) {
	if strings.HasSuffix(imageLink, ".svg") {
		return 0, 0, nil
	}

	log.Println("download image:", imageLink)
	start := uint64(time.Now().UnixNano())
	imageBody, err := si.httpGetter.get(imageLink)
	if err != nil {
		return 0, 0, fmt.Errorf("error while get body: %w", err)
	}
	elapsed := uint64(time.Now().UnixNano()) - start

	defer imageBody.Close()

	imageName, err := parseImageName(imageLink)
	if err != nil {
		return 0, 0, fmt.Errorf("error while get img name: %w", err)
	}

	file, err := os.Create(imagePath + imageName)
	if err != nil {
		return 0, 0, fmt.Errorf("error while create file: %w", err)
	}
	defer file.Close()

	// here is a very moment where we must pay some attention:
	// we use Copy, not ioutil.ReadAll
	written, err := io.Copy(file, imageBody)
	if err != nil {
		return 0, 0, fmt.Errorf("error while write file: %w", err)
	}

	// additional check that we wrote something to a file
	if written == 0 {
		return 0, 0, fmt.Errorf("error while write file: nothing saved")
	}

	return uint64(written), elapsed, nil
}
