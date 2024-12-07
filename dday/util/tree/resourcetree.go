package tree

import (
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
