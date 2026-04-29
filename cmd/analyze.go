package cmd

import (
	"storage-optimizer/internal/analyzer"
	"strings"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze [path]",
	Short: "Analyze directory storage usage",
	Long:  "Perform comprehensive analysis of storage usage in specified directory",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		outputFile, _ := cmd.Flags().GetString("output")
		format, _ := cmd.Flags().GetString("format")
		exclude, _ := cmd.Flags().GetString("exclude")
		extensions, _ := cmd.Flags().GetString("extensions")
		minSize, _ := cmd.Flags().GetInt64("min-size")
		maxAge, _ := cmd.Flags().GetInt64("max-age")

		a := analyzer.NewAnalyzer(path)
		
		if exclude != "" {
			a.SetExcludePatterns(strings.Split(exclude, ","))
		}
		if extensions != "" {
			a.SetIncludeExtensions(strings.Split(extensions, ","))
		}
		if minSize > 0 {
			a.SetMinSize(minSize)
		}
		if maxAge > 0 {
			a.SetMaxAge(maxAge)
		}
		
		report, err := a.AnalyzeDirectory()
		if err != nil {
			return err
		}

		return report.Print(format, outputFile)
	},
}

func init() {
	analyzeCmd.Flags().StringP("output", "o", "", "Output file (optional)")
	analyzeCmd.Flags().StringP("format", "f", "table", "Output format: table, json, csv")
	analyzeCmd.Flags().String("exclude", "", "Exclude patterns (comma-separated, e.g., node_modules,vendor,.git)")
	analyzeCmd.Flags().String("extensions", "", "Include only these file types (comma-separated, e.g., .go,.js,.txt)")
	analyzeCmd.Flags().Int64("min-size", 0, "Minimum file size in bytes")
	analyzeCmd.Flags().Int64("max-age", 0, "Maximum file age in days")
}
