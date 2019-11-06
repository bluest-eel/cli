package main

import (
	"os"

	"github.com/bluest-eel/state/cli/commands"
	"github.com/bluest-eel/state/components/logging"
	log "github.com/sirupsen/logrus"
)

// XXX Note that any of this which ends up being useful will be moved into the
//     bluest-eel/cli repo; the `tool` code here is just an experiment to get
//     familiar with the s2 library.
func main() {
	// Create the tool object and assign components to it
	cli := commands.NewCLI()
	// Bootstrap configuration and logging with defaults; this is to assist with
	// any debugging, e.g., logging output, etc.
	cli.Config = cli.BootstrapConfiguration()
	cli.Logger = logging.LoadClient(cli.Config)
	err := cli.Execute(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
