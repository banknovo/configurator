package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	envCmd = &cobra.Command{
		Use:   "env [flags]",
		Short: "Print the config in a format to export as environment variables",
		RunE:  runEnv,
	}
)

func init() {
	RootCmd.AddCommand(envCmd)
}

func runEnv(*cobra.Command, []string) error {
	var err error

	if convertMode != "flat" {
		return fmt.Errorf("only flat convert mode is supported with env")
	}

	configMap, err := getConfigs()
	if err != nil {
		return err
	}

	for key, value := range configMap {
		fmt.Printf("export %s=%v\n",
			strings.ToUpper(strings.Replace(key, "/", "_", -1)),
			value,
		)
	}

	return nil
}
