package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var inputFile = "./input/input.txt"

// input today was two sections so didnt feel like giving a "unique" name
// like I usually do.

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Could not open input file: %s\n", err)
		return
	}
	defer file.Close()

	var total int
	var totalPart2 int

	input, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Could not read input file: %s\n", err)
		return
	}
	inputString := strings.ReplaceAll(string(input), "\r", "")
	// GRAAAAHHHHHH!!! I HATE WINDOWS!!!!!!

	rules, updates := parseInput(inputString)

	for _, update := range updates {
		if isValidOrder(update, rules) {
			total += middleOfSlice(update)
			continue
		}
		reorderedUpdate := reorderUpdateLine(update, rules)
		totalPart2 += middleOfSlice(reorderedUpdate)
	}

	fmt.Printf("Sum of middle numbers for part 1: %d\n", total)
	fmt.Printf("Sum of middle numbers post re-order for part 2: %d\n", totalPart2)
}

// Returns "rules" and "updates"
func parseInput(input string) (map[int][]int, [][]int) {
	var rules = map[int][]int{}
	var updates [][]int

	sections := strings.Split(input, "\n\n")

	// Read "rules" section line by line, splitting by pipe and appending to
	// rules slice
	for _, line := range strings.Split(sections[0], "\n") {
		parts := strings.Split(line, "|")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		rules[a] = append(rules[a], b)
	}

	// Read "updates" section line by line and simply append the numbers as a 2D slice
	for _, line := range strings.Split(sections[1], "\n") {
		fields := strings.Split(line, ",")
		update := make([]int, len(fields))

		for i, field := range fields {
			update[i], _ = strconv.Atoi(field)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

// Check if "updates" line is in the correct order based on rules
func isValidOrder(update []int, rules map[int][]int) bool {
	indexMap := make(map[int]int)
	for i, page := range update {
		indexMap[page] = i
	}

	for left, numbersBefore := range rules {
		for _, number := range numbersBefore {
			indexA, okA := indexMap[left]
			indexB, okB := indexMap[number]
			if okA && okB && indexA >= indexB {
				return false
			}
		}
	}

	return true
}

// Reorder "updates" according to rules (in possibly the worst way ever)
func reorderUpdateLine(update []int, rules map[int][]int) []int {
	numbersBefore := make(map[int]int)
	numbersAfter := make(map[int][]int) // idk.

	// Initialize dependencies
	for _, page := range update {
		numbersBefore[page] = 0
	}

	for a, dependencies := range rules {
		for _, b := range dependencies {
			if contains(update, a) && contains(update, b) {
				numbersAfter[a] = append(numbersAfter[a], b)
				numbersBefore[b]++
			}
		}
	}
	var reordered []int
	var queue []int

	// Add pages with no dependencies to the queue
	for page, count := range numbersBefore {
		if count == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		page := queue[0]
		queue = queue[1:]
		reordered = append(reordered, page)

		for _, dependent := range numbersAfter[page] {
			numbersBefore[dependent]--
			if numbersBefore[dependent] == 0 {
				queue = append(queue, dependent)
			}
		}
	}

	for _, page := range update {
		if !contains(reordered, page) {
			reordered = append(reordered, page)
		}
	}

	return reordered
}

// Checks if a slice contains a specific value
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Return the middle index of the slice (im begging you to not be an odd number)
func middleOfSlice(slice []int) int {
	return slice[len(slice)/2]
}
