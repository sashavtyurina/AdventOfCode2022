package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var win = map[string]string{"A": "Y", "B": "Z", "C": "X"}
// var draw = map[string]string{"A": "X", "B": "Y", "C": "Z"}
// var lose = map[string]string{"A": "Z", "B": "X", "C": "Y"}

var winScore = map[string]int{"A": 2, "B": 3, "C": 1}
var drawScore = map[string]int{"A": 1, "B": 2, "C": 3}
var loseScore = map[string]int{"A": 3, "B": 1, "C": 2}

func judgePlay(round string) int {

	play := strings.Split(round, " ")
	var opponent, mine string = play[0], play[1]
	var score int = 0
	switch mine {
	case "X":
		score += 0 // lose
		score += loseScore[opponent]
		return score
	case "Y":
		score += 3 // draw
		score += drawScore[opponent]
		return score
	case "Z":
		score += 6 // win
		score += winScore[opponent]
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
