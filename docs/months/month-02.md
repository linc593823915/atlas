# Month 02

Last updated: 2026-06-29

Semester: [Semester 1](../semesters/semester-1.md)
Quarter: [Quarter 1](../quarters/quarter-1.md)
Week range: Week 05-Week 08
Day range: Day 029-Day 056

## Theme

Background execution and failure handling.

## Objective

Build dependable worker execution with retries, timeouts, idempotency, and operational drills.

## Monthly OKR

- Objective: Build dependable worker execution with retries, timeouts, idempotency, and operational drills.
- Key Result 1: Ship the planned Atlas milestone for the month.
- Key Result 2: Explain the main design trade-offs in writing and in a mock interview.
- Key Result 3: Produce evidence through tests, validation, or measurements.
- Key Result 4: Finish documentation, review, and memory artifacts for every weekly issue.

## Weekly Coverage

- [Week 05](../weeks/week-05.md): Job runner rfc and queue model
- [Week 06](../weeks/week-06.md): Worker loop, queue abstraction, and job execution
- [Week 07](../weeks/week-07.md): Retry, timeout, idempotency, and dead-letter strategy
- [Week 08](../weeks/week-08.md): Worker tests, drills, and month 2 review

## Atlas Milestone

- Atlas should visibly improve in the area of background execution and failure handling.

## Monthly Assessment Standard

- Score out of 100.
- Results vs monthly OKR: 20 points.
- Technical depth and judgment: 20 points.
- Implementation quality and code health: 15 points.
- Testing, reliability, and failure thinking: 15 points.
- System design clarity: 10 points.
- Documentation and communication: 10 points.
- Learning, peer feedback, and coaching loop: 10 points.
- Passing bar: 80+ and no missing core artifact.
- Strong bar: 90+ with evidence that Atlas improved beyond the assigned weekly tasks.

## Evidence Checklist

- Monthly objective and key results are written and graded.
- All weekly deliverables exist and are easy to locate.
- Tests, traces, or other validation evidence exist for the month’s core change.
- Documentation and runbooks reflect the current Atlas state.
- At least one peer or self-review summary exists with concrete improvement actions.

## Interview Questions

1. How would you explain the architecture choices behind background execution and failure handling to a senior reviewer?
2. What is the most failure-prone part of build dependable worker execution with retries, timeouts, idempotency, and operational drills. and how would you test it?
3. Which interface or contract introduced this month is most likely to change later, and how did you design for that?
4. If Atlas had to support 10x more usage after this month, where would the design break first?
5. What code-health debt was created or removed this month, and how do you know?
6. Which weekly artifact from this month best demonstrates technical judgment, and why?
7. How would you turn the themes Job runner rfc and queue model, Worker tests, drills, and month 2 review into one coherent system story?
8. What metrics, traces, or test evidence would you show to prove the month’s work is production-worthy?
