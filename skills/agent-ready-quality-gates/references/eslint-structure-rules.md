# ESLint structure rule sketch for JS/TS repos

Use this as a baseline sketch when adding hard structural guards.

## Example flat-config direction

```ts
import js from '@eslint/js';
import tseslint from 'typescript-eslint';

export default tseslint.config(
    js.configs.recommended,
    ...tseslint.configs.recommended,
    {
        files: ['src/**/*.{ts,tsx,js,jsx}'],
        rules: {
            'max-lines': [
                'error',
                { max: 450, skipBlankLines: true, skipComments: true },
            ],
            'max-lines-per-function': [
                'error',
                {
                    max: 90,
                    skipBlankLines: true,
                    skipComments: true,
                    IIFEs: true,
                },
            ],
        },
    },
    {
        files: ['**/*.test.{ts,tsx}', '**/*.spec.{ts,tsx}', '**/fixtures/**/*'],
        rules: {
            'max-lines': [
                'warn',
                { max: 800, skipBlankLines: true, skipComments: true },
            ],
            'max-lines-per-function': 'off',
        },
    },
    {
        ignores: ['**/generated/**', '**/*.gen.ts', '**/*.generated.ts'],
    },
);
```

## Practical guidance

- prefer `error`, not `warn`, for production source files
- keep overrides narrow and well justified
- do not exempt the whole `src/` tree to quiet the rule
- if a file repeatedly pushes the limit, split config/maps/helpers before touching UI first
