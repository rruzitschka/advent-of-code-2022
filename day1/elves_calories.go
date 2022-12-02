package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func sliceCaloriesList(calList []string) []int {

	result := []int{}
	var i_result = 0
	l := len(calList)
	for i, cal := range calList {
		if v, err := strconv.Atoi(cal); err == nil {
			i_result += v
		}
		if cal == "" || i == l-1 {
			result = append(result, i_result)
			i_result = 0
		}

	}

	return result
}

func maxCalories(computedCalList []int) int {
	var max = 0

	for _, cal := range computedCalList {
		fmt.Printf("Calorie value %d \n", cal)
		if cal > max {
			max = cal
		}
	}

	return max
}

func topThreeTotalCalories(calList []int) int {

	sort.Ints(calList)

	l := len(calList)

	if l >= 3 {
		return calList[l-1] + calList[l-2] + calList[l-3]
	}

	return 0

}

func main() {

	//elvesList := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	elvesList := readInput("/Users/wzhroru/Desktop/input.txt")

	fmt.Printf("The maximum amount of claories carried by an elf is %d! \n", maxCalories((sliceCaloriesList((elvesList)))))

	fmt.Printf("The top three calories combined are %d \n", topThreeTotalCalories((sliceCaloriesList(elvesList))))

}
