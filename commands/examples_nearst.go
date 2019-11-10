package commands

import (
	"github.com/bluest-eel/state/examples/nearest"
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

func init() {
	examplesCmd.AddCommand(nearestCmd)
	nearestCmd.Flags().Float64VarP(&capRadius, "radius", "r", 20,
		"The cap radius used to determine nearness, in kilometers")
	nearestCmd.Flags().StringVarP(&rawCenter, "center", "",
		nearest.Tolkien,
		"double-colon separated center point data of the format\n"+
			"'lat::long::description'; additionally, you may choose\n"+
			"to pass one of the following pre-defined points:\n"+
			"  * "+nearest.Helsinki+"\n"+
			"  * "+nearest.Tolkien+"\n")
	nearestCmd.Flags().IntVarP(&maxCells, "max-cells", "", 8,
		"The maximum desired number of cells in the approximation")
	nearestCmd.Flags().IntVarP(&maxLevel, "max-levels", "", 20,
		"The maximum cell level to be used")
	nearestCmd.Flags().StringVarP(&rawPoints, "points", "p",
		nearest.Oxford,
		"vertical-bar separated list of points of the format\n"+
			"'point|point|...' where each point has the same format\n"+
			"as the 'center' flag above; additionally, you may choose\n"+
			"to pass one of the following pre-defined example points:\n"+
			"  * "+nearest.Helsinki+"\n"+
			"  * "+nearest.Oxford+"\n"+
			"  * "+nearest.OxfordPubs+"\n")
}

var nearestCmd = &cobra.Command{
	Use:   "nearest",
	Short: "Get the nearest points",
	Long:  "Get the nearest points",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Args:", args)
		log.Debug("Center:", rawCenter)
		log.Tracef("cmd: %#v", cmd)
		nearest.Run(&nearest.RunOptions{
			CapRadius:   capRadius,
			CenterPoint: rawCenter,
			MaxCells:    maxCells,
			MaxLevel:    maxLevel,
			Points:      rawPoints,
		})
	},
}
