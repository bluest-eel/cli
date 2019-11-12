package commands

import (
	"fmt"
	"strconv"

	"github.com/google/hilbert"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	hilbertCmd.AddCommand(getIntCmd)
}

var getIntCmd = &cobra.Command{
	Use:   "get-int",
	Short: "Get the Hilbert Curve integer for a 2D coordinate",
	Long:  "Get the Hilbert Curve integer for a 2D coordinate",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		x, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		maxCoordValue := order - 1
		tooBigMsg := fmt.Sprintf(
			"Coordinate value cannot be greater than order, minus 1 (max: %d)",
			maxCoordValue)
		if x > maxCoordValue {
			log.Fatal(tooBigMsg)
		}
		y, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		if y > maxCoordValue {
			log.Fatal(tooBigMsg)
		}
		curve, err := hilbert.NewHilbert(order)
		t, err := curve.MapInverse(x, y)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
	},
}
