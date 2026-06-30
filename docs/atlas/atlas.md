# Atlas Engineering Track

Last updated: 2026-06-29

## Purpose

Atlas is the main engineering line that evolves together with the 52-week learning system.

Learning without a real system produces shallow understanding. Atlas exists to force every concept into architecture, interfaces, implementation, tests, documentation, and operational trade-offs.

This engineering track now follows a Google-style public-practices adaptation:

- monthly OKRs
- weekly shippable artifacts
- explicit code-health and reliability expectations
- quarterly calibration reviews
- interview-style architecture defense using Atlas as evidence

## Semester Evolution

### Semester 1

- Build bootstrap, config manager, logger, runtime shell, Postgres and Redis wiring
- Add one stable single-agent workflow with tool calling and structured output

### Semester 2

- Add tool gateway and contract validation
- Add MCP server surface
- Add multi-agent orchestration, handoffs, approvals, and guardrails

### Semester 3

- Add durable workflow and memory boundaries
- Add traces, graders, datasets, and release gates
- Add governance, policy, and audit controls

### Semester 4

- Add Kubernetes deployment, SLO model, and rollback path
- Add benchmark lab and capacity planning
- Integrate all subsystems into a graduation-platform architecture

## Core Subsystems

- Bootstrap
- Config center
- Logger and observability
- Runtime and workflow engine
- Tool gateway
- MCP integration layer
- Eval pipeline
- Approval and audit pipeline
- Deployment and operations layer

## Engineering Rules

- Atlas remains the main implementation target for all labs and projects
- No learning module is considered complete until it changes Atlas or explicitly proves why it should not
- Each semester ends with an Atlas architecture review
- Each month ends with an OKR review and mock interview
- Each quarter ends with a calibration review covering results, technical judgment, code health, reliability, communication, and learning

## Assessment Rhythm

- Daily: read, write, build, output, and evidence capture
- Weekly: one shippable artifact and five interview questions
- Monthly: one OKR review, one evidence pack, and one interview bank
- Quarterly: one architecture walkthrough, one mock panel, and one written retrospective

## Operating Artifacts

- [Governance Index](../governance/README.md)
- [Program Executive Summary](../governance/program-executive-summary.md)
- [Annual Readiness Matrix](../governance/annual-readiness-matrix.md)
- [Role Gap Analysis](../governance/role-gap-analysis.md)
- [Mentor and Reviewer Handbook](../governance/mentor-reviewer-handbook.md)
- [Steering Committee Pack](../governance/steering-committee-pack.md)
- [Annual Review Memo](../governance/annual-review-memo.md)
- [Promotion Committee Guide](../governance/promotion-committee-guide.md)
- [Quarter Trend Pack](../governance/quarter-trend-pack.md)
- [Exec Dashboard Pack](../governance/exec-dashboard-pack.md)
- [Promotion Packet Archive](../governance/promotion-packet-archive.md)
- [Committee Decision Log](../governance/committee-decision-log.md)
- [Quarterly Steering Review Template](../governance/quarterly-steering-review-template.md)
- [Exec Dashboard Sample](../governance/exec-dashboard-sample.md)
- [Promotion Packet Sample](../governance/promotion-packet-sample.md)
- [Committee Decision Examples](../governance/committee-decision-examples.md)
- [Quarterly Steering Review Sample](../governance/quarterly-steering-review-sample.md)
- [Month 01 End-to-End Sample](../governance/month-01-end-to-end-sample.md)
- [Quarter 1 End-to-End Sample](../governance/quarter-1-end-to-end-sample.md)
- [Sample Evidence Entries](../governance/sample-evidence-entries.md)
- [Sample Issue Board History](../governance/sample-issue-board-history.md)
- [Atlas Backlog Index](backlog/README.md)
- [Atlas Issue Card Index](issues/README.md)
- [Mentor Review Template](../reviews/mentor-review-template.md)
- [Promotion Rubric](../reviews/promotion-rubric.md)
- [Mock Panel Score Sheet](../reviews/mock-panel-score-sheet.md)
- [Hiring Bar Packet](../reviews/hiring-bar-packet.md)
- [Review Archive Index](../reviews/archive/README.md)
- [Issue Status Board](../operations/issue-status-board.md)
- [Review Calendar](../operations/review-calendar.md)
- [Evidence Registry](../operations/evidence-registry.md)
- [Operations Index](../operations/README.md)
- [Program Control Center](../operations/program-control-center.md)
- [Review Metrics](../operations/review-metrics.md)
- [Dashboard Export](../operations/dashboard-export.md)
- [Burnup and Burndown](../operations/burnup-burndown.md)
- [Program Weekly Report](../operations/program-weekly-report.md)
- [Monthly Scorecard](../operations/monthly-scorecard.md)
- [Quarterly KPI Sheet](../operations/quarterly-kpi-sheet.md)
- [Issue Board Snapshot Template](../operations/issue-board-snapshot-template.md)
- [Reporting Records Index](../reports/README.md)

## Backlog Rule

- Every roadmap month must map to one Atlas backlog file.
- Every monthly review must cite both roadmap evidence and Atlas backlog progress.
- Promotion or hiring readiness cannot be argued from interview answers alone; it must point back to Atlas artifacts.
- Every issue card should carry status, owner, dependency, risk, and evidence fields.

## Final State

At the end of the year, Atlas should look like an internal agent platform instead of a toy demo.

It should support:

- stable configuration
- tool contracts
- orchestrated agents
- durable workflow
- evaluation and regression gates
- approvals and audit
- production deployment
- benchmark evidence
