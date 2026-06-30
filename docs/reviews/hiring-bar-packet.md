# Hiring Bar Packet

Last updated: 2026-06-29

## Purpose

This packet defines how to judge whether the learner has reached a credible hiring bar for AI agent architect roles.

The goal is not to simulate a toy interview. The goal is to connect interview answers to real Atlas evidence.

## Required Evidence

- Atlas architecture docs
- RFC archive
- tests, traces, drills, or benchmark outputs
- monthly and quarterly review records
- at least one end-to-end platform walkthrough

## Interview Loop

### Loop 1: Coding and Code Health

- review one real Atlas module
- review one change proposal
- ask for test strategy and refactoring judgment

### Loop 2: System Design

- design a new agent subsystem or redesign an existing one
- discuss state, interfaces, tools, memory, safety, and failure paths

### Loop 3: Reliability and Operations

- evaluate SLOs, rollout, rollback, observability, and incident handling
- require benchmark or drill evidence where appropriate

### Loop 4: Behavioral and Growth

- ask about retrospectives, trade-offs, missed assumptions, and coaching behavior
- require concrete examples from weekly, monthly, or quarterly work

## Bar-Raising Questions

1. What evidence shows Atlas is more than a demo?
2. Which decision best shows architect-level judgment under uncertainty?
3. What part of the system is still under-designed, and what would you change next?
4. Where did you improve code health instead of only shipping faster?
5. How do you know your reliability story is real rather than theoretical?
6. What would break first if a second team started building on Atlas tomorrow?

## Decision Rules

- no pass without Atlas evidence
- no pass with missing documentation or missing review artifacts
- no pass if system design answers are disconnected from operational reality
- strong pass requires consistent evidence across coding, design, reliability, and communication

## Output Format

- overall recommendation
- strongest two signals
- weakest two signals
- evidence cited
- decision rationale
- next-stage development plan
