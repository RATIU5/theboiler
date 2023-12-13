/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Use a boilerplate to start your project.",
	Long: `Setup your project with a specific boilerplate.
This command will copy all the files and directories from the boilerplate to the current directory.`,
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
			fmt.Printf("boilerplate '%s' does not exist.\nadd one with 'boil create -b <name>'\n", boilerplateName)
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
			if file.IsDir {
				err := files.CreateDir(file.Path)
				if err != nil {
					log.Fatalf("error: failed to create directory. reason: %s\n", err)
				}
			} else {
				err := files.CreateFile(file.Path, file.Content)
				if err != nil {
					log.Fatalf("error: failed to create file. reason: %s\n", err)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	useCmd.Flags().StringP("boilerplate", "b", "", "The boilerplate to store the current directory to")
}
