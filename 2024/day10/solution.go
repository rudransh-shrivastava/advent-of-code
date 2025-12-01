package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	matrix := parseInput(file)
	totalTrails := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				visited := make([][]bool, len(matrix))
				for i := range visited {
					visited[i] = make([]bool, len(matrix[0]))
				}
				totalTrails += dfs(&matrix, i, j, 0, &visited)
			}
		}
	}
	fmt.Println(totalTrails)
}

func dfs(matrix *[][]int, i, j, height int, visited *[][]bool) int {
	if i < 0 || i >= len(*matrix) || j < 0 || j >= len((*matrix)[0]) {
		return 0
	}

	if (*matrix)[i][j] != height {
		return 0
	}
	if height == 9 && (*matrix)[i][j] == 9 && (*visited)[i][j] == false {
		// uncomment for part 1
		// (*visited)[i][j] = true
		return 1
	}
	return dfs(matrix, i+1, j, height+1, visited) + dfs(matrix, i-1, j, height+1, visited) + dfs(matrix, i, j+1, height+1, visited) + dfs(matrix, i, j-1, height+1, visited)
}

func parseInput(input string) [][]int {
	var result [][]int

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		var row []int
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				fmt.Println(fmt.Errorf("parsing input: %w", err))
			}
			row = append(row, num)
		}
		result = append(result, row)
	}

	return result
}
