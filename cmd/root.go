package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/avakarev/dotfiles-cli/internal/config"
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
			log.Println(err)
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	cobra.OnInitialize(config.Init)

	rootCmd.PersistentFlags().StringP("config", "c", "", "config file")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}
