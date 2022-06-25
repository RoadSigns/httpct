package commands

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/roadsigns/http-header-scanner/actions"
	"github.com/roadsigns/http-header-scanner/services"
	"time"
)

func ScanSecurityHeaders(args []string) int {
	if len(args) == 1 {
		// Needs to be changed to a nice format message
		panic("Domain argument must be provided")
	}

	domain := actions.GetHttpSecurityHeadersData{
		Domain: args[1],
	}

	if ok, err := domain.ValidatedUrl(); !ok {
		color.Set(color.FgRed, color.Bold)
		defer color.Unset()
		fmt.Println(err.Error())
		return 1
	}

	results := actions.CheckSecurityHttpHeaders(domain)

	figure.NewFigure("HTTP Header Scanner", "", true).Print()
	tableGenerator := services.CommandLineTableGenerator{}
	tableGenerator.OutputGenericInformationTable(domain.Domain, time.Now().Format("15:04:05 01 Jan 2006"))
	tableGenerator.OutputRawHttpHeaderTable(results.RawHeaders)
	if len(results.MissingHeaders) > 0 {
		figure.NewFigure("Missing Security Headers", "", true).Print()
		tableGenerator.OutputMissingHeaderTable(results.MissingHeaders)
	}
	return 0
}
