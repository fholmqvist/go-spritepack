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

	flag.StringVar(&input, "input", "", "Input file path.\nExample: -input=\"./testfile/dwarves.png\"")
	flag.StringVar(&output, "output", "", "Output file path.\nExample: -output=\"./testfile/dwarves_packed.png\"")
	flag.IntVar(&spritesize, "spritesize", 16, "Sprite size for input spritesheet.\nExample: -spritesize=16")

	flag.Parse()

	help := "Run with -help to display more information."

	if input == "" {
		fmt.Println("No input file given. Please supply one with -input=\"path\".")
		fmt.Println(help)
		return
	}

	if output == "" {
		fmt.Println("No output file given. Please supply one with -output=\"path\".")
		fmt.Println(help)
		return
	}

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
