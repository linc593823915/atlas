# Atlas Agent Guide

## Coding Conventions

Atlas permanently follows the Rule of 30.

- Any source file over 300 lines must be split.
- Any function over 30 lines must be split.
- Any struct with over 20 fields must be redesigned.

Run `go run ./scripts/ruleof30` before finishing code changes. `make test`
already includes this check.
