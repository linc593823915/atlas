package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
)

type evidenceSummary struct {
	Months []monthEvidence `json:"months"`
}

type monthEvidence struct {
	Month                       string `json:"month"`
	TotalIssues                 int    `json:"total_issues"`
	MissingCurrentProgress      int    `json:"missing_current_progress"`
	MissingNextAction           int    `json:"missing_next_action"`
	MissingValidationCheckpoint int    `json:"missing_validation_checkpoint"`
	MissingDocCheckpoint        int    `json:"missing_documentation_checkpoint"`
	ReviewMissingGrade          bool   `json:"review_missing_grade"`
	ScorecardMissingGrade       bool   `json:"scorecard_missing_grade"`
}

func runEvidence(command command) error {
	summary, err := buildEvidenceSummary(command.Root)
	if err != nil {
		return err
	}
	if command.JSON {
		return printEvidenceJSON(summary)
	}
	printEvidenceText(summary)
	return nil
}

func buildEvidenceSummary(root string) (evidenceSummary, error) {
	cards, err := loadIssueCards(root)
	if err != nil {
		return evidenceSummary{}, err
	}
	grouped := groupCardsByMonth(cards)
	months := sortMonthKeys(grouped)
	var summaries []monthEvidence
	for _, month := range months {
		item, itemErr := summarizeEvidenceMonth(root, month, grouped[month])
		if itemErr != nil {
			return evidenceSummary{}, itemErr
		}
		summaries = append(summaries, item)
	}
	return evidenceSummary{Months: summaries}, nil
}

func summarizeEvidenceMonth(root, month string, cards []issueCard) (monthEvidence, error) {
	summary := monthEvidence{Month: month, TotalIssues: len(cards)}
	for _, card := range cards {
		summary.MissingCurrentProgress += countSectionFieldMissing(card.Sections[sectionExecutionLog], "current progress")
		summary.MissingNextAction += countSectionFieldMissing(card.Sections[sectionExecutionLog], "next action")
		summary.MissingValidationCheckpoint += countSectionFieldMissing(card.Sections[sectionReviewCheckpoints], "validation checkpoint")
		summary.MissingDocCheckpoint += countSectionFieldMissing(card.Sections[sectionReviewCheckpoints], "documentation checkpoint")
	}
	reviewPath := filepath.Join(root, "docs", "reviews", "archive", "month-"+month+"-review.md")
	reviewText, err := osReadFileString(reviewPath)
	if err != nil {
		return monthEvidence{}, err
	}
	scorecardPath := filepath.Join(root, "docs", "reports", "monthly", "month-"+month+"-scorecard.md")
	scorecardText, err := osReadFileString(scorecardPath)
	if err != nil {
		return monthEvidence{}, err
	}
	summary.ReviewMissingGrade = strings.Contains(reviewText, "- Overall monthly grade:")
	summary.ScorecardMissingGrade = strings.Contains(scorecardText, "- overall grade:")
	return summary, nil
}

func countSectionFieldMissing(body, field string) int {
	fields := parseMetadataLines(body)
	if fields[field] == "" {
		return 1
	}
	return 0
}

func printEvidenceJSON(summary evidenceSummary) error {
	payload, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(payload))
	return nil
}

func printEvidenceText(summary evidenceSummary) {
	fmt.Println("program ops evidence")
	for _, month := range summary.Months {
		fmt.Printf(
			"%s issues=%d progress_gaps=%d next_action_gaps=%d validation_gaps=%d doc_gaps=%d review_grade_missing=%t scorecard_grade_missing=%t\n",
			month.Month,
			month.TotalIssues,
			month.MissingCurrentProgress,
			month.MissingNextAction,
			month.MissingValidationCheckpoint,
			month.MissingDocCheckpoint,
			month.ReviewMissingGrade,
			month.ScorecardMissingGrade,
		)
	}
}
