package sprite

import (
	"image"
	"image/color"
	_ "image/png"
	"os"
	"testing"
)

func TestFromImageSection(t *testing.T) {
	file, err := os.Open("../../samples/sample_001.png")
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatalf("%v", err)
	}
	sprite := FromImageSection(img, 0, 0, 8)
	if len(sprite) != 8 {
		t.Fatalf("expected length to be 8 was %v", len(sprite))
	}

	// Name clash between type RGBA and function
	// RGBA arbitrarily doesn't work on some lines,
	// hence this ... slightly less elegant implementation.
	want := color.RGBA{0, 0, 0, 255}
	wr, wg, wb, wa := want.RGBA()
	for _, row := range sprite {
		for _, color := range row {
			r, g, b, a := color.RGBA()
			if r != wr || g != wg || b != wb || a != wa {
				t.Fatalf("colors not identical\nSprite: %v\nwant: %v", color, want)
			}
		}
	}
}

func TestIdentical(t *testing.T) {
	sprite1, sprite2 := dummySprites()

	if !sprite1.Identical(sprite2) {
		t.Fatalf("should be identical, was:\n%v\n%v", sprite1, sprite2)
	}
}

func TestCopy(t *testing.T) {
	sprite1 := dummySprite()
	sprite2 := sprite1.Copy()
	sprite2.Rotate()

	if sprite1.Identical(sprite2) {
		t.Fatalf("\n%v\nshould not equal rotated\n%v", sprite1, sprite2)
	}
}

func TestRotate(t *testing.T) {
	sprite := Sprite{
		{color.Black, color.White, color.White},
		{color.White, color.White, color.White},
		{color.White, color.White, color.White},
	}

	rotations := Sprites{
		{
			{color.White, color.White, color.Black},
			{color.White, color.White, color.White},
			{color.White, color.White, color.White},
		},
		{
			{color.White, color.White, color.White},
			{color.White, color.White, color.White},
			{color.White, color.White, color.Black},
		},
		{
			{color.White, color.White, color.White},
			{color.White, color.White, color.White},
			{color.Black, color.White, color.White},
		},
		{
			{color.Black, color.White, color.White},
			{color.White, color.White, color.White},
			{color.White, color.White, color.White},
		},
	}

	for _, rot := range rotations {
		sprite.Rotate()
		if !sprite.Identical(rot) {
			t.Fatalf("\n%v\n%v\nand\n%v\n%v\nshould be equal when rotated",
				sprite[0], sprite[1], rot[0], rot[1])
		}
	}
}

func TestRotatedDuplicate(t *testing.T) {
	sprite1, sprite2 := dummySprites()
	sprite2.Rotate()

	if !sprite2.IdenticalIfRotated(sprite1) {
		t.Fatalf("%v should be identical to %v despite rotations", sprite2, sprite1)
	}
}

func TestHorizontalFlip(t *testing.T) {
	sprite := Sprite{
		{color.Black, color.White},
	}

	desired := Sprite{
		{color.White, color.Black},
	}

	sprite.FlipHorizontally()

	if !sprite.Identical(desired) {
		t.Fatalf("%v should be identical to %v", sprite, desired)
	}
}

func TestVerticalFlip(t *testing.T) {
	sprite := Sprite{
		{color.Black, color.White},
		{color.White, color.Black},
	}

	desired := Sprite{
		{color.White, color.Black},
		{color.Black, color.White},
	}

	sprite.FlipVertically()

	if !sprite.Identical(desired) {
		t.Fatalf("%v should be identical to %v", sprite, desired)
	}
}

func TestSpritesToSet(t *testing.T) {
	rotated := dummySprite()
	rotated.Rotate()

	Sprites := Sprites{
		dummySprite(),
		rotated,
		dummySprite(),
	}

	set := Sprites.ToSet()
	if len(set) != 1 {
		t.Fatalf("%v should be 1", len(set))
	}
}

func dummySprite() Sprite {
	return Sprite{
		{color.Black, color.Black, color.White},
		{color.Black, color.White, color.Black},
	}
}

func dummySprites() (Sprite, Sprite) {
	return dummySprite(), dummySprite()
}
