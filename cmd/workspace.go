package cmd

import (
	"fmt"

	"github.com/RATIU5/theboiler/pkg/db"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	isVerbose bool
	cmdWksp   = &cobra.Command{
		Use:   "wksp",
		Short: "Show or change the current workspace",
		Long:  "Groups of similar templates are split into workspaces. View or change the current workspace with this command",
		Run: func(cmd *cobra.Command, args []string) {
			res := handleCommand(cmd, args, isVerbose)
			fmt.Println(res)
		},
	}
	cmdWkspLs = &cobra.Command{
		Use:   "ls",
		Short: "List all workspaces",
		Long:  "View all the workspaces that have been used",
		Run: func(cmd *cobra.Command, args []string) {
			// res := handleLsCommand(cmd, args, isVerbose)
			// fmt.Println(res)
			fmt.Println("Test")
		},
	}
)

func init() {
	// The signature of BoolVarP is: BoolVarP(p *bool, name, shorthand string, value bool, usage string)
	cmdWksp.Flags().BoolVarP(&isVerbose, "verbose", "v", false, "Output verbose messages")
	cmdWksp.AddCommand(cmdWkspLs)
	root.AddCommand(cmdWksp)
}

func handleCommand(cmd *cobra.Command, args []string, isVerbose bool) string {
	var output string

	if len(args) > 1 {
		output = "Too many arguments provided"
	} else if len(args) == 1 {
		err := db.SetCurrentWorkspace(args[0])
		if err != nil {
			if isVerbose {
				output = fmt.Sprintf("%s: %s", color.RedString("error"), err)
			} else {
				output = "Failed to set workspace"
			}
		} else {
			output = fmt.Sprintf("Set workspace to '%s'", args[0])
		}
	} else {
		val, err := db.GetCurrentWorkspace()
		if err != nil {
			if isVerbose {
				output = fmt.Sprintf("%s: %s", color.RedString("error"), err)
			} else {
				output = "There is no workspace set. Set one with `boil wksp [name]`"
			}
		} else {
			output = fmt.Sprintf("%s", val)
		}
	}

	return output
}
