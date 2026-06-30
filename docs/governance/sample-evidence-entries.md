# Sample Evidence Entries

Last updated: 2026-06-29

## Purpose

These examples show how evidence registry entries might look in practice.

## Example 1: Design Evidence

- artifact id: EVD-M01-RFC-BOOTSTRAP
- linked issue or period: ATLAS-M01-01 / Month 01
- artifact type: RFC
- why it matters: defines lifecycle boundaries for Atlas bootstrap
- reviewer comment: good foundation, but health-check ownership should be explicit

## Example 2: Code Evidence

- artifact id: EVD-M01-LOGGER
- linked issue or period: ATLAS-M01-03 / Week 03
- artifact type: code change
- why it matters: standardizes logger entrypoint and source behavior
- reviewer comment: improves consistency and future maintainability

## Example 3: Reliability Evidence

- artifact id: EVD-M02-WORKER-TESTS
- linked issue or period: ATLAS-M02-05 / Month 02
- artifact type: validation suite
- why it matters: proves retry and dead-letter behavior
- reviewer comment: add one adversarial timeout case next month

## Example 4: Documentation Evidence

- artifact id: EVD-M03-RUNBOOK
- linked issue or period: Month 03
- artifact type: operator guide
- why it matters: lets another engineer run the first single-agent flow
- reviewer comment: sufficient for local operation, not yet for production handoff
