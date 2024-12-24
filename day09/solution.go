package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func main() {
	file, err := utils.FileAsString("given.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}

	var fileLadies []string
	var mtLadies []string
	for i, char := range strings.Split(file, "") {
		if i%2 == 0 {
			fileLadies = append(fileLadies, char)
		} else {
			mtLadies = append(mtLadies, char)
		}
	}

	list := createList(fileLadies, mtLadies)

	// partOne := partOne(list)
	partTwo := partTwo(list, len(fileLadies))
	// fmt.Println(calculateChecksum(partOne))
	fmt.Println(partTwo)
}
func partTwo(list []string, numsCount int) []string {
	start := 0
	end := len(list) - 1
	fmt.Println(list, start, end)
	lenEnd := 0
	for n := 0; n < numsCount; n++ {
		if start >= end {
			start = 0
			end = end - lenEnd
			if end < 0 {
				break
			}
			continue
		}
		for start < end && list[start] != "." {
			start++
		}
		for start < end && list[end] == "." {
			end--
		}
		if start >= end {
			continue
		}
		// Identify the number at the end
		number := list[end]
		lenEnd = 0

		// Calculate the length of the number sequence at the end
		for i := end; i >= 0; i-- {
			if list[i] != number {
				break
			}
			lenEnd++
		}

		// Try placing the number in the first block that fits
		placed := false
		for start < end {
			// Move start to the next available dot block
			for start < end && list[start] != "." {
				start++
			}

			// Calculate the length of the dot sequence at start
			lenStart := 0
			for i := start; i < len(list); i++ {
				if list[i] != "." {
					break
				}
				lenStart++
			}

			// Check if the current block can fit the number
			if lenStart >= lenEnd {
				for i := 0; i < lenEnd; i++ {
					list[start+i] = number
					list[end-i] = "."
				}
				fmt.Printf("Placed %s at position %d, updated list: %v\n", number, start, list)
				start += lenEnd
				end -= lenEnd
				placed = true
				break
			} else {
				// Move start to the next block
				start += lenStart
			}
		}

		// Skip the number if it couldn't be placed
		if !placed {
			fmt.Printf("Could not place %s, skipping it.\n", number)
			end -= lenEnd
		}

		fmt.Println("Current list:", list)
	}

	return list
}

func partOne(list []string) []string {
	start := 0
	end := len(list) - 1

	for start < end {
		if list[start] != "." {
			start++
		} else if list[end] == "." {
			end--
		} else {
			list[start], list[end] = list[end], list[start]
			start++
			end--
		}
	}
	return list
}

func createList(fileLadies, mtLadies []string) []string {
	var index int = 0
	var result []string
	for len(fileLadies) > 0 && len(mtLadies) > 0 {
		count, err := strconv.Atoi(fileLadies[0])
		if err != nil {
			fmt.Println(fmt.Errorf("input scanner: %w", err))
		}

		for j := 0; j < count; j++ {
			convIndex := strconv.Itoa(index)
			result = append(result, convIndex)
		}
		index++
		fileLadies = fileLadies[1:]

		mtCount, err := strconv.Atoi(mtLadies[0])
		if err != nil {
			fmt.Println(fmt.Errorf("input scanner: %w", err))
		}

		for j := 0; j < mtCount; j++ {
			result = append(result, ".")
		}

		mtLadies = mtLadies[1:]
	}
	for len(fileLadies) > 0 {
		count, err := strconv.Atoi(fileLadies[0])
		if err != nil {
			fmt.Println(fmt.Errorf("input scanner: %w", err))
		}

		for j := 0; j < count; j++ {
			convIndex := strconv.Itoa(index)
			result = append(result, convIndex)
		}
		index++
		fileLadies = fileLadies[1:]
	}
	for len(mtLadies) > 0 {
		mtCount, err := strconv.Atoi(mtLadies[0])
		if err != nil {
			fmt.Println(fmt.Errorf("input scanner: %w", err))
		}

		for j := 0; j < mtCount; j++ {
			result = append(result, ".")
		}

		mtLadies = mtLadies[1:]
	}

	return result
}

func calculateChecksum(nums []string) int {
	result := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] == "." {
			continue
		}
		num, err := strconv.Atoi(nums[index])
		if err != nil {
			fmt.Println(fmt.Errorf("input scanner: %w", err))
		}
		result += num * index
	}
	return result
}
