package cmd

import (
	"fmt"

	"github.com/RATIU5/theboiler/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of theBoiler",
	Long:  `All software has versions. This is the version for theBoiler`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("theBoiler v%s\n", version.GetVersion())
	},
}

func init() {
	root.AddCommand(versionCmd)
}
