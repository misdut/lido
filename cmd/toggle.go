package cmd

import (
	"todo/lido"

	"github.com/spf13/cobra"
)

var toggleCmd = &cobra.Command{
	Use:     "t",
	Aliases: []string{"toggle"},
	Short:   "Toggle checks task as done",
	Args:	 cobra.ExactArgs(1),
	Run:	func(cmd *cobra.Command, args []string) {
		lido.TodoList.ToggleChecked(args[0])
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
