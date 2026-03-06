import type { AuthStatus } from "./types.js";

export function getApiKey(env: NodeJS.ProcessEnv = process.env): string | null {
  const value = env.OPENAI_API_KEY?.trim();
  return value ? value : null;
}

export function getAuthStatus(env: NodeJS.ProcessEnv = process.env): AuthStatus {
  return {
    available: getApiKey(env) !== null,
    source: "OPENAI_API_KEY",
  };
}
