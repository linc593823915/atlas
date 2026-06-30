# Month 01 Backlog

Last updated: 2026-06-29

Theme: backend foundation and service skeleton

## Target Milestone

Atlas has a stable service skeleton with bootstrap, config, logger, Postgres and Redis wiring, Docker, and CI.

## Priority Issues

- `ATLAS-M01-01`
  Build bootstrap shell and runtime lifecycle boundaries.
- `ATLAS-M01-02`
  Harden `config.Manager` and sample config loading.
- `ATLAS-M01-03`
  Standardize logger initialization and source behavior.
- `ATLAS-M01-04`
  Wire Postgres and Redis clients with health checks.
- `ATLAS-M01-05`
  Add Docker, CI, and smoke-test entrypoints.

## Exit Criteria

- service boots cleanly
- config and logger are the only allowed entrypoints
- database and cache clients can be validated
- Docker and CI can execute baseline checks

## Issue Set

- [Month 01 Issue Set](../issues/month-01/README.md)

## Review Checkpoint

- demo startup and shutdown flow
- explain package layout and ownership boundaries
- show tests and run instructions
