package main

import (
	"fmt"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

// Part 1: 1h 15m 0s ; Part 2:

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	matrix := parseStringTo2DList(file)
	fmt.Println(checkAll(matrix))

}

func checkAll(matrix [][]string) int {
	count := 0
	for i, row := range matrix {
		for j, col := range row {
			if col == "X" {
				fmt.Println("checking... ", i, j)
				count += checkXMAS(i, j, matrix)
				fmt.Println("count...", count)
			}
		}
	}
	return count
}

func checkXMAS(i int, j int, matrix [][]string) int {
	count := 0
	if checkLeft(i, j, matrix) {
		fmt.Println("found left")
		count++
	}
	if checkRight(i, j, matrix) {
		fmt.Println("found right")
		count++
	}
	if checkUp(i, j, matrix) {
		fmt.Println("found up")
		count++
	}
	if checkDown(i, j, matrix) {
		fmt.Println("found down")
		count++
	}
	if checkTopLeftDiagonal(i, j, matrix) {
		fmt.Println("found top left")
		count++
	}
	if checkTopRightDiagonal(i, j, matrix) {
		fmt.Println("found top right")
		count++
	}
	if checkBottomRightDiagonal(i, j, matrix) {
		fmt.Println("found bottom right")
		count++
	}
	if checkBottomLeftDiagonal(i, j, matrix) {
		fmt.Println("found bottom left")
		count++
	}

	return count
}

func checkTopLeftDiagonal(i int, j int, matrix [][]string) bool {
	if i < 3 || j < 3 {
		return false
	}
	pattern := []string{"X", "M", "A", "S"}

	for k := i; k > i-4; k-- {
		if matrix[k][j-(i-k)] != pattern[i-k] {
			return false
		}
	}
	return true
}

func checkTopRightDiagonal(i int, j int, matrix [][]string) bool {
	if i < 3 || j > len(matrix[i])-4 {
		return false
	}
	pattern := []string{"X", "M", "A", "S"}
	for k := i; k > i-4; k-- {
		if matrix[k][j+(i-k)] != pattern[i-k] {
			return false
		}
	}
	return true
}

func checkBottomRightDiagonal(i int, j int, matrix [][]string) bool {
	if i > len(matrix)-4 || j > len(matrix[i])-4 {
		return false
	}
	pattern := []string{"X", "M", "A", "S"}
	for k := i; k < i+4; k++ {
		if matrix[k][j+(k-i)] != pattern[k-i] {
			return false
		}
	}
	return true
}

func checkBottomLeftDiagonal(i int, j int, matrix [][]string) bool {
	if i > len(matrix)-4 || j < 3 {
		return false
	}
	pattern := []string{"X", "M", "A", "S"}
	for k := i; k < i+4; k++ {
		if matrix[k][j-(k-i)] != pattern[k-i] {
			return false
		}
	}

	return true
}

func checkLeft(i int, j int, matrix [][]string) bool {
	if j < 3 {
		return false
	}

	pattern := []string{"S", "A", "M", "X"}

	for k := j - 3; k < j; k++ {
		if matrix[i][k] != pattern[k-(j-3)] {
			return false
		}
	}
	return true
}

func checkRight(i int, j int, matrix [][]string) bool {
	if j > len(matrix[i])-4 {
		return false
	}

	pattern := []string{"X", "M", "A", "S"}

	for k := j; k < j+4; k++ {
		if matrix[i][k] != pattern[k-j] {
			return false
		}
	}
	return true
}

func checkUp(i int, j int, matrix [][]string) bool {
	if i < 3 {
		return false
	}

	pattern := []string{"S", "A", "M", "X"}

	for k := i - 3; k < i; k++ {
		if matrix[k][j] != pattern[k-(i-3)] {
			return false
		}
	}
	return true
}

func checkDown(i int, j int, matrix [][]string) bool {
	if i > len(matrix)-4 {
		return false
	}

	pattern := []string{"X", "M", "A", "S"}

	for k := i; k < i+4; k++ {
		if matrix[k][j] != pattern[k-i] {
			return false
		}
	}
	return true
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
