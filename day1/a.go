package main

import (
	"fmt"
	"strings"
	"sort"
	"strconv"
	"math"

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

func distanceBetweenArrays(first []int, last []int) int {
	var sum int

	for i := 0; i < len(first); i++ {
		sum += int(math.Abs(float64(first[i] - last[i])))
	}
	return sum
}

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}

	first, last := parseStringToTwoArrays(file)
	sort.Ints(first)
	sort.Ints(last)

	fmt.Println(distanceBetweenArrays(first, last))
}