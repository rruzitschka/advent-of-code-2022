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

func splitBagContent(sourceString string) []string {

	// split sourceString in half - number of characters needd to be even

	l := len(sourceString)

	half1 := sourceString[0 : l/2]
	half2 := sourceString[l/2 : l]

	split := []string{half1, half2}

	return split

}

func findSameCharacterInStringList(stringList []string, characterList string) string {

	// this function checks if any of the characters that are part of characterList can be found in the slice of strings
	// stringList. The first character from the CharacterList that si found in all strings of the slice is returned.
	//characterString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var isInAll bool = false

	for i := 0; i < len(characterList); i++ {
		s := string(characterList[i])
		isInAll = true
		for _, st := range stringList {
			if !strings.ContainsAny(st, s) {
				isInAll = false
				break //no need to check any further
			}
		}
		if isInAll {
			return s
		}
	}
	return ""
}

func findPriority(c string) int {

	//assigns prio value based on position in characterString

	characterString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(characterString, c) + 1
}

func main() {
	characterString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sameCharacter string
	inputList := readInput("/Users/roru/Desktop/inputDay3.txt")

	// part 1
	var totalPriority int

	for _, s := range inputList {
		sp := splitBagContent(s)
		sameCharacter = findSameCharacterInStringList([]string{sp[0], sp[1]}, characterString)
		totalPriority += findPriority(sameCharacter)
	}

	fmt.Printf("Part 1: The total priority is %d \n", totalPriority)

	// part 1
	totalPriority = 0
	s := len(inputList)

	for i := 0; i < s; i += 3 {

		sameCharacter = findSameCharacterInStringList([]string{inputList[i], inputList[i+1], inputList[i+2]}, characterString)
		totalPriority += findPriority(sameCharacter)
	}

	fmt.Printf("Part 2: The total priority is %d \n", totalPriority)
}
