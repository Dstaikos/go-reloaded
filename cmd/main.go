package main

import (
	"fmt"
	"os"

	"go-reloaded/pkg"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// read text from input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// convert bytes -> string
	text := string(data)

	text = pkg.HexBin(text)
	text = pkg.UpLowCap(text)
	text = pkg.AnConvert(text)
	text = pkg.FixPunctuation(text)

	// write processed text to output file
	if err := os.WriteFile(outputFile, []byte(text), 0o644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Wrote output to", outputFile)
}
