package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	ID             int
	WinningNumbers []int
	YourNumbers    []int
	InstanceCount  int
}

const filename string = "input"

// solution of https://adventofcode.com/2023/day/3
func main() {
	lines, err := ParseFile(filename)
	if err != nil {
		return
	}

	cards, err := ParseData(lines)
	if err != nil {
		fmt.Println(err)
	}

	points := PartOne(cards)
	fmt.Println("Part One:", points)

	cardCount := PartTwo(cards)
	fmt.Println("Part Two:", cardCount)
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

func ParseData(lines []string) (cards []Card, err error) {
	for lineIndex, line := range lines {
		_, cardDataStr, _ := strings.Cut(line, ": ")
		currentCard := Card{
			ID:             lineIndex + 1,
			WinningNumbers: []int{},
			YourNumbers:    []int{},
			InstanceCount:  1,
		}
		winningNumRaw, yourNumRaw, _ := strings.Cut(cardDataStr, "| ")
		winningNumStrArr := strings.Split(winningNumRaw, " ")
		yourNumStrArr := strings.Split(yourNumRaw, " ")

		for _, numStr := range winningNumStrArr {
			if numStr == "" {
				continue
			}
			num, atoiErr := strconv.Atoi(numStr)
			if atoiErr != nil {
				err = atoiErr
				return
			}
			currentCard.WinningNumbers = append(currentCard.WinningNumbers, num)
		}

		for _, numStr := range yourNumStrArr {
			if numStr == "" {
				continue
			}
			num, atoiErr := strconv.Atoi(numStr)
			if atoiErr != nil {
				err = atoiErr
				return
			}
			currentCard.YourNumbers = append(currentCard.YourNumbers, num)
		}

		cards = append(cards, currentCard)
	}
	return
}

func PartOne(cards []Card) (totalPoint int) {
	for _, card := range cards {
		cardPoint := 0
		for _, yourNum := range card.YourNumbers {
			found := slices.Contains[[]int](card.WinningNumbers, yourNum)
			if found {
				if cardPoint == 0 {
					cardPoint = 1
				} else {
					cardPoint *= 2
				}
			}
		}
		totalPoint += cardPoint
	}
	return
}

func PartTwo(cards []Card) (cardCount int) {
	cardsLen := len(cards)
	for cardIndex, card := range cards {
		matched := 0
		for _, yourNum := range card.YourNumbers {
			found := slices.Contains[[]int](card.WinningNumbers, yourNum)
			if found {
				matched += 1
			}
		}
		for i := 1; i <= matched; i++ {
			if cardIndex+i < cardsLen {
				cards[cardIndex+i].InstanceCount += 1 * card.InstanceCount
			}
		}
		cardCount += card.InstanceCount
	}

	return
}
