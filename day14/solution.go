package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	input := parseInput(file)
	for i := 8260; i < 100000; i++ {
		fmt.Println(i)
		solve(input, i, 101, 103)
		// block loop until user enters something in the console
		fmt.Scanln()
	}
}

func solve(input [][]int, seconds, width, height int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	matrix := make([][]string, height)
	for i := range matrix {
		matrix[i] = make([]string, width)
		for j := range matrix[i] {
			matrix[i][j] = ". "
		}
	}
	for i := 0; i < len(input); i++ {
		x, y, vx, vy := input[i][0], input[i][1], input[i][2], input[i][3]
		x, y = pos(x, y, vx, vy, seconds, width, height)
		matrix[y][x] = "# "
		quadrant := findQuadrant(x, y, height, width)
		switch quadrant {
		case 1:
			q1++
		case 2:
			q2++
		case 3:
			q3++
		case 4:
			q4++
		}
	}
	prettyPrint(matrix)
	return q1 * q2 * q3 * q4
}

func prettyPrint(matrix [][]string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}
func findQuadrant(x, y, height, width int) int {
	if x == width/2 || y == height/2 {
		return 0
	}
	if x < width/2 && y < height/2 {
		return 1
	}
	if x >= width/2 && y < height/2 {
		return 2
	}
	if x < width/2 && y >= height/2 {
		return 3
	}
	return 4
}
func pos(x, y, vx, vy, seconds int, width, height int) (int, int) {
	for i := 0; i < seconds; i++ {
		// oldx := x
		// oldy := y
		x += vx
		y += vy
		if x < 0 {
			x = width + x
		}
		if y < 0 {
			y = height + y
		}
		if x >= width {
			x = x % width
		}
		if y >= height {
			y = y % height
		}
	}
	return x, y
}

// regex to parse p=0,4 v=3,-3 and get 0, 4, 3, -3
func parseInput(input string) [][]int {
	var data [][]int

	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`[-+]?\d+,[+-]?\d+`)
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}
		x, _ := strconv.Atoi(strings.Split(matches[0], ",")[0])
		y, _ := strconv.Atoi(strings.Split(matches[0], ",")[1])
		vx, _ := strconv.Atoi(strings.Split(matches[1], ",")[0])
		vy, _ := strconv.Atoi(strings.Split(matches[1], ",")[1])
		data = append(data, []int{x, y, vx, vy})
	}
	return data
}
