export class RepoKitError extends Error {
  readonly exitCode: number;

  constructor(message: string, exitCode = 1) {
    super(message);
    this.name = "RepoKitError";
    this.exitCode = exitCode;
  }
}
