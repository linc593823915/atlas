package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type monthMetadata struct {
	Month     string
	Path      string
	Theme     string
	Objective string
}

type quarterMetadata struct {
	Quarter   string
	Path      string
	Theme     string
	Objective string
}

func loadMonthMetadata(root string) (map[string]monthMetadata, error) {
	result := make(map[string]monthMetadata)
	for month := 1; month <= 12; month++ {
		item, err := readMonthMetadata(root, month)
		if err != nil {
			return nil, err
		}
		result[item.Month] = item
	}
	return result, nil
}

func loadQuarterMetadata(root string) (map[string]quarterMetadata, error) {
	result := make(map[string]quarterMetadata)
	for quarter := 1; quarter <= 4; quarter++ {
		item, err := readQuarterMetadata(root, quarter)
		if err != nil {
			return nil, err
		}
		result[item.Quarter] = item
	}
	return result, nil
}

func readMonthMetadata(root string, month int) (monthMetadata, error) {
	path := filepath.Join(root, "docs", "months", fmt.Sprintf("month-%02d.md", month))
	content, err := os.ReadFile(path)
	if err != nil {
		return monthMetadata{}, fmt.Errorf("read %s: %w", path, err)
	}
	sections := splitSections(string(content))
	return monthMetadata{
		Month:     fmt.Sprintf("%02d", month),
		Path:      path,
		Theme:     firstSectionLine(sections["Theme"]),
		Objective: firstSectionLine(sections["Objective"]),
	}, nil
}

func readQuarterMetadata(root string, quarter int) (quarterMetadata, error) {
	path := filepath.Join(root, "docs", "quarters", fmt.Sprintf("quarter-%d.md", quarter))
	content, err := os.ReadFile(path)
	if err != nil {
		return quarterMetadata{}, fmt.Errorf("read %s: %w", path, err)
	}
	sections := splitSections(string(content))
	return quarterMetadata{
		Quarter:   fmt.Sprintf("Q%d", quarter),
		Path:      path,
		Theme:     firstSectionLine(sections["Theme"]),
		Objective: firstSectionLine(sections["Quarter Objective"]),
	}, nil
}

func markdownPath(path string) string {
	normalized := filepath.ToSlash(path)
	return strings.TrimPrefix(normalized, "./")
}
