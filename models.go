package main

import (
	"image"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"
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
	return face.Image
}

func (fl *FaceList) Load(dir string) error {
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
