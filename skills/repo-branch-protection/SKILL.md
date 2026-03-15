---
name: repo-branch-protection
description: Standardize GitHub branch protection for main and dev, enforce PR-only changes even for admins, and keep a clean PR-first workflow.
license: See repository license
---

# repo-branch-protection

Use this skill when you want a fast, repeatable baseline for repository branch safety:

- Protect `main` so direct pushes are blocked and changes only land through reviewed pull requests.
- Protect `dev` the same way so it also stays PR-only, even for repository admins.
- Disallow force pushes on protected branches.
- Keep feature branches unprotected so stacked PR workflows stay practical.

## Prerequisites

- GitHub CLI installed (`gh --version`).
- Authenticated CLI session with `repo` scope (`gh auth status`).
- Repo owner/name known (for example `DotNaos/pluto`).

## 1) Verify current branch state

Check the default branch and whether branch protection already exists:

```bash
gh repo view <owner>/<repo> --json nameWithOwner,defaultBranchRef
gh api repos/<owner>/<repo>/branches/main/protection
gh api repos/<owner>/<repo>/branches/dev/protection
```

If branch protection is missing, GitHub returns `404 Branch not protected`.

## 2) Protect `main`

Use this baseline:

- 1 approving review required
- admins also enforced
- force push disabled
- branch deletion disabled
- conversation resolution enabled

```bash
gh api --method PUT repos/<owner>/<repo>/branches/main/protection --input - <<'JSON'
{
  "required_status_checks": null,
  "enforce_admins": true,
  "required_pull_request_reviews": {
    "dismiss_stale_reviews": false,
    "require_code_owner_reviews": false,
    "required_approving_review_count": 1,
    "require_last_push_approval": false
  },
  "restrictions": null,
  "allow_force_pushes": false,
  "allow_deletions": false,
  "required_conversation_resolution": true,
  "lock_branch": false,
  "allow_fork_syncing": false
}
JSON
```

## 3) Create `dev` if needed

Create `dev` from the current `main` commit when it does not exist:

```bash
MAIN_SHA=$(gh api repos/<owner>/<repo>/git/ref/heads/main --jq '.object.sha')
gh api repos/<owner>/<repo>/git/refs -f ref='refs/heads/dev' -f sha="$MAIN_SHA"
```

## 4) Protect `dev`

Apply the same direct-push policy to `dev`. The key point is `enforce_admins: true`, otherwise admins can still bypass the PR-first flow.

```bash
gh api --method PUT repos/<owner>/<repo>/branches/dev/protection --input - <<'JSON'
{
  "required_status_checks": null,
  "enforce_admins": true,
  "required_pull_request_reviews": {
    "dismiss_stale_reviews": false,
    "require_code_owner_reviews": false,
    "required_approving_review_count": 1,
    "require_last_push_approval": false
  },
  "restrictions": null,
  "allow_force_pushes": false,
  "allow_deletions": false,
  "required_conversation_resolution": true,
  "lock_branch": false,
  "allow_fork_syncing": false
}
JSON
```

## 5) Verify

```bash
gh api repos/<owner>/<repo>/branches/main/protection
gh api repos/<owner>/<repo>/branches/dev/protection
gh api repos/<owner>/<repo>/branches --paginate --jq '.[] | {name, protected}'
```

Expected outcome:

- `main` is protected
- `dev` is protected
- both require PR approval
- both enforce rules for admins
- feature branches remain unprotected by default

## Recommended workflow after setup

### Day-to-day

- `feature/*` -> PR into `dev`
- if using stacked PRs, later feature branches can target earlier feature branches
- `dev` -> `main` for release

### Hotfix flow

1. branch from `main`: `hotfix/<short-name>`
2. PR into `main`
3. deploy production
4. back-merge `main` -> `dev`

## Quick checklist

- `main` protected with admin enforcement
- `dev` protected with admin enforcement
- direct pushes blocked on both branches
- feature branches left unprotected
- PR-first policy documented for the repository
