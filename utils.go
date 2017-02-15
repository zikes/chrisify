package main

import (
	"image"
	"image/draw"
	"log"
	"os"
)

func rectMargin(pct float64, rect image.Rectangle) image.Rectangle {
	width := float64(rect.Max.X - rect.Min.X)
	height := float64(rect.Max.Y - rect.Min.Y)

	padding_width := int(pct * (width / 100) / 2)
	padding_height := int(pct * (height / 100) / 2)

	return image.Rect(
		rect.Min.X-padding_width,
		rect.Min.Y-padding_height*3,
		rect.Max.X+padding_width,
		rect.Max.Y+padding_height,
	)
}

func loadImage(file string) image.Image {
	reader, err := os.Open(file)
	if err != nil {
		log.Fatalf("error loading %s: %s", file, err)
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalf("error loading %s: %s", file, err)
	}
	return img
}

func canvasFromImage(i image.Image) *image.RGBA {
	bounds := i.Bounds()
	canvas := image.NewRGBA(bounds)
	draw.Draw(canvas, bounds, i, bounds.Min, draw.Src)

	return canvas
}
