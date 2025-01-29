package cmd

import (
	"todo/lido"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command {
    Use:    "del",
    Aliases: []string{"remove"},
    Short:  "Remove a task",
    Long:   "Remove a task you have already completed",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string){
        lido.TodoList.Del(args[0])
    },
}

func init() {
    rootCmd.AddCommand(delCmd)
}