package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/avakarev/dotfiles-cli/internal/config"
)

// envCmd implements the `env` command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Prints dotfiles env",
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(log.Writer(), 0, 0, 2, ' ', 0)
		f := "%s\t%s\t\n"

		fmt.Fprintf(w, f, "HomeDir", config.HomeDir)
		fmt.Fprintf(w, f, "WorkingDir", config.WorkingDir)
		fmt.Fprintf(w, f, "ConfigDir", config.ConfigDir)
		fmt.Fprintf(w, f, "ConfigFile", config.ConfigFile)

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
