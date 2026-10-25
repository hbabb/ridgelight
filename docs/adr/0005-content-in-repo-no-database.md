# 5. Content in the repo, embedded, no database

- **Status:** Accepted
- **Date:** 2026-10-24

## Context

The site's content is blog posts, project write-ups, and a resume. It changes occasionally and is written by one person.

## Decision

Keep all content as files in this repository. No database.

- Blog posts: `content/posts/<slug>.md`
- Project write-ups: `content/projects/<slug>.md`
- Resume data: `content/resume.yaml`
- Posts and projects are Markdown with a YAML front matter block (title, date, tags, and so on). The front matter fields are documented in the README.
- Markdown is rendered to HTML with [goldmark](https://github.com/yuin/goldmark).
- Projects are grouped on the projects page by their `category` field: `gis`, `cli`, or `web`. Moving a project to another group only means changing that field.
- Content and static files are compiled into the binary with Go's `embed` package (`embed.go`). Today only `static/` is embedded; `content/` will be added when it exists.

How front matter is parsed is still an open decision (see the README).

## Consequences

- Content goes through the same pull request, review, and history as code.
- There is no database to back up, migrate, or secure.
- Publishing a post or fixing a typo means a rebuild and deploy.
