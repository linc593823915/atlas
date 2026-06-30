# Month 07 Backlog

Last updated: 2026-06-29

Theme: durable workflow and memory boundaries

## Target Milestone

Atlas workflows can pause, resume, checkpoint, and support human intervention with clear memory boundaries.

## Priority Issues

- `ATLAS-M07-01`
  Define workflow state and checkpoint contract.
- `ATLAS-M07-02`
  Implement checkpoint persistence.
- `ATLAS-M07-03`
  Add pause and resume behavior.
- `ATLAS-M07-04`
  Add replay and manual intervention paths.
- `ATLAS-M07-05`
  Write workflow tests and operating notes.

## Exit Criteria

- state is persisted explicitly
- workflows can resume after interruption
- intervention paths are clear to operators
- tests cover replay or resume behavior

## Issue Set

- [Month 07 Issue Set](../issues/month-07/README.md)

## Review Checkpoint

- demonstrate pause/resume
- explain state ownership
- show one recovery path
