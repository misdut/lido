package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lido",
	Short:	"CLI program for manage and storage your tasks",
	Long: "CLI program for manage and storage your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Ooops. An error while executing LiDo! '%s'\n", err)
		os.Exit(1)
	}
}
