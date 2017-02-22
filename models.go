package main

import (
	"bytes"
	"image"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"sort"
	"time"

	"github.com/disintegration/imaging"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Face struct {
	image.Image
}

func (f *Face) LoadFile(file string) error {
	reader, err := os.Open(file)
	if err != nil {
		return err
	}
	f.Image, _, err = image.Decode(reader)
	if err != nil {
		return err
	}
	return nil
}

func NewFace(file string) (*Face, error) {
	face := &Face{}
	if err := face.LoadFile(file); err != nil {
		return face, err
	}
	return face, nil
}

func NewMustFace(file string) *Face {
	face, err := NewFace(file)
	if err != nil {
		panic(err)
	}
	return face
}

type FaceList []*Face

func (fl FaceList) Random() image.Image {
	i := rand.Intn(len(fl))
	face := fl[i]
	if rand.Intn(2) == 0 {
		return imaging.FlipH(face.Image)
	}
	return face.Image
}

func (fl *FaceList) loadInternal() {
	assets := AssetNames()
	sort.Strings(assets)
	for _, asset := range assets {
		img, _, err := image.Decode(bytes.NewReader(MustAsset(asset)))
		if err != nil {
			log.Fatalf("error decoding internal image: %s", err)
		}
		*fl = append(*fl, &Face{Image: img})
	}
}

func (fl *FaceList) Load(dir string) error {
	if dir == "" {
		fl.loadInternal()
		return nil
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".png" {
			f, err := NewFace(path.Join(dir, file.Name()))
			if err != nil {
				return err
			}
			*fl = append(*fl, f)
		}
	}
	return nil
}
