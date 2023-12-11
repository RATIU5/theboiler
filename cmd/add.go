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

		if !utils.IsDot(args[0]) {
			fmt.Println("Can only add the current directory.")
			return
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

		name, err := db.GetBytesInBucket(dbc, []byte(db.BUCKET_NAME_CORE), []byte(db.BUCKET_KEY_INIT))
		if err != nil {
			log.Fatalf("error: failed to read database. reason: %s\n", err)
		}

		encodedData, err := utils.Encode(fileContent)
		if err != nil {
			log.Fatalf("error: failed to encode data. reason: %s\n", err)
		}

		err = db.SetBytesInBucket(dbc, []byte(db.BUCKET_NAME_CORE), name, encodedData)
		if err != nil {
			log.Fatalf("error: failed to write database. reason: %s\n", err)
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
