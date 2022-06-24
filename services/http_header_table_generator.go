package services

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

type TableGenerator interface {
	OutputGenericInformationTable()
	OutputRawHttpHeaderTable()
	OutputMissingHeaderTable()
}

type CommandLineTableGenerator struct {
}

func (tableGenerator CommandLineTableGenerator) OutputRawHttpHeaderTable(headers []RawHeader) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Http Header", "Content"})
	for _, header := range headers {
		t.AppendRow([]interface{}{header.Title, header.Content})
		t.AppendSeparator()
	}
	t.Render()
}

func (tableGenerator CommandLineTableGenerator) OutputMissingHeaderTable(headers []MissingHeader) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Http Header", "Reason"})
	for _, header := range headers {
		t.AppendRow([]interface{}{header.Title, header.Reason})
		t.AppendSeparator()
	}
	t.Render()
}

func (tableGenerator CommandLineTableGenerator) OutputGenericInformationTable(url, time string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow([]interface{}{"Domain", url})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"Time", time})
	t.Render()
}
