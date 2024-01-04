/*
Copyright © 2024 Peter Leung
*/
package cmd

import (
	"aoc/cmd/year2023"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code",
	Long:  `aoc is a command line utility to show the answers of the solutions to Advent of Code`,
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
	rootCmd.AddCommand(year2023.Cmd)

}
