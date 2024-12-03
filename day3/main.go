package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var inputFile = "./input/memory.txt"
var memoryRegex = regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Could not open input file: %s\n", err)
		return
	}
	defer file.Close()

	var total int
	var totalPart2 int
	var isDoing = true

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Could not read input file: %s\n", err)
		return
	}
	// Get all matches in the entire file (lazy)
	data := memoryRegex.FindAllSubmatch(content, -1)

	for _, v := range data {
		// Check if match is disable (part 2)
		if string(v[0]) == "don't()" {
			isDoing = false
			continue
		}
		// Check if match is enable (part 2)
		if string(v[0]) == "do()" {
			isDoing = true
			continue
		}

		// Convert arguments to numbers
		number1, err := strconv.Atoi(string(v[1]))
		if err != nil {
			fmt.Printf("Failed to parse number 1 %s\n", err)
			continue
		}
		number2, err := strconv.Atoi(string(v[2]))
		if err != nil {
			fmt.Printf("Failed to parse number 2 %s\n", err)
			continue
		}

		// Multiply both numbers
		var mult = (number1 * number2)

		// Add to total for part 1
		total += mult
		// Check if multiplication should be abled & add to total for part 2
		if isDoing {
			totalPart2 += mult
		}
	}

	fmt.Printf("Total for part 1: %d\n", total)
	fmt.Printf("Total for part 2: %d\n", totalPart2)
}
