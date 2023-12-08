// https://adventofcode.com/2023/day/1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindDigits(line string) (first, last, firstIndex, lastIndex int) {
	first = -1
	last = -1
	for index, r := range line {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			continue
		}
		if first == -1 {
			first, last = digit, digit
			firstIndex, lastIndex = index, index
		} else {
			last = digit
			lastIndex = index
		}
	}
	return
}

func FindDigitLetters(line string) (first, last, firstIndex, lastIndex int) {
	numFirstIndexMap := make(map[int]int)
	numLastIndexMap := make(map[int]int)
	letters := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for letterIndex, letter := range letters {
		index := strings.Index(line, letter)
		if index != -1 {
			numFirstIndexMap[letterIndex+1] = index
		}
		index = strings.LastIndex(line, letter)
		if index != -1 {
			numLastIndexMap[letterIndex+1] = index
		}
	}
	first = -1
	last = -1
	firstIndex = len(line)
	lastIndex = 0
	for num, index := range numFirstIndexMap {
		if first == -1 || (first != -1 && index < firstIndex) {
			firstIndex = index
			first = num
		}
	}
	for num, index := range numLastIndexMap {
		if last == -1 || (last != -1 && index > lastIndex) {
			lastIndex = index
			last = num
		}
	}
	return
}

// Find sum by number digit only
func PartOne(filename string) (sum int) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		first, last, _, _ := FindDigits(line)
		sum += first*10 + last
	}

	return sum
}

// Find sum by number digit and letter digit
func PartTwo(filename string) (sum int) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lineCount := 0
	for scanner.Scan() {
		lineCount += 1
		line := scanner.Text()

		firstDigit, lastDigit, firstDigitIndex, lastDigitIndex := FindDigits(line)
		firstLetter, lastLetter, firstLetterIndex, lastLetterIndex := FindDigitLetters(line)

		indexNumMap := make(map[int]int)
		if firstDigit != -1 {
			indexNumMap[firstDigitIndex] = firstDigit
		}
		if lastDigit != -1 {
			indexNumMap[lastDigitIndex] = lastDigit
		}
		if firstLetter != -1 {
			indexNumMap[firstLetterIndex] = firstLetter
		}
		if lastLetter != -1 {
			indexNumMap[lastLetterIndex] = lastLetter
		}

		var first, last int = -1, -1
		var currentFirstIndex int = len(line)
		var currentLastIndex int = -1
		for index, num := range indexNumMap {
			if index < currentFirstIndex {
				first = num
				currentFirstIndex = index
			}
			if index > currentLastIndex {
				last = num
				currentLastIndex = index
			}
		}

		sum += first*10 + last
	}

	return sum
}

func main() {
	const filename = "day_01_input"

	partOne := PartOne(filename)
	partTwo := PartTwo(filename)
	fmt.Println(partOne)
	fmt.Println(partTwo)
}
