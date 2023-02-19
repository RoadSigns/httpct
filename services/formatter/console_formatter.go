package formatter

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/roadsigns/httpct/services"
	"time"
)

const CONSOLE_FORMATTER = "console"

type ConsoleFormatter struct {
}

func (c ConsoleFormatter) Format(results services.HttpHeaderScanResults) {
	figure.NewFigure("httpct", "", true).Print()
	tableGenerator := services.CommandLineTableGenerator{}
	tableGenerator.OutputGenericInformationTable("", time.Now().Format("15:04:05 01 Jan 2006"))
	tableGenerator.OutputRawHttpHeaderTable(results.RawHeaders)
	if len(results.MissingHeaders) > 0 {
		figure.NewFigure("Missing Security Headers", "", true).Print()
		tableGenerator.OutputMissingHeaderTable(results.MissingHeaders)

	}
}
