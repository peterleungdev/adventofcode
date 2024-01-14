package day02

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Game struct {
	GameID   int
	Sets     []Cubes
	MaxCubes Cubes
	Possible bool
}

type Cubes struct {
	red   int
	green int
	blue  int
}

var Cmd = &cobra.Command{
	Use:   "day02",
	Short: "day02",
	Long:  `day02`,
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

	partOne, partTwo, err := PartOneTwo(string(b), Cubes{
		red:   12,
		green: 13,
		blue:  14,
	})
	if err != nil {
		fmt.Println("Part one error: ", err)
	}

	logrus.Infof("part 1: %d", partOne)
	logrus.Infof("part 2: %d", partTwo)
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

func PartOneTwo(s string, maxCubes Cubes) (sumOfPossibleGameID int, sumOfPower int, err error) {

	lines := strings.Split(s, "\n")

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
			MaxCubes: Cubes{},
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

				if currentSet.red > game.MaxCubes.red {
					game.MaxCubes.red = currentSet.red
				}
				if currentSet.green > game.MaxCubes.green {
					game.MaxCubes.green = currentSet.green
				}
				if currentSet.blue > game.MaxCubes.blue {
					game.MaxCubes.blue = currentSet.blue
				}
			}
		}
		if game.Possible == true {
			sumOfPossibleGameID += game.GameID
		}
		power := game.MaxCubes.red * game.MaxCubes.green * game.MaxCubes.blue
		sumOfPower += power
	}
	return
}
