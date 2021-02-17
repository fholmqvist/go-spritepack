package sprite

import (
	"fmt"
	"image"
)

type Sprites []Sprite

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

func (sp Sprites) ToSet() (set Sprites) {
	if len(sp) == 0 {
		return
	}

	spriteMap := map[string]Sprite{}
	for i := 0; i < len(sp)-1; i++ {
		for ii := i + 1; ii < len(sp); ii++ {
			if sp[ii].IdenticalIfRotated(sp[i]) {
				continue
			}
			if sp[ii].IdenticalIfFlippedHorizontally(sp[i]) {
				continue
			}
			if sp[ii].IdenticalIfFlippedVertically(sp[i]) {
				continue
			}
			checksum := sp[ii].Checksum()
			_, ok := spriteMap[checksum]
			if !ok {
				spriteMap[checksum] = sp[ii]
				continue
			}
		}
	}

	for _, v := range spriteMap {
		set = append(set, v)
	}

	return
}

func (sp Sprites) Checksum() string {
	var checksum string
	for _, sprite := range sp {
		checksum += fmt.Sprintf("%v_", sprite)
	}
	return checksum[:len(checksum)-1]
}
