package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/RATIU5/theboiler/pkg/db"
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
	if len(args) == 0 {
		return "not enough arguments"
	}
	if len(args) > 1 {
		return "too many arguments"
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "failed to get current working directory"
	}
	var items []item.Item
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

			items = append(items, *itm)

			return nil
		})
	if err != nil {
		return fmt.Sprintf("error walking the path %q: %v\n", cwd, err)
	}

	err = db.StoreItems([]byte(args[0]), items)
	if err != nil {
		return fmt.Sprintf("error storing data: %v\n", err)
	}

	return ""
}
