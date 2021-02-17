# GO-SPRITEPACK!

A spritesheet packer, written in Go.

Checks for:
* Identical sprites.
* Identical sprites, but rotated.
* Identical sprites, but flipped horizontally.
* Identical sprites, but flipped vertically.

### **Example use.**
```
./go-spritepack.exe -input="path/file.png" -output="otherpath/file_packed.png" -tilesize=16
```

### **Supports:**
- PNG
- JPEG
- GIF