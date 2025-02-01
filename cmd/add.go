package cmd

import (
	"fmt"
	"time"
	"todo/lido"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition"},
	Short:   "Add the task to the list",
	Long:    "Add your task to the list",
	Args:	 cobra.ExactArgs(1),
	Run:	func(cmd *cobra.Command, args []string) {
		lido.TodoList.Add(args[0], time.Now().Local().Format("02/01/2006"), false)
		fmt.Printf("Task \"%v\" added successfully.\n", args[0])

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
