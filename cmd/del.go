package cmd

import (
	"lido/todo"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command {
    Use:    "del",
    Aliases: []string{"remove"},
    Short:  "Remove a task",
    Long:   "Remove a task you have already completed",
    Run: func(cmd *cobra.Command, args []string){
        todo.TodoList.Del(args[0])
    },
}

func init() {
    rootCmd.AddCommand(delCmd)
}