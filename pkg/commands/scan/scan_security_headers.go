package scan

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/roadsigns/httpct/pkg/actions"
	"github.com/roadsigns/httpct/pkg/formatter"
	"github.com/roadsigns/httpct/pkg/writer"
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

	factory := formatter.Factory{}
	writer := writer.Create(options.Output)
	factory.Create(options.Format).Format(results, writer)

	if options.Cli && len(results.MissingHeaders) > 0 {
		return 1
	}
	return 0
}
