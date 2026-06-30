package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/linc593823915/atlas/internal/logger"
)

const commandOutputLimit = 2000

const htmlTemplate = `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <style>
    @page {
      size: A4;
      margin: 18mm 16mm 20mm 16mm;
    }
    body {
      color: #202124;
      font-family: -apple-system, BlinkMacSystemFont, "PingFang SC", "Hiragino Sans GB", "Noto Sans CJK SC", "Microsoft YaHei", sans-serif;
      font-size: 12.5pt;
      line-height: 1.68;
    }
    h1, h2, h3, h4 {
      color: #111827;
      line-height: 1.3;
      page-break-after: avoid;
    }
    h1 {
      font-size: 24pt;
      border-bottom: 1px solid #d1d5db;
      padding-bottom: 8px;
    }
    h2 {
      font-size: 18pt;
      margin-top: 28px;
    }
    h3 {
      font-size: 14.5pt;
    }
    a {
      color: #1a73e8;
      text-decoration: none;
    }
    code {
      font-family: "SFMono-Regular", Consolas, "Liberation Mono", monospace;
      font-size: 0.9em;
      white-space: pre-wrap;
      overflow-wrap: anywhere;
    }
    pre {
      background: #f6f8fa;
      border: 1px solid #d0d7de;
      border-radius: 6px;
      padding: 10px 12px;
      white-space: pre-wrap;
      overflow-wrap: anywhere;
    }
    table {
      border-collapse: collapse;
      width: 100%;
      page-break-inside: avoid;
    }
    th, td {
      border: 1px solid #d1d5db;
      padding: 6px 8px;
      vertical-align: top;
    }
    blockquote {
      border-left: 4px solid #d1d5db;
      color: #4b5563;
      margin-left: 0;
      padding-left: 12px;
    }
    .source-path {
      color: #6b7280;
      font-size: 9.5pt;
      margin-bottom: 16px;
    }
    .page-break {
      break-before: page;
      page-break-before: always;
    }
  </style>
</head>
<body>
$body$
</body>
</html>
`

// run parses options, collects source files, and starts PDF generation.
func run(ctx context.Context, args []string) error {
	opts, err := parseOptions(args)
	if err != nil {
		return err
	}
	files, err := collectMarkdownFiles(opts.inputDir)
	if err != nil {
		return err
	}
	logger.Info(ctx, "collected cookbook markdown files", "count", len(files), "input", opts.inputDir)
	return generatePDF(ctx, opts, files)
}

// generatePDF writes intermediate files and renders the final PDF.
func generatePDF(ctx context.Context, opts options, files []markdownFile) error {
	workDir, cleanup, err := prepareWorkDir(opts)
	if err != nil {
		return err
	}
	defer cleanup()

	mergedPath := filepath.Join(workDir, "cookbook.md")
	htmlPath := filepath.Join(workDir, "cookbook.html")
	templatePath := filepath.Join(workDir, "template.html")

	if err := writeMergedMarkdown(mergedPath, files); err != nil {
		return err
	}
	if err := os.WriteFile(templatePath, []byte(htmlTemplate), 0644); err != nil {
		return fmt.Errorf("write html template: %w", err)
	}
	if err := renderHTML(ctx, mergedPath, templatePath, htmlPath); err != nil {
		return err
	}
	if err := renderPDF(ctx, htmlPath, opts.outputPDF); err != nil {
		return err
	}
	logger.Info(ctx, "cookbook pdf generated", "output", opts.outputPDF, "workdir", workDir)
	return nil
}

// prepareWorkDir creates or reuses the directory for intermediate artifacts.
func prepareWorkDir(opts options) (string, func(), error) {
	if opts.workDir != "" {
		if err := os.MkdirAll(opts.workDir, 0755); err != nil {
			return "", nil, fmt.Errorf("create workdir %s: %w", opts.workDir, err)
		}
		return opts.workDir, func() {}, nil
	}
	workDir, err := os.MkdirTemp("", "atlas-cookbookpdf-*")
	if err != nil {
		return "", nil, fmt.Errorf("create temporary workdir: %w", err)
	}
	cleanup := func() {
		if opts.keepWorkDir {
			return
		}
		_ = os.RemoveAll(workDir)
	}
	return workDir, cleanup, nil
}

