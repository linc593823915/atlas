package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	sectionBoardMetadata     = "Board Metadata"
	sectionTitle             = "Title"
	sectionExecutionLog      = "Execution Log"
	sectionReviewCheckpoints = "Review Checkpoints"
)

func loadIssueCards(root string) ([]issueCard, error) {
	paths, err := listIssueCardPaths(root)
	if err != nil {
		return nil, err
	}

	var cards []issueCard
	for _, path := range paths {
		card, parseErr := parseIssueCard(path)
		if parseErr != nil {
			return nil, parseErr
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func listIssueCardPaths(root string) ([]string, error) {
	base := filepath.Join(root, "docs", "atlas", "issues")
	var paths []string
	err := filepath.WalkDir(base, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
			return err
		}
		if filepath.Base(path) == "README.md" {
			return nil
		}
		paths = append(paths, path)
		return nil
	})
	sort.Strings(paths)
	return paths, err
}

func parseIssueCard(path string) (issueCard, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return issueCard{}, fmt.Errorf("read %s: %w", path, err)
	}
	return parseIssueCardContent(path, string(content))
}

func parseIssueCardContent(path, content string) (issueCard, error) {
	sections := splitSections(content)
	card := issueCard{
		Path:     path,
		ID:       strings.ToUpper(strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))),
		Month:    monthFromPath(path),
		Sections: sections,
	}
	card.Title = firstSectionLine(sections[sectionTitle])
	populateBoardMetadata(&card, sections[sectionBoardMetadata])
	return card, nil
}

func splitSections(content string) map[string]string {
	sections := make(map[string]string)
	parts := strings.Split(content, "\n## ")
	if len(parts) == 0 {
		return sections
	}
	sections["preamble"] = parts[0]
	for _, part := range parts[1:] {
		sectionsPart := strings.SplitN(part, "\n", 2)
		name := strings.TrimSpace(sectionsPart[0])
		body := ""
		if len(sectionsPart) == 2 {
			body = strings.TrimSpace(sectionsPart[1])
		}
		sections[name] = body
	}
	return sections
}

func monthFromPath(path string) string {
	dir := filepath.Base(filepath.Dir(path))
	return strings.TrimPrefix(dir, "month-")
}

func firstSectionLine(body string) string {
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			return strings.TrimSpace(line)
		}
	}
	return ""
}

func populateBoardMetadata(card *issueCard, body string) {
	fields := parseMetadataLines(body)
	card.Status = fields["status"]
	card.Owner = fields["owner"]
	card.TargetWeek = fields["target week"]
	card.StartDate = fields["start date"]
	card.FinishDate = fields["finish date"]
	card.Dependency = fields["dependencies"]
	card.RiskLevel = fields["risk level"]
}

func parseMetadataLines(body string) map[string]string {
	fields := make(map[string]string)
	for _, line := range strings.Split(body, "\n") {
		name, value, ok := parseMetadataLine(line)
		if ok {
			fields[name] = value
		}
	}
	return fields
}

func parseMetadataLine(line string) (string, string, bool) {
	trimmed := strings.TrimSpace(line)
	if !strings.HasPrefix(trimmed, "- ") {
		return "", "", false
	}
	parts := strings.SplitN(strings.TrimPrefix(trimmed, "- "), ":", 2)
	if len(parts) != 2 {
		return "", "", false
	}
	name := strings.TrimSpace(strings.ToLower(parts[0]))
	value := strings.TrimSpace(parts[1])
	return name, value, true
}
