/*
Copyright © 2023 RATIU5 contact@ratiu5.dev
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "theboiler",
	Short: "A boilerplate and template manager for your projects",
	Long: `theBoiler is a CLI tool that manages your boilerplates
and templates for all your coding projects. Get started with the
command 'boil init'.`,
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
}
