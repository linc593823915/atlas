# AI Agent Architect ROADMAP

Last updated: 2026-06-29

## Goal

This document is the 52-week master roadmap for becoming an enterprise AI agent architect using a Google-style public-practices adaptation.

The target is to design, build, evaluate, secure, operate, and evolve agent systems in production, using Atlas as the main evidence base.

## Interpretation Note

This roadmap is inspired by public Google material on OKRs, team effectiveness, coaching, code health, testing, and SRE. It is not a claim about Google’s private internal promotion ladder.

## Execution Rules

- Every issue follows: RFC -> Interface -> Implementation -> Unit Test -> Review -> Benchmark (when necessary) -> Documentation -> Memory
- Weekly study budget: 12-15 hours
- Daily execution format: Read -> Write -> Build -> Output
- Atlas evolves together with the course and remains the main engineering line

## Framework

- [Google-Style Capability Model](framework/google-style-capability-model.md)
- [Google Public Sources](framework/google-public-sources.md)
- [Governance Index](governance/README.md)
- [Program Executive Summary](governance/program-executive-summary.md)
- [Annual Readiness Matrix](governance/annual-readiness-matrix.md)
- [Role Gap Analysis](governance/role-gap-analysis.md)
- [Mentor and Reviewer Handbook](governance/mentor-reviewer-handbook.md)
- [Steering Committee Pack](governance/steering-committee-pack.md)
- [Annual Review Memo](governance/annual-review-memo.md)
- [Promotion Committee Guide](governance/promotion-committee-guide.md)
- [Quarter Trend Pack](governance/quarter-trend-pack.md)
- [Exec Dashboard Pack](governance/exec-dashboard-pack.md)
- [Promotion Packet Archive](governance/promotion-packet-archive.md)
- [Committee Decision Log](governance/committee-decision-log.md)
- [Quarterly Steering Review Template](governance/quarterly-steering-review-template.md)
- [Exec Dashboard Sample](governance/exec-dashboard-sample.md)
- [Promotion Packet Sample](governance/promotion-packet-sample.md)
- [Committee Decision Examples](governance/committee-decision-examples.md)
- [Quarterly Steering Review Sample](governance/quarterly-steering-review-sample.md)
- [Month 01 End-to-End Sample](governance/month-01-end-to-end-sample.md)
- [Quarter 1 End-to-End Sample](governance/quarter-1-end-to-end-sample.md)
- [Sample Evidence Entries](governance/sample-evidence-entries.md)
- [Sample Issue Board History](governance/sample-issue-board-history.md)
- [Mentor Review Template](reviews/mentor-review-template.md)
- [Promotion Rubric](reviews/promotion-rubric.md)
- [Mock Panel Score Sheet](reviews/mock-panel-score-sheet.md)
- [Hiring Bar Packet](reviews/hiring-bar-packet.md)
- [Review Archive Index](reviews/archive/README.md)
- [Issue Status Board](operations/issue-status-board.md)
- [Review Calendar](operations/review-calendar.md)
- [Evidence Registry](operations/evidence-registry.md)
- [Operations Index](operations/README.md)
- [Program Control Center](operations/program-control-center.md)
- [Review Metrics](operations/review-metrics.md)
- [Dashboard Export](operations/dashboard-export.md)
- [Burnup and Burndown](operations/burnup-burndown.md)
- [Program Weekly Report](operations/program-weekly-report.md)
- [Monthly Scorecard](operations/monthly-scorecard.md)
- [Quarterly KPI Sheet](operations/quarterly-kpi-sheet.md)
- [Issue Board Snapshot Template](operations/issue-board-snapshot-template.md)
- [Reporting Records Index](reports/README.md)

## Structure

- [Semester 1](semesters/semester-1.md)
- [Semester 2](semesters/semester-2.md)
- [Semester 3](semesters/semester-3.md)
- [Semester 4](semesters/semester-4.md)
- [Quarter Index](quarters/README.md)
- [Month Index](months/README.md)
- [Week Index](weeks/README.md)
- [Day Index](days/README.md)
- [Atlas Engineering Track](atlas/atlas.md)
- [Atlas Backlog Index](atlas/backlog/README.md)
- [Atlas Issue Card Index](atlas/issues/README.md)

## Quarter Overview

- [Quarter 1](quarters/quarter-1.md): foundation and first agent
- [Quarter 2](quarters/quarter-2.md): interfaces, MCP, and agent runtime
- [Quarter 3](quarters/quarter-3.md): durable execution, evals, and governance
- [Quarter 4](quarters/quarter-4.md): production operations and platform integration

## Month Overview

- [Month 01](months/month-01.md): backend foundation and service skeleton
- [Month 02](months/month-02.md): background execution and failure handling
- [Month 03](months/month-03.md): single-agent service and tool use
- [Month 04](months/month-04.md): tool gateway and interface quality
- [Month 05](months/month-05.md): MCP protocol engineering
- [Month 06](months/month-06.md): multi-agent runtime and approvals
- [Month 07](months/month-07.md): durable workflow and memory boundaries
- [Month 08](months/month-08.md): evals, traces, and release gates
- [Month 09](months/month-09.md): governance, safety, and audit
- [Month 10](months/month-10.md): production deployment and SLO operations
- [Month 11](months/month-11.md): benchmarking, optimization, and capacity planning
- [Month 12](months/month-12.md): Atlas platform integration and graduation defense

