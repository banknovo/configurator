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
