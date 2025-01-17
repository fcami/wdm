package wdm

import (
	"fmt"

	config "github.com/openshift-psap/wdm/pkg/config"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

const (
	DefaultConfigFile = "examples/prepare_cluster.yaml"
)

var log = logrus.New()

func GetLogger() *logrus.Logger {
	return log
}

type Flags struct {
	ConfigFile string
	DryRun     bool
	Install    bool
}

type Context struct {
	*cli.Context
	Flags *Flags
}

func BuildCommand() *cli.Command {
	// Create a flags struct to hold our flags
	wdmFlags := Flags{}

	wdm := cli.Command{}
	wdm.Name = "wdm"
	wdm.Usage = "Workload Dependency Manager"
	wdm.Action = func(c *cli.Context) error {
		return wdmWrapper(c, &wdmFlags)
	}

	// Setup the flags for this command
	wdm.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config-file",
			Aliases:     []string{"c"},
			Usage:       "Configuration file to use for WDM",
			Destination: &wdmFlags.ConfigFile,
			Value:       DefaultConfigFile,
			EnvVars:     []string{"WDM_CONFIG_FILE"},
		},
		&cli.BoolFlag{
			Name:        "dry-run",
			Aliases:     []string{"d"},
			Usage:       "Dry run only, do not test (mock all test results to false",
			Destination: &wdmFlags.DryRun,
			EnvVars:     []string{"WDM_DRY_RUN"},
		},
		&cli.BoolFlag{
			Name:        "install",
			Aliases:     []string{"i"},
			Usage:       "Enable missing dependency installation",
			Destination: &wdmFlags.Install,
			EnvVars:     []string{"WDM_DEP_INSTALL"},
		},
	}

	return &wdm
}

func wdmWrapper(c *cli.Context, f *Flags) error {
	dependencySpec, err := config.LoadYamlFile(f.ConfigFile)
	if err != nil {
		return fmt.Errorf("error parsing config file: %v", err)
	}

	log.Infof("Dependency spec: %v", dependencySpec)

	return nil
}
