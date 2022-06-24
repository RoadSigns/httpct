package main

import (
	"github.com/roadsigns/http-header-scanner/transport/commands"
	"os"
)

func main() {
	commands.ScanSecurityHeaders(os.Args)
}
