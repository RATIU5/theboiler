package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/RATIU5/theboiler/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	isRecursive    bool
	useDirectories bool
	useFiles       bool
	cmdAdd         = &cobra.Command{
		Use:   "add",
		Short: "Add files and folders as a template",
		Long:  "Add a curration of files and folders to be stored as a template",
		Run: func(cmd *cobra.Command, args []string) {
			res := handleAddCommand(cmd, args)
			fmt.Println(res)
		},
	}
)

func init() {
	root.AddCommand(cmdAdd)
}

type Item struct {
	itemType  string
	itemName  string
	itemValue string
}

func handleAddCommand(cmd *cobra.Command, args []string) string {
	var output string
	cwd, err := os.Getwd()
	if err != nil {
		output = "failed to get current working directory"
	}
	err = filepath.Walk(cwd,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if utils.IsExcludedDir(path, []string{
				".git",
			}) {
				return nil // Skip if excluded
			}

			if utils.IsExcludedFile(path, []string{}) {
				return nil
			}

			fmt.Println(path)
			return nil
		})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", cwd, err)
	}

	return output
}
