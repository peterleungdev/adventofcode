package day03

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Data struct {
	Lines []string
}

type Coordinate struct {
	lineIndex int
	charIndex int
}

var Cmd = &cobra.Command{
	Use:   "day03",
	Short: "day03",
	Long:  `day03`,
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

	lines := strings.Split(string(b), "\n")

	partOne, partTwo, err := PartOneTwo(lines)
	if err != nil {
		fmt.Println("Error - ", err)
	}

	logrus.Infof("part 1: %d", partOne)
	logrus.Infof("part 2: %d", partTwo)
}

func PartOneTwo(lines []string) (sumOfPartNumber int, sumOfGearRatio int, err error) {

	starMap := make(map[Coordinate][]int)

	for lineIndex, line := range lines {
		if len(line) == 0 {
			continue
		}
		currentNumber := ""
		readingNumber := false
		startIndex := -1
		endIndex := -1

		for charIndex := 0; ; charIndex++ {

			if charIndex >= len(line) {
				if readingNumber {
					isPartNumber, number := IsPartNumber(lines, currentNumber, lineIndex, startIndex, endIndex, starMap)
					if isPartNumber {
						sumOfPartNumber += number // sum of all of the part numbers in the engine schematic
					}
				}
				break
			}

			char := line[charIndex]
			if IsNumberChar(char) {
				currentNumber = currentNumber + string(char)
				if !readingNumber {
					readingNumber = true
					startIndex = charIndex
					endIndex = charIndex
				} else {
					endIndex = charIndex
				}
			} else {
				if readingNumber {
					isPartNumber, number := IsPartNumber(lines, currentNumber, lineIndex, startIndex, endIndex, starMap)
					if isPartNumber {
						sumOfPartNumber += number // sum of all of the part numbers in the engine schematic
					}
				}
				// reset state
				currentNumber = ""
				readingNumber = false
				startIndex = -1
				endIndex = -1
				continue
			}

		}
	}

	for _, numArr := range starMap {
		if len(numArr) == 2 {
			sumOfGearRatio += numArr[0] * numArr[1]
		}
	}
	return
}

func IsPartNumber(lines []string, currentNumber string, lineIndex int, charStartIndex int, charEndIndex int, starMap map[Coordinate][]int) (isPartNumber bool, number int) {
	number, atoiErr := strconv.Atoi(currentNumber)
	if atoiErr != nil {
		return
	}

	lineCheckStartIndex := max(lineIndex-1, 0)
	lineCheckEndIndex := min(len(lines)-1, lineIndex+1)

	for lineCheckIndex := lineCheckStartIndex; lineCheckIndex <= lineCheckEndIndex; lineCheckIndex++ {
		line := lines[lineCheckIndex]
		checkStartIndex := max(charStartIndex-1, 0)
		checkEndIndex := min(len(line)-1, charEndIndex+1)

		for i := checkStartIndex; i <= checkEndIndex; i++ {
			char := line[i]
			if IsSymbol(char) {
				isPartNumber = true
				if char == '*' {
					if (starMap[Coordinate{
						lineIndex: lineCheckIndex,
						charIndex: i,
					}] == nil) {
						starMap[Coordinate{
							lineIndex: lineCheckIndex,
							charIndex: i,
						}] = []int{}
					}
					starMap[Coordinate{
						lineIndex: lineCheckIndex,
						charIndex: i,
					}] = append(starMap[Coordinate{
						lineIndex: lineCheckIndex,
						charIndex: i,
					}], number)
				}
			} else {
				continue
			}
		}
	}

	return
}

func IsNumberChar(char byte) bool {
	return (char >= '0' && char <= '9')
}

func IsSymbol(char byte) bool {
	if (char >= '0' && char <= '9') || char == '.' {
		return false
	}
	return true
}

func ParseFile(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	lines = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}
