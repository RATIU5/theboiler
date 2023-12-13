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
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a boilerplate",
	Long:  `Create a new boilerplate from the current directory. The name of the boilerplate is required.`,
	Run: func(cmd *cobra.Command, args []string) {

		boilerplateName, err := cmd.Flags().GetString("boilerplate")
		if err != nil || boilerplateName == "" {
			fmt.Println("a boilerplate name was expected, none received.")
		}

		excludedFiles := []string{".git", ".DS_Store", "build"}
		res, err := files.GetFileList(files.GetWorkingDirectory(), excludedFiles)
		if err != nil {
			log.Fatalf("error: failed to get file list. reason: %s\n", err)
		}

		fileContent, err := files.GetFilesContent(res)
		if err != nil {
			log.Fatalf("error: failed to get file content. reason: %s\n", err)
		}

		dbc, err := db.OpenDB(files.GetDatabasePath())
		if err != nil {
			log.Fatalf("error: failed to read database. reason: %s\n", err)
		}

		if db.DoesBoilerplateExist(dbc, []byte(boilerplateName)) {
			fmt.Printf("boilerplate '%s' already exists.\n", boilerplateName)
			return
		}

		encodedData, err := utils.Encode(fileContent)
		if err != nil {
			log.Fatalf("error: failed to encode data. reason: %s\n", err)
		}

		err = db.WriteBoilerplate(dbc, []byte(boilerplateName), encodedData)
		if err != nil {
			log.Fatalf("error: failed to write database. reason: %s\n", err)
		}

		fmt.Printf("boilerplate '%s' created.\n", boilerplateName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().StringP("boilerplate", "b", "", "The boilerplate to store the current directory to")
}
