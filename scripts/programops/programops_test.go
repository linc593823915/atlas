package main

import "testing"

func TestParseIssueCardContent(t *testing.T) {
	card, err := parseIssueCardContent(sampleIssueCardPath(), sampleIssueCardContent())
	if err != nil {
		t.Fatalf("parseIssueCardContent returned error: %v", err)
	}
	assertParsedIssueCard(t, card)
}

func TestSummarizeProgram(t *testing.T) {
	cards := []issueCard{
		{Month: "01", Status: "planned", Owner: "learner", RiskLevel: "medium"},
		{Month: "01", Status: "done", Owner: "mentor", RiskLevel: "high", TargetWeek: "Week 01"},
		{Month: "04", Status: "blocked", Owner: "learner", RiskLevel: "critical"},
	}

	summary := summarizeProgram(cards)
	if len(summary.Months) != 2 {
		t.Fatalf("len(Months) = %d, want 2", len(summary.Months))
	}
	if len(summary.Quarters) != 2 {
		t.Fatalf("len(Quarters) = %d, want 2", len(summary.Quarters))
	}
	assertMonthSummary(t, summary.Months[0])
}

func TestValidateIssueCards(t *testing.T) {
	cards := []issueCard{{
		Path:      "sample.md",
		Status:    "planned",
		Owner:     "learner",
		RiskLevel: "medium",
		Sections: map[string]string{
			sectionBoardMetadata:     "",
			sectionTitle:             "",
			sectionExecutionLog:      "",
			sectionReviewCheckpoints: "",
		},
	}}

	issues := validateIssueCards(cards)
	if len(issues) != 0 {
		t.Fatalf("validateIssueCards returned %d issues, want 0", len(issues))
	}
}

func sampleIssueCardPath() string {
	return "docs/atlas/issues/month-01/atlas-m01-01.md"
}

func sampleIssueCardContent() string {
	return `# ATLAS-M01-01

## Board Metadata

- status: planned
- owner: learner
- target week: Week 01
- start date:
- finish date:
- dependencies:
- risk level: medium

## Title

Build bootstrap shell and runtime lifecycle boundaries.

## Execution Log

- planned approach:

## Review Checkpoints

- RFC checkpoint:
`
}

func assertParsedIssueCard(t *testing.T, card issueCard) {
	t.Helper()
	assertEqual(t, "Month", card.Month, "01")
	assertEqual(t, "Status", card.Status, "planned")
	assertEqual(t, "Owner", card.Owner, "learner")
	assertEqual(t, "TargetWeek", card.TargetWeek, "Week 01")
}

func assertMonthSummary(t *testing.T, summary monthSummary) {
	t.Helper()
	if summary.StatusCounts["planned"] != 1 {
		t.Fatalf("planned count = %d, want 1", summary.StatusCounts["planned"])
	}
	if summary.MissingTargetWeek != 1 {
		t.Fatalf("MissingTargetWeek = %d, want 1", summary.MissingTargetWeek)
	}
}

func assertEqual(t *testing.T, field, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("%s = %s, want %s", field, got, want)
	}
}
