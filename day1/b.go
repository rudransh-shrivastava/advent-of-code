package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func parseStringToTwoArrays(input string) ([]int, []int) {
	var first []int
	var last []int
	lines := strings.Split(input, "\n")

	for _, line:= range lines {
			fields := strings.Fields(line)
			if len(fields) == 2 {
				num1, err := strconv.Atoi(fields[0])
				if err != nil {
					fmt.Println(fmt.Errorf("parse int: %w", err))
				}

				num2, err := strconv.Atoi(fields[1])
				if err != nil {
					fmt.Println(fmt.Errorf("parse int: %w", err))
				}

				first = append(first, num1)
				last = append(last, num2)
			}
	}
	return first, last
}

func createMap(nums []int) map[int]int {
	m := make(map[int]int)

	for _, num := range nums {
		_, exists := m[num]
		if exists {
			m[num]++
		} else {
			m[num] = 1
		}
	}
	return m
}

func calcSimilarityScore(first []int, last []int) int {
	lastMap := createMap(last)
	score := 0

	for _, num := range first {
		count, exists := lastMap[num]
		if exists {
			score += num * count
		}
	}
	return score
}

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}

	first, last := parseStringToTwoArrays(file)
	fmt.Println(calcSimilarityScore(first, last))
}