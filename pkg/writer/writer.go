package writer

import (
	"fmt"
	"io"
	"os"
)

type Writer struct {
	writer io.Writer
}

func Create(output string) io.Writer {
	if output != "" {
		writer, err := os.Create(output)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		return writer
	}

	return os.Stdout
}
