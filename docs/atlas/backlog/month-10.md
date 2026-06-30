# Month 10 Backlog

Last updated: 2026-06-29

Theme: production deployment and SLO operations

## Target Milestone

Atlas runs in a production-like Kubernetes model with health checks, resource controls, alerts, and rollback readiness.

## Priority Issues

- `ATLAS-M10-01`
  Write deployment RFC and service topology.
- `ATLAS-M10-02`
  Add manifests or Helm chart for Atlas services.
- `ATLAS-M10-03`
  Add startup, readiness, and liveness probes.
- `ATLAS-M10-04`
  Define SLI, SLO, and alert rules.
- `ATLAS-M10-05`
  Write rollback and incident runbook.

## Exit Criteria

- Atlas can deploy in Kubernetes
- health checks exist
- alerting and rollback paths are defined
- operators have a usable runbook

## Issue Set

- [Month 10 Issue Set](../issues/month-10/README.md)

## Review Checkpoint

- show deployment layout
- explain one alert and one rollback trigger
- defend chosen SLO boundaries
