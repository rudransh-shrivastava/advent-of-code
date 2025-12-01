package main

import (
	"fmt"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	locks, keys := parseString(file)
	fmt.Println(solve(locks, keys))

}

func solve(locks, keys [][][]string) int {
	locksHeights := convertLocksToHeights(locks)
	keysHeights := convertKeysToHeights(keys)
	var count = 0
	for _, lock := range locksHeights {
		for _, key := range keysHeights {
			if compareLockandKey(lock, key) {
				count++
			}
		}
	}

	return count
}

func compareLockandKey(lock, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func convertKeysToHeights(keys [][][]string) [][]int {
	var heights [][]int
	for _, key := range keys {
		heights = append(heights, convertKeyToHeight(key))
	}
	return heights
}

func convertKeyToHeight(key [][]string) []int {
	var height []int
	for i := 0; i < len(key[i]); i++ {
		count := 0
		for j := 0; j < len(key); j++ {
			if key[len(key)-j-1][i] == "#" {
				count++
			}
		}
		height = append(height, count-1)
	}
	return height
}

func convertLocksToHeights(locks [][][]string) [][]int {
	var heights [][]int
	for _, lock := range locks {
		heights = append(heights, convertLockToHeight(lock))
	}
	return heights
}

func convertLockToHeight(lock [][]string) []int {
	var height []int
	for i := 0; i < len(lock[i]); i++ {
		count := 0
		for j := 0; j < len(lock); j++ {
			if lock[j][i] == "#" {
				count++
			}
		}
		height = append(height, count-1)
	}
	return height
}

// top filled = locks
func parseString(input string) ([][][]string, [][][]string) {
	var locks, keys [][][]string

	lines := strings.Split(input, "\n")
	var temp [][]string
	for _, line := range lines {
		if line == "" {
			if temp[0][0] == "#" {
				locks = append(locks, temp)
			} else {
				keys = append(keys, temp)
			}
			temp = [][]string{}

		} else {

			tempLine := []string{}
			for _, char := range line {
				tempLine = append(tempLine, string(char))
			}
			temp = append(temp, tempLine)
		}

	}
	return locks, keys
}

func prettyPrint(matrix [][][]string) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Println(col)
		}
		fmt.Println()
	}
}
