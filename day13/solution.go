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
	fmt.Println(solve(input))
}

func solve(input [][][]int) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		x1, x2, x3 := input[i][0][0], input[i][1][0], input[i][2][0]
		y1, y2, y3 := input[i][0][1], input[i][1][1], input[i][2][1]
		// comment for part 1
		x3 += 10000000000000
		y3 += 10000000000000
		//
		a, b := solveEquations(x1, x2, x3, y1, y2, y3)
		if a != -1 && b != -1 {
			sum += 3*a + b
		}
	}
	return sum
}

func solveEquations(x1, x2, x3, y1, y2, y3 int) (int, int) {
	d := x1*y2 - y1*x2
	A := (y2*x3 - x2*y3)
	B := (x1*y3 - y1*x3)
	// if A or B has decinals we return -1 -1
	if A%d != 0 || B%d != 0 {
		return -1, -1
	}
	return A / d, B / d
}

func parseInput(input string) [][][]int {
	var data [][][]int
	lines := strings.Split(input, "\n")
	machine := [][]int{}
	re := regexp.MustCompile(`(\d+).+?(\d+)`)
	for _, line := range lines {
		if line == "" {
			data = append(data, machine)
			machine = [][]int{}
			continue
		}
		match := re.FindAllStringSubmatch(line, -1)
		x, _ := strconv.Atoi(match[0][1])
		y, _ := strconv.Atoi(match[0][2])
		machine = append(machine, []int{x, y})
	}
	return data
}
