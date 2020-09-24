package cmd

import (
	"bytes"
	"fmt"
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

func TestOutput(t *testing.T) {
	tests := []map[string]interface{}{
		// simulates a flat config
		{
			"Key1": "Value1",
			"Key2": "Value2",
			"Key3": "Value3",
		},
		// simulates a hierarchal config
		{"Key1": map[string]interface{}{
			"Key2": map[string]interface{}{
				"Key3": "Value3",
				"Key4": "Value4",
			},
			"Key5": "Value5",
		}},
	}

	// inside the raw strings (``) use only spaces no tabs otherwise tests will fail
	expects := []struct {
		json, yaml []byte
	}{
		{
			json: []byte(`{
  "Key1": "Value1",
  "Key2": "Value2",
  "Key3": "Value3"
}
`),
			yaml: []byte(`Key1: Value1
Key2: Value2
Key3: Value3
`),
		},
		{
			json: []byte(`{
  "Key1": {
    "Key2": {
      "Key3": "Value3",
      "Key4": "Value4"
    },
    "Key5": "Value5"
  }
}
`),
			yaml: []byte(`Key1:
  Key2:
    Key3: Value3
    Key4: Value4
  Key5: Value5
`),
		},
	}

	indent = true

	for i, v := range tests {
		t.Run(fmt.Sprintf("JSON-%d", i+1), func(st *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			err := exportAsJSON(v, buf)
			if err != nil {
				st.Fatalf("An error occurred while formating JSON(%v): %v", v, err)
			}

			b := buf.Bytes()
			if bytes.Compare(expects[i].json, b) != 0 {
				st.Fatalf("Expected: %q\nbut\ngot: %q", expects[i].json, b)
			}
		})

		t.Run(fmt.Sprintf("YAML-%d", i+1), func(st *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			err := exportAsYAML(v, buf)
			if err != nil {
				st.Fatalf("An error occurred while formating YAML(%v): %v", v, err)
			}

			b := buf.Bytes()
			if bytes.Compare(expects[i].yaml, b) != 0 {
				st.Fatalf("Expected: %q\nbut\ngot: %q", expects[i].yaml, b)
			}
		})
	}
}
