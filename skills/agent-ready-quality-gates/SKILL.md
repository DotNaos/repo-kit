---
name: agent-ready-quality-gates
description: Set up hard repository quality gates for JS/TS projects so oversized files, oversized functions/components, and dependency hygiene regressions fail locally and in CI before they can slip into the codebase.
license: See repository license
---

## When to use

Use this skill when a repository should become agent-ready instead of agent-fragile:

- large files such as `App.tsx` should fail fast instead of depending on human vigilance
- oversized functions or components should be blocked before they become maintenance traps
- unused dependencies and dead exports should be detected automatically
- agents should see architecture hints while reading files, but enforcement must still come from tooling
- local hooks and CI should agree on the same non-negotiable structure rules

This skill is intentionally **JS/TS first** in v1.

## Goals

1. Make structural problems visible early while the agent is still reading and editing files.
2. Make structural violations fail deterministically through tooling.
3. Ensure `build` is not treated as proof of code health when architecture and lint checks are already broken.
4. Keep exceptions explicit and narrow for tests, fixtures, and generated files.

## Workflow

1. Inspect the target repository before making changes.
    - Check `package.json`, workspace layout, existing lint config, TypeScript config, CI workflows, and any current hook setup.
    - Identify whether the repo already uses ESLint, Biome, oxlint, or another primary linter.
    - Do not introduce a second overlapping lint system unless there is a strong reason and the repository clearly lacks a usable one.

2. Install or extend the quality scripts.
    - Ensure the repository has a deterministic `lint` script.
    - Add a dependency hygiene script such as `check:deps` or `knip`.
    - Keep `build` separate from structure checks; do not hide lint or dependency failures inside a vague all-in-one script.

3. Enforce hard limits for source structure.
    - Add file-length rules for production source files.
    - Add function or component length rules for production source files.
    - Prefer explicit per-folder or per-file overrides over one permissive global limit.
    - Document exemptions for tests, fixtures, stories, generated files, and machine-written snapshots.

4. Add readable guidance for agents.
    - Use short file headers only in high-risk hotspots such as `App.tsx`, shell components, large screens, or central orchestrators.
    - Keep headers short and architectural, not chatty TODO dumps.
    - Make sure the header complements the rules instead of replacing them.

5. Wire local and remote enforcement.
    - Keep `pre-commit` fast and deterministic.
    - Put the full hard gate into CI or PR workflows as well.
    - Ensure the branch cannot appear healthy when lint or dependency hygiene is failing.

## Recommended baseline for JS/TS

Use the smallest setup that reliably blocks regressions:

- ESLint for structural rules
- Knip for unused dependencies and exports
- Husky for local hooks when the repo accepts local Git hooks
- CI job that runs `lint`, `knip`, and `build` separately

See `references/js-ts-thresholds.md`, `references/file-header-template.md`, and `references/eslint-structure-rules.md`.

## Guardrails

- Do not add multiple competing linters just because they are fashionable.
- Do not make `pre-commit` so heavy that developers bypass it.
- Do not allow `build` to be the only required status check.
- Do not silently disable rules for `src/**` just to get green builds.
- Do not turn file headers into freeform planning notes or long TODO journals.
- Do not ignore existing repository conventions if the current setup is already sound.

## Definition of done

- the repo has explicit scripts for structure and dependency hygiene
- oversized files and oversized functions/components fail deterministically
- dead dependencies or exports fail deterministically
- CI runs quality gates independently from `build`
- high-risk files can include short architectural headers, but the true enforcement lives in tooling
