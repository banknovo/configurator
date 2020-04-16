package export

import (
	"testing"

	"github.com/banknovo/configurator/core"
	"github.com/stretchr/testify/require"
)

func TestHierarchicalExport(t *testing.T) {
	configs := make([]*core.Config, 0)

	configs = append(configs, &core.Config{
		Key:   "Key1/Key2/Key3",
		Value: "value1",
	})
	configs = append(configs, &core.Config{
		Key:   "Key4/Key5",
		Value: "value2",
	})
	configs = append(configs, &core.Config{
		Key:   "Key6",
		Value: "1",
	})
	configs = append(configs, &core.Config{
		Key:   "Key7",
		Value: "true",
	})
	configs = append(configs, &core.Config{
		Key:   "Key8",
		Value: "false",
	})

	var e Exporter = &HierarchicalExporter{
		Separator: "/",
	}
	configMap, err := e.Export(configs)

	require.Empty(t, err, "got error while export")

	m, ok := configMap["Key1"].(map[string]interface{})
	require.True(t, ok)
	m, ok = m["Key2"].(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, "value1", m["Key3"])

	m, ok = configMap["Key4"].(map[string]interface{})
	require.True(t, ok)

	require.Equal(t, "value2", m["Key5"])

	require.Equal(t, int64(1), configMap["Key6"])
	require.Equal(t, true, configMap["Key7"])
	require.Equal(t, false, configMap["Key8"])
}
