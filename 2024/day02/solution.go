package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	reports := parseStringToReportsLists(file)
	fmt.Println("safe reports:", countSafeReports(reports))
	fmt.Println("safe reports after using module:", countSafeReportsAfterUsingModule(reports))
}

func countSafeReports(reports [][]int) int {
	var count int
	for _, report := range reports {
		isSafe := isSafeReport(report)
		if isSafe {
			count++
		}
	}
	return count
}

func countSafeReportsAfterUsingModule(reports [][]int) int {
	var count int
	for _, report := range reports {
		isSafe := isSafeReport(report)
		if isSafe {
			count++
		} else {
			for i := 0; i < len(report); i++ {
				removedLevelList := removeElement(report, i)
				isSafe := isSafeReport(removedLevelList)
				if isSafe {
					count++
					break
				}
			}
		}
	}
	return count
}

func removeElement(list []int, i int) []int {
	removedLevelList := make([]int, 0)
	for j := 0; j < len(list); j++ {
		if j != i {
			removedLevelList = append(removedLevelList, list[j])
		}
	}
	return removedLevelList
}
func isSafeReport(report []int) bool {
	if len(report) == 1 {
		return true
	}
	isIncreasing := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		first := report[i]
		second := report[i+1]
		difference := math.Abs(float64(second - first))
		if isIncreasing && first >= second {
			return false
		}
		if !isIncreasing && first <= second {
			return false
		}
		if difference < 1 {
			return false
		}
		if difference > 3 {
			return false
		}
	}
	return true
}
func parseStringToReportsLists(input string) [][]int {
	var list [][]int
	reports := strings.Split(input, "\n")
	for _, report := range reports {
		levels := strings.Split(report, " ")
		levelsList := make([]int, 0)
		for _, level := range levels {
			// parsing int: strconv.Atoi: parsing "1\r": invalid syntax
			level = strings.TrimSuffix(level, "\r")
			num, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println(fmt.Errorf("parsing int: %w", err))
			}
			levelsList = append(levelsList, num)
		}

		list = append(list, levelsList)
	}
	return list
}
