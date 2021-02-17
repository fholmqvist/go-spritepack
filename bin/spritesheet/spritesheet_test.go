package spritesheet

import (
	_ "image/png"
	"testing"
)

func TestFromFile(t *testing.T) {
	sh := LoadFromFile(t)
	if len(sh.Sprites) != 9 {
		t.Fatalf("length was expected to be 4 was %v", len(sh.Sprites))
	}
	sh.FilterUnique()
	if len(sh.Sprites) != 2 {
		t.Fatalf("length was expected to be 2 was %v", len(sh.Sprites))
	}
}

func TestUnique(t *testing.T) {
	sh := LoadFromFile(t)
	sh.FilterUnique()
	if len(sh.Sprites) != 2 {
		t.Fatalf("length was expected to be 2 was %v", len(sh.Sprites))
	}
}

func LoadFromFile(t *testing.T) *Spritesheet {
	sh, err := FromPath("../../testfile/dwarves.png", 16)
	if err != nil {
		t.Fatal(err)
	}
	return sh
}
