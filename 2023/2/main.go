package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	GameID   int
	Sets     []Cubes
	Possible bool
}

type Cubes struct {
	red   int
	green int
	blue  int
}

func ParseFile(filename string) (data []Game, err error) {
	return
}

func GetCubeCount(cubeString string, color string) (count int, err error) {
	countStr, _ := strings.CutSuffix(cubeString, " "+color)
	count, atoiErr := strconv.Atoi(countStr)
	if atoiErr != nil {
		err = fmt.Errorf("convert %s cube count err - %s", color, atoiErr.Error())
		return
	}
	return
}

func IsPossibleSet(currentSet Cubes, maxCubes Cubes) (possible bool) {
	if currentSet.red > maxCubes.red {
		return false
	}
	if currentSet.green > maxCubes.green {
		return false
	}
	if currentSet.blue > maxCubes.blue {
		return false
	}
	return true
}

func PartOne(filename string, maxCubes Cubes) (ans int, err error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("fail to open file %s\n", filename)
		return
	}
	str := string(fileBytes)
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		gameInfo, gameSets, foundColon := strings.Cut(line, ": ")
		if !foundColon {
			err = errors.New("unexpected syntax - no colon in line")
			return
		}

		// could be using index+1 as id but just for fun and parse from line
		gameId, found := strings.CutPrefix(gameInfo, "Game ")
		if !found {
			err = errors.New("unexpected syntax - game ID")
			return
		}
		id, atoiErr := strconv.Atoi(gameId)
		if atoiErr != nil {
			err = errors.New("unexpected syntax - game ID not int")
			return
		}
		setsString := strings.Split(gameSets, "; ")

		game := Game{
			GameID:   id,
			Sets:     []Cubes{},
			Possible: true,
		}
		for _, set := range setsString {
			cubesStrings := strings.Split(set, ", ")
			for _, cubeString := range cubesStrings {
				currentSet := Cubes{}
				if strings.Contains(cubeString, "red") {
					count, countErr := GetCubeCount(cubeString, "red")
					if countErr != nil {
						err = countErr
						return
					}
					currentSet.red = count
				} else if strings.Contains(cubeString, "green") {
					count, countErr := GetCubeCount(cubeString, "green")
					if countErr != nil {
						err = countErr
						return
					}
					currentSet.green = count
				} else if strings.Contains(cubeString, "blue") {
					count, countErr := GetCubeCount(cubeString, "blue")
					if countErr != nil {
						err = countErr
						return
					}
					currentSet.blue = count
				}

				game.Sets = append(game.Sets, currentSet)
				possible := IsPossibleSet(currentSet, maxCubes)
				if !possible {
					game.Possible = false
				}
			}
		}
		if game.Possible == true {
			ans += game.GameID
		}
	}
	return
}

func main() {

	ans, err := PartOne("input", Cubes{
		red:   12,
		green: 13,
		blue:  14,
	})
	if err != nil {
		fmt.Println("Part one error: ", err)
	}
	fmt.Println("Part one answer:", ans)
}
