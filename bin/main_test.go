package main

import (
	"image/color"
	"testing"
)

func dummyTile() Tile {
	return Tile{
		{color.Black, color.Black, color.White},
		{color.Black, color.White, color.Black},
	}
}

func TestIdentical(t *testing.T) {
	tile := dummyTile()
	tile2 := tile.Copy()

	if !tile.Identical(tile2) {
		t.Fatalf("should be identical, was:\n%v\n%v", tile, tile2)
	}
}

func TestRotateAndCopy(t *testing.T) {
	tile := dummyTile()
	tile2 := tile.Copy()
	tile.Rotate()

	if tile.Identical(tile2) {
		t.Fatalf("%v should not equal rotated %v", tile, tile2)
	}
}

func TestRotatedDuplicate(t *testing.T) {
	tile := dummyTile()
	tile2 := tile.Copy()
	tile2.Rotate()

	if !tile2.IdenticalInRotations(tile) {
		t.Fatalf("%v is not recognized as identical to %v after rotation even though it should", tile2, tile)
	}
}

func TestTilesSet(t *testing.T) {
	rotated := dummyTile()
	rotated.Rotate()

	tiles := Tiles{
		dummyTile(),
		rotated,
		dummyTile(),
	}

	set := tiles.ToSet()
	if len(set) != 1 {
		t.Fatalf("%v should be 1", len(set))
	}
}
