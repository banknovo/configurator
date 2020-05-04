package convert

import (
	"regexp"
	"strconv"

	"github.com/banknovo/configurator/config"
)

// Converter converts the config values into desired format
type Converter interface {
	// Convert converts the array of configs into a map
	Convert(configs []*config.Config) (map[string]interface{}, error)
}

// getTypedValue attempts to convert rawValue to its data type
func getTypedValue(rawValue string) interface{} {
	var parsedValue interface{}
	var err error

	// check if it is a bool
	if rawValue == "true" {
		return true
	}
	if rawValue == "false" {
		return false
	}

	// check if it is an int, if so convert to int64
	matched, err := regexp.Match("^(0|[1-9][0-9]*)$", []byte(rawValue))
	if err == nil && matched {
		parsedValue, err = strconv.ParseInt(rawValue, 10, 64)
		if err == nil {
			return parsedValue
		}
	}

	// for all others return rawValue as string
	return rawValue
}