// writeMergedMarkdown concatenates source documents with source markers and page breaks.
func writeMergedMarkdown(path string, files []markdownFile) error {
	var buffer bytes.Buffer
	buffer.WriteString("% AI Agent 架构工程师一年自学教材\n")
	buffer.WriteString("% Atlas\n")
	buffer.WriteString("% 2026-06-30\n\n")

	for index, file := range files {
		if index > 0 {
			buffer.WriteString("\n<div class=\"page-break\"></div>\n\n")
		}
		buffer.WriteString(fmt.Sprintf("<div class=\"source-path\">来源：%s</div>\n\n", file.RelPath))
		content, err := os.ReadFile(file.AbsPath)
		if err != nil {
			return fmt.Errorf("read markdown %s: %w", file.AbsPath, err)
		}
		buffer.Write(content)
		buffer.WriteString("\n")
	}

	if err := os.WriteFile(path, buffer.Bytes(), 0644); err != nil {
		return fmt.Errorf("write merged markdown %s: %w", path, err)
	}
	return nil
}

// renderHTML invokes pandoc to convert the merged Markdown into standalone HTML.
func renderHTML(ctx context.Context, markdownPath, templatePath, htmlPath string) error {
	args := []string{
		"--from", "gfm+tex_math_dollars",
		"--to", "html5",
		"--standalone",
		"--toc",
		"--toc-depth=3",
		"--template", templatePath,
		"--metadata", "title=AI Agent 架构工程师一年自学教材",
		"--output", htmlPath,
		markdownPath,
	}
	if err := runCommand(ctx, "pandoc", args...); err != nil {
		return err
	}
	return nil
}

// renderPDF invokes Chrome headless to print the generated HTML to PDF.
func renderPDF(ctx context.Context, htmlPath, outputPath string) error {
	chromePath, err := findChrome()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}
	absHTML, err := filepath.Abs(htmlPath)
	if err != nil {
		return fmt.Errorf("resolve html path: %w", err)
	}
	absOutput, err := filepath.Abs(outputPath)
	if err != nil {
		return fmt.Errorf("resolve output path: %w", err)
	}
	args := []string{
		"--headless",
		"--disable-gpu",
		"--no-sandbox",
		"--print-to-pdf=" + absOutput,
		"file://" + absHTML,
	}
	return runCommand(ctx, chromePath, args...)
}

// runCommand executes an external renderer and records bounded command output.
func runCommand(ctx context.Context, name string, args ...string) error {
	logger.Info(ctx, "running command", "name", name, "args", strings.Join(args, " "))
	command := exec.CommandContext(ctx, name, args...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run %s: %w\n%s", name, err, string(output))
	}
	if len(output) > 0 {
		logger.Info(ctx, "command output", "name", name, "output", trimCommandOutput(output))
	}
	return nil
}

// trimCommandOutput bounds noisy renderer output before logging it.
func trimCommandOutput(output []byte) string {
	text := strings.TrimSpace(string(output))
	if len(text) <= commandOutputLimit {
		return text
	}
	return text[:commandOutputLimit] + "... output truncated"
}

// findChrome locates a Chrome or Chromium executable supported by the script.
func findChrome() (string, error) {
	candidates := []string{
		"google-chrome",
		"chromium",
		"chromium-browser",
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
	}
	for _, candidate := range candidates {
		if strings.Contains(candidate, "/") {
			if _, err := os.Stat(candidate); err == nil {
				return candidate, nil
			}
			continue
		}
		path, err := exec.LookPath(candidate)
		if err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("chrome or chromium executable not found")
}
