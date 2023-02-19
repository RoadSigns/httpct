package formatter

import (
	"github.com/roadsigns/httpct/services"
	"io"
)

type Formatter interface {
	Format(results services.HttpHeaderScanResults, writer io.Writer) error
}

type Factory struct {
}

func (f Factory) Create(format string) Formatter {
	if format == JsonFormatterFlag {
		return JsonFormatter{}
	}

	return ConsoleFormatter{}
}
