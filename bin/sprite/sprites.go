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

// Removes all duplicate sprites in collection.
func (sp Sprites) ToSet() (set Sprites) {
	if len(sp) == 0 {
		return
	}

	spriteMap := map[string]Sprite{}

outer:
	for fst := 0; fst < len(sp)-1; fst++ {
		for snd := fst + 1; snd < len(sp); snd++ {
			if sp[snd].IdenticalIfRotated(sp[fst]) {
				continue outer
			}
			if sp[snd].IdenticalIfFlippedHorizontally(sp[fst]) {
				continue outer
			}
			if sp[snd].IdenticalIfFlippedVertically(sp[fst]) {
				continue outer
			}

			id := sp[snd].ID()

			_, ok := spriteMap[id]
			if !ok {
				// Completely unique
				// sprite. Add it.
				spriteMap[id] = sp[snd]
				continue outer
			}
		}
	}

	// Wow, such filter. Much removed everything.
	if len(spriteMap) == 0 {
		spriteMap[sp[0].ID()] = sp[0]
	}

	for _, v := range spriteMap {
		set = append(set, v)
	}

	return
}

// Generates an aggregate ID based
// on the underlying sprite IDs.
func (sp Sprites) ID() string {
	var checksum string
	for _, sprite := range sp {
		checksum += fmt.Sprintf("%v_", sprite)
	}
	return checksum[:len(checksum)-1]
}
