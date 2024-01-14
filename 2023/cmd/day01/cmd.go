// https://adventofcode.com/2023/day/1

package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day01",
	Short: "day01",
	Long:  `day01`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	filename := fmt.Sprintf("cmd/%s/input.txt", command)
	b, err := os.ReadFile(filename)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("part 1: %d", PartOne(string(b)))
	logrus.Infof("part 2: %d", PartTwo(string(b)))
}

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
func PartOne(s string) (ans int64) {

	lines := strings.Split(s, "\n")
	var sum int

	for _, line := range lines {
		if line == "" {
			continue
		}
		first, last, _, _ := FindDigits(line)
		sum += first*10 + last
	}

	return int64(sum)
}

// Find sum by number digit and letter digit
func PartTwo(s string) (ans int64) {

	lines := strings.Split(s, "\n")

	var sum int
	for _, line := range lines {

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

	return int64(sum)
}
