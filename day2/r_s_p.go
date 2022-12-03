package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) []string {

	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines

}

func mapTool(t string) string {
	if t == "A" || t == "X" {
		return "R"
	}

	if t == "B" || t == "Y" {
		return "P"
	}

	if t == "C" || t == "Z" {
		return "S"
	}

	return ""
}

func RoundScore(p1, p2 string) int {

	playScore := winRound(p1, p2)
	toolScore := 0

	switch p2 {
	case "R":
		toolScore = 1
	case "P":
		toolScore = 2
	case "S":
		toolScore = 3
	}


	return playScore + toolScore
}

func winRound(p1, p2 string) int {
	if p1 == p2 {
		return 3
	}

	if p1 == "S" {
		if p2 == "R" {
			return 6
		}
		if p2 == "P" {
			return 0
		}
	}

	if p1 == "R" {
		if p2 == "P" {
			return 6
		}
		if p2 == "S" {
			return 0
		}
	}

	if p1 == "P" {
		if p2 == "R" {
			return 0
		}
		if p2 == "S" {
			return 6
		}
	}

	return 0
}

func player2ToolNewStrategy(player1, strategy string) string {

	if strategy == "Y" {
		return player1
	}

	if strategy == "X" {
		if player1 == "R" {
			return "S"
		}
		if player1 == "S" {
			return "P"
		}
		if player1 == "P" {
			return "R"
		}

	}

	if strategy == "Z" {
		if player1 == "R" {
			return "P"
		}
		if player1 == "S" {
			return "R"
		}
		if player1 == "P" {
			return "S"
		}

	}

	return ""

}

func main() {

	resultsList := readInput("/Users/wzhroru/Desktop/input2.txt")
	roundScore := 0
	totalScore := 0

	for i, results := range resultsList {

		input := strings.Split(results, " ")
		player1 := mapTool(input[0])
		player2 := player2ToolNewStrategy(player1, input[1])
		roundScore = RoundScore(player1, player2)
		totalScore += roundScore
		fmt.Printf("Roundscore = %d for round %d\n", roundScore, i)

	}

	fmt.Printf("The total score after all rounds is %d! \n", totalScore)

}
