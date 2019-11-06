package commands

import (
	"github.com/bluest-eel/state/tool"
	"github.com/golang/geo/s2"
	"github.com/spf13/cobra"
)

// Location data
type Location struct {
	Lat    float32
	Lon    float32
	Name   string
	CellID s2.CellID
}

// Nearest utility data
type Nearest struct {
	Center *Location
}

func init() {
	rootCmd.AddCommand(nearestCmd)
}

var nearestCmd = &cobra.Command{
	Use:   "nearest",
	Short: "Get the nearest points",
	Long:  "Get the nearest points",
	Run: func(cmd *cobra.Command, args []string) {
		tool.Run()
	},
}
