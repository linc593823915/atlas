package main

type issueCard struct {
	ID         string
	Month      string
	Path       string
	Title      string
	Status     string
	Owner      string
	TargetWeek string
	StartDate  string
	FinishDate string
	Dependency string
	RiskLevel  string
	Sections   map[string]string
}

type monthSummary struct {
	Month             string         `json:"month"`
	TotalIssues       int            `json:"total_issues"`
	StatusCounts      map[string]int `json:"status_counts"`
	OwnerCounts       map[string]int `json:"owner_counts"`
	RiskCounts        map[string]int `json:"risk_counts"`
	MissingTargetWeek int            `json:"missing_target_week"`
	MissingStartDate  int            `json:"missing_start_date"`
	MissingFinishDate int            `json:"missing_finish_date"`
}

type quarterSummary struct {
	Quarter           string         `json:"quarter"`
	TotalIssues       int            `json:"total_issues"`
	StatusCounts      map[string]int `json:"status_counts"`
	RiskCounts        map[string]int `json:"risk_counts"`
	MissingTargetWeek int            `json:"missing_target_week"`
}

type programSummary struct {
	Months   []monthSummary   `json:"months"`
	Quarters []quarterSummary `json:"quarters"`
}

type validationIssue struct {
	Path    string
	Message string
}

type rollupState struct {
	GeneratedAt string          `json:"generated_at"`
	Months      []monthRollup   `json:"months"`
	Quarters    []quarterRollup `json:"quarters"`
}

type monthRollup struct {
	Month                 string         `json:"month"`
	GeneratedAt           string         `json:"generated_at"`
	Theme                 string         `json:"theme"`
	Objective             string         `json:"objective"`
	TotalIssues           int            `json:"total_issues"`
	StatusCounts          map[string]int `json:"status_counts"`
	RiskCounts            map[string]int `json:"risk_counts"`
	MissingTargetWeek     int            `json:"missing_target_week"`
	MissingStartDate      int            `json:"missing_start_date"`
	MissingFinishDate     int            `json:"missing_finish_date"`
	LinkedEvidenceIssues  int            `json:"linked_evidence_issues"`
	MissingEvidenceIssues int            `json:"missing_evidence_issues"`
	HighestRiskIssue      string         `json:"highest_risk_issue"`
	HighestRiskLevel      string         `json:"highest_risk_level"`
	CompletionRate        string         `json:"completion_rate"`
	ValidationCoverage    string         `json:"validation_coverage"`
	EvidenceCompleteness  string         `json:"evidence_completeness"`
	MilestoneStatus       string         `json:"milestone_status"`
	PaceAssessment        string         `json:"pace_assessment"`
	BiggestRisk           string         `json:"biggest_risk"`
}

type quarterRollup struct {
	Quarter           string         `json:"quarter"`
	GeneratedAt       string         `json:"generated_at"`
	Theme             string         `json:"theme"`
	Objective         string         `json:"objective"`
	TotalIssues       int            `json:"total_issues"`
	StatusCounts      map[string]int `json:"status_counts"`
	MissingTargetWeek int            `json:"missing_target_week"`
	CompletionTrend   string         `json:"completion_trend"`
	BlockedTrend      string         `json:"blocked_trend"`
	ValidationTrend   string         `json:"validation_trend"`
	QuarterStatus     string         `json:"quarter_status"`
	BiggestRisk       string         `json:"biggest_risk"`
}
