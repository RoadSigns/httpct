package formatter

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/roadsigns/httpct/pkg/scanners"
	"io"
)

type TableGenerator interface {
	OutputGenericInformationTable()
	OutputRawHttpHeaderTable()
	OutputMissingHeaderTable()
}

type CommandLineTableGenerator struct {
	Writer io.Writer
}

func (tableGenerator CommandLineTableGenerator) OutputRawHttpHeaderTable(headers []scanners.RawHeader) {
	t := table.NewWriter()
	t.SetOutputMirror(tableGenerator.Writer)
	t.AppendHeader(table.Row{"Http Header", "Content"})
	for _, header := range headers {
		t.AppendRow([]interface{}{header.Title, header.Content})
		t.AppendSeparator()
	}
	t.Render()
}

func (tableGenerator CommandLineTableGenerator) OutputMissingHeaderTable(headers []scanners.MissingHeader) {
	t := table.NewWriter()
	t.SetOutputMirror(tableGenerator.Writer)
	t.AppendHeader(table.Row{"Http Header", "Reason", "Guide"})
	for _, header := range headers {
		t.AppendRow([]interface{}{header.Title, header.Reason, header.Guide})
		t.AppendSeparator()
	}
	t.Render()
}

func (tableGenerator CommandLineTableGenerator) OutputGenericInformationTable(url, time string) {
	t := table.NewWriter()
	t.SetOutputMirror(tableGenerator.Writer)
	t.AppendRow([]interface{}{"Domain", url})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"Time", time})
	t.Render()
}
