package spritesheet

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/holmqvist1990/go-spritepack/bin/sprite"
)

type Spritesheet struct {
	Sprites    sprite.Sprites
	bounds     image.Rectangle
	spritesize int
}

func FromPath(filepath string, spritesize int) (*Spritesheet, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return FromFile(file, spritesize)
}

func FromFile(file *os.File, spritesize int) (*Spritesheet, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("tileset.FromFile: %w", err)
	}
	return &Spritesheet{
		Sprites:    sprite.NewSpritesFromImage(img, spritesize),
		bounds:     img.Bounds(),
		spritesize: spritesize,
	}, nil
}

func (sp *Spritesheet) FilterUnique() {
	sp.Sprites = sp.Sprites.ToSet()

	spriteMap := make(map[string]sprite.Sprite)
	for _, sprite := range sp.Sprites {
		checksum := sprite.Checksum()
		_, ok := spriteMap[checksum]
		if !ok {
			spriteMap[checksum] = sprite
		}
	}

	newSprites := sprite.Sprites{}
	for _, v := range spriteMap {
		newSprites = append(newSprites, v)
	}

	sp.Sprites = newSprites
}

func (sp *Spritesheet) SaveToFile(filename string) error {
	if len(filename) < 4 || !strings.Contains(filename, ".") {
		return fmt.Errorf("invalid filename: %v", filename)
	}

	img := image.NewRGBA(sp.bounds)
	sp.spritesToImage(img)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	switch strings.Split(filename, ".")[1] {
	case "png":
		err = png.Encode(file, img)
	case "gif":
		err = gif.Encode(file, img, nil)
	case "jpeg":
		err = jpeg.Encode(file, img, nil)
	}
	if err != nil {
		return nil
	}
	return nil
}

func (sp *Spritesheet) spritesToImage(img *image.RGBA) {
	var x, y int
	for _, sprite := range sp.Sprites {
		sp.spriteToImage(img, sprite, x, y)
		x += sp.spritesize
		if x >= sp.bounds.Max.X {
			x = 0
			y += sp.spritesize
		}
	}
}

func (sp *Spritesheet) spriteToImage(img *image.RGBA, sprite sprite.Sprite, xOffset, yOffset int) {
	for x, row := range sprite {
		for y, col := range row {
			img.Set(x+xOffset, y+yOffset, col)
		}
	}
}
