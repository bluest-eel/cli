package commands

import (
	"fmt"
	"strconv"

	"github.com/google/hilbert"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	hilbertCmd.AddCommand(getCoordCmd)
}

var getCoordCmd = &cobra.Command{
	Use:   "get-coord",
	Short: "Get the 2D coordinate for the provided Hilbert Curve integer",
	Long:  "Get the 2D coordinate for the provided Hilbert Curve integer",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		maxValue := (order * order) - 1
		tooBigMsg := fmt.Sprintf(
			"Hilbert Curve integer cannot be greater than order squared, minus 1 (max: %d)",
			maxValue)
		if t > maxValue {
			log.Fatal(tooBigMsg)
		}
		curve, err := hilbert.NewHilbert(order)
		x, y, err := curve.Map(t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d, %d\n", x, y)
	},
}
