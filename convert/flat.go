package convert

import (
	"github.com/banknovo/configurator/config"
)

// ensure FlatConverter confirms to Converter interface
var _ Converter = &FlatConverter{}

// FlatConverter converts the configs into flat key-value pairs
type FlatConverter struct{}

// Convert converts the configs into a flat map with key value pairs
func (e *FlatConverter) Convert(configs []*config.Config) (map[string]interface{}, error) {
	configMap := make(map[string]interface{})
	for _, config := range configs {
		configMap[config.Key] = getTypedValue(config.Value)
	}
	return configMap, nil
}
