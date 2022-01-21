package main

import (
	"os"

	wdm "github.com/openshift-psap/wdm/cmd/wdm"
	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

// flags
// - dry-run
// - dep display
// - dep install
// NB: knobs are not mutually exclusive.
type Flags struct {
	// missing: cmd or API operation
	Verbose bool
	DryRun  bool
	Install bool
}

func main() {
	// Create a flags struct to hold our flags
	flags := Flags{}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	app.Usage = "Workload Dependency Manager"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "verbose",
			Aliases:     []string{"V"},
			Usage:       "Enable verbose logging",
			Destination: &flags.Verbose,
			EnvVars:     []string{"WDM_VERBOSE"},
		},
	}

	app.Commands = []*cli.Command{
		wdm.BuildCommand(),
	}

	// Set log-level for all subcommands
	app.Before = func(app *cli.Context) error {
		/*
			logLevel := log.InfoLevel
			if flags.Verbose {
				logLevel = log.DebugLevel
			}

			configLog := config.GetLogger()
			configLog.SetLevel(logLevel)
		*/

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
