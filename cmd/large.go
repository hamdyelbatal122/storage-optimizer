package cmd

import (
	"storage-optimizer/internal/analyzer"
	"strings"
	"github.com/spf13/cobra"
)

var findLargeCmd = &cobra.Command{
	Use:   "large [path]",
	Short: "Find large files",
	Long:  "Find largest files consuming storage space",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		limit, _ := cmd.Flags().GetInt("limit")
		outputFile, _ := cmd.Flags().GetString("output")
		format, _ := cmd.Flags().GetString("format")
		minSize, _ := cmd.Flags().GetInt64("min-size")
		exclude, _ := cmd.Flags().GetString("exclude")

		a := analyzer.NewAnalyzer(path)
		
		if minSize > 0 {
			a.SetMinSize(minSize)
		}
		if exclude != "" {
			a.SetExcludePatterns(strings.Split(exclude, ","))
		}
		
		largeFiles, err := a.FindLargeFiles(limit)
		if err != nil {
			return err
		}

		return largeFiles.Print(format, outputFile)
	},
}

func init() {
	findLargeCmd.Flags().IntP("limit", "l", 50, "Number of files to display")
	findLargeCmd.Flags().StringP("output", "o", "", "Output file")
	findLargeCmd.Flags().StringP("format", "f", "table", "Output format: table, json, csv")
	findLargeCmd.Flags().Int64("min-size", 0, "Minimum file size in bytes to display")
	findLargeCmd.Flags().String("exclude", "", "Exclude patterns (comma-separated)")
}
