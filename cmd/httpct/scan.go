package httpct

import (
	"github.com/roadsigns/httpct/transport/commands"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:     "scan",
	Aliases: []string{"scan"},
	Short:   "Scan a URL",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commands.ScanSecurityHeaders(args)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
