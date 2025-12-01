package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

// Part 1: 42m 16s ; Part 2: 36m 02s

func main() {
	pages_file, err := utils.FileAsString("input_pages.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	orders_file, err := utils.FileAsString("input_order.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	orders := parseStringToOrders(orders_file)
	updates := parseStringTo2DArray(pages_file)
	results := checkUpdates(updates, orders)
	rightUpdates := filterUpdates(updates, results, true)
	fmt.Println(addMiddlePages(rightUpdates))
	wrongUpdates := filterUpdates(updates, results, false)
	fixedUpdates := fixAllWrongUpdates(wrongUpdates, orders)
	fmt.Println(addMiddlePages(fixedUpdates))
}

func addMiddlePages(updates [][]int) int {
	middlePagesSum := 0
	for i := 0; i < len(updates); i++ {
		middlePagesSum += updates[i][len(updates[i])/2]

	}
	return middlePagesSum

}

func fixAllWrongUpdates(updates [][]int, orders [][]int) [][]int {
	var fixedUpdates = [][]int{}
	for i := 0; i < len(updates); i++ {
		fixedUpdates = append(fixedUpdates, fixWrongUpdate(updates[i], orders))
	}
	return fixedUpdates
}

func fixWrongUpdate(update []int, orders [][]int) []int {
	correctUpdate := update
	isCorrect, firstIndex, secondIndex := isCorrectUpdate(update, orders)
	for isCorrect == false {
		correctUpdate[firstIndex], correctUpdate[secondIndex] = correctUpdate[secondIndex], correctUpdate[firstIndex]
		isCorrect, firstIndex, secondIndex = isCorrectUpdate(correctUpdate, orders)
	}
	return correctUpdate
}

func filterUpdates(updates [][]int, results []bool, right bool) [][]int {
	var wrongUpdates = [][]int{}
	for i := 0; i < len(updates); i++ {
		if results[i] == right {
			wrongUpdates = append(wrongUpdates, updates[i])
		}
	}
	return wrongUpdates
}

func checkUpdates(updates [][]int, orders [][]int) []bool {
	var results []bool
	for i := 0; i < len(updates); i++ {
		isCorrect, _, _ := isCorrectUpdate(updates[i], orders)
		results = append(results, isCorrect)
	}
	return results
}

func isCorrectUpdate(update []int, orders [][]int) (bool, int, int) {

	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if isCorrectOrder(update[i], update[j], orders) == false {
				return false, i, j
			}
		}
	}
	return true, -1, -1
}

func isCorrectOrder(first, second int, orders [][]int) bool {
	for i := 0; i < len(orders); i++ {
		if orders[i][0] == second && orders[i][1] == first {
			return false
		}
	}
	return true
}

func parseStringToOrders(input string) [][]int {
	var orders [][]int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		temp := strings.Split(line, "|")
		var order []int
		for _, orderString := range temp {
			num, err := strconv.Atoi(orderString)
			if err != nil {
				fmt.Println(fmt.Errorf("parsing int: %w", err))
			}
			order = append(order, num)
		}
		orders = append(orders, order)
	}
	return orders
}

func parseStringTo2DArray(input string) [][]int {
	var list [][]int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		temp := strings.Split(line, ",")
		var update []int
		for _, page := range temp {
			num, err := strconv.Atoi(page)
			if err != nil {
				fmt.Println(fmt.Errorf("parsing int: %w", err))
			}
			update = append(update, num)
		}
		list = append(list, update)
	}
	return list
}
