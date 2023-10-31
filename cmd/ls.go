package cmd

import (
	"fmt"

	"github.com/RATIU5/theboiler/pkg/db"
	"github.com/spf13/cobra"
)

var cmdLs = &cobra.Command{
	Use:   "ls",
	Short: "List contents of a boilerplate",
	Long:  "View the the files and directories of a given boilerplate",
	Run: func(cmd *cobra.Command, args []string) {
		handleLsCommand(cmd, args)
	},
}

func init() {
	root.AddCommand(cmdLs)
}

func handleLsCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("not enough arguments; name expected")
		return
	}

	if len(args) > 1 {
		fmt.Println("too many arguments")
		return
	}

	name := []byte(args[0])

	items, err := db.ReadItem(name)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, item := range items {
		item.Print()
	}
}
