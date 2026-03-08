# JS/TS thresholds and exceptions

Use these as starting points, not dogma. The point is to force decomposition before files turn into emergency archaeology.

## Suggested baseline

### Production source files

- `max-lines`: 400 to 600
- `max-lines-per-function`: 75 to 120
- prefer the stricter side for React screens and container components

### Hotspot files

For files such as `App.tsx`, route-level screens, or orchestration shells:

- aim for `max-lines` around 300 to 400
  Break logic into:

- feature components
- hooks
- services or view-model helpers
- config maps and constants

### Tests and fixtures

Allow broader limits where aggregation is normal, but keep them finite:

- tests: relaxed file-length limits
- fixtures and snapshots: ignore or heavily relax
- generated files: ignore fully when regenerated from source

## What to split first

When a file grows too large, extract in this order:

1. static config, mappings, and constants
2. data transformation helpers
3. side-effectful services or adapters
4. feature-specific hooks
5. presentational subcomponents

## Knip baseline

A practical first pass for JS/TS repos is:

- fail on unused dependencies
- fail on unused exports unless there is a documented public API reason
- add narrow ignore patterns only after verifying the false positive

## Good exceptions

Allowed exceptions should be explicit and reviewed:

- generated SDK files
- migration snapshots
- storybook test harness files
- intentionally aggregated test helpers

## Bad exceptions

These are usually smells, not reasons:

- "the file is easier to scroll this way"
- "the component knows the whole page already"
- "the build still passes"
- "the agent wrote it like that"
