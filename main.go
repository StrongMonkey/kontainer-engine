package main

import (
	"os"

	"github.com/rancher/kontainer-engine/cmd"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

var AppHelpTemplate = `{{.Usage}}

Usage: {{.Name}} {{if .Flags}}[DRIVER_OPTIONS] {{end}}COMMAND [arg...]

Version: {{.Version}}
{{if .Flags}}
Options:
  {{range .Flags}}{{if .Hidden}}{{else}}{{.}}
  {{end}}{{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`

var CommandHelpTemplate = `{{.Usage}}
{{if .Description}}{{.Description}}{{end}}
Usage: kontainer-engine [global options] {{.Name}} {{if .Flags}}[OPTIONS] {{end}}{{if ne "None" .ArgsUsage}}{{if ne "" .ArgsUsage}}{{.ArgsUsage}}{{else}}[arg...]{{end}}{{end}}

{{if .Flags}}Options:{{range .Flags}}
	 {{.}}{{end}}{{end}}
`

func main() {
	cli.AppHelpTemplate = AppHelpTemplate
	cli.CommandHelpTemplate = CommandHelpTemplate

	app := cli.NewApp()
	app.Name = "kontainer-engine"
	app.Version = VERSION
	app.Usage = "CLI tool for creating and managing kubernetes clusters"
	app.Before = func(ctx *cli.Context) error {
		if ctx.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		logrus.Debugf("kontainer-engine version: %v", VERSION)
		return nil
	}
	app.Author = "Rancher Labs, Inc."
	app.Commands = []cli.Command{
		cmd.CreateCommand(),
		cmd.UpdateCommand(),
		cmd.InspectCommand(),
		cmd.LsCommand(),
		cmd.RmCommand(),
		cmd.EnvCommand(),
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable verbose logging",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
