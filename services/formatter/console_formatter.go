package formatter

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/roadsigns/httpct/services"
	"io"
	"time"
)

const ConsoleFormatterFlag = "console"

type ConsoleFormatter struct {
}

func (c ConsoleFormatter) Format(results services.HttpHeaderScanResults, writer io.Writer) error {
	writer.Write([]byte(figure.NewFigure("httpct", "", true).String()))
	tableGenerator := services.CommandLineTableGenerator{
		Writer: writer,
	}
	tableGenerator.OutputGenericInformationTable(results.Url, time.Now().Format("15:04:05 01 Jan 2006"))
	tableGenerator.OutputRawHttpHeaderTable(results.RawHeaders)
	if len(results.MissingHeaders) > 0 {
		writer.Write([]byte(figure.NewFigure("Missing Security Headers", "", true).String()))
		tableGenerator.OutputMissingHeaderTable(results.MissingHeaders)
	}

	return nil
}
