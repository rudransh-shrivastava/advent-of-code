package main

import (
	"fmt"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

// Part 1 : 1h 07m 02s ; Part 2 :

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	puzzleMap := parseStringTo2DList(file)
	uniqueChars := uniqueCharacters(puzzleMap)
	// fmt.Println(uniqueChars)
	antiNodeMap := createAntiNodes(puzzleMap, uniqueChars)
	// prettyPrint(antiNodeMap)
	fmt.Println(countHashes(antiNodeMap))
}

func createAntiNodes(matrix [][]string, uniqueChars []string) [][]string {
	height := len(matrix)
	width := len(matrix[0])

	antiNodeMatrix := make([][]string, height)
	for i := range antiNodeMatrix {
		antiNodeMatrix[i] = make([]string, width)
	}
	for _, char := range uniqueChars {
		// fmt.Println("for char: ", char)
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if matrix[i][j] == char {
					// fmt.Println("found char at ", i, j, matrix[i][j])
					for k := 0; k < height; k++ {
						for l := 0; l < width; l++ {
							if matrix[k][l] == char {
								if i == k && j == l {
									continue
								}
								// fmt.Println("found next char at: ", k, l, matrix[k][l])
								h := k - i
								w := l - j
								x1 := i - h
								y1 := j - w
								x2 := k + h
								y2 := l + w
								if x1 >= 0 && x1 < height && y1 >= 0 && y1 < width {
									// fmt.Println("plot 1: ", x1, y1)
									antiNodeMatrix[x1][y1] = "#"
								}
								if x2 >= 0 && x2 < height && y2 >= 0 && y2 < width {
									// fmt.Println("plot 2: ", x2, y2)
									antiNodeMatrix[x2][y2] = "#"
								}
							}
						}
					}
				}
			}
		}
	}
	return antiNodeMatrix
}

func countHashes(matrix [][]string) int {
	count := 0
	for _, row := range matrix {
		for _, char := range row {
			if char == "#" {
				count++
			}
		}
	}
	return count
}

func parseStringTo2DList(input string) [][]string {
	var list [][]string

	for _, line := range strings.Split(input, "\n") {
		for _, word := range strings.Split(line, " ") {
			list = append(list, strings.Split(word, ""))
		}
	}
	return list
}

func uniqueCharacters(matrix [][]string) []string {
	charSet := make(map[string]bool) // Track unique characters
	charArray := []string{}          // Store result

	for _, row := range matrix {
		for _, char := range row {
			if char != "." && !charSet[char] {
				charSet[char] = true // Mark as seen
				charArray = append(charArray, char)
			}
		}
	}

	return charArray
}

func prettyPrint(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
