package analyzer

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
)

type Analyzer struct {
	path             string
	files            []*FileInfo
	excludePatterns  []string
	includeExtensions []string
	minSize          int64
	maxAge           int64
	verbose          bool
	progressBar      *progressbar.ProgressBar
}

type FileInfo struct {
	Path      string
	Size      int64
	Modified  int64
	Hash      string
	Extension string
}

type DuplicateGroup struct {
	Hash  string
	Count int
	Files []string
	Size  int64
}

func NewAnalyzer(path string) *Analyzer {
	return &Analyzer{
		path:             path,
		files:            []*FileInfo{},
		excludePatterns:  []string{},
		includeExtensions: []string{},
		minSize:          0,
		maxAge:           0,
		verbose:          false,
	}
}

func (a *Analyzer) SetExcludePatterns(patterns []string) *Analyzer {
	a.excludePatterns = patterns
	return a
}

func (a *Analyzer) SetIncludeExtensions(extensions []string) *Analyzer {
	a.includeExtensions = extensions
	return a
}

func (a *Analyzer) SetMinSize(size int64) *Analyzer {
	a.minSize = size
	return a
}

func (a *Analyzer) SetMaxAge(days int64) *Analyzer {
	a.maxAge = days * 86400
	return a
}

func (a *Analyzer) SetVerbose(v bool) *Analyzer {
	a.verbose = v
	return a
}

func (a *Analyzer) AnalyzeDirectory() (*AnalysisReport, error) {
	fmt.Println(color.CyanString("\nAnalyzing directory..."))

	err := a.walkDirectory()
	if err != nil {
		return nil, err
	}

	report := &AnalysisReport{
		TotalFiles:    len(a.files),
		TotalSize:     a.calculateTotalSize(),
		FilesByType:   a.groupByExtension(),
		LargestFiles:  a.getTopLargeFiles(10),
		SizeByType:    a.calculateSizeByType(),
	}

	return report, nil
}

func (a *Analyzer) FindDuplicates() (*DuplicateReport, error) {
	fmt.Println(color.CyanString("\nSearching for duplicate files..."))

	err := a.walkDirectory()
	if err != nil {
		return nil, err
	}

	duplicates := a.findDuplicatesByHash()

	return &DuplicateReport{
		TotalDuplicates: len(duplicates),
		Groups:          duplicates,
		PotentialSaved:  a.calculateDuplicateSpace(duplicates),
	}, nil
}

func (a *Analyzer) FindLargeFiles(limit int) (*LargeFilesReport, error) {
	fmt.Println(color.CyanString("\nSearching for large files..."))

	err := a.walkDirectory()
	if err != nil {
		return nil, err
	}

	sort.Slice(a.files, func(i, j int) bool {
		return a.files[i].Size > a.files[j].Size
	})

	files := a.files
	if len(a.files) > limit {
		files = a.files[:limit]
	}

	return &LargeFilesReport{
		Files:     files,
		TotalSize: a.calculateTotalSize(),
	}, nil
}

func (a *Analyzer) CleanupDirectory(dryRun, backup bool) (*CleanupReport, error) {
	fmt.Println(color.YellowString("\nPreparing cleanup..."))

	duplicates, err := a.FindDuplicates()
	if err != nil {
		return nil, err
	}

	result := &CleanupReport{
		DryRun:        dryRun,
		DeletedFiles:  0,
		FreedSpace:    0,
		BackupCreated: backup,
	}

	for _, group := range duplicates.Groups {
		if len(group.Files) > 1 {
			for i := 1; i < len(group.Files); i++ {
				if dryRun {
					result.DeletedFiles++
					result.FreedSpace += group.Size
				} else {
					os.Remove(group.Files[i])
					result.DeletedFiles++
					result.FreedSpace += group.Size
				}
			}
		}
	}

	return result, nil
}

