package formatter

import (
	"github.com/roadsigns/httpct/pkg/scanners"
	"gopkg.in/yaml.v3"
	"io"
)

const YamlFormatterType = "yaml"
const YmlFormatterType = "yml"

type YamlFormatter struct {
}

func (y YamlFormatter) Format(results scanners.HttpHeaderScanResults, writer io.Writer) error {
	result, err := yaml.Marshal(results)
	if err != nil {
		return err
	}

	_, err = writer.Write(result)
	if err != nil {
		return err
	}

	return nil
}
