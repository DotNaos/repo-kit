---
name: pr-quality-pipeline
description: Set up a pull request quality pipeline that reruns hard local quality gates in CI and adds an independent review lane so code is checked before merge, not only before build.
license: See repository license
---

## When to use

Use this skill when a repository should enforce quality after every commit in a branch or pull request:

- local hooks exist or are planned, but server-side enforcement is still missing
- teams need a PR pipeline that reruns `lint`, dependency hygiene, and `build`
- clean-code review should be independent from the original implementation run
- the repository should not depend on a human noticing structural decay in code review

## Goals

1. Re-run local quality expectations in CI.
2. Keep quality checks separate so failures are obvious.
3. Add an independent review lane for clean-code findings.
4. Make the resulting checks suitable for branch protection.

## Workflow

1. Inspect the repository's current automation.
    - Check existing workflows under `.github/workflows/`.
    - Check whether the project already has `lint`, `test`, `build`, and dependency hygiene scripts.
    - Check whether an agent reviewer file exists, especially `.github/agents/reviewer.md`.

2. Establish the baseline PR jobs.
    - Add or update jobs for `lint`, dependency hygiene (`knip` or equivalent), and `build`.
    - Keep them visible as separate steps or separate jobs; do not collapse all failures into one opaque command.

3. Add an independent review lane.
    - The review lane must not simply reuse the same implementation context and call that "review".
    - Prefer a separate reviewer configuration or a dedicated agent invocation path.
    - Make the review output actionable: architecture findings, layering concerns, responsibility overlap, abstraction leaks.

4. Keep local and remote policies aligned.
    - If Husky is used locally, the PR pipeline should still re-run the authoritative checks.
    - Document which checks are fast local gates and which checks are CI-only or manual local review flows.

5. Prepare for branch protection.
    - Name jobs clearly so they can be required.
    - Avoid optional wording for required structural checks.

## Local gate strategy

A good baseline split is:

- `prepare`: validate that the quality toolchain and checklist are wired correctly
- `pre-commit`: fast deterministic checks only
- manual local review command: trigger the independent clean-code review on demand
- PR pipeline: rerun hard checks plus the independent review lane

See `references/quality-checklist.md` for a concrete baseline checklist that `prepare` can validate, `references/local-gate-layout.md` for the local Husky split, and `templates/pr-quality.yml.j2` for a starter PR workflow.

## Guardrails

- Do not rely on `build` alone as a proxy for code quality.
- Do not make `pre-commit` the only enforcement point.
- Do not call a self-review by the same execution flow an independent review.
- Do not hide failing jobs behind optional informational workflows.
- Do not add slow AI-heavy checks to every local commit by default.

## Definition of done

- PRs or main-branch commits run explicit quality jobs
- `lint`, dependency hygiene, and `build` are all enforced independently
- an independent review lane exists for clean-code findings
- job names are suitable for branch protection and repository policy
- local hooks and CI have clear, documented roles instead of overlapping chaos
