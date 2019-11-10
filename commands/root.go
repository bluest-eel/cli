package commands

import (
	"strings"

	"github.com/bluest-eel/cli/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Variables ...
var (
	Authors = strings.Join([]string{
		"Bluest Eel Team <bluesteel@billo.systems>"}, "\n  ")
	Copyright = strings.Join([]string{
		"(c) 2019 Antonio Macías Ojeda",
		"(c) 2019 BilloSystems, Ltd. Co."}, "\n  ")
	Support = "https://github.com/bluest-eel/cli/issues/new"
	Website = "https://github.com/bluest-eel/cli"
)

var cliInstance *CLI
var rootCmd = &cobra.Command{
	Use:   "bluest-eel",
	Short: "A bluest-eel CLI",
	Long: `This tool is a placeholder for a larger effort that will live ` +
		`in aseparate repository.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Tracef("Command: %#v", cmd)
		log.Debugf("Args: %#v", args)
		cliInstance.PostSetupPreRun()
	},
	Version: common.VersionString(),
}

var helpTemplate = `NAME
  {{.Name}} - {{.Short}}

USAGE{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}

DESCRIPTION
  {{.Long}}{{if gt (len .Aliases) 0}}

ALIASES
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

EXAMPLES
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

AVAILABLE COMMANDS{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

FLAGS
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

GLOBAL FLAGS
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

TOPICS{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

COPYRIGHT
  {{Copyright}}

AUTHORS
  {{Authors}}

WEBSITE
  {{Website}}

SUPPORT
  {{Support}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
