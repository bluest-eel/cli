package commands

import (
	"github.com/spf13/cobra"
)

var (
	order int
)

func init() {
	rootCmd.AddCommand(hilbertCmd)
	hilbertCmd.Flags().IntVarP(&order, "order", "o", 16,
		"Set the Hibert curve order; must be a power of two\n"+
			"squaring this provides the total number of cells")
}

var hilbertCmd = &cobra.Command{
	Use:   "hilbert",
	Short: "Perform Hilbert Curve calculations",
	Long:  "Perform Hilbert Curve calculations",
}
