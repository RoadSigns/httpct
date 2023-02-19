package httpct

import (
	"fmt"
	"github.com/roadsigns/httpct/transport/commands/scan"
	"github.com/spf13/cobra"
	"os"
)

var scanCmd = &cobra.Command{
	Use:     "scan",
	Aliases: []string{"scan"},
	Short:   "Scan a URL",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		options := scan.Options{
			Cli: getCliOption(cmd),
		}

		domain := args[len(args)-1]
		result := scan.SecurityHeaders(domain, options)
		if result != 0 {
			fmt.Fprintf(os.Stderr, "Erroring due to CLI flag being set and missing headers found")
			os.Exit(result)
		}
	},
}

func getCliOption(cmd *cobra.Command) bool {
	cli := cmd.Flag("cli").Value.String()
	if cli == "true" {
		return true
	}
	return false
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().Bool("cli", false, "Returns an error code if any headers are missing")
}
