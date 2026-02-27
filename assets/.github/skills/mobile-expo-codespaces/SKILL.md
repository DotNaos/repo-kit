---
name: mobile-expo-codespaces
description: Setup and run Expo apps in GitHub Codespaces using tunnel for iPhone testing, and trigger production EAS builds later from CI.
license: See repository license
---

## When to use

Use this skill when a repository contains an Expo app and you need a repeatable Codespaces workflow for:

- Running the app from GitHub Codespaces
- Testing on a physical iPhone with Expo Go via QR code
- Preparing production builds later in CI with EAS

## Workflow

1. Run environment checks first:

   ```bash
   ./scripts/doctor.sh
   ```

2. Start Expo in tunnel mode from Codespaces:

   ```bash
   ./scripts/start-tunnel.sh
   ```

   This runs `expo start --tunnel` (via `bunx` when Bun is used in the repo, otherwise `npx`).

3. Test on iPhone:

   - Open **Expo Go** on your iPhone.
   - Scan the QR code shown in the terminal.

4. If your app uses `expo-dev-client`, start with dev client tunnel mode:

   ```bash
   npx expo start --dev-client --tunnel
   ```

   If the repo uses Bun, the equivalent is:

   ```bash
   bunx expo start --dev-client --tunnel
   ```

5. Notes:

   - Tunnel mode is usually slower than LAN/Local, but works reliably for Codespaces + phone testing.
   - Port forwarding is optional, but tunnel is typically the simplest path for Expo Go on iPhone.

## Production builds in CI (later)

Use the reference workflow in `references/eas-build-github-actions.md` to trigger EAS builds with `EXPO_TOKEN` from GitHub Actions.
