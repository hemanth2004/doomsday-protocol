package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
)

// The default model of the resource tree in the downloads view
func InitResourceTree() []Node {
	return []Node{
		{
			Value: "Default Resources",
			Desc:  styles.TreeDescriptionTitle.Render("Default Resources:") + "\nResources that the protocol recommends. Curated selection of resources that cover general aspects of survival",
			Children: []Node{
				{
					Value:    "Tier 0",
					Desc:     styles.TreeDescriptionTitle.Render("Tier 0:") + "\nAbsolutely necessary resources.",
					Children: []Node{},
				},
				{
					Value:    "Tier 1",
					Desc:     styles.TreeDescriptionTitle.Render("Tier 1:") + "\nSlightly less absolutely necessary. Still very much recommended.",
					Children: []Node{},
				},
			},
		},
		{
			Value:    "Custom Resources",
			Desc:     styles.TreeDescriptionTitle.Render("Custom Resources:") + "\nResources that the user has added.",
			Children: []Node{},
		},
	}
}

func ExampleResourceTree() []Node {
	return []Node{
		{
			Value: "Example",
			Desc:  "",
			Children: []Node{
				{
					Value: "Example",
					Desc:  "",
				},
				{
					Value: "Example",
					Desc:  "",
				},
			},
		},
	}
}

func GenerateResourceTree(m core.ResourceList) []Node {
	// Initialize the tree structure
	a := InitResourceTree()

	// Populate default resources with their respective tiers based on their Tier field
	for _, resource := range m.DefaultResources {
		newNode := Node{
			Value:    resource.Name,
			Desc:     styles.TreeDescriptionTitle.Render(resource.Name+":") + "\n" + resource.Description,
			Children: []Node{},
		}

		// Check the Tier and add to the corresponding tier node
		a[0].Children[resource.Tier].Children = append(a[0].Children[resource.Tier].Children, newNode)
	}

	// If no resources are added under a specific tier, add an empty node
	if len(a[0].Children[0].Children) == 0 {
		a[0].Children[0].Children = append(a[0].Children[0].Children, Node{
			Value: "-",
			Desc:  "No resources under here.",
		})
	}
	if len(a[0].Children[1].Children) == 0 {
		a[0].Children[1].Children = append(a[0].Children[1].Children, Node{
			Value: "-",
			Desc:  "No resources under here.",
		})
	}

	// Populate custom resources if there are any, otherwise add an empty node
	if len(m.CustomResources) > 0 {
		for _, resource := range m.CustomResources {
			newNode := Node{
				Value:    resource.Name,
				Children: []Node{},
			}
			a[1].Children = append(a[1].Children, newNode)
		}
	} else {
		a[1].Children = append(a[1].Children, Node{
			Value: "-",
			Desc:  "No custom resources.",
		})
	}

	return a
}

func GenerateGuideTree(path string) []Node {
	// Recursively build the tree
	root := buildTreeRecursively(path)

	return []Node{root}
}

// Helper function to recursively build the tree
func buildTreeRecursively(currentPath string) Node {
	// Get file or directory information
	info, err := os.Stat(currentPath)
	if err != nil {
		// Return an error node if there's an issue reading the path
		return Node{
			Value: filepath.Base(currentPath),
			Desc:  styles.TreeDescriptionTitle.Render("Error:") + "\n" + err.Error(),
		}
	}

	// Create a new node for the current file/directory
	node := Node{
		Value:    info.Name(),
		Desc:     currentPath,
		Children: []Node{},
	}

	// If it's a directory, recurse into its children
	if info.IsDir() {
		node.Value = node.Value + string(os.PathSeparator)

		entries, err := os.ReadDir(currentPath)
		if err != nil {
			// Add an error node if the directory can't be read
			node.Children = append(node.Children, Node{
				Value: "Error",
				Desc:  "Error reading directory: " + err.Error(),
			})
			return node
		}

		// Create separate slices for files and directories
		var files, directories []Node

		// Process each entry in the directory
		for _, entry := range entries {
			childPath := filepath.Join(currentPath, entry.Name())

			childNode := buildTreeRecursively(childPath)

			// Add to appropriate slice based on whether it's a directory
			if strings.HasSuffix(childNode.Value, string(os.PathSeparator)) {
				directories = append(directories, childNode)
			} else {
				files = append(files, childNode)
			}
		}

		// Combine files first, then directories
		node.Children = append(files, directories...)
	}

	return node
}

// Helper function to generate description for files/directories
func generateFileDescription(path string, info os.FileInfo) string {
	desc := styles.TreeDescriptionTitle.Render(info.Name() + ":")
	if info.IsDir() {
		desc += "\nDirectory"
	} else {
		desc += fmt.Sprintf("\nFile Size: %d bytes\nLast Modified: %s\nExtension: %s",
			info.Size(),
			info.ModTime().Format("2006-01-02 15:04:05"),
			filepath.Ext(path))
	}
	return desc
}

// Helper function to find the parent node based on relative path
func findParentNode(nodes []Node, relPath string) *Node {
	if relPath == "." || relPath == "" {
		return nil
	}

	netcode := strings.Split(relPath, string(os.PathSeparator))
	current := nodes

	for _, part := range netcode {
		found := false
		for i := range current {
			if current[i].Value == part {
				if i < len(current) {
					current = current[i].Children
					found = true
					break
				}
			}
		}
		if !found {
			return nil
		}
	}

	if len(current) > 0 {
		return &current[len(current)-1]
	}
	return nil
}
