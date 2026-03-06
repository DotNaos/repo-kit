export interface SkillSummary {
  id: string;
  title: string | null;
  valid: boolean;
  format: "normalized" | "legacy" | "markdown-fallback" | "invalid";
  rootDir: string;
  errors: string[];
}

export interface LoadedSkill extends SkillSummary {
  title: string;
  prompt: string;
  description?: string;
  promptPath?: string;
  metadataPath?: string;
}

export interface RepoContextFilePreview {
  path: string;
  preview: string;
}

export interface RepoContext {
  cwd: string;
  gitRoot: string | null;
  gitBranch: string | null;
  topLevelEntries: string[];
  filePreviews: RepoContextFilePreview[];
}

export interface AuthStatus {
  available: boolean;
  source: "OPENAI_API_KEY";
}

export interface ExecutionUsage {
  inputTokens: number;
  cachedInputTokens: number;
  outputTokens: number;
}

export interface PlannedChanges {
  summary: string;
  steps: string[];
  risks: string[];
}

export interface ExecutionSummary {
  threadId: string | null;
  finalResponse: string;
  changedFiles: string[];
  commands: string[];
  usage: ExecutionUsage | null;
}

export interface PlanExecutionResult extends ExecutionSummary {
  mode: "plan";
  plan: PlannedChanges | null;
}

export interface TaskExecutionResult extends ExecutionSummary {
  mode: "run";
}

export type AdapterExecutionResult = PlanExecutionResult | TaskExecutionResult;

export interface SkillExecutionInput {
  skill: LoadedSkill;
  repoContext: RepoContext;
}

export interface AgentAdapter {
  plan(input: SkillExecutionInput): Promise<PlanExecutionResult>;
  run(input: SkillExecutionInput): Promise<TaskExecutionResult>;
}
