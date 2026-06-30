package main

import (
	"fmt"
	"strings"
)

func renderGeneratedIndex(state rollupState) string {
	lines := []string{
		"# Generated Reports Index",
		"",
		fmt.Sprintf("Last updated: %s", state.GeneratedAt),
		"",
		"## Generated State",
		"",
		"- [program-state.json](program-state.json)",
		"",
		"## Generated Monthly Snapshots",
		"",
	}
	for _, month := range state.Months {
		lines = append(lines, fmt.Sprintf("- [Month %s Snapshot](month-%s-snapshot.md)", month.Month, month.Month))
	}
	lines = append(lines, "", "## Generated Monthly Scorecard Drafts", "")
	for _, month := range state.Months {
		lines = append(lines, fmt.Sprintf("- [Month %s Scorecard Draft](month-%s-scorecard-draft.md)", month.Month, month.Month))
	}
	lines = append(lines, "", "## Generated Quarterly KPI Drafts", "")
	for _, quarter := range state.Quarters {
		lines = append(lines, fmt.Sprintf("- [%s KPI Draft](%s-kpi-draft.md)", quarter.Quarter, strings.ToLower(quarter.Quarter)))
	}
	return strings.Join(lines, "\n") + "\n"
}

func renderMonthSnapshot(month monthRollup) string {
	lines := snapshotHeader(month)
	lines = append(lines, snapshotBoardCounts(month)...)
	lines = append(lines, snapshotMetadataGaps(month)...)
	lines = append(lines, snapshotEvidence(month)...)
	lines = append(lines, snapshotConclusion(month)...)
	return strings.Join(lines, "\n") + "\n"
}

func snapshotHeader(month monthRollup) []string {
	return []string{
		fmt.Sprintf("# Month %s Generated Snapshot", month.Month),
		"",
		fmt.Sprintf("Last updated: %s", month.GeneratedAt),
		"",
		fmt.Sprintf("- Source Month: [Month %s](../../months/month-%s.md)", month.Month, month.Month),
		fmt.Sprintf("- Source Backlog: [Month %s Backlog](../../atlas/backlog/month-%s.md)", month.Month, month.Month),
		"",
		"## Snapshot",
		"",
		fmt.Sprintf("- snapshot date: %s", month.GeneratedAt),
		fmt.Sprintf("- scope: Month %s", month.Month),
		"- owner: learner",
	}
}

func snapshotBoardCounts(month monthRollup) []string {
	return []string{
		"",
		"## Board Counts",
		"",
		fmt.Sprintf("- planned: %d", month.StatusCounts["planned"]),
		fmt.Sprintf("- ready: %d", month.StatusCounts["ready"]),
		fmt.Sprintf("- in_progress: %d", month.StatusCounts["in_progress"]),
		fmt.Sprintf("- in_review: %d", month.StatusCounts["in_review"]),
		fmt.Sprintf("- blocked: %d", month.StatusCounts["blocked"]),
		fmt.Sprintf("- validated: %d", month.StatusCounts["validated"]),
		fmt.Sprintf("- done: %d", month.StatusCounts["done"]),
	}
}

func snapshotMetadataGaps(month monthRollup) []string {
	return []string{
		"",
		"## Metadata Gaps",
		"",
		fmt.Sprintf("- missing target week: %d", month.MissingTargetWeek),
		fmt.Sprintf("- missing start date: %d", month.MissingStartDate),
		fmt.Sprintf("- missing finish date: %d", month.MissingFinishDate),
		"",
		"## Highest-Risk Issue",
		"",
		fmt.Sprintf("- issue id: %s", month.HighestRiskIssue),
		fmt.Sprintf("- risk: %s", month.HighestRiskLevel),
	}
}

func snapshotEvidence(month monthRollup) []string {
	return []string{
		"",
		"## Evidence Health",
		"",
		fmt.Sprintf("- issues with linked evidence: %d", month.LinkedEvidenceIssues),
		fmt.Sprintf("- issues missing evidence: %d", month.MissingEvidenceIssues),
	}
}