## Week Overview

### Quarter 1

- [Week 01](weeks/week-01.md): Learning system setup and backend skeleton rfc
- [Week 02](weeks/week-02.md): Cli, http bootstrap, and service shell
- [Week 03](weeks/week-03.md): Config, logger, postgres, and redis integration
- [Week 04](weeks/week-04.md): Docker, ci, tests, and month 1 review
- [Week 05](weeks/week-05.md): Job runner rfc and queue model
- [Week 06](weeks/week-06.md): Worker loop, queue abstraction, and job execution
- [Week 07](weeks/week-07.md): Retry, timeout, idempotency, and dead-letter strategy
- [Week 08](weeks/week-08.md): Worker tests, drills, and month 2 review
- [Week 09](weeks/week-09.md): Single-agent rfc and tool inventory
- [Week 10](weeks/week-10.md): Responses api, structured output, and first tool
- [Week 11](weeks/week-11.md): Multi-tool routing and error handling
- [Week 12](weeks/week-12.md): Eval baseline, docs, and month 3 review
- [Week 13](weeks/week-13.md): Semester 1 retrospective and semester 2 prep

### Quarter 2

- [Week 14](weeks/week-14.md): Tool gateway rfc
- [Week 15](weeks/week-15.md): Schema validation and registry model
- [Week 16](weeks/week-16.md): Tool auth, timeout, and audit hooks
- [Week 17](weeks/week-17.md): Tool gateway tests and review
- [Week 18](weeks/week-18.md): Mcp server rfc
- [Week 19](weeks/week-19.md): Mcp bootstrap and transport model
- [Week 20](weeks/week-20.md): Mcp tools, resources, and compatibility
- [Week 21](weeks/week-21.md): Mcp contract tests and review
- [Week 22](weeks/week-22.md): Agents sdk rfc and role map
- [Week 23](weeks/week-23.md): Agent definitions and run context
- [Week 24](weeks/week-24.md): Handoffs, approvals, and guardrails
- [Week 25](weeks/week-25.md): Multi-agent tests and review
- [Week 26](weeks/week-26.md): Semester 2 retrospective and semester 3 prep

### Quarter 3

- [Week 27](weeks/week-27.md): Durable workflow rfc
- [Week 28](weeks/week-28.md): Graph state, checkpoint, and resume model
- [Week 29](weeks/week-29.md): Pause, replay, and human intervention
- [Week 30](weeks/week-30.md): Workflow tests and review
- [Week 31](weeks/week-31.md): Eval loop rfc
- [Week 32](weeks/week-32.md): Graders, traces, and dataset model
- [Week 33](weeks/week-33.md): Metrics, regression gates, and release policy
- [Week 34](weeks/week-34.md): Eval review and operational docs
- [Week 35](weeks/week-35.md): Governance rfc and threat model
- [Week 36](weeks/week-36.md): Policy engine and action classification
- [Week 37](weeks/week-37.md): Injection, leakage, and misuse testing
- [Week 38](weeks/week-38.md): Governance docs and review
- [Week 39](weeks/week-39.md): Semester 3 retrospective and semester 4 prep

### Quarter 4

- [Week 40](weeks/week-40.md): Kubernetes deployment rfc
- [Week 41](weeks/week-41.md): Manifests, health checks, and resource controls
- [Week 42](weeks/week-42.md): Alerts, rollback, and slo model
- [Week 43](weeks/week-43.md): Benchmark rfc
- [Week 44](weeks/week-44.md): Latency, throughput, and token cost baseline
- [Week 45](weeks/week-45.md): Optimization and capacity planning
- [Week 46](weeks/week-46.md): Graduation platform rfc
- [Week 47](weeks/week-47.md): Config center, tool gateway, and mcp integration
- [Week 48](weeks/week-48.md): Eval, approval, and audit integration
- [Week 49](weeks/week-49.md): End-to-end tests and failure drills
- [Week 50](weeks/week-50.md): Architecture and operator documentation
- [Week 51](weeks/week-51.md): Final benchmark and review
- [Week 52](weeks/week-52.md): Year-end retrospective and next roadmap

## Day Navigation

- Daily plans are stored in [Day Index](days/README.md).
- Semester 1: Day 001-Day 091
- Semester 2: Day 092-Day 182
- Semester 3: Day 183-Day 273
- Semester 4: Day 274-Day 364
- Final capstone day: [Day 365](days/day-365.md)

## Exit Standard

- Choose between Responses API and Agents SDK with clear reasoning.
- Design tool schemas, MCP contracts, state models, release gates, and approval boundaries.
- Build agent systems with traces, evals, approvals, audit, and production operations.
- Explain Atlas architecture with evidence from tests, drills, benchmarks, and documentation.
- Perform credibly in coding, design, reliability, and behavioral interview loops for AI agent architecture roles.
