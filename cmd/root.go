package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "boil",
	Short: "A boilerplate manager",
	Long:  "A boilerplate manager that handles all your favorite project boilerplates and templates",
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
