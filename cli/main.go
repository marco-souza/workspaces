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
			Run:     addWorspaceHandler,
		},
		&cobra.Command{
			Use:     "remove",
			Aliases: []string{"rm", "r"},
			Short:   "remove workspace",
			Run:     removeWorspaceHandler,
		},
		&cobra.Command{
			Use:     "open",
			Aliases: []string{"o"},
			Short:   "open workspace",
			Run:     openWorspaceHandler,
		},
	)

	return rootCmd
}
