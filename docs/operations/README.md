# Operations Index

Last updated: 2026-06-29

## Purpose

This index groups the operating controls that make the roadmap measurable, schedulable, and auditable.

For sponsor-facing and evaluator-facing governance, see [Governance Index](../governance/README.md).

## Control Plane

- [Program Control Center](program-control-center.md)
- [Issue Status Board](issue-status-board.md)
- [Review Calendar](review-calendar.md)
- [Evidence Registry](evidence-registry.md)

## Metrics and Reporting

- [Review Metrics](review-metrics.md)
- [Dashboard Export](dashboard-export.md)
- [Burnup and Burndown](burnup-burndown.md)
- [Program Weekly Report](program-weekly-report.md)
- [Monthly Scorecard](monthly-scorecard.md)
- [Quarterly KPI Sheet](quarterly-kpi-sheet.md)
- [Issue Board Snapshot Template](issue-board-snapshot-template.md)
- [Reporting Records Index](../reports/README.md)

## Automation

- Run `go run ./scripts/programops validate` to validate issue-card metadata and structure.
- Run `go run ./scripts/programops checkrefs` to validate local markdown references across the program docs.
- Run `go run ./scripts/programops summary` to print month and quarter status summaries.
- Run `go run ./scripts/programops summary --json` for structured output.
- Run `go run ./scripts/programops evidence` to summarize evidence gaps and empty review fields.
- Run `go run ./scripts/programops evidence --json` for structured evidence-gap output.
- Run `go run ./scripts/programops rollup` to generate derived reports under `docs/reports/generated`.
- Use `--date YYYY-MM-DD` when generated output should reflect an explicit reporting date.
- Generated outputs are indexed in [reports/generated/README.md](../reports/generated/README.md).
