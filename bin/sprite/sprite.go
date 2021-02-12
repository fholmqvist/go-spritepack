package sprite

import (
	"image"
	"image/color"
)

type Sprite [][]color.Color

func FromImageSection(image image.Image, startX, startY, tileSize int) Sprite {
	t := Sprite{}
	for x := 0; x < tileSize; x++ {
		t = append(t, []color.Color{})
		for y := 0; y < tileSize; y++ {
			t[x] = append(t[x], image.At(x+startX, y+startY))
		}
	}
	return t
}

func (s Sprite) Identical(b Sprite) bool {
	if len(s) != len(b) {
		return false
	}
	for x, row := range s {
		for y := range row {
			if s[x][y] != b[x][y] {
				return false
			}
		}
	}
	return true
}

func (s Sprite) IdenticalInRotations(b Sprite) bool {
	for i := 0; i < 4; i++ {
		if s.Identical(b) {
			return true
		}

		b.Rotate()
	}
	return false
}

func (s Sprite) Copy() Sprite {
	var copy [][]color.Color
	for x, row := range s {
		copy = append(copy, []color.Color{})
		for y := range row {
			copy[x] = append(copy[x], s[x][y])
		}
	}
	return copy
}

func (s Sprite) Rotate() {
	n := len(s)

	// Transpose.
	for i := 0; i < n-1; i++ {
		for ii := i; ii < n; ii++ {
			s[i][ii], s[ii][i] = s[ii][i], s[i][ii]
		}
	}

	// Flip.
	for _, row := range s {
		for i := 0; i < n/2; i++ {
			row[i], row[n-1-i] = row[n-1-i], row[i]
		}
	}
}
