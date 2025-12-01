package main

import (
	"fmt"
	"github.com/ryntak94/advent-of-code-go-starter/utils"
	"regexp"
	"strconv"
)

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}

	indices := findMulIndices(file)
	fmt.Println(calculateAllMul(indices, file))
}

func findMulIndices(code string) [][]int {

	pattern := `mul\(\d+,\d+\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error:", err)
	}

	matches := re.FindAllStringIndex(code, -1)

	return matches
}

func calculateAllMul(indices [][]int, code string) int {
	// add 4 and add 5
	answer := 0
	for i := range indices {
		index := indices[i][0]
		// first num can be more than one digit

		firstNum := 0
		secondNum := 0
		for j := index + 4; code[j] != ','; j++ {
			firstNum = firstNum*10 + int(code[j]-'0')
		}
		for j := index + 5 + len(strconv.Itoa(firstNum)); code[j] != ')'; j++ {
			secondNum = secondNum*10 + int(code[j]-'0')
		}
		answer += firstNum * secondNum
		//fmt.Println("First:", firstNum, "Second:", secondNum)
		//fmt.Println("Answer:", answer)
	}
	return answer
}
