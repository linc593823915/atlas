# Mentor Review Template

Last updated: 2026-06-29

## Purpose

Use this template for monthly mentor reviews and quarterly development conversations.

This is a Google-style public-practices adaptation:

- objective and key result review
- evidence-based performance discussion
- code health and reliability calibration
- growth and coaching loop

## Review Inputs

- monthly OKR document
- weekly deliverables
- Atlas code and docs changed during the review period
- tests, traces, drills, or benchmark evidence
- self-review note from the learner

## Scoring Scale

- `Not Yet`
  The learner is missing core artifacts, weak on explanation, or not yet independently dependable.
- `Meets`
  The learner reliably ships the expected scope, explains trade-offs, and maintains acceptable quality.
- `Strong`
  The learner improves Atlas beyond the assigned scope and shows architect-level judgment in design and operations.

## Review Dimensions

### 1. Deliver Results

- Did the learner finish the planned monthly scope?
- Were key results explicit and graded honestly?
- Did the learner close the loop with documentation and memory?

Evidence:

- shipped code or docs
- completed RFCs
- monthly OKR grading

### 2. Technical Judgment

- Were interfaces and abstractions chosen for good reasons?
- Did the learner explain trade-offs using evidence instead of preference?
- Did the learner simplify the design where possible?

Evidence:

- RFC decisions
- architecture notes
- review comments

### 3. Code Health and Simplicity

- Are changes reviewable and maintainable?
- Did the learner reduce accidental complexity?
- Did the learner protect future changeability?

Evidence:

- change structure
- naming and flow clarity
- follow-up debt notes

### 4. Reliability and Testing

- Were tests or validation added at the right level?
- Were failure paths considered early instead of late?
- Does the learner think in terms of rollback, observability, and recovery?

Evidence:

- test artifacts
- drill notes
- metrics, traces, or benchmark outputs

### 5. Communication and Documentation

- Can another engineer understand the work quickly?
- Are docs, runbooks, and review notes current?
- Can the learner explain the system clearly in conversation?

Evidence:

- docs
- runbooks
- self-review
- mock interview notes

### 6. Learning and Community

- Did the learner request feedback effectively?
- Did the learner coach, review, or document in ways that help others?
- Did the learner convert mistakes into reusable lessons?

Evidence:

- peer-review notes
- retrospective quality
- memory entries

## Monthly Review Questions

1. What was the highest-value result this month?
2. Which architectural choice was strongest, and which was weakest?
3. What code-health debt was created or removed?
4. What failure mode is still under-protected?
5. What evidence best proves production thinking?
6. What should change next month in scope, habits, or standards?

## Review Summary Template

- Month:
- Reviewer:
- Results rating:
- Technical judgment rating:
- Code health rating:
- Reliability rating:
- Communication rating:
- Learning rating:
- Overall:
- Key strengths:
- Key concerns:
- Next-month focus:
