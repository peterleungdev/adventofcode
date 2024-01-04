package year2023

import (
	"aoc/cmd/year2023/day01"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2023",
	Short: "Advent of Code 2023",
	Long:  `Advent of Code 2023`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() {
// 	err := Cmd.Execute()
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }

func init() {
	Cmd.AddCommand(day01.Cmd)
}
