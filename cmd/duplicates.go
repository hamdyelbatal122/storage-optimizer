package cmd

import (
	"storage-optimizer/internal/analyzer"
	"strings"
	"github.com/spf13/cobra"
)

var findDuplicatesCmd = &cobra.Command{
	Use:   "duplicates [path]",
	Short: "Find duplicate files",
	Long:  "Find duplicate and identical files using secure hash comparison",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		outputFile, _ := cmd.Flags().GetString("output")
		format, _ := cmd.Flags().GetString("format")
		exclude, _ := cmd.Flags().GetString("exclude")
		extensions, _ := cmd.Flags().GetString("extensions")
		minSize, _ := cmd.Flags().GetInt64("min-size")

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
		
		duplicates, err := a.FindDuplicates()
		if err != nil {
			return err
		}

		return duplicates.Print(format, outputFile)
	},
}

func init() {
	findDuplicatesCmd.Flags().StringP("output", "o", "", "Output file")
	findDuplicatesCmd.Flags().StringP("format", "f", "table", "Output format: table, json, csv")
	findDuplicatesCmd.Flags().String("exclude", "", "Exclude patterns (comma-separated)")
	findDuplicatesCmd.Flags().String("extensions", "", "Include only these file types (comma-separated)")
	findDuplicatesCmd.Flags().Int64("min-size", 0, "Minimum file size in bytes")
}
