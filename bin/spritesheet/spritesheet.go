package spritesheet

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/holmqvist1990/go-spritepack/bin/sprite"
)

type Spritesheet struct {
	Sprites  sprite.Sprites
	bounds   image.Rectangle
	tileSize int
}

func FromFile(file *os.File, tileSize int) (*Spritesheet, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("tileset.FromFile: %w", err)
	}
	return &Spritesheet{
		Sprites:  sprite.NewSpritesFromImage(img, tileSize),
		bounds:   img.Bounds(),
		tileSize: tileSize,
	}, nil
}

func (sp *Spritesheet) FilterUnique() {
	sp.Sprites = sp.Sprites.ToSet()

	spriteMap := make(map[string]sprite.Sprite)
	for _, sprite := range sp.Sprites {
		_, ok := spriteMap[sprite.Checksum()]
		if !ok {
			spriteMap[sprite.Checksum()] = sprite
		}
	}

	newSprites := sprite.Sprites{}
	for _, v := range spriteMap {
		newSprites = append(newSprites, v)
	}

	sp.Sprites = newSprites
}

func (sp *Spritesheet) SaveToFile(filename string) error {
	img := image.NewRGBA(sp.bounds)
	sp.spritesToImage(img)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = png.Encode(file, img)
	if err != nil {
		return nil
	}
	return nil
}

func (sp *Spritesheet) spritesToImage(img *image.RGBA) {
	var x, y int
	for _, sprite := range sp.Sprites {
		sp.spriteToImage(img, sprite, x, y)
		x += sp.tileSize
		if x >= sp.bounds.Max.X {
			x = 0
			y += sp.tileSize
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
