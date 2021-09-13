package main

import "time"

const (
	imageSavePath          = "imgs/"
	downloadPagesQueueSize = 100000

	// some html attributes we use -------------------
	img = "img"
	src = "src"

	a    = "a"
	href = "href"
	// -----------------------------------------------

	// client client settings here -------------------
	clientDialTimeout            = time.Second * 5
	clientHandshakeTimeout       = time.Second * 5
	clientResponseHeadersTimeout = time.Second * 5
	clientTimeout                = time.Second * 15
	clientIdleConnectionTimeout  = time.Second * 20
	idleConns                    = 1000
	// -----------------------------------------------
)
