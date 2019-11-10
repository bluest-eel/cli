package commands

import (
	"os"

	"github.com/bluest-eel/cli/common"
	"github.com/bluest-eel/cli/components"
	"github.com/bluest-eel/cli/components/config"
	"github.com/bluest-eel/cli/components/logging"
	cfglib "github.com/bluest-eel/common/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CLI ...
type CLI struct {
	components.Base
	components.BaseApp
	components.BaseCLI
}

// NewCLI ...
func NewCLI() *CLI {
	cli := CLI{}
	cli.AppName = common.AppName
	cli.AppAbbv = common.AppAbbreviation
	cli.ProjectPath = common.CallerPaths().DotPath
	cli.RawArgs = os.Args
	return &cli
}

// BootstrapConfiguration ...
func (cli *CLI) BootstrapConfiguration() *config.Config {
	cfglib.Setup(cli.AppAbbv, cli.ProjectPath, config.ConfigFile)
	return config.NewConfig()
}

// // Close the gRPC connection
// func (cli *CLI) Close() {
// 	t.APIConn.Close()
// }

// Execute executes the root command, kicking off whatever setup is needed
// first.
func (cli *CLI) Execute(args []string) error {
	cli.RawArgs = args
	cli.Setup()
	return rootCmd.Execute()
}

// PostSetupPreRun ...
func (cli *CLI) PostSetupPreRun() {
	// Setup gets called after bootstrapping has taken place
	cliInstance.SetEnvVars()
	// We now may need to make changes to the setup; let's redo it with any of
	// the new inputs that may affect things:
	cliInstance.Config = cliInstance.SetupConfiguration()
	cliInstance.Logger = logging.LoadClient(cliInstance.Config)
	// cliInstance.SetupDBConnection()
}

// SetEnvVars looks for specific values in the envionment that are not part of
// configuration and pulls them in. Note that 99% of the time, you'll actually
// want to update the config and not mess with this. ParseEnv is only needed
// for pulling data out of the environment that can impact how configuration
// is read.
func (cli *CLI) SetEnvVars() {
	cfgFile := cfglib.EnvConfigFile()
	if cfgFile != "" {
		log.Info("Overwriting config file flag with ENV var ...")
		cli.ConfigFile = cfgFile
	}
}

// Setup ...
func (cli *CLI) Setup() {
	cobra.OnInitialize(func() {
		cliInstance = cli
	})
	cobra.AddTemplateFunc("Authors", func() string { return Authors })
	cobra.AddTemplateFunc("Copyright", func() string { return Copyright })
	cobra.AddTemplateFunc("Support", func() string { return Support })
	cobra.AddTemplateFunc("Website", func() string { return Website })
	rootCmd.PersistentFlags().StringVarP(
		&cli.ConfigFile,
		"config", "c",
		"",
		"configuration file to use")
	rootCmd.SetHelpTemplate(helpTemplate)
}

// SetupConfiguration ...
func (cli *CLI) SetupConfiguration() *config.Config {
	if cli.ConfigFile != "" {
		log.Debug("Updating configuration ...")
		log.Debugf("Using project path '%s' ...", cli.ProjectPath)
		log.Debugf("Using config file '%s' ...", cli.ConfigFile)
		cfglib.Setup(cli.AppAbbv, cli.ProjectPath, cli.ConfigFile)
		return config.NewConfig()
	}
	log.Debug("No config file passed; using bootstrapped config")
	return cli.Config
}

// SetupDBConnection ...
// func (cli *CLI) SetupDBConnection() {
// 	log.Debug("Setting up database connection ...")
// 	db, err := db.Open(cli.Config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	cli.DB = &db
// }

// SetupGRPCConnection ...
func (cli *CLI) SetupGRPCConnection() {
	// connectionOpts := c.Config.GRPCConnectionString()
	// conn, err := grpc.Dial(connectionOpts, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect to gRPC server: %v", err)
	// }
	// c.KVConn = conn
	// c.KVClient = api.NewKVServiceClient(conn)
}
