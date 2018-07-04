package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jcelliott/lumber"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config   string
	showVers bool

	name     string // Deployment Name
	replicas int32  // Number of Desired Pod Replicas
	port     int32  // Container Port
	atype    string // Artefact type

	// Populated during compile
	version string
	commit  string

	// RootCmd defineds the entry point command for the CLI
	RootCmd = &cobra.Command{
		Use:               "kstub",
		Short:             "KStub is a very fast generator for Kubernetes manifests",
		Long:              ``,
		SilenceErrors:     true,
		SilenceUsage:      true,
		PersistentPreRunE: readConfig,
		PreRunE:           preFlight,
		RunE:              runKStub,
	}
)

func preFlight(cmd *cobra.Command, args []string) error {
	// if --version is passed print the version info
	if showVers {
		fmt.Printf("kstube %s (%s)\n", version, commit)
		return fmt.Errorf("")
	}

	// if --server is not passed, print help
	if !viper.GetBool("server") {
		cmd.HelpFunc()(cmd, args)
		return fmt.Errorf("") // no error, just exit
	}

	return nil
}

func runKStub(cmd *cobra.Command, args []string) error {
	// convert the log level
	logLvl := lumber.LvlInt(viper.GetString("log-level"))

	// configure the logger
	lumber.Prefix("[kstub]")
	lumber.Level(logLvl)

	// fall back on default help if no args/flags are passed
	cmd.HelpFunc()(cmd, args)
	return nil
}

func readConfig(cmd *cobra.Command, args []string) error {
	// if --config is passed, attempt to parse the config file
	if config != "" {

		// get the filepath
		abs, err := filepath.Abs(config)
		if err != nil {
			lumber.Error("Error reading filepath: ", err.Error())
		}

		// get the config name
		base := filepath.Base(abs)

		// get the path
		path := filepath.Dir(abs)

		//
		viper.SetConfigName(strings.Split(base, ".")[0])
		viper.AddConfigPath(path)

		// Find and read the config file; Handle errors reading the config file
		if err := viper.ReadInConfig(); err != nil {
			lumber.Fatal("Failed to read config file: ", err.Error())
			os.Exit(1)
		}
	}
	return nil
}

func init() {
	// set config defaults
	logLevel := "INFO"
	viper.SetDefault("garbage-collect", false)

	// CLI flags
	RootCmd.PersistentFlags().String("log-level", logLevel, "Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL)")

	// Local flags
	RootCmd.Flags().StringVar(&config, "config", "", "/path/to/config.yml")
	RootCmd.Flags().BoolVarP(&showVers, "version", "v", false, "Display the current version of this CLI")

	// Commands
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(deploymentCmd)
	RootCmd.AddCommand(serviceCmd)
	RootCmd.AddCommand(ingressCmd)
}
