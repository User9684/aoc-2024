package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var inputFile = "./input/reports.txt"

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Could not open input file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safe int
	var safePart2 int

	// Read input file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Convert string fields to integers
		numbers := make([]int, len(fields))
		for i, field := range fields {
			number, err := strconv.Atoi(field)
			if err != nil {
				fmt.Printf("Invalid number in line: %s\n", line)
				continue
			}
			numbers[i] = number
		}

		// Check if the sequence is safe for part 1
		if isSafe(numbers) {
			safe++
		}
		// Check if the sequence is safe for part 2
		if isSafe(numbers) || isSafeWithDampener(numbers) {
			safePart2++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %s\n", err)
	}

	fmt.Println(safe)
	fmt.Println(safePart2)
}

// Check if "report" is safe for part 1
func isSafe(numbers []int) bool {
	increasing := true
	decreasing := true

	for i := 1; i < len(numbers); i++ {
		diff := math.Abs(float64(numbers[i] - numbers[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
		if numbers[i] > numbers[i-1] {
			decreasing = false
		} else if numbers[i] < numbers[i-1] {
			increasing = false
		}
	}

	return increasing || decreasing
}

// Check if "report" is safe for part 2
func isSafeWithDampener(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		// Remove 1 of the values and check if safe
		// idk why but slices.Remove just completely fucks it so im doing it
		// the long way tm.
		modified := append([]int{}, numbers[:i]...)
		modified = append(modified, numbers[i+1:]...)

		if isSafe(modified) {
			return true
		}
	}

	return false
}
