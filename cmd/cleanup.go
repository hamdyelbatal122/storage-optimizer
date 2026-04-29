package cmd

import (
	"storage-optimizer/internal/analyzer"
	"strings"
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup [path]",
	Short: "Safe cleanup of files",
	Long:  "Safely delete duplicate and temporary files",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		backup, _ := cmd.Flags().GetBool("backup")
		backupDir, _ := cmd.Flags().GetString("backup-dir")
		exclude, _ := cmd.Flags().GetString("exclude")
		extensions, _ := cmd.Flags().GetString("extensions")

		a := analyzer.NewAnalyzer(path)
		
		if exclude != "" {
			a.SetExcludePatterns(strings.Split(exclude, ","))
		}
		if extensions != "" {
			a.SetIncludeExtensions(strings.Split(extensions, ","))
		}
		
		result, err := a.CleanupDirectory(dryRun, backup)
		if err != nil {
			return err
		}
		
		if backupDir != "" {
			_ = backupDir
		}

		return result.Print("table", "")
	},
}

func init() {
	cleanupCmd.Flags().Bool("dry-run", true, "Preview without actual deletion")
	cleanupCmd.Flags().Bool("backup", true, "Create backup before deletion")
	cleanupCmd.Flags().String("backup-dir", "", "Custom backup directory")
	cleanupCmd.Flags().String("exclude", "", "Exclude patterns (comma-separated)")
	cleanupCmd.Flags().String("extensions", "", "Include only these file types (comma-separated)")
}
