package main

import (
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/avakarev/dotfiles-cli/internal/buildmeta"
)

// versionCmd implements the `version` command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints dotfiles version and build information",
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(log.Writer(), 0, 0, 2, ' ', 0)
		f := "%s\t%s\t\n"

		fmt.Fprintf(w, f, "GitCommit", buildmeta.GitCommit)
		fmt.Fprintf(w, f, "BuildDate", buildmeta.BuildDate)
		fmt.Fprintf(w, f, "Version", buildmeta.Version)
		fmt.Fprintf(w, f, "Compiler", buildmeta.Compiler)

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
