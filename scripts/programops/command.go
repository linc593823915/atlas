package main

import (
	"flag"
	"fmt"
)

const (
	commandSummary  = "summary"
	commandValidate = "validate"
	commandRollup   = "rollup"
	commandCheckref = "checkrefs"
	commandEvidence = "evidence"
)

type command struct {
	Name      string
	JSON      bool
	Root      string
	OutputDir string
	Date      string
}

func parseCommand(args []string) (command, error) {
	if len(args) == 0 {
		return parseSummaryCommand(nil)
	}
	switch args[0] {
	case commandSummary:
		return parseSummaryCommand(args[1:])
	case commandValidate:
		return parseValidateCommand(args[1:])
	case commandRollup:
		return parseRollupCommand(args[1:])
	case commandCheckref:
		return parseCheckrefsCommand(args[1:])
	case commandEvidence:
		return parseEvidenceCommand(args[1:])
	default:
		return command{}, fmt.Errorf("unknown command: %s", args[0])
	}
}

func parseSummaryCommand(args []string) (command, error) {
	flags := flag.NewFlagSet(commandSummary, flag.ContinueOnError)
	json := flags.Bool("json", false, "print JSON output")
	root := flags.String("root", ".", "project root")
	date := flags.String("date", "", "override output date (YYYY-MM-DD)")
	if err := flags.Parse(args); err != nil {
		return command{}, err
	}
	return command{Name: commandSummary, JSON: *json, Root: *root, Date: *date}, nil
}

func parseValidateCommand(args []string) (command, error) {
	flags := flag.NewFlagSet(commandValidate, flag.ContinueOnError)
	root := flags.String("root", ".", "project root")
	date := flags.String("date", "", "override output date (YYYY-MM-DD)")
	if err := flags.Parse(args); err != nil {
		return command{}, err
	}
	return command{Name: commandValidate, Root: *root, Date: *date}, nil
}

func parseRollupCommand(args []string) (command, error) {
	flags := flag.NewFlagSet(commandRollup, flag.ContinueOnError)
	root := flags.String("root", ".", "project root")
	outputDir := flags.String("output-dir", "docs/reports/generated", "output directory")
	date := flags.String("date", "", "override output date (YYYY-MM-DD)")
	if err := flags.Parse(args); err != nil {
		return command{}, err
	}
	return command{
		Name:      commandRollup,
		Root:      *root,
		OutputDir: *outputDir,
		Date:      *date,
	}, nil
}

func parseCheckrefsCommand(args []string) (command, error) {
	flags := flag.NewFlagSet(commandCheckref, flag.ContinueOnError)
	root := flags.String("root", ".", "project root")
	if err := flags.Parse(args); err != nil {
		return command{}, err
	}
	return command{Name: commandCheckref, Root: *root}, nil
}

func parseEvidenceCommand(args []string) (command, error) {
	flags := flag.NewFlagSet(commandEvidence, flag.ContinueOnError)
	json := flags.Bool("json", false, "print JSON output")
	root := flags.String("root", ".", "project root")
	if err := flags.Parse(args); err != nil {
		return command{}, err
	}
	return command{Name: commandEvidence, JSON: *json, Root: *root}, nil
}