func (a *Analyzer) walkDirectory() error {
	fmt.Println("Scanning files...")

	var fileCount int64 = 0
	err := filepath.Walk(a.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			if a.shouldExclude(path) {
				return filepath.SkipDir
			}
			return nil
		}

		if a.shouldExclude(path) {
			return nil
		}

		if !a.matchesExtension(path) {
			return nil
		}

		if info.Size() < a.minSize {
			return nil
		}

		if a.isTooOld(info.ModTime()) {
			return nil
		}

		fileCount++

		ext := filepath.Ext(path)
		a.files = append(a.files, &FileInfo{
			Path:      path,
			Size:      info.Size(),
			Modified:  info.ModTime().Unix(),
			Extension: ext,
		})

		return nil
	})

	fmt.Printf("Found %d files\n", fileCount)
	return err
}

func (a *Analyzer) shouldExclude(path string) bool {
	for _, pattern := range a.excludePatterns {
		if strings.Contains(path, string(filepath.Separator)+pattern+string(filepath.Separator)) || strings.HasPrefix(filepath.Base(path), pattern) {
			return true
		}
	}
	return false
}

func (a *Analyzer) matchesExtension(path string) bool {
	if len(a.includeExtensions) == 0 {
		return true
	}
	ext := filepath.Ext(path)
	for _, e := range a.includeExtensions {
		if e == ext {
			return true
		}
	}
	return false
}

func (a *Analyzer) isTooOld(modTime time.Time) bool {
	if a.maxAge == 0 {
		return false
	}
	ageSeconds := int64(time.Since(modTime).Seconds())
	return ageSeconds > a.maxAge
}

func (a *Analyzer) calculateHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func (a *Analyzer) findDuplicatesByHash() []*DuplicateGroup {
	hashMap := make(map[string]*DuplicateGroup)

	fmt.Print(color.CyanString("\nCalculating file hashes: "))
	count := 0

	for _, file := range a.files {
		count++
		if count%50 == 0 {
			fmt.Print(".")
		}

		hash, err := a.calculateHash(file.Path)
		if err != nil {
			continue
		}

		if group, exists := hashMap[hash]; exists {
			group.Files = append(group.Files, file.Path)
			group.Count++
			group.Size = file.Size
		} else {
			hashMap[hash] = &DuplicateGroup{
				Hash:  hash,
				Count: 1,
				Files: []string{file.Path},
				Size:  file.Size,
			}
		}
	}

	fmt.Println(" Done")

	var duplicates []*DuplicateGroup
	for _, group := range hashMap {
		if group.Count > 1 {
			duplicates = append(duplicates, group)
		}
	}

	sort.Slice(duplicates, func(i, j int) bool {
		return duplicates[i].Count > duplicates[j].Count
	})

	return duplicates
}

func (a *Analyzer) calculateTotalSize() int64 {
	var total int64
	for _, file := range a.files {
		total += file.Size
	}
	return total
}

func (a *Analyzer) calculateDuplicateSpace(duplicates []*DuplicateGroup) int64 {
	var total int64
	for _, group := range duplicates {
		total += group.Size * int64(group.Count-1)
	}
	return total
}

func (a *Analyzer) groupByExtension() map[string]int {
	result := make(map[string]int)
	for _, file := range a.files {
		ext := file.Extension
		if ext == "" {
			ext = "no-extension"
		}
		result[ext]++
	}
	return result
}

func (a *Analyzer) calculateSizeByType() map[string]int64 {
	result := make(map[string]int64)
	for _, file := range a.files {
		ext := file.Extension
		if ext == "" {
			ext = "no-extension"
		}
		result[ext] += file.Size
	}
	return result
}

func (a *Analyzer) getTopLargeFiles(count int) []*FileInfo {
	sort.Slice(a.files, func(i, j int) bool {
		return a.files[i].Size > a.files[j].Size
	})

	if len(a.files) > count {
		return a.files[:count]
	}
	return a.files
}
