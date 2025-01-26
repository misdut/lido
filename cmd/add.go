package cmd

import (
	"fmt"
	"todo/lido"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition"},
	Short:   "Add the task to the list",
	Long:    "Add title and date to task list",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		lido.TodoList.Add(args[0], args[1])
		fmt.Println(lido.TodoList)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
