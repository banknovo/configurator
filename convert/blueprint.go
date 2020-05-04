package convert

import (
	"fmt"

	"github.com/banknovo/configurator/config"
)

// ensure BlueprintBasedConverter confirms to Converter interface
var _ Converter = &BlueprintBasedConverter{}

// BlueprintBasedConverter takes a blueprint of the required structure and
// aims to map the configs to it based on the lookup of keys as values
type BlueprintBasedConverter struct {
	Blueprint map[string]interface{}
}

// Convert replaces the values in blueprint with their actual values from config based on key matching
func (e *BlueprintBasedConverter) Convert(configs []*config.Config) (map[string]interface{}, error) {
	// use FlatConverter to get map of configs
	f := &FlatConverter{}
	configMap, err := f.Convert(configs)
	if err != nil {
		return nil, err
	}

	// create a copy of blueprint and create the final version using configMap
	err = createMap(configMap, e.Blueprint)
	if err != nil {
		return nil, err
	}

	return e.Blueprint, nil
}

// createMap is a recursive function that iterates through the blueprint and
// replaces its values, which are key names, from the configMap
func createMap(configMap map[string]interface{}, blueprint map[string]interface{}) error {
	for key, val := range blueprint {
		v, ok := val.(map[string]interface{})
		if ok {
			err := createMap(configMap, v)
			if err != nil {
				return err
			}
		} else {
			v, ok := val.(string)
			if !ok {
				return fmt.Errorf("expected %s to have string value", val)
			}
			value := configMap[v]
			if value == nil {
				return fmt.Errorf("config value %s missing", v)
			}
			blueprint[key] = value
		}
	}
	return nil
}
