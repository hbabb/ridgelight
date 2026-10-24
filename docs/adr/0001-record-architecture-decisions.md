# 1. Record architecture decisions

- **Status:** Accepted
- **Date:** 2026-10-24

## Context

Ridgelight's design choices have been made across several planning sessions. Without a written record, the reasons behind them get lost, and the same questions get reopened later.

## Decision

Record each significant decision as an architecture decision record (ADR) in `docs/adr/`.

- One decision per file.
- Files are numbered in order: `0001-...md`, `0002-...md`, and so on.
- Each ADR has a **Status**, **Date**, **Context**, **Decision**, and **Consequences** section.
- When a decision changes, the old ADR is marked **Superseded** and links to the new one. It is not deleted.

Questions that are still open are listed in the README under **Open decisions** until they are settled and get their own ADR.

## Consequences

- The repo holds both what was decided and why.
- Writing an ADR is an extra step for each significant change.
