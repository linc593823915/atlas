package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func runSummary(command command) error {
	cards, err := loadIssueCards(command.Root)
	if err != nil {
		return err
	}
	summary := summarizeProgram(cards)
	if command.JSON {
		return printJSONSummary(summary)
	}
	printTextSummary(summary)
	return nil
}

func summarizeProgram(cards []issueCard) programSummary {
	months := summarizeMonths(cards)
	quarters := summarizeQuarters(months)
	return programSummary{Months: months, Quarters: quarters}
}

func summarizeMonths(cards []issueCard) []monthSummary {
	grouped := groupCardsByMonth(cards)
	months := sortMonthKeys(grouped)
	var summaries []monthSummary
	for _, month := range months {
		summaries = append(summaries, summarizeMonth(month, grouped[month]))
	}
	return summaries
}

func groupCardsByMonth(cards []issueCard) map[string][]issueCard {
	grouped := make(map[string][]issueCard)
	for _, card := range cards {
		grouped[card.Month] = append(grouped[card.Month], card)
	}
	return grouped
}

func sortMonthKeys(grouped map[string][]issueCard) []string {
	var months []string
	for month := range grouped {
		months = append(months, month)
	}
	sort.Strings(months)
	return months
}

func summarizeMonth(month string, cards []issueCard) monthSummary {
	summary := monthSummary{
		Month:        month,
		TotalIssues:  len(cards),
		StatusCounts: make(map[string]int),
		OwnerCounts:  make(map[string]int),
		RiskCounts:   make(map[string]int),
	}
	for _, card := range cards {
		accumulateMonthSummary(&summary, card)
	}
	return summary
}

func accumulateMonthSummary(summary *monthSummary, card issueCard) {
	increment(summary.StatusCounts, card.Status)
	increment(summary.OwnerCounts, card.Owner)
	increment(summary.RiskCounts, card.RiskLevel)
	summary.MissingTargetWeek += countMissing(card.TargetWeek)
	summary.MissingStartDate += countMissing(card.StartDate)
	summary.MissingFinishDate += countMissing(card.FinishDate)
}

func increment(counter map[string]int, key string) {
	if key == "" {
		counter["<missing>"]++
		return
	}
	counter[key]++
}

func countMissing(value string) int {
	if value == "" {
		return 1
	}
	return 0
}

func summarizeQuarters(months []monthSummary) []quarterSummary {
	grouped := groupMonthsByQuarter(months)
	keys := sortedQuarterKeys(grouped)
	var summaries []quarterSummary
	for _, key := range keys {
		summaries = append(summaries, summarizeQuarter(key, grouped[key]))
	}
	return summaries
}

func groupMonthsByQuarter(months []monthSummary) map[string][]monthSummary {
	grouped := make(map[string][]monthSummary)
	for _, month := range months {
		grouped[quarterFromMonth(month.Month)] = append(grouped[quarterFromMonth(month.Month)], month)
	}
	return grouped
}

func quarterFromMonth(month string) string {
	switch month {
	case "01", "02", "03":
		return "Q1"
	case "04", "05", "06":
		return "Q2"
	case "07", "08", "09":
		return "Q3"
	default:
		return "Q4"
	}
}

func sortedQuarterKeys(grouped map[string][]monthSummary) []string {
	var keys []string
	for key := range grouped {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func summarizeQuarter(quarter string, months []monthSummary) quarterSummary {
	summary := quarterSummary{
		Quarter:      quarter,
		StatusCounts: make(map[string]int),
		RiskCounts:   make(map[string]int),
	}
	for _, month := range months {
		accumulateQuarterSummary(&summary, month)
	}
	return summary
}

func accumulateQuarterSummary(summary *quarterSummary, month monthSummary) {
	summary.TotalIssues += month.TotalIssues
	summary.MissingTargetWeek += month.MissingTargetWeek
	mergeCounts(summary.StatusCounts, month.StatusCounts)
	mergeCounts(summary.RiskCounts, month.RiskCounts)
}

func mergeCounts(target, source map[string]int) {
	for key, value := range source {
		target[key] += value
	}
}

func printJSONSummary(summary programSummary) error {
	payload, err := json.MarshalIndent(summary, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(payload))
	return nil
}

func printTextSummary(summary programSummary) {
	fmt.Println("program ops summary")
	for _, month := range summary.Months {
		printMonthSummary(month)
	}
	for _, quarter := range summary.Quarters {
		printQuarterSummary(quarter)
	}
}

func printMonthSummary(month monthSummary) {
	fmt.Printf(
		"%s issues=%d planned=%d in_progress=%d done=%d missing_target_week=%d\n",
		month.Month,
		month.TotalIssues,
		month.StatusCounts["planned"],
		month.StatusCounts["in_progress"],
		month.StatusCounts["done"],
		month.MissingTargetWeek,
	)
}

func printQuarterSummary(quarter quarterSummary) {
	fmt.Printf(
		"%s issues=%d planned=%d done=%d missing_target_week=%d\n",
		quarter.Quarter,
		quarter.TotalIssues,
		quarter.StatusCounts["planned"],
		quarter.StatusCounts["done"],
		quarter.MissingTargetWeek,
	)
}
