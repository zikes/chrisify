package main

import (
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"gocv.io/x/gocv"
)

var haarCascade = flag.String("haar", "haarcascade_frontalface_alt.xml", "how to make things work")
var facesDir = flag.String("faces", "faces", "The directory to search for faces.")

func detect(img gocv.Mat, xmlFile *string) []image.Rectangle {
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	var deref = *xmlFile
	classifier.Load(deref)

	faces := classifier.DetectMultiScale(img)

	return faces
}

func main() {
	flag.Parse()

	var chrisFaces FaceList

	var facesPath string
	var err error

	if *facesDir != "" {
		facesPath, err = filepath.Abs(*facesDir)
		if err != nil {
			panic(err)
		}
	}

	err = chrisFaces.Load(facesPath)
	if err != nil {
		panic(err)
	}
	if len(chrisFaces) == 0 {
		panic("no faces found")
	}

	file := flag.Arg(0)
	baseImg := loadImage(file)
	detectImg := gocv.IMRead(file, 4)

	faces := detect(detectImg, haarCascade)

	bounds := baseImg.Bounds()

	canvas := canvasFromImage(baseImg)

	for _, face := range faces {
		rect := rectMargin(30.0, face)

		newFace := chrisFaces.Random()
		if newFace == nil {
			panic("nil face")
		}
		chrisFace := imaging.Fit(newFace, rect.Dx(), rect.Dy(), imaging.Lanczos)

		draw.Draw(
			canvas,
			rect,
			chrisFace,
			bounds.Min,
			draw.Over,
		)
	}

	if len(faces) == 0 {
		face := imaging.Resize(
			chrisFaces[0],
			bounds.Dx()/3,
			0,
			imaging.Lanczos,
		)
		faceBounds := face.Bounds()
		draw.Draw(
			canvas,
			bounds,
			face,
			bounds.Min.Add(image.Pt(-bounds.Max.X/2+faceBounds.Max.X/2, -bounds.Max.Y+int(float64(faceBounds.Max.Y)/1.9))),
			draw.Over,
		)
	}

	jpeg.Encode(os.Stdout, canvas, &jpeg.Options{jpeg.DefaultQuality})
}
