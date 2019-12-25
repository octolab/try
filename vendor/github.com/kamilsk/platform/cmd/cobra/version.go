package cobra

import (
	"runtime"

	"github.com/spf13/cobra"
)

// NewVersionCommand returns new version command.
// Deprecated: use go.octolab.org/toolkit/cli/cobra.NewVersionCommand instead.
func NewVersionCommand(commit, date, release string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show application version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf(
				"Version %s (commit: %s, build date: %s, go version: %s, compiler: %s, platform: %s/%s)\n",
				release, commit, date, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
			)
		},
		Version: release,
	}
}
