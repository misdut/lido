package cmd

import (
	"fmt"
	"os"
	"todo/lido"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "LiDo",
	Short:	"CLI program for manage and storage your tasks",
	Long: "CLI program for manage and storage your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(lido.TodoList.Show())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Ooops. An error while executing LiDo! '%s'\n", err)
		os.Exit(1)
	}
}
