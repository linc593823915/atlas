package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSortMarkdownFilesUsesCookbookOrder(t *testing.T) {
	files := []markdownFile{
		{RelPath: "days-semester-1/day-001.md"},
		{RelPath: "week-02-cli-http-bootstrap.md"},
		{RelPath: "README.md"},
		{RelPath: "month-01-backend-foundation.md"},
		{RelPath: "how-to-use.md"},
		{RelPath: "semester-1-foundation.md"},
	}

	sortMarkdownFiles(files)

	got := make([]string, 0, len(files))
	for _, file := range files {
		got = append(got, file.RelPath)
	}
	want := []string{
		"README.md",
		"how-to-use.md",
		"semester-1-foundation.md",
		"month-01-backend-foundation.md",
		"week-02-cli-http-bootstrap.md",
		"days-semester-1/day-001.md",
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("sorted[%d] = %s, want %s; got %#v", i, got[i], want[i], got)
		}
	}
}

func TestCollectMarkdownFilesRejectsEmptyDirectory(t *testing.T) {
	dir := t.TempDir()
	if _, err := collectMarkdownFiles(dir); err == nil {
		t.Fatal("collectMarkdownFiles expected error for empty directory")
	}
}

func TestCollectMarkdownFilesRecurses(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "README.md"), []byte("# Root"), 0644); err != nil {
		t.Fatal(err)
	}
	nested := filepath.Join(dir, "days-semester-1")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nested, "day-001.md"), []byte("# Day"), 0644); err != nil {
		t.Fatal(err)
	}

	files, err := collectMarkdownFiles(dir)
	if err != nil {
		t.Fatalf("collectMarkdownFiles returned error: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("len(files) = %d, want 2", len(files))
	}
	if files[0].RelPath != "README.md" || files[1].RelPath != "days-semester-1/day-001.md" {
		t.Fatalf("unexpected order: %#v", files)
	}
}
