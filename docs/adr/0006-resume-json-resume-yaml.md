# 6. Resume data in JSON Resume structure, parsed by Ridgelight

- **Status:** Accepted
- **Date:** 2026-10-24

## Context

The resume page should come from one data file that is easy to edit. [JSON Resume](https://jsonresume.org) is an existing open schema for resume data.

resumekiln, a separate planned CLI for turning `resume.yaml` into JSON and PDF, has not been built yet.

## Decision

- Resume data lives in `content/resume.yaml`, following the JSON Resume structure.
- Ridgelight defines its own Go structs for that data in `internal/resume/`. It does not import code from resumekiln.
- The downloadable PDF is a static file at `static/resume.pdf`. The site only serves it; producing it is outside the site's code.

## Consequences

- Ridgelight can be built and run without resumekiln existing.
- If resumekiln later provides its own structs, the two copies can drift. Sharing a package can be reconsidered then, in a new ADR.
- The resume page and `static/resume.pdf` must be updated together so they don't disagree.
