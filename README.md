# GO-SPRITEPACK!

![Example](https://github.com/Holmqvist1990/go-spritepack/blob/master/example.png?raw=true)

A spritesheet packer, written in Go.

Checks for:
* Identical sprites.
* Identical sprites, but rotated.
* Identical sprites, but flipped horizontally.
* Identical sprites, but flipped vertically.

### **Example use.**
```
$ ./go-spritepack.exe -input="path/file.png" -output="otherpath/file_packed.png" -spritesize=16
```

### **Supports:**
- PNG
- JPEG
- GIF