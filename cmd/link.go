package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/avakarev/dotfiles-cli/internal/op"
	"github.com/avakarev/dotfiles-cli/pkg/dotfiles"
)

var linkCmdGroups bool

// linkCmd implements the `link` command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link dotfiles",
	Run: func(cmd *cobra.Command, args []string) {
		dfiles, err := dotfiles.New()
		if err != nil {
			log.Fatalln(err)
		}

		dfiles.Sort()
		for _, g := range dfiles.Filter(args) {
			if statusCmdGroups {
				fmt.Println("[" + g.Name + "]")
			}
			for i := range g.Symlinks {
				op.Println(op.Link(&g.Symlinks[i]))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
	linkCmd.Flags().BoolVar(&linkCmdGroups, "groups", true, "Show groups")
}
