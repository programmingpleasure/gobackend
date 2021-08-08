package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

// The agenda:

// struct
// errors
// panics (Aaaaaaaaa!)
// arrays & slices
// maps
// pointers

const (
	// Const values can not be changed.
	// Don't use snake-case
	// Don't use UPPER_CASE as we do in C or Python
	width  = 6
	height = 6

	fileName = "image.png"
)

// Image is actually array of arrays
// 0 1 2 3
// 0 1 2 3
// 0 1 2 3
// 0 1 2 3

var (
	// Here is a map. Generally, dictionary
	// As we have in Google translate:
	// f.e.
	// English-German:
	// Airplane: Flugzeug
	// Here we do map number to color:
	// 0 -> Cyan
	// 1 -> Yellow
	// etc

	colors = map[int]*color.RGBA{
		// RGBa: RED, GREEN, BLUE, ALPHA ch. (transparency)
		// It is just a set of colors I use in draw func
		0: {R: 100, G: 200, B: 200, A: 0xff},
		1: {R: 70, G: 70, B: 21, A: 0xff},
		2: {R: 207, G: 70, B: 110, A: 0xff},
		3: {R: 78, G: 70, B: 207, A: 0xff},
		4: {R: 207, G: 205, B: 70, A: 0xff},
		5: {R: 177, G: 37, B: 180, A: 0xff},
	}
)

func init() {
	// We use random seed to take some seed for pseudo-random numbers algorithm
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// create an empty image with some size
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	// draw func that fill the image with some colors
	draw(img, colors)

	// create a file
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// write an image to a file
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func draw(img *image.RGBA, colors map[int]*color.RGBA) {
	// line by line go through the every pixel and fill that with some random color
	for x := 0; x < (width / 2); x++ {
		for y := 0; y < height; y++ {
			color := rand.Intn(5)
			// fill the left side firstly
			img.Set(x, y, colors[color])
			// fill the right side to make the image symmetric
			img.Set(width-x-1, y, colors[color])
		}
	}
}
