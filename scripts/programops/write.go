package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func writeJSONState(outputDir string, state rollupState) error {
	path := filepath.Join(outputDir, "program-state.json")
	content, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(content, '\n'), 0o644)
}

func writeGeneratedIndex(outputDir string, state rollupState) error {
	path := filepath.Join(outputDir, "README.md")
	return os.WriteFile(path, []byte(renderGeneratedIndex(state)), 0o644)
}

func writeMonthRollups(outputDir string, months []monthRollup) error {
	for _, month := range months {
		if err := writeMonthSnapshot(outputDir, month); err != nil {
			return err
		}
		if err := writeMonthScorecardDraft(outputDir, month); err != nil {
			return err
		}
	}
	return nil
}

func writeMonthSnapshot(outputDir string, month monthRollup) error {
	path := filepath.Join(outputDir, "month-"+month.Month+"-snapshot.md")
	return os.WriteFile(path, []byte(renderMonthSnapshot(month)), 0o644)
}

func writeMonthScorecardDraft(outputDir string, month monthRollup) error {
	path := filepath.Join(outputDir, "month-"+month.Month+"-scorecard-draft.md")
	return os.WriteFile(path, []byte(renderMonthScorecardDraft(month)), 0o644)
}

func writeQuarterRollups(outputDir string, quarters []quarterRollup) error {
	for _, quarter := range quarters {
		path := filepath.Join(outputDir, strings.ToLower(quarter.Quarter)+"-kpi-draft.md")
		if err := os.WriteFile(path, []byte(renderQuarterKPIDraft(quarter)), 0o644); err != nil {
			return err
		}
	}
	return nil
}
