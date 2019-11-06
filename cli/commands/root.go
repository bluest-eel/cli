package commands

import (
	"strings"

	"github.com/bluest-eel/state/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Variables ...
var (
	Authors = strings.Join([]string{
		"Bluest Eel Team <bluesteel@billo.systems>"}, "\n   ")
	Copyright = strings.Join([]string{
		"(c) 2019 Antonio Macías Ojeda",
		"(c) 2019 BilloSystems, Ltd. Co."}, "\n   ")
	Support     = "https://github.com/bluest-eel/state/issues/new"
	Website     = "https://github.com/bluest-eel/state"
	cliInstance *CLI
	rootCmd     = &cobra.Command{
		Use:   "state",
		Short: "A bluest-eel CLI",
		Long: `This tool is a placeholder for a larger effort that will live in a
separate repository.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.Tracef("Command: %#v", cmd)
			log.Debugf("Args: %#v", args)
			cliInstance.PostSetupPreRun()
		},
		Version: common.VersionString(),
	}
)
