package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var displayAsFlag bool
	var height, width int
	var filename string

	flag.BoolVar(&displayAsFlag, "f", false, "display as a flag of differing colors")
	flag.IntVar(&height, "h", 6, "height of the resulting image, defaults to 6")
	flag.IntVar(&width, "w", 6, "width of the resulting image, defaults to 6")
	flag.StringVar(&filename, "o", "image.png", "file to write the image to, defaults to image.png")
	flag.Parse()

	// create an empty image with some size
	var img *image.RGBA

	size := largestSide(height, width)

	if displayAsFlag {
		//colors := generateColorMap(size * 2)
		img = drawFlag(width, height)
	} else {
		img = drawFace(size)
	}

	writeImgToFile(filename, img)
}

func writeImgToFile(filename string, img *image.RGBA) {

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	// write an image to a file
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func largestSide(height, width int) int {
	return int(math.Max(float64(width), float64(height)))
}

func generateColorMap(size int) []*color.RGBA {
	generatedColors := make(map[uint32]*color.RGBA)
	var outcome []*color.RGBA = make([]*color.RGBA, size)

	const mask uint32 = 0x000000ff
	for len(generatedColors) < size {
		randomColor := rand.Uint32() | 255
		generatedColors[randomColor] = &color.RGBA{R: byte(randomColor >> 24 & mask), G: byte(randomColor >> 16 & mask), B: byte(randomColor >> 8 & mask), A: 0xff}
	}

	c := 0
	for _, i := range generatedColors {
		outcome[c] = i
		c++

	}

	return outcome
}

func drawFlag(width int, height int) *image.RGBA {
	var size int = largestSide(height, width)

	switch rand.Intn(3) {
	case 0:
		return drawFlagHorizontal(generateColorMap(size), width, height)
	case 1:
		return drawFlagVertical(generateColorMap(size), width, height)
	default:
		return drawFlagDiagonal(generateColorMap(size*2), width, height)
	}
}

func drawFlagHorizontal(colors []*color.RGBA, width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for x := 0; x < (width); x++ {
		color := rand.Intn(len(colors))
		for y := 0; y < height; y++ {
			// fill the left side firstly
			img.Set(y, x, colors[color])
			// fill the right side to make the image symmetric

		}
		colors = deleteColorElement(colors, color)
	}

	return img

}
func drawFlagVertical(colors []*color.RGBA, width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for x := 0; x < (width); x++ {
		color := rand.Intn(len(colors))
		for y := 0; y < height; y++ {
			// fill the left side firstly
			img.Set(x, y, colors[color])
			// fill the right side to make the image symmetric

		}
		colors = deleteColorElement(colors, color)
	}

	return img

}
func drawFlagDiagonal(colors []*color.RGBA, width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	for x := 0; x < width; x++ {
		color := rand.Intn(len(colors))
		for y, x2 := height-1, x; y >= 0 && x2 >= 0; y, x2 = y-1, x2-1 {
			img.Set(x2, y, colors[color])
		}

		colors = deleteColorElement(colors, color)

		color = rand.Intn(len(colors))
		for y, x2 := 0, width-x; y < height && x2 < width; y, x2 = y+1, x2+1 {
			img.Set(x2, y, colors[color])
		}

		colors = deleteColorElement(colors, color)
	}

	return img

}

func drawFace(size int) *image.RGBA {
	var colors []*color.RGBA = generateColorMap(size)
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{size, size}})

	// line by line go through the every pixel and fill that with some random color
	for x := 0; x < (size / 2); x++ {
		for y := 0; y < size; y++ {
			color := rand.Intn(size)
			// fill the left side firstly
			img.Set(x, y, colors[color])
			// fill the right side to make the image symmetric
			img.Set(size-x-1, y, colors[color])
		}
	}

	return img
}
func deleteColorElement(elements []*color.RGBA, index int) []*color.RGBA {
	return append(elements[:index], elements[index+1:]...)
}

func palindromeP(sentence string) bool {
	var sanitized string = sanitize(sentence)

	for s, e := 0, utf8.RuneCountInString(sanitized); s <= e; s, e = s+1, e-1 {

		startingRune, _ := utf8.DecodeRuneInString(sanitized[s:])
		if startingRune == utf8.RuneError {
			return false
		}

		endingRune, _ := utf8.DecodeLastRuneInString(sanitized[:e])
		if endingRune != startingRune {
			return false
		}
	}

	return true
}

func sanitize(source string) string {
	return strings.Map(removeWhitespaceAndInterpunction, strings.ToLower(source))
}

func removeWhitespaceAndInterpunction(letter rune) rune {
	toRemove := ",.;:?!&()[]{}-\"' \t\n\r"

	if strings.ContainsRune(toRemove, letter) {
		return -1
	}

	return letter
}
