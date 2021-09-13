package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

type (
	// httpGetter used to exec GET HTTP request and return ReadCloser for body
	// Attention: body must be closed after every usage
	// The question:
	// Why do we use this interface?
	httpGetter interface {
		get(URL string) (io.ReadCloser, error)
	}

	httpGetterImpl struct {
		client *http.Client
	}
)

func newHTTPGetter() httpGetter {
	return &httpGetterImpl{
		client: httpClient(),
	}
}

func (hg *httpGetterImpl) get(URL string) (io.ReadCloser, error) {
	res, err := hg.client.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("error while do get: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while scrap page :%s, status code is %d", URL, res.StatusCode)
	}

	return res.Body, nil
}

func httpClient() *http.Client {
	return &http.Client{
		Timeout: clientTimeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: clientDialTimeout,
			}).DialContext,
			TLSHandshakeTimeout:   clientHandshakeTimeout,
			MaxIdleConns:          idleConns,
			MaxIdleConnsPerHost:   idleConns,
			IdleConnTimeout:       clientIdleConnectionTimeout,
			ResponseHeaderTimeout: clientResponseHeadersTimeout,
		},
	}
}
