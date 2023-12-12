/*
Copyright © 2023 RATIU5 <contact@ratiu5.dev>
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

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		boilerplateName, err := cmd.Flags().GetString("boilerplate")
		if err != nil || boilerplateName == "" {
			fmt.Println("a boilerplate name was expected, none received.")
			return
		}

		dbc, err := db.OpenDB(files.GetDatabasePath())
		if err != nil {
			log.Fatalf("error: failed to read database. reason: %s\n", err)
			return
		}

		if !db.DoesCoreBucketExist(dbc) {
			db.CreateCoreBucket(dbc)
		}

		if !db.DoesBoilerplateExist(dbc, []byte(boilerplateName)) {
			fmt.Printf("boilerplate '%s' does not exist.\nAdd one with 'boil create -b <name>'", boilerplateName)
			return
		}

		encodedData, err := db.ReadBoilerplate(dbc, []byte(boilerplateName))
		if err != nil {
			log.Fatalf("error: failed to read database. reason: %s\n", err)
			return
		}

		fileContent, err := utils.Decode[[]files.FileContent](encodedData)
		if err != nil {
			fmt.Printf("error: failed to decode data. reason: %s\n", err)
			return
		}

		for _, file := range fileContent {
			fmt.Println(file.String())
		}

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	viewCmd.Flags().StringP("boilerplate", "b", "", "The boilerplate to store the current directory to")
}
