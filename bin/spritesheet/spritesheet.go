package spritesheet

import (
	"fmt"
	"image"
	"os"

	"github.com/holmqvist1990/go-spritepack/bin/sprite"
)

type Spritesheet struct {
	Sprites sprite.Sprites
}

func FromFile(file *os.File, tileSize int) (*Spritesheet, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("tileset.FromFile: %w", err)
	}
	return &Spritesheet{
		Sprites: sprite.NewSpritesFromImage(img, tileSize),
	}, nil
}

func (sp *Spritesheet) FilterUnique() {
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

	// TODO: Test! :)
	sp.Sprites = newSprites
}
