package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/lazywei/go-opencv/opencv"
	"github.com/nfnt/resize"
)

func rectMargin(pct int, rect image.Rectangle) image.Rectangle {
	width := float64(rect.Max.X - rect.Min.X)
	height := float64(rect.Max.Y - rect.Min.Y)

	padding_width := int(float64(pct) * (width / 100) / 2)
	padding_height := int(float64(pct) * (height / 100) / 2)

	return image.Rect(
		rect.Min.X-padding_width,
		rect.Min.Y-padding_height*3,
		rect.Max.X+padding_width,
		rect.Max.Y+padding_height,
	)
}

func main() {
	currentfile, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("could not determine executable filepath: %s", err)
	}
	rand.Seed(time.Now().UnixNano())
	chris_faces := []image.Image{
		loadImage(path.Join(currentfile, "data/chris_face1.png")),
		loadImage(path.Join(currentfile, "data/chris_face2.png")),
		loadImage(path.Join(currentfile, "data/chris_face3.png")),
		loadImage(path.Join(currentfile, "data/chris_face4.png")),
		loadImage(path.Join(currentfile, "data/chris_face5.png")),
		loadImage(path.Join(currentfile, "data/chris_face6.png")),
		loadImage(path.Join(currentfile, "data/chris_face7.png")),
	}
	file := os.Args[1]
	base_opencv_image := opencv.LoadImage(file)
	base_image := loadImage(file)

	cascade := opencv.LoadHaarClassifierCascade(path.Join(currentfile, "data/haarcascade_frontalface_alt.xml"))
	faces := cascade.DetectObjects(base_opencv_image)

	bounds := base_image.Bounds()
	new_image := image.NewRGBA(bounds)

	draw.Draw(new_image, bounds, base_image, bounds.Min, draw.Src)

	for _, face := range faces {
		rect := rectMargin(
			30,
			image.Rectangle{
				image.Point{face.X(), face.Y()},
				image.Point{face.X() + face.Width(), face.Y() + face.Height()},
			},
		)

		rect_width := rect.Max.X - rect.Min.X
		rect_height := rect.Max.Y - rect.Min.Y

		target_wide := rect_width > rect_height

		chris_face := chris_faces[rand.Intn(len(chris_faces))]

		var m image.Image

		if target_wide {
			m = resize.Resize(
				uint(rect.Max.X-rect.Min.X),
				0,
				chris_face,
				resize.Lanczos3,
			)
		} else {
			m = resize.Resize(
				0,
				uint(rect.Max.Y-rect.Min.Y),
				chris_face,
				resize.Lanczos3,
			)
		}
		//Rect(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y, 2, new_image)
		draw.Draw(
			new_image,
			rect,
			m,
			bounds.Min,
			draw.Over,
		)
	}

	if len(faces) == 0 {
		chris_face := resize.Resize(
			uint(float64(bounds.Max.X-bounds.Min.X)/3),
			0,
			chris_faces[0],
			resize.Lanczos3,
		)
		face_bounds := chris_face.Bounds()
		draw.Draw(
			new_image,
			bounds,
			chris_face,
			bounds.Min.Add(image.Pt(-bounds.Max.X/2+face_bounds.Max.X/2, -bounds.Max.Y+(face_bounds.Max.Y/2))),
			draw.Over,
		)
	}

	jpeg.Encode(os.Stdout, new_image, &jpeg.Options{jpeg.DefaultQuality})
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

func Rect(x1, y1, x2, y2, thickness int, img *image.RGBA) {
	col := color.RGBA{0, 0, 0, 255}

	for t := 0; t < thickness; t++ {
		// draw horizontal lines
		for x := x1; x <= x2; x++ {
			img.Set(x, y1+t, col)
			img.Set(x, y2-t, col)
		}
		// draw vertical lines
		for y := y1; y <= y2; y++ {
			img.Set(x1+t, y, col)
			img.Set(x2-t, y, col)
		}
	}
}
