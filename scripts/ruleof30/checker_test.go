package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheckSourceReportsFunctionAndStructViolations(t *testing.T) {
	source := strings.Join([]string{
		"package sample",
		"",
		"type oversized struct {",
		buildStructFields(21),
		"}",
		"",
		"func tooLong() {",
		buildBodyLines(29),
		"}",
		"",
	}, "\n")

	violations, err := checkSource("sample.go", []byte(source))
	if err != nil {
		t.Fatalf("checkSource returned error: %v", err)
	}

	assertContainsMessage(t, violations, "function tooLong has 31 lines (max 30)")
	assertContainsMessage(t, violations, "struct has 21 fields (max 20)")
}

func TestCheckSourceReportsFileLengthViolation(t *testing.T) {
	source := buildLongFile(301)

	violations, err := checkSource("sample.go", []byte(source))
	if err != nil {
		t.Fatalf("checkSource returned error: %v", err)
	}

	assertContainsMessage(t, violations, "file has 301 lines (max 300)")
}

func buildStructFields(count int) string {
	lines := make([]string, 0, count)
	for idx := 0; idx < count; idx++ {
		lines = append(lines, fmt.Sprintf("\tField%d string", idx))
	}
	return strings.Join(lines, "\n")
}

func buildBodyLines(count int) string {
	lines := make([]string, 0, count)
	for idx := 0; idx < count; idx++ {
		lines = append(lines, fmt.Sprintf("\t_ = %d", idx))
	}
	return strings.Join(lines, "\n")
}

func buildLongFile(lines int) string {
	body := make([]string, 0, lines)
	body = append(body, "package sample")
	for idx := 1; idx < lines; idx++ {
		body = append(body, fmt.Sprintf("// line %d", idx))
	}
	return strings.Join(body, "\n")
}

func assertContainsMessage(t *testing.T, violations []violation, want string) {
	t.Helper()
	for _, violation := range violations {
		if violation.Message == want {
			return
		}
	}
	t.Fatalf("violation %q not found in %#v", want, violations)
}
