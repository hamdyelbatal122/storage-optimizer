package analyzer

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type AnalysisReport struct {
	TotalFiles   int
	TotalSize    int64
	FilesByType  map[string]int
	LargestFiles []*FileInfo
	SizeByType   map[string]int64
}

type DuplicateReport struct {
	TotalDuplicates int
	Groups          []*DuplicateGroup
	PotentialSaved  int64
}

type LargeFilesReport struct {
	Files     []*FileInfo
	TotalSize int64
}

type CleanupReport struct {
	DryRun        bool
	DeletedFiles  int
	FreedSpace    int64
	BackupCreated bool
}

func formatBytes(bytes int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	value := float64(bytes)

	for _, unit := range units {
		if value < 1024 {
			return fmt.Sprintf("%.2f %s", value, unit)
		}
		value /= 1024
	}

	return fmt.Sprintf("%.2f PB", value)
}

func (r *AnalysisReport) Print(format, outputFile string) error {
	if format == "json" {
		return r.printJSON(outputFile)
	} else if format == "csv" {
		return r.printCSV(outputFile)
	}

	fmt.Println(color.GreenString("\nComprehensive Analysis Report"))
	fmt.Println(strings.Repeat("─", 50))

	fmt.Printf("Total Files: %s\n", color.BlueString("%d", r.TotalFiles))
	fmt.Printf("Storage Used: %s\n", color.CyanString(formatBytes(r.TotalSize)))

	fmt.Println(color.YellowString("\nFile Distribution by Type:"))
	for ext, count := range r.FilesByType {
		size := r.SizeByType[ext]
		fmt.Printf("  %s: %d files (%s)\n", ext, count, formatBytes(size))
	}

	fmt.Println(color.YellowString("\nTop 10 Largest Files:"))
	for i, file := range r.LargestFiles {
		fmt.Printf("  %d. %s (%s)\n", i+1, file.Path, formatBytes(file.Size))
	}

	return nil
}

func (r *AnalysisReport) printCSV(outputFile string) error {
	var output *os.File
	var err error

	if outputFile != "" {
		output, err = os.Create(outputFile)
		if err != nil {
			return err
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	writer := csv.NewWriter(output)
	defer writer.Flush()

	writer.Write([]string{"File Type", "Count", "Total Size"})
	for ext, count := range r.FilesByType {
		size := r.SizeByType[ext]
		writer.Write([]string{ext, strconv.Itoa(count), strconv.FormatInt(size, 10)})
	}

	return nil
}

func (r *AnalysisReport) printJSON(outputFile string) error {
	data := map[string]interface{}{
		"totalFiles":   r.TotalFiles,
		"totalSize":    r.TotalSize,
		"filesByType":  r.FilesByType,
		"sizeByType":   r.SizeByType,
		"largestFiles": r.LargestFiles,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if outputFile != "" {
		return os.WriteFile(outputFile, jsonData, 0644)
	}

	fmt.Println(string(jsonData))
	return nil
}

func (r *DuplicateReport) Print(format, outputFile string) error {
	if format == "json" {
		return r.printJSON(outputFile)
	} else if format == "csv" {
		return r.printCSV(outputFile)
	}

	fmt.Println(color.GreenString("\nDuplicate Files Report"))
	fmt.Println(strings.Repeat("─", 50))

	fmt.Printf("Number of Duplicate Groups: %s\n", color.BlueString("%d", r.TotalDuplicates))
	fmt.Printf("Recoverable Space: %s\n", color.CyanString(formatBytes(r.PotentialSaved)))

	fmt.Println(color.YellowString("\nDuplicate Details:"))
	for _, group := range r.Groups {
		fmt.Printf("\n  Group (Hash: %s)\n", color.RedString(group.Hash[:8]))
		fmt.Printf("     Number of Copies: %d | Copy Size: %s\n", group.Count, formatBytes(group.Size))
		for i, file := range group.Files {
			fmt.Printf("     %d. %s\n", i+1, file)
		}
	}

	return nil
}

func (r *DuplicateReport) printCSV(outputFile string) error {
	var output *os.File
	var err error

	if outputFile != "" {
		output, err = os.Create(outputFile)
		if err != nil {
			return err
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	writer := csv.NewWriter(output)
	defer writer.Flush()

	writer.Write([]string{"Hash", "Copies", "Size", "Files"})
	for _, group := range r.Groups {
		files := strings.Join(group.Files, "; ")
		writer.Write([]string{group.Hash[:8], strconv.Itoa(group.Count), strconv.FormatInt(group.Size, 10), files})
	}

	return nil
}

func (r *DuplicateReport) printJSON(outputFile string) error {
	data := map[string]interface{}{
		"totalDuplicates": r.TotalDuplicates,
		"potentialSaved":  r.PotentialSaved,
		"groups":          r.Groups,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if outputFile != "" {
		return os.WriteFile(outputFile, jsonData, 0644)
	}

	fmt.Println(string(jsonData))
	return nil
}

func (r *LargeFilesReport) Print(format, outputFile string) error {
	if format == "json" {
		return r.printJSON(outputFile)
	} else if format == "csv" {
		return r.printCSV(outputFile)
	}

	fmt.Println(color.GreenString("\nLarge Files Report"))
	fmt.Println(strings.Repeat("─", 50))

	fmt.Printf("Total Files: %d\n", len(r.Files))
	fmt.Printf("Total Space: %s\n", formatBytes(r.TotalSize))

	fmt.Println(color.YellowString("\nLargest Files:"))
	for i, file := range r.Files {
		percentage := float64(file.Size) * 100 / float64(r.TotalSize)
		fmt.Printf("%2d. %s\n    %s (%.1f%%)\n",
			i+1, file.Path, formatBytes(file.Size), percentage)
	}

	return nil
}

func (r *LargeFilesReport) printCSV(outputFile string) error {
	var output *os.File
	var err error

	if outputFile != "" {
		output, err = os.Create(outputFile)
		if err != nil {
			return err
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	writer := csv.NewWriter(output)
	defer writer.Flush()

	writer.Write([]string{"Path", "Size (bytes)", "Percentage"})
	for _, file := range r.Files {
		percentage := fmt.Sprintf("%.2f", float64(file.Size)*100/float64(r.TotalSize))
		writer.Write([]string{file.Path, strconv.FormatInt(file.Size, 10), percentage})
	}

	return nil
}

func (r *LargeFilesReport) printJSON(outputFile string) error {
	data := map[string]interface{}{
		"files":      r.Files,
		"totalSize":  r.TotalSize,
		"fileCount":  len(r.Files),
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if outputFile != "" {
		return os.WriteFile(outputFile, jsonData, 0644)
	}

	fmt.Println(string(jsonData))
	return nil
}

func (r *CleanupReport) Print(format, outputFile string) error {
	status := "Preview"
	if !r.DryRun {
		status = "Executed"
	}

	fmt.Println(color.GreenString(fmt.Sprintf("\nCleanup Report [%s]", status)))
	fmt.Println(strings.Repeat("─", 50))

	fmt.Printf("Deleted Files: %d\n", r.DeletedFiles)
	fmt.Printf("Space Freed: %s\n", formatBytes(r.FreedSpace))
	fmt.Printf("Backup Created: %v\n", r.BackupCreated)

	return nil
}
