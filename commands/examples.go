package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(examplesCmd)
}

var examplesCmd = &cobra.Command{
	Use:   "examples",
	Short: "Run Bluest Eel examples",
	Long:  "Run Bluest Eel examples",
}
