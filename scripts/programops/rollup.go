package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

func runRollup(command command) error {
	state, err := buildRollupState(command.Root)
	if err != nil {
		return err
	}
	return writeRollupState(command.OutputDir, state)
}

func buildRollupState(root string) (rollupState, error) {
	return buildRollupStateWithDate(root, "")
}

func buildRollupStateWithDate(root, date string) (rollupState, error) {
	cards, err := loadIssueCards(root)
	if err != nil {
		return rollupState{}, err
	}
	months, err := buildMonthRollups(root, cards, resolveDate(date))
	if err != nil {
		return rollupState{}, err
	}
	quarters, err := buildQuarterRollups(root, months, resolveDate(date))
	if err != nil {
		return rollupState{}, err
	}
	return rollupState{GeneratedAt: resolveDate(date), Months: months, Quarters: quarters}, nil
}

func writeRollupState(outputDir string, state rollupState) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}
	if err := writeJSONState(outputDir, state); err != nil {
		return err
	}
	if err := writeGeneratedIndex(outputDir, state); err != nil {
		return err
	}
	if err := writeMonthRollups(outputDir, state.Months); err != nil {
		return err
	}
	if err := writeQuarterRollups(outputDir, state.Quarters); err != nil {
		return err
	}
	fmt.Printf("program ops rollup: wrote %s\n", outputDir)
	return nil
}

func buildMonthRollups(root string, cards []issueCard, date string) ([]monthRollup, error) {
	metadata, err := loadMonthMetadata(root)
	if err != nil {
		return nil, err
	}
	grouped := groupCardsByMonth(cards)
	keys := sortMonthKeys(grouped)
	var rollups []monthRollup
	for _, month := range keys {
		rollups = append(rollups, newMonthRollup(metadata[month], grouped[month], date))
	}
	return rollups, nil
}

func newMonthRollup(metadata monthMetadata, cards []issueCard, date string) monthRollup {
	summary := summarizeMonth(metadata.Month, cards)
	return monthRollup{
		Month:                 metadata.Month,
		GeneratedAt:           date,
		Theme:                 metadata.Theme,
		Objective:             metadata.Objective,
		TotalIssues:           summary.TotalIssues,
		StatusCounts:          summary.StatusCounts,
		RiskCounts:            summary.RiskCounts,
		MissingTargetWeek:     summary.MissingTargetWeek,
		MissingStartDate:      summary.MissingStartDate,
		MissingFinishDate:     summary.MissingFinishDate,
		LinkedEvidenceIssues:  0,
		MissingEvidenceIssues: summary.TotalIssues,
		HighestRiskIssue:      highestRiskIssue(cards),
		HighestRiskLevel:      highestRiskLevel(cards),
		CompletionRate:        ratioString(summary.StatusCounts["done"], summary.TotalIssues),
		ValidationCoverage:    ratioString(summary.StatusCounts["validated"]+summary.StatusCounts["done"], summary.TotalIssues),
		EvidenceCompleteness:  ratioString(0, summary.TotalIssues),
		MilestoneStatus:       milestoneStatus(summary),
		PaceAssessment:        paceAssessment(summary),
		BiggestRisk:           monthBiggestRisk(summary),
	}
}

func highestRiskIssue(cards []issueCard) string {
	sorted := sortCardsByRisk(cards)
	if len(sorted) == 0 {
		return ""
	}
	return sorted[0].ID
}

func highestRiskLevel(cards []issueCard) string {
	sorted := sortCardsByRisk(cards)
	if len(sorted) == 0 {
		return ""
	}
	return sorted[0].RiskLevel
}

func sortCardsByRisk(cards []issueCard) []issueCard {
	cloned := append([]issueCard(nil), cards...)
	sort.Slice(cloned, func(idx, jdx int) bool {
		return riskPriority(cloned[idx].RiskLevel) > riskPriority(cloned[jdx].RiskLevel)
	})
	return cloned
}

func riskPriority(level string) int {
	switch level {
	case "critical":
		return 4
	case "high":
		return 3
	case "medium":
		return 2
	case "low":
		return 1
	default:
		return 0
	}
}

func ratioString(done, total int) string {
	if total == 0 {
		return "0/0"
	}
	return fmt.Sprintf("%d/%d", done, total)
}

func milestoneStatus(summary monthSummary) string {
	if summary.StatusCounts["done"] == summary.TotalIssues {
		return "complete"
	}
	if summary.StatusCounts["in_progress"] > 0 || summary.StatusCounts["in_review"] > 0 {
		return "in progress"
	}
	return "planned"
}

func paceAssessment(summary monthSummary) string {
	if summary.StatusCounts["done"] == summary.TotalIssues {
		return "complete"
	}
	if summary.StatusCounts["blocked"] > 0 {
		return "at risk"
	}
	return "not started"
}

func monthBiggestRisk(summary monthSummary) string {
	if summary.MissingTargetWeek > 0 {
		return "missing target week metadata"
	}
	if summary.StatusCounts["blocked"] > 0 {
		return "blocked issues present"
	}
	return "evidence not linked yet"
}

func today() string {
	return time.Now().Format("2006-01-02")
}

func resolveDate(date string) string {
	if date != "" {
		return date
	}
	return today()
}
