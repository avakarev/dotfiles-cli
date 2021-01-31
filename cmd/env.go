package main

import (
	"github.com/spf13/cobra"

	"github.com/avakarev/dotfiles-cli/internal/config"
	"github.com/avakarev/dotfiles-cli/internal/printer"
)

// envCmd implements the `env` command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Prints dotfiles env",
	Run: func(cmd *cobra.Command, args []string) {
		p := printer.Default()
		p.Addln("HomeDir", config.HomeDir)
		p.Addln("WorkingDir", config.WorkingDir)
		p.Addln("ConfigDir", config.ConfigDir)
		p.Addln("ConfigFile", config.ConfigFile)
		p.Flush()
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
