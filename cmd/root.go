package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Utility to manage dotfiles symlinks",
}

func main() {
	Execute()
}

// Execute adds all child commands to the root command
func Execute() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
