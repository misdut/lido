package cmd

import (
	"fmt"
	"lido/todo"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition"},
	Short:   "Add the task to the list",
	Long:    "Add your task to the list",

	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < len(args); i++ {
			fmt.Println(todo.TodoList.Add(args[i], time.Now().Local().Format("02/01/2006"), false))
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
