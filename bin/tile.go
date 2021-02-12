package main

import "image/color"

type Tile [][]color.Color

func (t Tile) Identical(b Tile) bool {
	if len(t) != len(b) {
		return false
	}

	for x, row := range t {
		for y := range row {
			if t[x][y] != b[x][y] {
				return false
			}
		}
	}

	return true
}

func (t Tile) IdenticalInRotations(b Tile) bool {
	for i := 0; i < 4; i++ {
		if t.Identical(b) {
			return true
		}
		b.Rotate()
	}
	return false
}

func (t Tile) Copy() Tile {
	var copy [][]color.Color
	for x, row := range t {
		copy = append(copy, []color.Color{})
		for y := range row {
			copy[x] = append(copy[x], t[x][y])
		}
	}
	return copy
}

func (t Tile) Rotate() {
	n := len(t)

	// Transpose.
	for i := 0; i < n-1; i++ {
		for ii := i; ii < n; ii++ {
			t[i][ii], t[ii][i] = t[ii][i], t[i][ii]
		}
	}

	// Flip.
	for _, row := range t {
		for i := 0; i < n/2; i++ {
			row[i], row[n-1-i] = row[n-1-i], row[i]
		}
	}
}
