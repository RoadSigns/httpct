package formatter

import (
	"encoding/json"
	"github.com/roadsigns/httpct/pkg/scanners"
	"io"
)

const JsonFormatterFlag = "json"

type JsonFormatter struct {
}

func (j JsonFormatter) Format(results scanners.HttpHeaderScanResults, writer io.Writer) error {
	json, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		return err
	}

	_, err = writer.Write(json)
	if err != nil {
		return err
	}

	return nil
}
