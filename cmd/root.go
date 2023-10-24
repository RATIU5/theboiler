package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "boil",
	Short: "Boilerplate and template manager",
	Long:  "A fancy boilerplate and template manager for all your development needs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TheBoiler: boilerplate and template management with ease")
		fmt.Printf("Type `%v` for a list of commands\n", color.GreenString("boil help"))
	},
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
