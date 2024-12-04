package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputFile = "./input/wordsearch.txt"

var directions = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
}

// Offsets for checking "X-MAS" patterns from the center ("A")
var xMasOffsets = [][3][2]int{
	{{-1, -1}, {0, 0}, {1, 1}},
	{{1, -1}, {0, 0}, {-1, 1}},
	{{-1, 1}, {0, 0}, {1, -1}},
	{{1, 1}, {0, 0}, {-1, -1}},
}

func main() {
	grid, err := readGridFromFile(inputFile)
	if err != nil {
		fmt.Printf("Failed to create a grid from file: %v\n", err)
		return
	}

	// Count occurrences of "XMAS" and X-MAS patterns
	p1 := countWordOccurrences(grid, "XMAS")
	p2 := countXMasPatterns(grid)

	// Print the results
	fmt.Printf("XMAS occurances for part 1: %d\n", p1)
	fmt.Printf("X-MAS occurances for part 2: %d\n", p2)
}

// Checks if a word starts at a specific position and direction
func wordStartsPositive(wordRunes []rune, grid [][]rune, row, col, dr, dc int) bool {
	wordLen := len(wordRunes)
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < wordLen; i++ {
		r, c := row+i*dr, col+i*dc
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != wordRunes[i] {
			return false
		}
	}
	return true
}

// Count occurrences of a word in all directions
func countWordOccurrences(grid [][]rune, word string) int {
	wordRunes := []rune(word)
	count := 0

	// Iterate through each cell in the grid
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			for _, dir := range directions {
				// Check if the word starts at the current position and direction
				if wordStartsPositive(wordRunes, grid, row, col, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

// Checks if an X-MAS pattern exists at a specific 'A' (center of X)
func isXMas(grid [][]rune, row, col int, offsets [3][2]int) bool {
	rows := len(grid)
	cols := len(grid[0])

	for i, offset := range offsets {
		r, c := row+offset[0], col+offset[1]
		// Check if offset places it outside bounds
		if c < 0 || c >= cols || r < 0 || r >= rows {
			return false
		}
		// Check if the character matches the X-MAS pattern
		if grid[r][c] != []rune("MAS")[i] {
			return false
		}
	}
	return true
}

// Count X-MAS patterns in the grid by searching for 'A' and checking the X-MAS pattern
func countXMasPatterns(grid [][]rune) int {
	count := 0

	// Iterate through each cell in the grid
	for row := 0; row < len(grid)-1; row++ {
		for col := 0; col < len(grid[0])-1; col++ {
			// Check if the center of the pattern is 'A'
			if grid[row][col] == 'A' {
				var found int
				// Check all X-MAS patterns from the current 'A'
				for _, offsets := range xMasOffsets {
					if isXMas(grid, row, col, offsets) {
						found++
					}
					if found == 2 {
						count++
						break
					}
				}
			}
		}
	}

	return count
}

// Read the grid from a file
func readGridFromFile(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	// Read each line of the file and convert it to a rune slice
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return grid, nil
}
