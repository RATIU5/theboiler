package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/RATIU5/theboiler/pkg/item"
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

			var itm *item.Item
			if info.IsDir() {
				itm = item.New(true, path, "")
			} else {
				content, err := utils.ReadFile(path)
				if err != nil {
					return err
				}
				itm = item.New(false, path, content)
			}

			if info.IsDir() {
				itm.Print()
			}

			return nil
		})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", cwd, err)
	}

	return output
}
