package main

import (
	"fmt"

	"github.com/ryntak94/advent-of-code-go-starter/utils"
)

func main() {
	file, err := utils.FileAsString("input.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("input scanner: %w", err))
	}
	fmt.Println(file)
}
