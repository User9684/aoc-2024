package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputFile = "./input/lists.txt"

func main() {
	var leftList = []int{}
	var rightList = []int{}
	var rightUnique = map[int]int{}

	var totalDistance = 0
	var totalSimilarity = 0

	// Open input file & create scanner
	file, err := os.OpenFile(inputFile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("Could not open input file! %s\n", err)
	}

	scanner := bufio.NewScanner(file)

	// Read input file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Separate line by the spaces inbetween the numbers
		fields := strings.Fields(line)

		// Convert left side to int and push to slice
		leftNumber, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Printf("Failed to convert left side of line \"%s\"! %s", line, err)
		}
		leftList = append(leftList, leftNumber)
		// Convert right side to int and push to slice
		rightNumber, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Printf("Failed to convert right side of line \"%s\"! %s", line, err)
		}
		rightList = append(rightList, rightNumber)
	}
	// Close file as we no longer need it
	file.Close()

	// Order left & right list by smallest to largest
	// I would use sort.Ints but it literally just does the exact same thing and I
	// want to avoid ambiguity
	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	if len(leftList) != len(rightList) {
		fmt.Println("Length of two lists are not equal!")
		return
	}

	// Iterate over the two lists, adding the absolute value of their difference
	for i := 0; i < len(leftList); i += 1 {
		leftNumber := leftList[i]
		rightNumber := rightList[i]
		distance := int(math.Abs(float64(leftNumber) - float64(rightNumber)))

		// Add to total distance for part 1
		totalDistance += distance
		// Add right numbers for part 2
		if _, ok := rightUnique[rightNumber]; !ok {
			rightUnique[rightNumber] = 0
		}
		rightUnique[rightNumber] += 1
	}

	// Iterate over unique
	for _, v := range leftList {
		similarity := v * rightUnique[v]

		totalSimilarity += similarity
	}

	// Print final distance & similarity
	fmt.Printf("Final distance: %d\nFinal similarity: %d\n", totalDistance, totalSimilarity)
}
