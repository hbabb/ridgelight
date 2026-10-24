# 4. Plain CSS, no build step

- **Status:** Accepted
- **Date:** 2026-10-24

## Context

Styling tools such as Tailwind, Sass, or PostCSS need their own build step, and usually a Node.js toolchain alongside Go.

## Decision

Write styles in plain CSS.

- Stylesheets live in `static/css/`, starting with `main.css`.
- They are embedded in the binary with the rest of `static/` and served at `/static/css/`.
- No CSS framework, preprocessor, or bundler.

## Consequences

- The only build tools are Go and templ.
- Anything a framework would provide (resets, layout helpers, utility classes) is written by hand.
