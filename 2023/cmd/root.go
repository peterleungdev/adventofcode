/*
Copyright © 2024 Peter Leung
*/
package cmd

import (
	"aoc2023/cmd/day01"
	"aoc2023/cmd/day02"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc2023",
	Short: "Advent of Code 2023",
	Long:  `aoc is a command line utility to show the answers of the solutions to Advent of Code 2023`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(day01.Cmd)
	rootCmd.AddCommand(day02.Cmd)
}
