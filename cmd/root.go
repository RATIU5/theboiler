package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:   "boil [cmd]",
	Short: "Run boil command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TheBoiler: boilerplate and template management with ease")
		fmt.Printf("Type `%v` for a list of commands\n", color.GreenString("boil help"))
	},
}
