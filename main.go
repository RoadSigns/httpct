package main

import (
	"github.com/roadsigns/http-header-scanner/transport/commands"
	"os"
)

func main() {
	os.Exit(commands.ScanSecurityHeaders(os.Args))
}
