package cmd

import (
	"fmt"
	"os"
	"path"
)

type Project struct {
	Name           string
	Repository     string
	PackageManager string
}

type Workspace struct {
	Projects []Project
	Path     string
	Name     string
}

type WorkspaceManager struct {
	Workspaces []Workspace
	RootFolder string
}

var wsm = WorkspaceManager{
	Workspaces: make([]Workspace, 0),
	// TODO: get path from config
	RootFolder: os.Getenv("WORKSPACE"),
}

func AddWorkspace(name string) (*Workspace, error) {
	// TODO: Check if workspace already exists
	fmt.Printf("Adding %s workspace...\n", name)

	newWorkspace := Workspace{
		Projects: make([]Project, 0),
		Path:     path.Join(wsm.RootFolder, name),
		Name:     name,
	}

	// TODO: create folder with fs
	fmt.Println("✅ Added!")
	wsm.Workspaces = append(wsm.Workspaces, newWorkspace)

	return &newWorkspace, nil
}

func ListWorkspaces() {
	if len(wsm.Workspaces) == 0 {
		fmt.Println("😢 No workspace found")
		return
	}

	for _, ws := range wsm.Workspaces {
		fmt.Printf("  - 📦 %s\n", ws.Name)
	}
}

func RemoveWorkspace(name string) {
	index, ws := findWorkspaceByName(name)

	if ws != nil {
		// TODO: config deletion
		// TODO: delete fs workspace
		wsm.Workspaces = append(wsm.Workspaces[:index], wsm.Workspaces[index+1:]...)
		fmt.Printf("✅ Workspace %s deleted!\n", ws.Name)
		return
	}

	fmt.Println("😢 Workspace not found!")
}

func OpenWorkspace(name string) {
	_, ws := findWorkspaceByName(name)

	if ws == nil {
		fmt.Println("😢 Workspace not found!")
		return
	}

	// TODO: change current shell work dir
	workspacePath := path.Join(ws.Path)
	fmt.Println("Going to ", workspacePath)
}

func findWorkspaceByName(name string) (int, *Workspace) {
	for index, ws := range wsm.Workspaces {
		if ws.Name == name {
			return index, &ws
		}
	}
	return -1, nil
}
