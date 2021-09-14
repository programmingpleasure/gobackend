package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
)

func parseImageName(img string) (string, error) {
	URL, err := url.Parse(img)
	if err != nil {
		return "", fmt.Errorf("error while parse image URL: %w", err)
	}

	imgPath := strings.Split(URL.Path, "/")
	if len(imgPath) == 0 {
		return "", errors.New("img path is empty")
	}

	return imgPath[len(imgPath)-1], nil
}

// too complicated here
func wait() {
	var wg sync.WaitGroup
	wg.Add(1)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		wg.Done()
	}()

	wg.Wait()
}
