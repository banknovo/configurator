package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	app              string
	environment      string
	additionalConfig []string
	format           string

	// RootCmd is the root cli command
	RootCmd = &cobra.Command{
		Use:   "configurator",
		Short: "CLI for fetching config values from AWS parameters store",
		SilenceUsage: true,
	}
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&app, "app", "a", "", "App for which config need to be fetched")
	RootCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "",
		`Environment for which config needs to be fetched
	development
	production`)
	RootCmd.PersistentFlags().StringArrayVarP(&additionalConfig, "additional", "t", []string{}, "Any additional config values which need to be fetched, accepts comma separated strings")

	RootCmd.MarkPersistentFlagRequired("app")
	RootCmd.MarkPersistentFlagRequired("environment")
}

// Execute parses command line flags and starts the program
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getEnvironment() (string, error) {
	var env string
	switch environment {
	case "development":
		env = "Dev"
	case "production":
		env = "Prod"
	default:
		return "", fmt.Errorf("Invalid environment `%s`", environment)
	}
	return env, nil
}
