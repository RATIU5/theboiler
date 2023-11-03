/*
Copyright © 2023 RATIU5 contact@ratiu5.dev
*/
package cmd

import (
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
			log.Fatal("You didn't specify any additional arguments. A name for the boilerplate was expected.")
		}
		name := args[0]

		if !files.DoesPathExist(files.GetDatabasePath()) {
			err := files.CreatePath(files.GetDatabasePath())
			if err != nil {
				log.Fatalf("error: failed to create the database. reason: %s", err)
			}
		}

		dtbs, err := db.OpenDB(files.GetDatabasePath())
		if err != nil {
			log.Fatalf("error: failed to open the database. reason: %s", err)
		}
		defer dtbs.Close()

		err = db.CreateBucketIfNotExist(dtbs, db.BUCKET_NAME_CORE)
		if err != nil {
			log.Fatalf("error: failed to create bucket. reason: %s", err)
		}

		err = db.SetStringInBucket(dtbs, db.BUCKET_NAME_CORE, db.BUCKET_KEY_INIT, name)
		if err != nil {
			log.Fatalf("error: failed to assign value in bucket. reason: %s", err)
		}
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
