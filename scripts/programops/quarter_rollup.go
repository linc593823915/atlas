package main

import (
	"fmt"
	"sort"
)

func buildQuarterRollups(root string, months []monthRollup, date string) ([]quarterRollup, error) {
	metadata, err := loadQuarterMetadata(root)
	if err != nil {
		return nil, err
	}
	grouped := groupMonthRollupsByQuarter(months)
	keys := sortedQuarterKeysFromRollups(grouped)
	var rollups []quarterRollup
	for _, key := range keys {
		rollups = append(rollups, newQuarterRollup(metadata[key], grouped[key], date))
	}
	return rollups, nil
}

func groupMonthRollupsByQuarter(months []monthRollup) map[string][]monthRollup {
	grouped := make(map[string][]monthRollup)
	for _, month := range months {
		grouped[quarterFromMonth(month.Month)] = append(grouped[quarterFromMonth(month.Month)], month)
	}
	return grouped
}

func sortedQuarterKeysFromRollups(grouped map[string][]monthRollup) []string {
	var keys []string
	for key := range grouped {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func newQuarterRollup(metadata quarterMetadata, months []monthRollup, date string) quarterRollup {
	rollup := quarterRollup{
		Quarter:      metadata.Quarter,
		GeneratedAt:  date,
		Theme:        metadata.Theme,
		Objective:    metadata.Objective,
		StatusCounts: make(map[string]int),
	}
	for _, month := range months {
		accumulateQuarterRollup(&rollup, month)
	}
	rollup.CompletionTrend = quarterCompletionTrend(months)
	rollup.BlockedTrend = quarterBlockedTrend(months)
	rollup.ValidationTrend = quarterValidationTrend(months)
	rollup.QuarterStatus = quarterStatus(rollup)
	rollup.BiggestRisk = quarterBiggestRisk(rollup)
	return rollup
}

func accumulateQuarterRollup(rollup *quarterRollup, month monthRollup) {
	rollup.TotalIssues += month.TotalIssues
	rollup.MissingTargetWeek += month.MissingTargetWeek
	mergeCounts(rollup.StatusCounts, month.StatusCounts)
}

func quarterCompletionTrend(months []monthRollup) string {
	return trendString(months, func(item monthRollup) string {
		return fmt.Sprintf("M%s %s", item.Month, item.CompletionRate)
	})
}

func quarterBlockedTrend(months []monthRollup) string {
	return trendString(months, func(item monthRollup) string {
		return fmt.Sprintf("M%s %d blocked", item.Month, item.StatusCounts["blocked"])
	})
}

func quarterValidationTrend(months []monthRollup) string {
	return trendString(months, func(item monthRollup) string {
		return fmt.Sprintf("M%s %s", item.Month, item.ValidationCoverage)
	})
}

func trendString(items []monthRollup, mapper func(monthRollup) string) string {
	var parts []string
	for _, item := range items {
		parts = append(parts, mapper(item))
	}
	return joinWithComma(parts)
}

func joinWithComma(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for _, part := range parts[1:] {
		result += ", " + part
	}
	return result
}

func quarterStatus(rollup quarterRollup) string {
	if rollup.StatusCounts["done"] == rollup.TotalIssues {
		return "complete"
	}
	if rollup.StatusCounts["blocked"] > 0 {
		return "at risk"
	}
	return "in progress"
}

func quarterBiggestRisk(rollup quarterRollup) string {
	if rollup.MissingTargetWeek > 0 {
		return "missing target week metadata across issue cards"
	}
	if rollup.StatusCounts["blocked"] > 0 {
		return "blocked issues remain in the quarter"
	}
	return "evidence linkage remains incomplete"
}
