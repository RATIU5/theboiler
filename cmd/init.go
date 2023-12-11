/*
Copyright © 2023 RATIU5 contact@ratiu5.dev
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/RATIU5/theboiler/internal/db"
	"github.com/RATIU5/theboiler/internal/files"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new boilerplate/template",
	Long: `Initialize a new boilerplate/template with
a given name. This will setup a new boilerplate/template
environment behind the scenes.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You didn't specify any additional arguments. A name for the boilerplate was expected.")
			return
		}
		if len(args) > 1 {
			fmt.Println("Too many arguments were provided. Only a name for the boilerplate was expected.")
			return
		}
		name := args[0]

		if !files.DoesPathExist(files.GetApplicationPath()) {
			err := files.CreateDirPath(files.GetApplicationPath())
			if err != nil {
				log.Fatalf("error: failed to create the database. reason: %s", err)
			}
		}

		dtbs, err := db.OpenDB(files.GetDatabasePath())
		if err != nil {
			log.Fatalf("error: failed to open the database. reason: %s", err)
		}
		defer dtbs.Close()

		err = db.CreateBucketIfNotExist(dtbs, []byte(db.BUCKET_NAME_CORE))
		if err != nil {
			log.Fatalf("error: failed to create bucket. reason: %s", err)
		}

		err = db.SetBytesInBucket(dtbs, []byte(db.BUCKET_NAME_CORE), []byte(db.BUCKET_KEY_INIT), []byte(name))
		if err != nil {
			log.Fatalf("error: failed to assign value in bucket. reason: %s", err)
		}

		fmt.Printf("Initialized a new boilerplate '%s'\n", name)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
