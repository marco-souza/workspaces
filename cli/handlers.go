package cli

import (
	"fmt"
	"os"

	"github.com/marco-souza/workspaces/cmd"
	"github.com/spf13/cobra"
)

func helpHandler(command *cobra.Command, args []string) {
	command.Help()
}

func listWorspaceHandler(command *cobra.Command, args []string) {
	cmd.ListWorkspaces()
}

func addWorspaceHandler(command *cobra.Command, args []string) {
	validatingWorkspaceInArgs(args)
	cmd.AddWorkspace(args[0])
}

func removeWorspaceHandler(command *cobra.Command, args []string) {
	validatingWorkspaceInArgs(args)
	cmd.RemoveWorkspace(args[0])
}

func validatingWorkspaceInArgs(args []string) {
	if len(args) == 0 {
		fmt.Printf(MISSING_WORKSPACE_TEMPLATE, "remove")
		os.Exit(0)
	}
}

const MISSING_WORKSPACE_TEMPLATE = "You need to specify the workspace to %s\n"