func snapshotConclusion(month monthRollup) []string {
	return []string{
		"",
		"## Snapshot Conclusion",
		"",
		fmt.Sprintf("- pace assessment: %s", month.PaceAssessment),
		fmt.Sprintf("- biggest risk: %s", month.BiggestRisk),
	}
}

func renderMonthScorecardDraft(month monthRollup) string {
	lines := scorecardHeader(month)
	lines = append(lines, scorecardDeliveryMetrics(month)...)
	lines = append(lines, scorecardAtlasOutcome(month)...)
	return strings.Join(lines, "\n") + "\n"
}

func scorecardHeader(month monthRollup) []string {
	return []string{
		fmt.Sprintf("# Month %s Generated Scorecard Draft", month.Month),
		"",
		fmt.Sprintf("Last updated: %s", month.GeneratedAt),
		"",
		fmt.Sprintf("- Source Month: [Month %s](../../months/month-%s.md)", month.Month, month.Month),
		fmt.Sprintf("- Source Snapshot: [Month %s Generated Snapshot](month-%s-snapshot.md)", month.Month, month.Month),
		"",
		"## Scorecard",
		"",
		fmt.Sprintf("- month: Month %s", month.Month),
		fmt.Sprintf("- objective: %s", month.Objective),
		"- overall grade: draft",
		"- promotion signal: pending review",
	}
}

func scorecardDeliveryMetrics(month monthRollup) []string {
	return []string{
		"",
		"## Delivery Metrics",
		"",
		fmt.Sprintf("- planned issues: %d", month.TotalIssues),
		fmt.Sprintf("- done issues: %d", month.StatusCounts["done"]),
		fmt.Sprintf("- blocked issues: %d", month.StatusCounts["blocked"]),
		fmt.Sprintf("- completion rate: %s", month.CompletionRate),
		fmt.Sprintf("- validation coverage: %s", month.ValidationCoverage),
		fmt.Sprintf("- evidence completeness: %s", month.EvidenceCompleteness),
	}
}

func scorecardAtlasOutcome(month monthRollup) []string {
	return []string{
		"",
		"## Atlas Outcome",
		"",
		fmt.Sprintf("- milestone shipped: %s", month.MilestoneStatus),
		fmt.Sprintf("- strongest subsystem improvement: %s", month.Theme),
		fmt.Sprintf("- most fragile subsystem: %s", month.BiggestRisk),
	}
}

func renderQuarterKPIDraft(quarter quarterRollup) string {
	lines := []string{
		fmt.Sprintf("# %s Generated KPI Draft", quarter.Quarter),
		"",
		fmt.Sprintf("Last updated: %s", quarter.GeneratedAt),
		"",
		fmt.Sprintf("- Source Quarter: [Quarter %s](../../quarters/quarter-%s.md)", quarter.Quarter[1:], quarter.Quarter[1:]),
		"",
		"## KPI Sheet",
		"",
		fmt.Sprintf("- quarter: %s", quarter.Quarter),
		fmt.Sprintf("- theme: %s", quarter.Theme),
		fmt.Sprintf("- quarter objective: %s", quarter.Objective),
		"",
		"## Outcome KPIs",
		"",
		fmt.Sprintf("- issue completion trend: %s", quarter.CompletionTrend),
		fmt.Sprintf("- blocked aging trend: %s", quarter.BlockedTrend),
		fmt.Sprintf("- validation coverage trend: %s", quarter.ValidationTrend),
		fmt.Sprintf("- metadata completeness risk: %d missing target week entries", quarter.MissingTargetWeek),
		"",
		"## Draft Judgment",
		"",
		fmt.Sprintf("- quarter overall: %s", quarter.QuarterStatus),
		fmt.Sprintf("- biggest risk: %s", quarter.BiggestRisk),
	}
	return strings.Join(lines, "\n") + "\n"
}
