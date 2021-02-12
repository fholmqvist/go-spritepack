package sprite

import (
	"image"
)

type Sprites []Sprite

func (sp Sprites) ToSet() (set Sprites) {
	if len(sp) == 0 {
		return set
	}

	set = append(set, sp[0])
	for i := 0; i < len(sp)-1; i++ {
		for ii := i + 1; ii < len(sp); ii++ {
			if !sp[ii].IdenticalInRotations(sp[i]) {
				set = append(set, sp[ii])
			}
		}
	}
	return set
}

func NewSpritesFromImage(image image.Image, spriteSize int) Sprites {
	tt := Sprites{}

	maxX, maxY := image.Bounds().Max.X, image.Bounds().Max.Y
	for x := 0; x < maxX; x += spriteSize {
		for y := 0; y < maxY; y += spriteSize {
			tt = append(tt, FromImageSection(image, x, y, spriteSize))
		}
	}
	return tt
}
