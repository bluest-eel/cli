package commands

import (
	"github.com/bluest-eel/state/tool"
	"github.com/golang/geo/s2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	capRadius float64
	maxCells  int
	maxLevel  int
	rawCenter string
	rawPoints string
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
	nearestCmd.Flags().Float64VarP(&capRadius, "radius", "r", 20,
		"The cap radius used to determine nearness, in kilometers")
	nearestCmd.Flags().StringVarP(&rawCenter, "center", "",
		tool.Tolkien,
		"double-colon separated center point data of the format\n"+
			"'lat::long::description'; additionally, you may choose\n"+
			"to pass one of the following pre-defined points:\n"+
			"  * "+tool.Helsinki+"\n"+
			"  * "+tool.Tolkien+"\n")
	nearestCmd.Flags().IntVarP(&maxCells, "max-cells", "", 8,
		"The maximum desired number of cells in the approximation")
	nearestCmd.Flags().IntVarP(&maxLevel, "max-levels", "", 20,
		"The maximum cell level to be used")
	nearestCmd.Flags().StringVarP(&rawPoints, "points", "p",
		tool.Oxford,
		"vertical-bar separated list of points of the format\n"+
			"'point|point|...' where each point has the same format\n"+
			"as the 'center' flag above; additionally, you may choose\n"+
			"to pass one of the following pre-defined example points:\n"+
			"  * "+tool.Helsinki+"\n"+
			"  * "+tool.Oxford+"\n"+
			"  * "+tool.OxfordPubs+"\n")
}

var nearestCmd = &cobra.Command{
	Use:   "nearest",
	Short: "Get the nearest points",
	Long:  "Get the nearest points",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Args:", args)
		log.Debug("Center:", rawCenter)
		log.Tracef("cmd: %#v", cmd)
		tool.Run(&tool.RunOptions{
			CapRadius:   capRadius,
			CenterPoint: rawCenter,
			MaxCells:    maxCells,
			MaxLevel:    maxLevel,
			Points:      rawPoints,
		})
	},
}
