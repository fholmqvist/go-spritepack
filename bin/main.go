package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/holmqvist1990/go-spritepack/bin/spritesheet"
)

var (
	input      string
	output     string
	spritesize int
)

func main() {
	for _, line := range logo {
		fmt.Println(line)
	}

	flag.StringVar(&input, "input", "./samples/sample_001.png", "input file path")
	flag.StringVar(&output, "output", "./samples/sample_001_packed.png", "output file path")
	flag.IntVar(&spritesize, "spritesize", 8, "sprite size for input spritesheet")

	flag.Parse()

	fmt.Printf("Loading file: %v.\n", input)

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Generating spritesheet.")

	sp, err := spritesheet.FromFile(file, spritesize)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Removing duplicates.")

	sp.FilterUnique()

	fmt.Printf("Saving to: %v.\n", output)

	err = sp.SaveToFile(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done.")
}

var logo = []string{
	"",
	" ██████   ██████          ███████ ██████  ██████  ██ ████████ ███████ ██████   █████   ██████ ██   ██  ██",
	"██       ██    ██         ██      ██   ██ ██   ██ ██    ██    ██      ██   ██ ██   ██ ██      ██  ██   ██",
	"██   ███ ██    ██  ████   ███████ ██████  ██████  ██    ██    █████   ██████  ███████ ██      █████    ██",
	"██    ██ ██    ██              ██ ██      ██   ██ ██    ██    ██      ██      ██   ██ ██      ██  ██  ",
	" ██████   ██████          ███████ ██      ██   ██ ██    ██    ███████ ██      ██   ██  ██████ ██   ██  ██",
	"",
	"Copyright (c) 2021 Fredrik Holmqvist",
	"",
}
