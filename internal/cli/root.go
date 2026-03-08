package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/DotNaos/project-toolkit/internal/nodebridge"
	"github.com/DotNaos/project-toolkit/internal/projectinit"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "pkit",
		Short:         "project-toolkit hybrid CLI frontend",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(
		newSkillsCommand(),
		newProjectCommand(),
		newPlanCommand(),
		newRunCommand(),
		newDevCommand(),
		newAuthCommand(),
		newCompletionCommand(rootCmd),
	)

	return rootCmd
}

func newSkillsCommand() *cobra.Command {
	skillsCmd := &cobra.Command{
		Use:   "skills",
		Short: "Manage bundled skills",
	}

	skillsCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List discovered skills",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			return nodebridge.Run([]string{"skills", "list"})
		},
	})

	return skillsCmd
}

func newProjectCommand() *cobra.Command {
	projectCmd := &cobra.Command{
		Use:   "project",
		Short: "Project/workspace/worktree operations",
	}

	projectCmd.AddCommand(newProjectInitCommand())
	projectCmd.AddCommand(newProjectWorkspaceCommand())
	projectCmd.AddCommand(newProjectWorktreeCommand())

	return projectCmd
}

func newProjectInitCommand() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize project-toolkit scaffold files",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to resolve working directory: %w", err)
			}

			result, err := projectinit.Initialize(cwd, force)
			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Initialized project-toolkit scaffold%s.\n", forceSuffix(force))
			printFileGroup(cmd.OutOrStdout(), "Created", result.Created)
			printFileGroup(cmd.OutOrStdout(), "Updated", result.Updated)
			printFileGroup(cmd.OutOrStdout(), "Skipped", result.Skipped)
			return nil
		},
	}

	cmd.Flags().BoolVar(&force, "force", false, "rewrite scaffold files even if they already exist")
	return cmd
}

func newProjectWorkspaceCommand() *cobra.Command {
	workspaceCmd := &cobra.Command{
		Use:   "workspace",
		Short: "Generate and manage workspaces",
	}

	var name string
	var root string
	var output string

	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a workspace file for a target root",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			bridgeArgs := []string{"project", "workspace", "generate"}
			if name != "" {
				bridgeArgs = append(bridgeArgs, "--name", name)
			}
			if root != "" {
				bridgeArgs = append(bridgeArgs, "--root", root)
			}
			if output != "" {
				bridgeArgs = append(bridgeArgs, "--output", output)
			}

			return nodebridge.Run(bridgeArgs)
		},
	}

	generateCmd.Flags().StringVar(&name, "name", "", "logical workspace name")
	generateCmd.Flags().StringVar(&root, "root", "", "target repository or worktree root")
	generateCmd.Flags().StringVar(&output, "output", "", "target workspace file path")
	workspaceCmd.AddCommand(generateCmd)

	return workspaceCmd
}

func newProjectWorktreeCommand() *cobra.Command {
	worktreeCmd := &cobra.Command{
		Use:   "worktree",
		Short: "Create and manage worktrees",
	}

	var branch string
	var base string
	var workspace string
	var output string

	createCmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a managed worktree and matching workspace",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			bridgeArgs := []string{"project", "worktree", "create", args[0]}
			if branch != "" {
				bridgeArgs = append(bridgeArgs, "--branch", branch)
			}
			if base != "" {
				bridgeArgs = append(bridgeArgs, "--base", base)
			}
			if workspace != "" {
				bridgeArgs = append(bridgeArgs, "--workspace", workspace)
			}
			if output != "" {
				bridgeArgs = append(bridgeArgs, "--output", output)
			}

			return nodebridge.Run(bridgeArgs)
		},
	}

	createCmd.Flags().StringVar(&branch, "branch", "", "override Git branch name")
	createCmd.Flags().StringVar(&base, "base", "", "base ref for new branches")
	createCmd.Flags().StringVar(&workspace, "workspace", "", "override generated workspace name")
	createCmd.Flags().StringVar(&output, "output", "", "override generated workspace file path")
	worktreeCmd.AddCommand(createCmd)

	return worktreeCmd
}

func newPlanCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "plan <skill-id>",
		Short: "Plan an agent run through the TypeScript Codex adapter",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return nodebridge.Run([]string{"plan", args[0]})
		},
	}
}

func newRunCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "run <skill-id>",
		Short: "Execute an agent run through the TypeScript Codex adapter",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return nodebridge.Run([]string{"run", args[0]})
		},
	}
}

func newDevCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "dev [--] <command...>",
		Short: "Run a local development command with logging",
		Args:  cobra.ArbitraryArgs,
		RunE: func(_ *cobra.Command, args []string) error {
			bridgeArgs := []string{"dev"}
			bridgeArgs = append(bridgeArgs, args...)
			return nodebridge.Run(bridgeArgs)
		},
	}
}

func newAuthCommand() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication helpers",
	}

	authCmd.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "Show authentication status for the current adapter",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			return nodebridge.Run([]string{"auth", "status"})
		},
	})

	return authCmd
}

func newCompletionCommand(rootCmd *cobra.Command) *cobra.Command {
	completionCmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate shell completions from the Go command tree",
	}

	completionCmd.AddCommand(&cobra.Command{
		Use:   "zsh",
		Short: "Print zsh completion",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			return rootCmd.GenZshCompletion(io.Writer(rootCmd.OutOrStdout()))
		},
	})

	completionCmd.AddCommand(&cobra.Command{
		Use:   "bash",
		Short: "Print bash completion",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			return rootCmd.GenBashCompletion(io.Writer(rootCmd.OutOrStdout()))
		},
	})

	completionCmd.AddCommand(&cobra.Command{
		Use:   "bridge-zsh",
		Short: "Print the current TypeScript-generated zsh completion during migration",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			return nodebridge.Run([]string{"completion", "zsh"})
		},
	})

	completionCmd.Long = fmt.Sprintf(
		"The Go CLI owns shell completions going forward. During migration, bridge-zsh still exposes the existing TypeScript completion output.",
	)

	return completionCmd
}

func forceSuffix(force bool) string {
	if force {
		return " (force mode)"
	}

	return ""
}

func printFileGroup(writer io.Writer, label string, files []string) {
	if len(files) == 0 {
		return
	}

	fmt.Fprintf(writer, "%s:\n", label)
	for _, file := range files {
		fmt.Fprintf(writer, "- %s\n", file)
	}
}
