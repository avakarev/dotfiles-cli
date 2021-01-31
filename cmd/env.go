package main

import (
	"log"

	"github.com/spf13/cobra"
)

// envCmd implements the `env` command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Prints dotfiles env",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("@TODO: implement `env` command")
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
