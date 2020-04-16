package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
