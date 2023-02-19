package scan

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/roadsigns/httpct/actions"
	"github.com/roadsigns/httpct/services"
	"time"
)

func SecurityHeaders(domainPath string, options Options) int {
	domain := actions.GetHttpSecurityHeadersData{
		Domain: domainPath,
	}

	if ok, err := domain.ValidatedUrl(); !ok {
		color.Set(color.FgRed, color.Bold)
		defer color.Unset()
		fmt.Println(err.Error())
		return 1
	}

	results := actions.CheckSecurityHttpHeaders(domain)

	figure.NewFigure("httpct", "", true).Print()
	tableGenerator := services.CommandLineTableGenerator{}
	tableGenerator.OutputGenericInformationTable(domain.Domain, time.Now().Format("15:04:05 01 Jan 2006"))
	tableGenerator.OutputRawHttpHeaderTable(results.RawHeaders)
	if len(results.MissingHeaders) > 0 {
		figure.NewFigure("Missing Security Headers", "", true).Print()
		tableGenerator.OutputMissingHeaderTable(results.MissingHeaders)

		if options.Cli {
			return 1
		}
	}

	return 0
}
