package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func main() {
	input, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	given := "125 17"

	givenArray := parseString(given)
	inputArray := parseString(input)

	fmt.Println(blink(givenArray, 25, len(givenArray)))
	fmt.Println(blink(inputArray, 75, len(inputArray)))
}

func blink(input []int, times int, inputLength int) int {
	var wg sync.WaitGroup
	timeNow := time.Now()
	stones := []int{}

	var mu sync.Mutex
	for i := 0; i < inputLength; i++ {
		wg.Add(1)
		startTime := time.Now()
		go func(id int) {
			localStones := []int{input[i]}
			for j := 0; j < times; j++ {
				iterationStart := time.Now()
				newStones := []int{}
				for _, stone := range localStones {
					newStones = append(newStones, blinkOnce(stone)...)
				}
				localStones = newStones

				duration := time.Since(iterationStart)
				fmt.Printf("Goroutine %d, iteration %d took %v\n", id, j+1, duration)
			}
			mu.Lock()
			stones = append(stones, localStones...)
			mu.Unlock()
			wg.Done()
			fmt.Printf("Goroutine %d completed in %v\n", id, time.Since(startTime))
		}(i)
	}
	wg.Wait()
	fmt.Println("total time taken:", time.Since(timeNow))
	return len(stones)
}

func blinkOnce(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	stoneStr := strconv.Itoa(stone)
	if len(stoneStr)%2 == 0 {
		mid := len(stoneStr) / 2
		left, err := strconv.Atoi(stoneStr[:mid])
		if err != nil {
			fmt.Println("Error parsing string:", err)
		}
		right, err := strconv.Atoi(stoneStr[mid:])
		if err != nil {
			fmt.Println("Error parsing string:", err)
		}
		return []int{left, right}
	}
	return []int{stone * 2024}
}

func parseString(input string) []int {
	list := strings.Split(input, " ")
	intList := make([]int, len(list))
	for i := range list {
		number, err := strconv.Atoi(list[i])
		if err != nil {
			fmt.Println("Error parsing string:", err)
		}
		intList[i] = number
	}
	return intList
}
