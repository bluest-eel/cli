package main

import (
	"github.com/bluest-eel/cli/app"
	"github.com/bluest-eel/cli/common"
	"github.com/bluest-eel/cli/components/config"
	"github.com/bluest-eel/cli/components/logging"

	// "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Create the application object and assign components to it
	cli := new(app.Application)
	cli.Config = config.NewConfig()

	cli.Logger = logging.Load(cli.Config)
	log.Infof("Running bluest-eel version: %s", common.VersionString())
	log.Infof("Running bluest-eel build: %s", common.BuildString())

	// If at any future point CLI actions require one or more of the services
	// to be set up, we'll need to move this lower down in the setup sequence.
	cli.ProcessCLIOptions()
	cli.HandleSignals()

	// Add the remaining components to the application data structure - do
	// these next ones after processing CLI/signal interrupts, but before
	// starting the services.
	// cli.GRPCD = reverb.New()

	// Start the core servers
	cli.Start()
}
