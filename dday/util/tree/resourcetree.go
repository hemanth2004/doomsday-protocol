package tree

import "github.com/hemanth2004/doomsday-protocol/dday/ui/styles"

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
