package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	violations, err := collectViolations(root)
	if err != nil {
		return err
	}
	if len(violations) == 0 {
		fmt.Println("rule of 30: ok")
		return nil
	}

	printViolations(violations)
	return fmt.Errorf("rule of 30: %d violation(s)", len(violations))
}
