package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

// Part 1: 45m 21s ; Part 2:

func main() {
	file, err := utils.FileAsString("given.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	finalSum := 0
	for _, line := range strings.Split(file, "\n") {
		result, numbers := parseLine(line)
		results := computeResults(numbers)
		for i := 0; i < len(results); i++ {
			if results[i] == result {
				finalSum += result
				break
			}
		}
	}
	fmt.Println(finalSum)
}

func computeResults(numbers []int) []int {
	k := len(numbers)
	totalCombinations := 1 << (k - 1) // 2^(k-1)

	var results []int

	for i := 0; i < totalCombinations; i++ {
		expression := generateExpression(numbers, i)

		computedResult := evaluateExpression(expression)

		results = append(results, computedResult)
	}

	return results
}

// 0 = +, 1 = *
func generateExpression(numbers []int, binary int) string {
	k := len(numbers)
	expression := strconv.Itoa(numbers[0])

	for i := 0; i < k-1; i++ {
		// Check if the i-th bit is 1 (use *), if 0 (use +)
		if (binary & (1 << (k - 2 - i))) != 0 {
			expression += " * " + strconv.Itoa(numbers[i+1])
		} else {
			expression += " + " + strconv.Itoa(numbers[i+1])
		}
	}

	return expression
}

func evaluateExpression(expr string) int {
	tokens := strings.Fields(expr)

	result, _ := strconv.Atoi(tokens[0])

	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		operand, _ := strconv.Atoi(tokens[i+1])

		if operator == "+" {
			result += operand
		} else if operator == "*" {
			result *= operand
		}
	}

	return result
}

func parseLine(line string) (int, []int) {
	parts := strings.Split(line, ":")
	result, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println(fmt.Errorf("parseLine: %w", err))
	}
	equationNumbers := strings.Split(parts[1], " ")
	var numbers []int
	for _, num := range equationNumbers[1:] {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(fmt.Errorf("parseLine: %w", err))
		}
		numbers = append(numbers, n)
	}
	return result, numbers
}
