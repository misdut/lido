package cmd

import (
	"lido/todo"

	"github.com/spf13/cobra"
)

var toggleCmd = &cobra.Command{
	Use:     "t",
	Aliases: []string{"toggle"},
	Short:   "Toggle checks task as done",
	Run:	func(cmd *cobra.Command, args []string) {
		todo.TodoList.ToggleChecked(args[0])
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
