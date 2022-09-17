package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/avakarev/dotfiles-cli/internal/op"
	"github.com/avakarev/dotfiles-cli/pkg/dotfiles"
)

var unlinkCmdGroups bool

// unlinkCmd implements the `unlink` command
var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink dotfiles",
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
				op.MustPrintln(op.Unlink(&g.Symlinks[i]))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(unlinkCmd)
	unlinkCmd.Flags().BoolVar(&unlinkCmdGroups, "groups", true, "Show groups")
}
