---
name: clean-code-review-refactor
description: Review and refactor existing code against explicit clean code principles using an independent reviewer mindset, with concrete before/after examples and structural improvements instead of cosmetic commentary.
license: See repository license
---

## When to use

Use this skill when a repository has code that technically works but is becoming structurally expensive:

- responsibilities are mixed inside large files or large functions
- UI or API layers expose implementation details that should stay hidden
- components contain too much nesting and too many layers at once
- configuration, constants, or mappings live inside logic-heavy files without a good reason
- code changes need an independent clean-code review instead of self-approval by the same execution flow

## Review stance

This skill should behave like an independent reviewer and refactoring partner.

- assume the current code may work while still being expensive to maintain
- prioritize structural clarity over cosmetic formatting changes
- focus on separation of concerns, coupling, boundaries, naming, and decomposition
- provide concrete refactorings, not abstract lectures

## Core principles

1. **Tight coupling inside, loose coupling outside**
    - Keep internal implementation details cohesive.
    - Expose stable interfaces or ports to the outside.
    - Avoid leaking transport details, persistence details, or widget internals into public APIs or top-level UI composition.

2. **Hide abstraction in the implementation, not in the user-facing surface**
    - User-facing APIs and UI should communicate domain intent, not internal plumbing.
    - Move orchestration complexity behind helpers, adapters, hooks, or service modules.

3. **Single Responsibility Principle**
    - Split files, modules, and components that handle orchestration, data transformation, configuration, and presentation simultaneously.

4. **Configuration belongs outside logic-heavy files when practical**
    - Extract static mappings, constants, route metadata, and feature flags from components or use-case implementations when doing so improves clarity.

5. **Work in layers, especially in React**
    - Avoid giant trees that mix layout, primitives, widgets, business logic, data loading, and application orchestration in a single component.
    - Separate route/screen composition, stateful hooks, presentational components, and low-level primitives.

## Workflow

1. Inspect the changed code or target area.
    - Identify where responsibilities are mixed.
    - Look for leaking abstractions, oversized modules, oversized components, and excessive nesting.
    - Distinguish between a real structural issue and a reasonable local convenience.

2. Prioritize refactors by structural gain.
    - First extract hard architectural seams: interfaces, adapters, config modules, hooks, services, presentational pieces.
    - Then simplify naming and local control flow.
    - Leave purely stylistic changes for last.

3. Show the reasoning in concrete form.
    - For each major recommendation, explain what responsibility was split and why.
    - Use before/after examples from `references/before-after-examples.md` as the tone and level of specificity to emulate.

4. Keep the external surface clean.
    - Domain-focused props and APIs should stay readable.
    - Technical plumbing should move inward.

## Guardrails

- Do not produce a fake review that only comments on naming and formatting.
- Do not move complexity around without reducing coupling or responsibility overlap.
- Do not introduce abstraction layers that make the call site harder to understand.
- Do not split files mechanically when the result worsens navigability.
- Do not hide domain concepts behind generic utility wrappers.

## Definition of done

- major structural problems are identified with concrete reasoning
- proposed or applied refactors reduce responsibility overlap
- public UI or API surfaces become simpler, not more technical
- React or UI composition becomes more layered and less deeply entangled
- examples or explanations are concrete enough that another engineer or agent can continue the pattern
