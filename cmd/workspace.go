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
	RootFolder string
}

var wsm = WorkspaceManager{
	RootFolder: os.Getenv("WORKSPACE"),
}

func AddWorkspace(name string) (*Workspace, error) {
	ws := findWorkspaceByName(name)

	if ws != nil {
		fmt.Println("😆 Workspace already exists")
		os.Exit(0)
	}

	fmt.Printf("Adding %s workspace...\n", name)

	newWorkspace := Workspace{
		Projects: []Project{},
		Path:     path.Join(wsm.RootFolder, name),
		Name:     name,
	}

	// create folder with fs
	if err := os.Mkdir(newWorkspace.Path, 0755); err != nil {
		fmt.Println("😢 Failed creating workspace", err)
		os.Exit(1)
	}

	fmt.Println("✅ Added!")
	return &newWorkspace, nil
}

func ListWorkspaces() {
	workspaces := loadWorkspaces()

	if len(workspaces) == 0 {
		fmt.Println("😢 No workspace found")
		os.Exit(0)
	}

	for _, ws := range workspaces {
		fmt.Printf("  - 📦 %s\n", ws.Name)
	}
}

func RemoveWorkspace(name string) {
	ws := findWorkspaceByName(name)

	if ws == nil {
		fmt.Println("😅 Workspace not found!")
		os.Exit(0)
	}

	// delete fs workspace
	if err := os.Remove(ws.Path); err != nil {
		fmt.Println("😢 Failed deleting workspace", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Workspace %s deleted!\n", ws.Name)
}

func OpenWorkspace(name string) {
	ws := findWorkspaceByName(name)

	if ws == nil {
		fmt.Println("😢 Workspace not found!")
		os.Exit(0)
	}

	// TODO: change current shell work dir
	workspacePath := path.Join(ws.Path)
	fmt.Println("Going to ", workspacePath)
}

func findWorkspaceByName(name string) *Workspace {
	workspaces := loadWorkspaces()
	for _, ws := range workspaces {
		if ws.Name == name {
			return &ws
		}
	}
	return nil
}

func loadWorkspaces() []Workspace {
	files, err := os.ReadDir(wsm.RootFolder)
	if err != nil {
		fmt.Println("😢 failed to load workspaces", err)
		os.Exit(1)
	}

	workspaces := []Workspace{}
	for _, file := range files {
		newWorkspace := Workspace{
			Projects: []Project{},
			Path:     path.Join(wsm.RootFolder, file.Name()),
			Name:     file.Name(),
		}

		workspaces = append(workspaces, newWorkspace)
	}

	return workspaces
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
