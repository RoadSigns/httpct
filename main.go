package main

import (
	"github.com/roadsigns/httpct/cmd/httpct"
)

func main() {
	os.Exit(commands.ScanSecurityHeaders(os.Args))
}
