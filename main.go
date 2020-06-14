package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"

	"wallpapersolid/resolution"
)

func hextorgb(hex string) (uint8, uint8, uint8, error) {
	r, err := strconv.ParseUint(hex[0:2], 16, 32)

	if err != nil {
		return 0, 0, 0, err
	}

	g, err := strconv.ParseUint(hex[2:4], 16, 32)

	if err != nil {
		return 0, 0, 0, err
	}

	b, err := strconv.ParseUint(hex[4:6], 16, 32)

	if err != nil {
		return 0, 0, 0, err
	}

	return uint8(r), uint8(g), uint8(b), nil
}

func createWallpaper(resolution resolution.Resolution, r uint8, g uint8, b uint8) {
	width := resolution.X
	heigth := resolution.Y

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, heigth}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	cyan := color.RGBA{r, g, b, 0xff}

	for x := 0; x < width; x++ {
		for y := 0; y < heigth; y++ {
			img.Set(x, y, cyan)
		}
	}

	f, _ := os.Create("wallpaper.png")
	png.Encode(f, img)
}

var (
	hex = flag.String("h", "", "hexcolore for wallpaper.")
)

func init() {
	flag.Parse()
}

func main() {
	if flag.Arg(0) == "" {
		log.Fatal("\nMissed monitor resolution.")

	}

	resolution, err := resolution.Parse(flag.Arg(0))
	if err != nil {
		log.Fatal("\nCorrect Format: XxY\nExample: 1366x786")
	}

	if *hex == "" {
		log.Fatal("\nCorrect Format: --hex color XxY\nExample -h 282828 1366x786")
	}

	r, g, b, err := hextorgb(*hex)
	if err != nil {
		log.Fatal("\nIt was not possible to use that color!")
	}

	createWallpaper(resolution, r, g, b)
}
