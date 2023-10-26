package cmd

import (
	"fmt"

	"github.com/RATIU5/theboiler/internal/db"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cmdWksp = &cobra.Command{
	Use:   "wksp",
	Short: "List or change the current workspace",
	Long:  "Groups of similar templates are split into workspaces. View or change the current workspace with this command",
	Run: func(cmd *cobra.Command, args []string) {
		val, err := db.GetCurrentWorkspace()
		if err != nil {
			fmt.Printf("%s: %v\n", color.RedString("error"), err)
			return
		} else {
			fmt.Printf("%s", val)
		}
	},
}

func init() {
	root.AddCommand(cmdWksp)
}
