package convert

import (
	"encoding/json"
	"testing"

	"github.com/banknovo/configurator/config"
	"github.com/stretchr/testify/require"
)

func TestBlueprintConvert(t *testing.T) {
	configs := make([]*config.Config, 0)

	configs = append(configs, &config.Config{
		Key:   "Key1/Key2/Key3",
		Value: "value1",
	})
	configs = append(configs, &config.Config{
		Key:   "Key4/Key5",
		Value: "value2",
	})
	configs = append(configs, &config.Config{
		Key:   "Key6",
		Value: "1",
	})
	configs = append(configs, &config.Config{
		Key:   "Key7",
		Value: "true",
	})
	configs = append(configs, &config.Config{
		Key:   "Key8",
		Value: "false",
	})

	// create a blueprint map from JSON
	str := `{
		"k1": {
		  "k2": {
			"k3": "Key1/Key2/Key3"
		  }
		},
		"k4": {
		  "k5": "Key4/Key5"
		},
		"k6": "Key6",
		"k7": "Key7",
		"k8": "Key8"
	  }`
	var blueprintMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &blueprintMap)
	require.Empty(t, err, "got error reading JSON")

	var e Converter = &BlueprintBasedConverter{
		Blueprint: blueprintMap,
	}
	configMap, err := e.Convert(configs)

	require.Empty(t, err, "got error while export")

	m, ok := configMap["k1"].(map[string]interface{})
	require.True(t, ok)
	m, ok = m["k2"].(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, "value1", m["k3"])

	m, ok = configMap["k4"].(map[string]interface{})
	require.True(t, ok)

	require.Equal(t, "value2", m["k5"])

	require.Equal(t, int64(1), configMap["k6"])
	require.Equal(t, true, configMap["k7"])
	require.Equal(t, false, configMap["k8"])
}

func TestBlueprintConvertReturnsErrorOnMissingKey(t *testing.T) {
	configs := make([]*config.Config, 0)

	configs = append(configs, &config.Config{
		Key:   "Key1/Key2/Key3",
		Value: "value1",
	})
	configs = append(configs, &config.Config{
		Key:   "Key4/Key5",
		Value: "value2",
	})
	configs = append(configs, &config.Config{
		Key:   "Key6",
		Value: "1",
	})
	configs = append(configs, &config.Config{
		Key:   "Key7",
		Value: "true",
	})
	configs = append(configs, &config.Config{
		Key:   "Key8",
		Value: "true",
	})

	// create a blueprint map from JSON
	str := `{
		"k1": {
		  "k2": {
			"k3": "Key1/Key2/Key3"
		  }
		},
		"k4": {
		  "k5": "Key4/Key5"
		},
		"k6": "Key6",
		"k7": "Key7",
		"k8": "Key8",
		"k9": "Key9"
	  }`
	var blueprintMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &blueprintMap)
	require.Empty(t, err, "got error reading JSON")

	var e Converter = &BlueprintBasedConverter{
		Blueprint: blueprintMap,
	}
	_, err = e.Convert(configs)

	require.NotEmpty(t, err, "Expected err because key is missing from configs")
	require.Contains(t, err.Error(), "Key9")
}
