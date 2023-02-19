package formatter

import (
	"github.com/roadsigns/httpct/services"
	"gopkg.in/yaml.v3"
	"io"
)

const YamlFormatterType = "yaml"
const YmlFormatterType = "yml"

type YamlFormatter struct {
}

func (y YamlFormatter) Format(results services.HttpHeaderScanResults, writer io.Writer) error {
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
