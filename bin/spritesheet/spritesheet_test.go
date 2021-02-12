package spritesheet

import (
	_ "image/png"
	"os"
	"testing"
)

func TestFromFile(t *testing.T) {
	file, err := os.Open("../../samples/sample_001.png")
	if err != nil {
		t.Fatal(err)
	}
	sh, err := FromFile(file, 8)
	if err != nil {
		t.Fatal(err)
	}
	if len(sh.Sprites) != 4 {
		t.Fatalf("length was expected to be 4 was %v", len(sh.Sprites))
	}
}
