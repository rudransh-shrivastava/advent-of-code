package main

import (
	"fmt"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

// Part 1: 1h 55m 51s ; Part 2:
func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	lab := parseStringTo2DList(file)
	turn := "up"
	walking := true
	i, j := findStart(lab)
	fmt.Println("found starting index: ", i, j)

	for walking {
		if turn == "up" {
			if i > 0 && lab[i-1][j] != "#" {
				lab[i][j] = "X"
				i--
			} else {
				turn = "right"
			}
		} else if turn == "right" {
			if j+1 < len(lab[0]) && lab[i][j+1] != "#" {
				lab[i][j] = "X"
				j++
			} else {
				turn = "down"
			}
		} else if turn == "down" {
			if i+1 < len(lab) && lab[i+1][j] != "#" {
				lab[i][j] = "X"
				i++
			} else {
				turn = "left"
			}
		} else if turn == "left" {
			if j > 0 && lab[i][j-1] != "#" {
				lab[i][j] = "X"
				j--
			} else {
				turn = "up"
			}
		}

		if (i == 0 || j == 0 || i == len(lab)-1 || j == len(lab[0])-1) || lab[i][j] == "#" {
			walking = false
		}
	}
	lab[i][j] = "X"
	fmt.Println(countX(lab))
}

func moveUp(lab [][]string, i, j int) ([][]string, int, int, bool) {
	if isBorder(lab, i, j) {
		return lab, i, j, false
	}
	for i-1 > 0 && lab[i-1][j] != "#" {
		lab[i][j] = "X"
		i--
	}
	return lab, i, j, true
}

func moveRight(lab [][]string, i, j int) ([][]string, int, int, bool) {
	if isBorder(lab, i, j) {
		return lab, i, j, false
	}
	for j+1 < len(lab[0]) && lab[i][j+1] != "#" {
		lab[i][j] = "X"
		j++
	}
	return lab, i, j, true
}

func moveDown(lab [][]string, i, j int) ([][]string, int, int, bool) {
	if isBorder(lab, i, j) {
		return lab, i, j, false
	}
	for i+1 < len(lab) && lab[i+1][j] != "#" {
		lab[i][j] = "X"
		i++
	}
	return lab, i, j, true
}

func moveLeft(lab [][]string, i, j int) ([][]string, int, int, bool) {
	if isBorder(lab, i, j) {
		return lab, i, j, false
	}
	for j-1 > 0 && lab[i][j-1] != "#" {
		lab[i][j] = "X"
		j--
	}
	return lab, i, j, true
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
func isBorder(lab [][]string, i, j int) bool {
	labHeight := len(lab)
	labWidth := len(lab[0])
	if i == 0 || j == 0 || i == labHeight-1 || j == labWidth-1 {
		return true
	}
	return false
}

func prettyPrint(lab [][]string) {
	for line := range lab {
		fmt.Println(lab[line])
	}
}

func countX(lab [][]string) int {
	var count = 0
	for i := 0; i < len(lab); i++ {
		for j := 0; j < len(lab[0]); j++ {
			if lab[i][j] == "X" {
				count++
			}
		}
	}
	return count
}

func findStart(lab [][]string) (int, int) {
	for i := 0; i < len(lab); i++ {
		for j := 0; j < len(lab[0]); j++ {
			if lab[i][j] == "^" {
				return i, j
			}
		}
	}
	return 0, 0
}
