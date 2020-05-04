package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnvironment(t *testing.T) {
	tests := []struct {
		environment string
		expected    string
	}{
		{"development", "Dev"},
		{"production", "Prod"},
	}
	for _, test := range tests {
		environment = test.environment
		actual, err := getEnvironment()
		require.Empty(t, err, "got error")
		require.Equal(t, test.expected, actual)
	}
}

func TestRemovePrefix(t *testing.T) {
	tests := []struct {
		key          string
		prefixLength int
		expected     string
	}{
		{"/Key1/Key2/Key3/Key4", 0, "Key1/Key2/Key3/Key4"},
		{"/Key1/Key2/Key3/Key4", 1, "Key2/Key3/Key4"},
		{"/Key1/Key2/Key3/Key4", 2, "Key3/Key4"},
		{"/Key1/Key2/Key3/Key4", 3, "Key4"},
	}
	for _, test := range tests {
		actual := removePrefix(test.key, test.prefixLength)
		require.Equal(t, test.expected, actual)
	}
}
