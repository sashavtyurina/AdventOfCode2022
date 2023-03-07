package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(a [3]int) (minEl int) {
	minEl = a[0]

	for _, value := range a {
		if value < minEl {
			minEl = value
		}
	}
	return minEl
}

func replaceArrayElement(arr [3]int, oldElement int, newElement int) (updatedArr [3]int) {
	// fmt.Println("Replace %d with %d", oldElement, newElement)
	// fmt.Println("---")
	// fmt.Println(oldElement)
	// fmt.Println(newElement)
	// fmt.Println(arr)
	updatedArr = arr

	for ind, value := range arr {
		// fmt.Println(value)
		if value == oldElement {
			updatedArr[ind] = newElement
			// fmt.Println("A: replaced with new element")
			// fmt.Println(updatedArr)
			// fmt.Println("---")
			return updatedArr
		}
		// else {
		// 	fmt.Println("Value and old element do not match")
		// 	fmt.Println(value)
		// 	fmt.Println(oldElement)
		// }
	}
	return updatedArr

}

func main() {
	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("This is Arg0: %s,\nThis is Arg1: %s", os.Args[0], os.Args[1])
	// fmt.Println(os.Args[1])
	threeTopCals := [3]int{0, 0, 0}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var cumulCalories int = 0

	for fileScanner.Scan() {
		// fmt.Println(threeTopCals)
		if fileScanner.Text() == "" {
			var minCals = min(threeTopCals)
			if cumulCalories > minCals {
				threeTopCals = replaceArrayElement(threeTopCals, minCals, cumulCalories)
			}
			cumulCalories = 0
		} else {
			var calories int = 0
			calories, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			cumulCalories += calories
		}

	}

	fmt.Println("result is")
	fmt.Println(threeTopCals)
	fmt.Println(threeTopCals[0] + threeTopCals[1] + threeTopCals[2])
}
