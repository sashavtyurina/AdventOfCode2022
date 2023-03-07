package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func judgePlay(round string) int {

	play := strings.Split(round, " ")
	var mine string = play[1]

	var score int = 0
	switch mine {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	if (round == "A X") || (round == "B Y") || (round == "C Z") {
		score += 3 // draw
		return score
	}

	if (round == "A Y") || (round == "B Z") || (round == "C X") {
		score += 6 // win
		return score
	}

	return score

}

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var cumScore int = 0
	for fileScanner.Scan() {
		var nextRound string = fileScanner.Text()
		nextRound = strings.TrimSpace(nextRound)
		cumScore += judgePlay(nextRound)
	}

	fmt.Println(cumScore)
}
