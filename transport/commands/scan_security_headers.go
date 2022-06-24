package commands

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/roadsigns/http-header-scanner/actions"
	"os"
)

func ScanSecurityHeaders(args []string) {
	if len(args) == 1 {
		// Needs to be changed to a nice format message
		panic("Domain argument must be provided")
	}

	domain := actions.GetHttpSecurityHeadersData{
		Domain: args[1],
	}

	results := actions.CheckSecurityHttpHeaders(domain)

	// Create an Interface to write table so if package dies
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Http Header", "Reason"})
	for _, header := range results.MissingHeaders {
		t.AppendRow([]interface{}{header.Title, header.Reason})
		t.AppendSeparator()
	}
	t.Render()

	if len(results.AdditionalInformationHeader) > 0 {
		t = table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Http Header", "Reason"})
		for _, header := range results.AdditionalInformationHeader {
			t.AppendRow([]interface{}{header.Title, header.Reason})
			t.AppendSeparator()
		}
		t.Render()
	}

	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Http Header", "Content"})
	for _, header := range results.RawHeaders {
		t.AppendRow([]interface{}{header.Title, header.Content})
		t.AppendSeparator()
	}
	t.Render()
}
