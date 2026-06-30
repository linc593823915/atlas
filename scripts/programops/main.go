package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	command, err := parseCommand(args)
	if err != nil {
		return err
	}
	switch command.Name {
	case commandSummary:
		return runSummary(command)
	case commandValidate:
		return runValidate(command)
	case commandRollup:
		return runRollup(command)
	case commandCheckref:
		return runCheckrefs(command)
	case commandEvidence:
		return runEvidence(command)
	default:
		return fmt.Errorf("unsupported command: %s", command.Name)
	}
}
