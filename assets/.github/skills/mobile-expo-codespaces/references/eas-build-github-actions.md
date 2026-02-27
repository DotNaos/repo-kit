## GitHub Actions example: EAS Build

Use this workflow to trigger a production build on EAS from CI.

```yaml
name: EAS Build

on:
  workflow_dispatch:
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v5

      - uses: actions/setup-node@v6
        with:
          node-version: 22

      - uses: oven-sh/setup-bun@v2

      - name: Setup Expo and EAS
        uses: expo/expo-github-action@v8
        with:
          eas-version: latest
          token: ${{ secrets.EXPO_TOKEN }}
          packager: bun

      - name: Install dependencies
        run: bun install --frozen-lockfile

      - name: Trigger EAS build
        run: eas build --platform all --non-interactive --no-wait
```

If your project does not use Bun, use npm/yarn and omit `packager: bun`.

## Official docs

- https://docs.github.com/en/copilot/concepts/agents/about-agent-skills
- https://docs.expo.dev/get-started/start-developing/
- https://docs.expo.dev/build/building-on-ci/
- https://github.com/expo/expo-github-action
