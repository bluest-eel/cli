package commands

import (
	"github.com/bluest-eel/state/tool"
	"github.com/golang/geo/s2"
	log "github.com/sirupsen/logrus"
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
	nearestCmd.Flags().StringVarP(&rawCenter, "center", "",
		"60.1699::24.9384::Helsinki Center",
		"double-colon separated center point data of the format\n"+
			"'lat::long::description'; additionally, you may choose\n"+
			"to pass one of the following pre-defined points:\n"+
			"  * "+tool.Tolkien+"\n")
	nearestCmd.Flags().StringVarP(&rawPoints, "points", "p",
		"helskini",
		"vertical-bar separated list of points of the format\n"+
			"'point|point|...' where each point has the same format\n"+
			"as the 'center' flag above; additionally, you may choose\n"+
			"to pass one of the following pre-defined example points:\n"+
			"  * "+tool.Helsinki+"\n"+
			"  * "+tool.Oxford+"\n")
}

var rawCenter string
var rawPoints string
var nearestCmd = &cobra.Command{
	Use:   "nearest",
	Short: "Get the nearest points",
	Long:  "Get the nearest points",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Args:", args)
		log.Debug("Center:", rawCenter)
		log.Tracef("cmd: %#v", cmd)
		tool.Run(rawCenter, rawPoints)
	},
}
