package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/banknovo/configurator/config"
	"github.com/banknovo/configurator/convert"
	"github.com/banknovo/configurator/store"
)

var (
	paths         []string
	convertMode   string
	excludePrefix int

	// RootCmd is the root cli command
	RootCmd = &cobra.Command{
		Use:          "configurator",
		Short:        "CLI for fetching config values from AWS parameters store",
		SilenceUsage: true,
	}
)

func init() {
	RootCmd.PersistentFlags().StringSliceVarP(&paths, "paths", "p", []string{},
		"Parameter Store path for which values need to be fetched, accepts comma separated strings")

	RootCmd.PersistentFlags().StringVarP(&convertMode, "mode", "m", "flat",
		`Mode of convert required
	flat: Keys are exported as-is in a flat structure
	hierarchical: Keys are exported in a hierarchical structure, keys broken down by separator
	blueprint: A blueprint file is taken which has keys are values, and real values are replaced in it`)

	RootCmd.PersistentFlags().IntVarP(&excludePrefix, "excludePrefix", "x", 0,
		"The number of prefixes to exclude from the final export")

	err := RootCmd.MarkPersistentFlagRequired("paths")
	if err != nil {
		panic(err)
	}
}

// Execute parses command line flags and starts the program
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// getConverter returns the converter based on mode
func getConverter(mode string) (convert.Converter, error) {
	var e convert.Converter
	switch mode {
	case "flat":
		e = &convert.FlatConverter{}
	case "hierarchical":
		e = &convert.HierarchicalConverter{
			Separator: "/",
		}
	case "blueprint":
		{
			data, err := getBlueprint()
			if err != nil {
				return nil, err
			}

			var blueprintMap map[string]interface{}
			err = json.Unmarshal(data, &blueprintMap)
			if err != nil {
				return nil, err
			}

			e = &convert.BlueprintBasedConverter{
				Blueprint: blueprintMap,
			}
		}
	default:
		return nil, fmt.Errorf("invalid convert mode `%s`", mode)
	}
	return e, nil
}

// getConfig fetches the configs from the store
func getConfigs() (map[string]interface{}, error) {
	c, err := getConverter(convertMode)
	if err != nil {
		return nil, err
	}

	s, err := store.NewAWSPMStore()
	if err != nil {
		return nil, err
	}

	allConfigs := make([]*config.Config, 0)
	for _, path := range paths {
		configs, err := s.FetchAll(path)
		if err != nil {
			return nil, err
		}
		// remove prefix from the key name
		for _, c := range configs {
			c.Key = removePrefix(c.Key, excludePrefix)
		}
		allConfigs = append(allConfigs, configs...)
	}

	configMap, err := c.Convert(allConfigs)
	if err != nil {
		return nil, err
	}

	return configMap, nil
}

func removePrefix(key string, prefixLength int) string {
	splitAfter := prefixLength + 2
	idx := prefixLength + 1
	return strings.SplitAfterN(key, "/", splitAfter)[idx]
}
