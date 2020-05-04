package convert

import (
	"strings"

	"github.com/banknovo/configurator/config"
)

// ensure HierarchicalConverter confirms to Converter interface
var _ Converter = &HierarchicalConverter{}

// HierarchicalConverter converts the configs into a hierarchical structure
type HierarchicalConverter struct {
	Separator string
}

// Convert converts the configs into a nested map of key value pairs
func (e *HierarchicalConverter) Convert(configs []*config.Config) (map[string]interface{}, error) {
	configMap := make(map[string]interface{})
	for _, config := range configs {
		keys := strings.Split(config.Key, e.Separator)
		value := getTypedValue(config.Value)
		addToMap(configMap, keys, value)
	}
	return configMap, nil
}

// addToMap is a recursive function that creates the hierarchical map of keys and values
func addToMap(configMap map[string]interface{}, keys []string, value interface{}) map[string]interface{} {
	if len(keys) == 1 {
		configMap[keys[0]] = value
		return configMap
	}
	currKey := keys[0]
	var newMap map[string]interface{}
	if configMap[currKey] != nil {
		newMap = configMap[currKey].(map[string]interface{})
	} else {
		newMap = make(map[string]interface{})
	}
	configMap[currKey] = addToMap(newMap, keys[1:], value)
	return configMap
}
