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
			Cli:    getCliOption(cmd),
			Format: getFormatOption(cmd),
			Output: getOutputOption(cmd),
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

func getFormatOption(cmd *cobra.Command) string {
	format := cmd.Flag("format").Value.String()
	if format == "" {
		return cmd.Flag("format").DefValue
	}
	return format
}

func getOutputOption(cmd *cobra.Command) string {
	output := cmd.Flag("output").Value.String()
	if cmd.Flag("output").Value.String() != "" {
		return output
	}

	return ""
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().Bool("cli", false, "Returns an error code if any headers are missing")
	scanCmd.Flags().StringP("format", "f", "cli", "Returns the results in the format requested")
	scanCmd.Flags().StringP("output", "o", "", "Outputs the file to the location requested")
}
