package main

import "fmt"

var requiredSections = []string{
	sectionBoardMetadata,
	sectionTitle,
	sectionExecutionLog,
	sectionReviewCheckpoints,
}

var validStatuses = map[string]struct{}{
	"planned":     {},
	"ready":       {},
	"in_progress": {},
	"in_review":   {},
	"blocked":     {},
	"validated":   {},
	"done":        {},
}

func runValidate(command command) error {
	cards, err := loadIssueCards(command.Root)
	if err != nil {
		return err
	}
	issues := validateIssueCards(cards)
	if len(issues) == 0 {
		fmt.Println("program ops validate: ok")
		return nil
	}
	printValidationIssues(issues)
	return fmt.Errorf("program ops validate: %d issue(s)", len(issues))
}

func validateIssueCards(cards []issueCard) []validationIssue {
	var issues []validationIssue
	for _, card := range cards {
		issues = append(issues, validateRequiredSections(card)...)
		issues = append(issues, validateStatus(card)...)
		issues = append(issues, validateOwner(card)...)
		issues = append(issues, validateRiskLevel(card)...)
	}
	return issues
}

func validateRequiredSections(card issueCard) []validationIssue {
	var issues []validationIssue
	for _, section := range requiredSections {
		if _, ok := card.Sections[section]; !ok {
			issues = append(issues, validationIssue{
				Path:    card.Path,
				Message: fmt.Sprintf("missing section: %s", section),
			})
		}
	}
	return issues
}

func validateStatus(card issueCard) []validationIssue {
	if _, ok := validStatuses[card.Status]; ok {
		return nil
	}
	return []validationIssue{{
		Path:    card.Path,
		Message: fmt.Sprintf("invalid status: %s", card.Status),
	}}
}

func validateOwner(card issueCard) []validationIssue {
	if card.Owner != "" {
		return nil
	}
	return []validationIssue{{
		Path:    card.Path,
		Message: "missing owner",
	}}
}

func validateRiskLevel(card issueCard) []validationIssue {
	switch card.RiskLevel {
	case "low", "medium", "high", "critical":
		return nil
	default:
		return []validationIssue{{
			Path:    card.Path,
			Message: fmt.Sprintf("invalid risk level: %s", card.RiskLevel),
		}}
	}
}

func printValidationIssues(issues []validationIssue) {
	for _, issue := range issues {
		fmt.Printf("%s: %s\n", issue.Path, issue.Message)
	}
}
