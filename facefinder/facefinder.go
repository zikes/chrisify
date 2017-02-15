package facefinder

import (
	"image"

	"github.com/lazywei/go-opencv/opencv"
)

var faceCascade *opencv.HaarCascade

type Finder struct {
	cascade *opencv.HaarCascade
}

func NewFinder(xml string) *Finder {
	return &Finder{
		cascade: opencv.LoadHaarClassifierCascade(xml),
	}
}

func (f *Finder) Detect(i image.Image) []image.Rectangle {
	var output []image.Rectangle

	faces := f.cascade.DetectObjects(opencv.FromImage(i))
	for _, face := range faces {
		output = append(output, image.Rectangle{
			image.Point{face.X(), face.Y()},
			image.Point{face.X() + face.Width(), face.Y() + face.Height()},
		})
	}

	return output
}
