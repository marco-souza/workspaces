package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	workspaceCmd := makeWorkspaceCmd()
	if err := workspaceCmd.Execute(); err != nil {
		fmt.Printf("😵 %s\n", err)
		os.Exit(1)
	}
}

func makeWorkspaceCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ws",
		Short: "Workspace manager",
		Run:   helpHandler,
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:     "list",
			Aliases: []string{"ls", "l"},
			Short:   "list workspaces",
			Run:     listWorspaceHandler,
		},
		&cobra.Command{
			Use:     "add",
			Aliases: []string{"new"},
			Short:   "add new workspace",
			Run:     rootHandler,
		},
		&cobra.Command{
			Use:     "remove",
			Aliases: []string{"rm", "r"},
			Short:   "remove workspace",
			Run:     rootHandler,
		},
	)

	return rootCmd
}
