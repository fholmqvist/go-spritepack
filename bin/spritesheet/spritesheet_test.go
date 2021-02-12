package spritesheet

import (
	_ "image/png"
	"os"
	"testing"
)

func TestFromFile(t *testing.T) {
	sh := LoadFromFile(t)
	if len(sh.Sprites) != 4 {
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
	file, err := os.Open("../../samples/sample_001.png")
	if err != nil {
		t.Fatal(err)
	}
	sh, err := FromFile(file, 8)
	if err != nil {
		t.Fatal(err)
	}
	return sh
}
