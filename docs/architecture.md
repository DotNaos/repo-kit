# Architecture

`repo-kit` v1 is a small TypeScript CLI with one agent backend and filesystem-based skill discovery.

## Source Layout

- `src/cli/`
  - command parsing
  - terminal output
- `src/core/`
  - shared types
  - skill discovery and loading
  - minimal repository context collection
  - authentication checks
- `src/adapters/codex-sdk/`
  - the only agent backend in v1
  - maps `plan` and `run` into Codex SDK thread execution
- `skills/`
  - bundled skills and supporting assets
- `docs/`
  - usage and publishing notes

## Execution Flow

1. The CLI resolves the packaged `skills/` directory.
2. `skills list` inspects each first-level skill directory and reports whether it is runnable.
3. `plan` and `run` load a single skill, collect minimal context from the current working directory, and require a Git repository.
4. The Codex adapter starts a thread with the current repository as the working directory and the selected skill directory exposed via `additionalDirectories`.
5. `plan` uses read-only sandboxing and a structured JSON plan schema.
6. `run` uses workspace-write sandboxing and returns the agent response plus basic execution summaries.

## Skill Compatibility

The loader prefers normalized skills:

- `skill.yaml`
- `prompt.md`
- optional `assets/`, `templates/`, `scripts/`

To avoid rewriting the current catalog, v1 also accepts:

- `SKILL.md`
- a single markdown file fallback for older prompt-only entries

Directories without a prompt definition are intentionally left invalid instead of guessing at runnable behavior from raw assets.
