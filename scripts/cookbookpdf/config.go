package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

const (
	defaultInputDir  = "docs/cookbook-zh"
	defaultOutputPDF = "docs/cookbook-zh/atlas-cookbook-zh.pdf"
)

type options struct {
	inputDir    string
	outputPDF   string
	workDir     string
	keepWorkDir bool
}

// parseOptions converts command-line flags into validated script options.
func parseOptions(args []string) (options, error) {
	flags := flag.NewFlagSet("cookbookpdf", flag.ContinueOnError)
	flags.SetOutput(discardWriter{})

	var opts options
	flags.StringVar(&opts.inputDir, "input", defaultInputDir, "Markdown input directory")
	flags.StringVar(&opts.outputPDF, "output", defaultOutputPDF, "PDF output path")
	flags.StringVar(&opts.workDir, "workdir", "", "temporary working directory")
	flags.BoolVar(&opts.keepWorkDir, "keep-workdir", false, "keep generated intermediate files")

	if err := flags.Parse(args); err != nil {
		return options{}, err
	}
	if flags.NArg() > 0 {
		return options{}, fmt.Errorf("unexpected positional arguments: %v", flags.Args())
	}
	if opts.inputDir == "" {
		return options{}, fmt.Errorf("input directory is required")
	}
	if opts.outputPDF == "" {
		return options{}, fmt.Errorf("output path is required")
	}
	opts.inputDir = filepath.Clean(opts.inputDir)
	opts.outputPDF = filepath.Clean(opts.outputPDF)
	if opts.workDir != "" {
		opts.workDir = filepath.Clean(opts.workDir)
	}
	return opts, nil
}

type discardWriter struct{}

// Write discards flag package output while preserving the writer contract.
func (discardWriter) Write(p []byte) (int, error) {
	return len(p), nil
}
