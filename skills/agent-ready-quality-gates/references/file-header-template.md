# File header template for hotspots

Use a short header only in files that are historically risky or central to orchestration.

## Example

```ts
/**
 * Purpose: Application shell that composes routing and top-level providers.
 * Boundaries: Keep feature logic, data fetching flows, and screen-specific UI out of this file.
 * Refactor trigger: If this file grows beyond the agreed threshold, split into route modules, hooks, or provider composition helpers.
 */
```

## Rules

- keep it under 4 lines of actual guidance
- describe purpose and boundaries, not ticket history
- avoid vague TODO language such as "clean up later"
- prefer architectural cues over implementation trivia
