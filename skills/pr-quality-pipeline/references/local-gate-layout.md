# Local gate layout with Husky

Use this split to keep local enforcement fast and predictable.

## Suggested package scripts

```json
{
    "scripts": {
        "prepare": "husky",
        "lint": "eslint . --max-warnings 0",
        "check:deps": "knip",
        "build": "vite build",
        "quality:quick": "npm run lint",
        "quality:full": "npm run lint && npm run check:deps && npm run build"
    }
}
```

## Suggested hook split

### `.husky/pre-commit`

Run only quick deterministic checks, for example:

```sh
npm run quality:quick
```

### Optional `.husky/pre-push`

If the team wants a heavier local gate:

```sh
npm run quality:full
```

## Manual local independent review

Keep the independent clean-code review as a manual command or IDE task when possible.

Examples:

- IDE command palette entry
- repo-specific review script
- explicit agent run triggered by the developer

This preserves fast commits while still allowing strict local review before pushing.
