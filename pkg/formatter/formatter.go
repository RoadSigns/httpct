package formatter

import (
	"github.com/roadsigns/httpct/pkg/scanners"
	"io"
)

type Formatter interface {
	Format(results scanners.HttpHeaderScanResults, writer io.Writer) error
}

type Factory struct {
}

func (f Factory) Create(format string) Formatter {
	if format == JsonFormatterFlag {
		return JsonFormatter{}
	}

	if format == YamlFormatterType || format == YmlFormatterType {
		return YamlFormatter{}
	}

	return ConsoleFormatter{}
}
