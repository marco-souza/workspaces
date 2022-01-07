package main

import (
	"fmt"
	workspace "marco-souza/workspaces/cmd"
)

func main() {
	fmt.Print("test")
	workspace.ListWorkspaces()

	workspace.AddWorkspace("testolino")
	workspace.AddWorkspace("paeja")

	workspace.ListWorkspaces()

	workspace.RemoveWorkspace("paeja")
	workspace.ListWorkspaces()

	workspace.RemoveWorkspace("paeja")
	workspace.ListWorkspaces()

	workspace.OpenWorkspace("paeja")
	workspace.OpenWorkspace("testolino")
}
