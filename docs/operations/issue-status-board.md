# Issue Status Board

Last updated: 2026-06-29

## Purpose

This board defines the canonical workflow state model for Atlas learning issues and engineering cards.

## Status Model

- `planned`
  The issue exists and is prioritized, but implementation has not started.
- `ready`
  Dependencies are clear and the issue can be started immediately.
- `in_progress`
  Active implementation or document work is underway.
- `in_review`
  The main artifact exists and is being reviewed for design, code health, tests, or docs.
- `blocked`
  Progress is halted by a real dependency, missing decision, or external state.
- `validated`
  Tests, drills, traces, or benchmark evidence confirm the issue outcome.
- `done`
  The issue is closed with review, docs, and evidence complete.

## Board Rules

- No issue moves to `in_progress` without a clear owner and definition of done.
- No issue moves to `in_review` without an implementation or document artifact.
- No issue moves to `done` without evidence.
- `blocked` issues must state blocker, owner of unblock, and expected next review date.

## Monthly Board Check

For each month, track:

- planned issues count
- in-progress issues count
- blocked issues count
- validated issues count
- done issues count

## Minimal Board Row Template

- issue id:
- title:
- owner:
- status:
- target week:
- dependency summary:
- evidence summary:
- next action:
