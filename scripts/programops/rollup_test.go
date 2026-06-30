package main

import "testing"

func TestNewMonthRollup(t *testing.T) {
	metadata := monthMetadata{
		Month:     "01",
		Theme:     "backend foundation",
		Objective: "Build a clean backend base",
	}
	cards := []issueCard{
		{ID: "a", Month: "01", Status: "planned", RiskLevel: "medium"},
		{ID: "b", Month: "01", Status: "done", RiskLevel: "high", TargetWeek: "Week 01"},
	}

	rollup := newMonthRollup(metadata, cards, "2026-06-30")
	if rollup.TotalIssues != 2 {
		t.Fatalf("TotalIssues = %d, want 2", rollup.TotalIssues)
	}
	if rollup.HighestRiskIssue != "b" {
		t.Fatalf("HighestRiskIssue = %s, want b", rollup.HighestRiskIssue)
	}
	if rollup.MissingTargetWeek != 1 {
		t.Fatalf("MissingTargetWeek = %d, want 1", rollup.MissingTargetWeek)
	}
}

func TestNewQuarterRollup(t *testing.T) {
	metadata := quarterMetadata{
		Quarter:   "Q1",
		Theme:     "foundation",
		Objective: "Deliver the first agent",
	}
	months := []monthRollup{
		{Month: "01", TotalIssues: 5, StatusCounts: map[string]int{"planned": 5}, ValidationCoverage: "0/5", CompletionRate: "0/5", MissingTargetWeek: 5},
		{Month: "02", TotalIssues: 5, StatusCounts: map[string]int{"done": 2}, ValidationCoverage: "2/5", CompletionRate: "2/5"},
	}

	rollup := newQuarterRollup(metadata, months, "2026-06-30")
	if rollup.TotalIssues != 10 {
		t.Fatalf("TotalIssues = %d, want 10", rollup.TotalIssues)
	}
	if rollup.QuarterStatus != "in progress" {
		t.Fatalf("QuarterStatus = %s, want in progress", rollup.QuarterStatus)
	}
	if rollup.MissingTargetWeek != 5 {
		t.Fatalf("MissingTargetWeek = %d, want 5", rollup.MissingTargetWeek)
	}
}
