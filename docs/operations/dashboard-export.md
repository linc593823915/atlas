# Dashboard Export

Last updated: 2026-06-29

## Purpose

This document defines the minimum export shape for a dashboard or reporting layer built on top of the learning system.

## Export Views

### 1. Monthly Program View

Fields:

- month
- monthly objective
- planned issues
- done issues
- blocked issues
- validation coverage
- review status
- OKR grade
- promotion signal

### 2. Quarterly Program View

Fields:

- quarter
- quarter theme
- monthly pass rate
- Atlas milestone status
- mock panel result
- strongest evidence
- biggest risk

### 3. Issue Board View

Fields:

- issue id
- month
- owner
- status
- target week
- risk level
- dependency summary
- evidence summary
- next action

### 4. Evidence View

Fields:

- artifact id
- linked issue
- linked month or quarter
- artifact type
- reviewer comment

## Minimal JSON Shape

```json
{
  "month": "01",
  "objective": "Build a clean backend base for Atlas",
  "issues": {
    "planned": 5,
    "done": 4,
    "blocked": 1
  },
  "validation_coverage": 0.8,
  "okr_grade": "Meets",
  "promotion_signal": "L4-like emerging"
}
```

## Export Rule

- Do not export vanity metrics without linked evidence.
- Every dashboard number should trace back to a review packet, issue card, or evidence entry.
