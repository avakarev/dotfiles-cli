package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/avakarev/dotfiles-cli/internal/op"
	"github.com/avakarev/dotfiles-cli/pkg/dotfiles"
)

var statusCmdGroups bool

// statusCmd implements the `status` command
var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"st"},
	Short:   "Prints dotfiles status",
	Run: func(cmd *cobra.Command, args []string) {
		dfiles, err := dotfiles.New()
		if err != nil {
			log.Fatalln(err)
		}

		for _, g := range dfiles.Sort().Filter(args) {
			if statusCmdGroups {
				fmt.Println("[" + g.Name + "]")
			}
			for i := range g.Symlinks {
				op.MustPrintln(op.Read(&g.Symlinks[i]))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().BoolVar(&statusCmdGroups, "groups", true, "Show groups")
}
