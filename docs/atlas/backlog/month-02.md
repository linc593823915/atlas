# Month 02 Backlog

Last updated: 2026-06-29

Theme: background execution and failure handling

## Target Milestone

Atlas supports queue-backed background execution with retries, timeouts, idempotency, and drill-ready failure handling.

## Priority Issues

- `ATLAS-M02-01`
  Define queue abstraction and job state model.
- `ATLAS-M02-02`
  Implement worker loop and graceful shutdown behavior.
- `ATLAS-M02-03`
  Add retry policy, timeout handling, and dead-letter flow.
- `ATLAS-M02-04`
  Add idempotency keys and duplicate suppression strategy.
- `ATLAS-M02-05`
  Write failure drill notes and validation tests.

## Exit Criteria

- workers can start and stop safely
- transient failures retry correctly
- hard failures are visible and inspectable
- tests prove failure-path behavior

## Issue Set

- [Month 02 Issue Set](../issues/month-02/README.md)

## Review Checkpoint

- walk through worker lifecycle
- explain one drill scenario
- justify retry and timeout defaults
