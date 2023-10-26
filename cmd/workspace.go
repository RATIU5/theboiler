package cmd

import (
	"fmt"

	"github.com/RATIU5/theboiler/pkg/db"
	"github.com/spf13/cobra"
)

var cmdWksp = &cobra.Command{
	Use:   "wksp",
	Short: "List or change the current workspace",
	Long:  "Groups of similar templates are split into workspaces. View or change the current workspace with this command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			val, err := db.GetCurrentWorkspace()
			if err != nil {
				fmt.Println("There is no workspace set. Set one with `boil wksp [name]`")
			} else {
				fmt.Printf("%s\n", val)
			}
		} else {
			fmt.Printf("Added workspace '%s'", args[0])
		}
	},
}

func init() {
	root.AddCommand(cmdWksp)
}
