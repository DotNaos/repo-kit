#!/usr/bin/env node

import path from "node:path";
import process from "node:process";
import { CodexSdkAdapter } from "../adapters/codex-sdk/adapter.js";
import { getAuthStatus } from "../core/auth.js";
import { RepoKitError } from "../core/errors.js";
import { getPackageRoot } from "../core/package-root.js";
import { collectRepoContext } from "../core/repo-context.js";
import { discoverSkills, loadSkill } from "../core/skills.js";
import type {
  LoadedSkill,
  PlanExecutionResult,
  RepoContext,
  SkillSummary,
  TaskExecutionResult,
} from "../core/types.js";

async function main(): Promise<void> {
  const args = process.argv.slice(2);

  if (args.length === 0 || args.includes("--help") || args.includes("-h")) {
    printUsage();
    process.exit(args.length === 0 ? 1 : 0);
  }

  const packageRoot = getPackageRoot();
  const skillsRoot = path.join(packageRoot, "skills");

  switch (args[0]) {
    case "skills":
      await handleSkillsCommand(args.slice(1), skillsRoot);
      return;
    case "plan":
      await handleAgentCommand("plan", args.slice(1), skillsRoot);
      return;
    case "run":
      await handleAgentCommand("run", args.slice(1), skillsRoot);
      return;
    case "auth":
      await handleAuthCommand(args.slice(1));
      return;
    default:
      throw new RepoKitError(`Unknown command: ${args[0]}`);
  }
}

async function handleSkillsCommand(args: string[], skillsRoot: string): Promise<void> {
  if (args[0] !== "list" || args.length !== 1) {
    throw new RepoKitError("Usage: repo-kit skills list");
  }

  const skills = await discoverSkills(skillsRoot);
  printSkillSummaries(skills);
}

async function handleAgentCommand(mode: "plan" | "run", args: string[], skillsRoot: string): Promise<void> {
  const skillId = args[0];
  if (!skillId || args.length !== 1) {
    throw new RepoKitError(`Usage: repo-kit ${mode} <skill-id>`);
  }

  const skill = await loadSkill(skillsRoot, skillId);
  const repoContext = await collectRepoContext(process.cwd());

  if (!repoContext.gitRoot) {
    throw new RepoKitError("Current working directory must be inside a Git repository");
  }

  const adapter = new CodexSdkAdapter();
  if (mode === "plan") {
    const result = await adapter.plan({ skill, repoContext });
    printPlanResult(skill, repoContext, result);
    return;
  }

  const result = await adapter.run({ skill, repoContext });
  printRunResult(skill, repoContext, result);
}

async function handleAuthCommand(args: string[]): Promise<void> {
  if (args[0] !== "status" || args.length !== 1) {
    throw new RepoKitError("Usage: repo-kit auth status");
  }

  const auth = getAuthStatus();
  console.log(`Auth source: ${auth.source}`);
  console.log(`Status: ${auth.available ? "available" : "missing"}`);
  console.log(
    auth.available
      ? "repo-kit can authenticate to Codex with the configured API key."
      : "Set OPENAI_API_KEY to enable plan/run commands.",
  );
}

function printSkillSummaries(skills: SkillSummary[]): void {
  if (skills.length === 0) {
    console.log("No skills found.");
    return;
  }

  const rows = skills.map((skill) => ({
    id: skill.id,
    title: skill.title ?? "-",
    valid: skill.valid ? "yes" : "no",
    notes: skill.errors.join("; ") || "-",
  }));

  const widths = {
    id: Math.max("ID".length, ...rows.map((row) => row.id.length)),
    title: Math.max("TITLE".length, ...rows.map((row) => row.title.length)),
    valid: Math.max("VALID".length, ...rows.map((row) => row.valid.length)),
  };

  console.log(
    `${pad("ID", widths.id)}  ${pad("TITLE", widths.title)}  ${pad("VALID", widths.valid)}  NOTES`,
  );
  for (const row of rows) {
    console.log(
      `${pad(row.id, widths.id)}  ${pad(row.title, widths.title)}  ${pad(row.valid, widths.valid)}  ${row.notes}`,
    );
  }
}

function printPlanResult(skill: LoadedSkill, repoContext: RepoContext, result: PlanExecutionResult): void {
  console.log(`Skill: ${skill.id}`);
  console.log(`Working directory: ${repoContext.cwd}`);
  if (result.plan) {
    console.log("");
    console.log(`Summary: ${result.plan.summary}`);
    console.log("");
    console.log("Steps:");
    for (const step of result.plan.steps) {
      console.log(`- ${step}`);
    }
    console.log("");
    console.log("Risks:");
    if (result.plan.risks.length === 0) {
      console.log("- none");
    } else {
      for (const risk of result.plan.risks) {
        console.log(`- ${risk}`);
      }
    }
  } else {
    console.log("");
    console.log(result.finalResponse);
  }

  printExecutionFooter(result);
}

function printRunResult(skill: LoadedSkill, repoContext: RepoContext, result: TaskExecutionResult): void {
  console.log(`Skill: ${skill.id}`);
  console.log(`Working directory: ${repoContext.cwd}`);
  console.log("");
  console.log(result.finalResponse);
  printExecutionFooter(result);
}

function printExecutionFooter(result: PlanExecutionResult | TaskExecutionResult): void {
  console.log("");
  console.log(`Thread ID: ${result.threadId ?? "n/a"}`);

  if (result.changedFiles.length > 0) {
    console.log("Changed files:");
    for (const changedFile of result.changedFiles) {
      console.log(`- ${changedFile}`);
    }
  }

  if (result.commands.length > 0) {
    console.log("Commands:");
    for (const command of result.commands) {
      console.log(`- ${command}`);
    }
  }

  if (result.usage) {
    console.log(
      `Usage: input=${result.usage.inputTokens}, cached=${result.usage.cachedInputTokens}, output=${result.usage.outputTokens}`,
    );
  }
}

function printUsage(): void {
  console.log(`repo-kit skills list
repo-kit plan <skill-id>
repo-kit run <skill-id>
repo-kit auth status`);
}

function pad(value: string, width: number): string {
  return value.padEnd(width, " ");
}

main().catch((error: unknown) => {
  const message = error instanceof Error ? error.message : String(error);
  console.error(`error: ${message}`);
  process.exit(error instanceof RepoKitError ? error.exitCode : 1);
});
