package main

import "testing"

func TestParseOptionsDefaults(t *testing.T) {
	opts, err := parseOptions(nil)
	if err != nil {
		t.Fatalf("parseOptions returned error: %v", err)
	}
	if opts.inputDir != defaultInputDir {
		t.Fatalf("inputDir = %s, want %s", opts.inputDir, defaultInputDir)
	}
	if opts.outputPDF != defaultOutputPDF {
		t.Fatalf("outputPDF = %s, want %s", opts.outputPDF, defaultOutputPDF)
	}
}

func TestParseOptionsRejectsUnexpectedArgs(t *testing.T) {
	if _, err := parseOptions([]string{"extra"}); err == nil {
		t.Fatal("parseOptions expected error for positional argument")
	}
}
