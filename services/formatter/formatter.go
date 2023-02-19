package formatter

import "github.com/roadsigns/httpct/services"

type Formatter interface {
	Format(results services.HttpHeaderScanResults)
}

type Factory struct {
}

func (f Factory) Create(format string) Formatter {
	if format == JSON_FORMATTER {
		return JsonFormatter{}
	}

	return ConsoleFormatter{}
}
