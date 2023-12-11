/*
Copyright © 2023 RATIU5 contact@ratiu5.dev
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/RATIU5/theboiler/internal/db"
	"github.com/RATIU5/theboiler/internal/files"
	"github.com/RATIU5/theboiler/internal/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file(s) to an initialized boilerplate",
	Long: `Add any number of files and folders to an
initialized boilerplate. 
This command takes at least 1 argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Nothing was specified to add. A file or folder was expected.")
			return
		}

		if utils.IsDot(args[0]) {
			fmt.Println("Added directory")
			return
		}

		_, err := db.OpenDB(files.GetDatabasePath())
		if err != nil {
			log.Fatalf("error: failed to read database. reason: %s\n", err)
		}

		
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
