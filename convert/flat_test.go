package convert

import (
	"testing"

	"github.com/banknovo/configurator/config"
	"github.com/stretchr/testify/require"
)

func TestFlatConvert(t *testing.T) {
	configs := make([]*config.Config, 0)

	configs = append(configs, &config.Config{
		Key:   "Key1",
		Value: "value1",
	})
	configs = append(configs, &config.Config{
		Key:   "Key2",
		Value: "value2",
	})
	configs = append(configs, &config.Config{
		Key:   "Key3",
		Value: "1",
	})
	configs = append(configs, &config.Config{
		Key:   "Key4",
		Value: "true",
	})
	configs = append(configs, &config.Config{
		Key:   "Key5",
		Value: "false",
	})

	var e Converter = &FlatConverter{}
	configMap, err := e.Convert(configs)

	require.Empty(t, err, "got error while export")
	require.Equal(t, "value1", configMap["Key1"])
	require.Equal(t, "value2", configMap["Key2"])
	require.Equal(t, int64(1), configMap["Key3"])
	require.Equal(t, true, configMap["Key4"])
	require.Equal(t, false, configMap["Key5"])
}
