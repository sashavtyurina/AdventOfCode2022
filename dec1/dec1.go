package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var maxCalories, cumulCalories int = 0, 0

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			if cumulCalories > maxCalories {
				maxCalories = cumulCalories
			}
			cumulCalories = 0
		} else {
			calories, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			cumulCalories += calories
		}
	}
	fmt.Println(maxCalories)
}
