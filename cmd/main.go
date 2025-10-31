package main

import (
	"fmt"
	"os"

	piscine "go-reloaded/pkg"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// convert bytes -> stringggg
	text := string(data)

	text = piscine.HexBin(text)
	text = piscine.UpLowCap(text)
	text = piscine.FixPunctuation(text)
	text = piscine.AnConvert(text)

	// write processed text to output file
	if err := os.WriteFile(outputFile, []byte(text), 0o644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Wrote output to", outputFile)
}
