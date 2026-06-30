package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var markdownLinkPattern = regexp.MustCompile(`\[[^\]]+\]\(([^)]+)\)`)

type refIssue struct {
	Path   string
	Target string
	Reason string
}

func runCheckrefs(command command) error {
	issues, err := collectReferenceIssues(command.Root)
	if err != nil {
		return err
	}
	if len(issues) == 0 {
		fmt.Println("program ops checkrefs: ok")
		return nil
	}
	printReferenceIssues(issues)
	return fmt.Errorf("program ops checkrefs: %d issue(s)", len(issues))
}

func collectReferenceIssues(root string) ([]refIssue, error) {
	paths, err := listMarkdownPaths(root)
	if err != nil {
		return nil, err
	}
	var issues []refIssue
	for _, path := range paths {
		fileIssues, fileErr := checkMarkdownRefs(path)
		if fileErr != nil {
			return nil, fileErr
		}
		issues = append(issues, fileIssues...)
	}
	return issues, nil
}

func listMarkdownPaths(root string) ([]string, error) {
	var paths []string
	paths = append(paths, filepath.Join(root, "README.md"))
	base := filepath.Join(root, "docs")
	err := filepath.WalkDir(base, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".md" {
			return err
		}
		paths = append(paths, path)
		return nil
	})
	sort.Strings(paths)
	return paths, err
}

func checkMarkdownRefs(path string) ([]refIssue, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	refs := extractLocalRefs(string(content))
	var issues []refIssue
	for _, ref := range refs {
		targetPath := resolveRefTarget(path, ref)
		if _, statErr := os.Stat(targetPath); statErr != nil {
			issues = append(issues, refIssue{
				Path:   path,
				Target: ref,
				Reason: "target does not exist",
			})
		}
	}
	return issues, nil
}

func extractLocalRefs(content string) []string {
	matches := markdownLinkPattern.FindAllStringSubmatch(content, -1)
	var refs []string
	for _, match := range matches {
		ref := strings.TrimSpace(match[1])
		if isLocalRef(ref) {
			refs = append(refs, ref)
		}
	}
	return refs
}

func isLocalRef(ref string) bool {
	if ref == "" || strings.HasPrefix(ref, "#") {
		return false
	}
	if strings.Contains(ref, "://") || strings.HasPrefix(ref, "mailto:") {
		return false
	}
	return true
}

func resolveRefTarget(path, ref string) string {
	clean := strings.SplitN(ref, "#", 2)[0]
	if filepath.IsAbs(clean) {
		return clean
	}
	return filepath.Clean(filepath.Join(filepath.Dir(path), clean))
}

func printReferenceIssues(issues []refIssue) {
	for _, issue := range issues {
		fmt.Printf("%s -> %s: %s\n", issue.Path, issue.Target, issue.Reason)
	}
}
