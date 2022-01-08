package cli

import (
	"fmt"

	"github.com/marco-souza/workspaces/cmd"
	"github.com/spf13/cobra"
)

func rootHandler(command *cobra.Command, args []string) {
	fmt.Println("Hello from root handler")

	cmd.AddWorkspace("testolino")
	cmd.AddWorkspace("paeja")

	cmd.ListWorkspaces()

	cmd.RemoveWorkspace("paeja")
	cmd.ListWorkspaces()

	cmd.RemoveWorkspace("paeja")
	cmd.ListWorkspaces()

	cmd.OpenWorkspace("paeja")
	cmd.OpenWorkspace("testolino")
}

func helpHandler(command *cobra.Command, args []string) {
	command.Help()
}

func listWorspaceHandler(command *cobra.Command, args []string) {
	cmd.ListWorkspaces()
}
