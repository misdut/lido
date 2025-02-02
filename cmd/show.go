package cmd 

import (
	"fmt"
	"lido/todo"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:	"ls",
	Aliases: []string{"list"},
	Short:	"Show all your tasks",
	Long:	"List all your tasks",
	Run:	func (cmd *cobra.Command, args []string) {
		fmt.Println(todo.TodoList.Show())
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}