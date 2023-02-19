package scan

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/roadsigns/httpct/actions"
	"github.com/roadsigns/httpct/services/formatter"
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
	factory.Create(options.Format).Format(results)

	if options.Cli {
		return 1
	}
	return 0
}
