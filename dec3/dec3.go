package main

import (
	"bufio"
	"fmt"
	"os"
)

func splitIntoCompartments(rucksack string) (compartments [2]string) {
	var compartmentLength int = len(rucksack) / 2
	var comp1 string = rucksack[:compartmentLength]
	var comp2 string = rucksack[compartmentLength:]
	// var compartments [2]string = [comp1, comp2]  -- not the right syntax?
	compartments = [2]string{comp1, comp2}
	return compartments
}

func findRepeatedElement(compartments [2]string) (repeatedElement rune) {
	var elementCounter = make(map[rune]int)

	var comp1, comp2 string = compartments[0], compartments[1]

	for _, c := range comp1 {
		elementCounter[c] = 1
	}
	for _, c := range comp2 {
		_, exists := elementCounter[c]
		if exists {
			return c
		}
	}
	// assuming that a repeated element always exists
	// if not, return zero rune value
	var c rune = 0
	return c
}

func getPriority(c rune) (priority int) {
	// var priority int = 0
	if (96 <= int(c)) && (int(c) <= 122) {
		priority = int(c) - 96
	} else {
		priority = int(c) - 38
	}
	return
}

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var prioritySum int = 0
	for fileScanner.Scan() {
		var rucksack string = fileScanner.Text()
		//var comps string[] = splitIntoCompartments(rucksack)
		comps := splitIntoCompartments(rucksack)
		fmt.Println(comps)
		var repeatedElement rune = findRepeatedElement(comps)
		fmt.Println(string(repeatedElement))
		fmt.Println(repeatedElement)
		var priority int = getPriority(repeatedElement)
		fmt.Println(priority)
		prioritySum += priority
	}
	fmt.Println(prioritySum)
}
