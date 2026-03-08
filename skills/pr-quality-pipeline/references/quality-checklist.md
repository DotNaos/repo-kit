# Quality checklist for prepare and PR setup

Use this checklist as the baseline that a `prepare`-style setup validation can enforce.

## Required repository capabilities

### Scripts

The repository should expose explicit scripts instead of one opaque catch-all command:

- `lint`
- `build`
- `check:deps` or `knip`
- optional but recommended: `test`
- optional local review launcher if the repo integrates an agent review command

### Local hooks

If Husky is used, the expected split is:

- `prepare` installs or validates Husky wiring
- `pre-commit` runs only fast deterministic checks
- optional `pre-push` runs heavier local checks

Do not require a full independent AI review inside every `pre-commit` by default.

### Reviewer guidance

The repository should contain reviewer guidance when independent agent review is part of the policy:

- `.github/agents/reviewer.md`

### PR automation

The repository should contain at least one workflow that reruns quality gates in CI:

- `lint`
- dependency hygiene check
- `build`
- independent review lane when configured

## Example prepare validations

A setup validator can fail when any of these are missing:

- no `lint` script
- no dependency hygiene command
- no CI workflow for PRs or main-branch pushes
- no reviewer guidance file when independent review is required
- Husky declared but no `pre-commit` hook present

## Recommended policy split

| Stage               | Purpose                                                  |
| ------------------- | -------------------------------------------------------- |
| `prepare`           | verify quality wiring and required files/scripts         |
| `pre-commit`        | run fast deterministic local checks                      |
| manual local review | run the independent clean-code review on demand          |
| PR pipeline         | rerun authoritative quality gates and independent review |

## Anti-patterns

- `prepare` silently succeeding when required quality scripts are absent
- `build` existing without any structural or dependency hygiene gate
- `pre-commit` trying to do every expensive review task
- PR workflow that only echoes placeholder text and never enforces quality
