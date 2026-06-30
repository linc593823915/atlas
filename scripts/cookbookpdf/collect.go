package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type markdownFile struct {
	AbsPath string
	RelPath string
}

// collectMarkdownFiles recursively collects Markdown files and applies cookbook ordering.
func collectMarkdownFiles(inputDir string) ([]markdownFile, error) {
	var files []markdownFile
	if err := validateInputDir(inputDir); err != nil {
		return nil, err
	}
	if err := filepath.WalkDir(inputDir, collectMarkdownFile(inputDir, &files)); err != nil {
		return nil, fmt.Errorf("collect markdown files from %s: %w", inputDir, err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no markdown files found in %s", inputDir)
	}
	sortMarkdownFiles(files)
	return files, nil
}

// validateInputDir ensures the configured input path exists and is a directory.
func validateInputDir(inputDir string) error {
	info, err := os.Stat(inputDir)
	if err != nil {
		return fmt.Errorf("stat input directory %s: %w", inputDir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("input path is not a directory: %s", inputDir)
	}
	return nil
}

// collectMarkdownFile returns a WalkDir callback that appends Markdown files.
func collectMarkdownFile(inputDir string, files *[]markdownFile) fs.WalkDirFunc {
	return func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil || entry.IsDir() {
			return walkErr
		}
		if !strings.EqualFold(filepath.Ext(entry.Name()), ".md") {
			return nil
		}
		file, err := newMarkdownFile(inputDir, path)
		if err != nil {
			return err
		}
		*files = append(*files, file)
		return nil
	}
}

// newMarkdownFile builds absolute and slash-normalized relative file metadata.
func newMarkdownFile(inputDir, path string) (markdownFile, error) {
	relPath, err := filepath.Rel(inputDir, path)
	if err != nil {
		return markdownFile{}, err
	}
	return markdownFile{
		AbsPath: path,
		RelPath: filepath.ToSlash(relPath),
	}, nil
}

// sortMarkdownFiles orders files by cookbook reading flow before path fallback.
func sortMarkdownFiles(files []markdownFile) {
	sort.SliceStable(files, func(i, j int) bool {
		left := sortKey(files[i].RelPath)
		right := sortKey(files[j].RelPath)
		if left.group != right.group {
			return left.group < right.group
		}
		if left.rank != right.rank {
			return left.rank < right.rank
		}
		return files[i].RelPath < files[j].RelPath
	})
}

type fileSortKey struct {
	group int
	rank  int
}

// sortKey maps a relative path into a stable cookbook section rank.
func sortKey(relPath string) fileSortKey {
	name := filepath.Base(relPath)
	if rank, ok := explicitRootOrder[name]; ok && !strings.Contains(relPath, "/") {
		return fileSortKey{group: 0, rank: rank}
	}
	if strings.HasPrefix(name, "semester-") {
		return fileSortKey{group: 1, rank: numberAfterPrefix(name, "semester-")}
	}
	if strings.HasPrefix(name, "month-") {
		return fileSortKey{group: 2, rank: numberAfterPrefix(name, "month-")}
	}
	if strings.HasPrefix(name, "week-index-") {
		return fileSortKey{group: 3, rank: numberAfterPrefix(name, "week-index-semester-")}
	}
	if strings.HasPrefix(name, "week-") {
		return fileSortKey{group: 4, rank: numberAfterPrefix(name, "week-")}
	}
	if strings.HasPrefix(name, "day-index-") {
		return fileSortKey{group: 5, rank: dayIndexRank(name)}
	}
	if strings.HasPrefix(relPath, "days-") {
		return fileSortKey{group: 7, rank: dayDirectoryRank(relPath)}
	}
	return fileSortKey{group: 6, rank: fallbackRank(name)}
}

var explicitRootOrder = map[string]int{
	"README.md":             1,
	"how-to-use.md":         2,
	"year-overview.md":      3,
	"glossary.md":           4,
	"source-reading-map.md": 5,
}

// dayIndexRank keeps day index pages in their intended reading sequence.
func dayIndexRank(name string) int {
	switch name {
	case "day-index-month-01.md":
		return 1
	case "day-index-semester-1-part-1.md":
		return 2
	case "day-index-semester-1.md":
		return 3
	case "day-index-semester-2-part-1.md":
		return 4
	case "day-index-semester-2.md":
		return 5
	case "day-index-semester-3-part-1.md":
		return 6
	case "day-index-semester-3.md":
		return 7
	case "day-index-semester-4.md":
		return 8
	default:
		return fallbackRank(name)
	}
}

// dayDirectoryRank keeps day pages grouped by their source day directory.
func dayDirectoryRank(relPath string) int {
	parts := strings.Split(relPath, "/")
	if len(parts) < 2 {
		return fallbackRank(relPath)
	}
	dir := parts[0]
	name := parts[len(parts)-1]
	dirRank := 99
	switch {
	case strings.HasPrefix(dir, "days-month-"):
		dirRank = numberAfterPrefix(dir, "days-month-")
	case strings.HasPrefix(dir, "days-semester-"):
		dirRank = 10 + numberAfterPrefix(dir, "days-semester-")
	}
	return dirRank*1000 + numberAfterPrefix(name, "day-")
}

// fallbackRank provides deterministic ordering for uncategorized root documents.
func fallbackRank(name string) int {
	rank := 0
	for _, char := range name {
		rank += int(char)
	}
	return rank
}

// numberAfterPrefix extracts a leading numeric rank from a known filename prefix.
func numberAfterPrefix(name, prefix string) int {
	if !strings.HasPrefix(name, prefix) {
		return 999
	}
	value := strings.TrimPrefix(name, prefix)
	value = strings.TrimSuffix(value, ".md")
	digits := ""
	for _, char := range value {
		if char < '0' || char > '9' {
			break
		}
		digits += string(char)
	}
	if digits == "" {
		return 999
	}
	number := 0
	for _, char := range digits {
		number = number*10 + int(char-'0')
	}
	return number
}
