package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "storage-optimizer",
	Short: "Storage analysis and optimization tool",
	Long: `Storage Optimizer: A professional tool for analyzing and optimizing disk storage usage
	
Features:
	- Find duplicate and similar files
	- Detect large and heavy files
	- Analyze file distribution by type
	- Detailed reports in JSON and CSV formats
	- Safe deletion with optional backup`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	rootCmd.AddCommand(findDuplicatesCmd)
	rootCmd.AddCommand(findLargeCmd)
	rootCmd.AddCommand(cleanupCmd)
}
