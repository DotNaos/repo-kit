package nodebridge

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type ExitError struct {
	Code int
	Err  error
}

func (e *ExitError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return fmt.Sprintf("command exited with code %d", e.Code)
}

func Run(args []string) error {
	entryPath, err := resolveTypeScriptEntry()
	if err != nil {
		return err
	}

	commandArgs := append([]string{entryPath}, args...)
	cmd := exec.Command("node", commandArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		var execErr *exec.ExitError
		if ok := AsExitError(err, &execErr); ok {
			return &ExitError{Code: execErr.ExitCode(), Err: err}
		}

		return err
	}

	return nil
}

func resolveTypeScriptEntry() (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to resolve bridge source path")
	}

	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(currentFile), "..", ".."))
	entryPath := filepath.Join(repoRoot, "dist", "cli", "index.js")

	if _, err := os.Stat(entryPath); err != nil {
		return "", fmt.Errorf("typescript entrypoint not found at %s: %w", entryPath, err)
	}

	return entryPath, nil
}

func AsExitError(err error, target **exec.ExitError) bool {
	if err == nil {
		return false
	}

	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		return false
	}

	*target = exitErr
	return true
}
