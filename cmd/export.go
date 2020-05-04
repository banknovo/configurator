package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	blueprintFile string
	exportFormat  string
	exportOutput  string

	exportCmd = &cobra.Command{
		Use:   "export [flags]",
		Short: "Export the config",
		RunE:  runExport,
	}
)

func init() {
	exportCmd.Flags().StringVarP(&blueprintFile, "blueprint-file", "b", "",
		"Path to blueprint file (only required if convertMode is blueprint)")

	exportCmd.Flags().StringVarP(&exportFormat, "format", "f", "json",
		`Format of the export
	json`)

	exportCmd.Flags().StringVarP(&exportOutput, "output-file", "o", "",
		"Output file (default is standard output)")

	RootCmd.AddCommand(exportCmd)
}

func runExport(*cobra.Command, []string) error {
	var err error

	configMap, err := getConfigs()
	if err != nil {
		return err
	}

	file := os.Stdout
	if exportOutput != "" {
		if file, err = os.OpenFile(exportOutput, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
			return errors.Wrap(err, "Failed to open output file for writing")
		}
		defer file.Sync()
		defer file.Close()
	}

	w := bufio.NewWriter(file)
	defer w.Flush()

	switch exportFormat {
	case "json":
		err = exportAsJSON(configMap, w)
	default:
		err = fmt.Errorf("unknown export format: %s", exportFormat)
	}

	if err != nil {
		return errors.Wrap(err, "Unable to export configs")
	}

	return nil
}

func exportAsJSON(configMap map[string]interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(configMap)
}

func getBlueprint() ([]byte, error) {
	if blueprintFile == "" {
		return nil, fmt.Errorf("blueprint-file is required when convert mode is blueprint")
	}
	return ioutil.ReadFile(blueprintFile)
}
