package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lido",
	Short: "lido is a cli program to store your tasks",
	Long:  "Lido is a cli program to store your tasks in a simple and intuitive way",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Ooops. An error while executing LiDo! '%s'\n", err)
		os.Exit(1)
	}
}
