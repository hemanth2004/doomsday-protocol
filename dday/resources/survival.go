package resources

import (
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/parts"
)

var (
	SurvivalBooks = []core.Resource{
		{
			Name:        "The Ultimate Survival Guide (txt)",
			Description: "Whether you’re lost in the woods, facing an armed insurrection, or preparing for a hurricane, the experts at Outdoor Life magazine are the people you want on your side. This book is the one you need if you want to protect your family, save yourself, and prevail over any danger.",
			Note:        "",
			Tier:        0,
			UrlGetter: core.UrlGetter{
				Key:           "tusg-txt",
				UpdatedURLURL: []string{},
				DefaultURLs: []string{
					"https://archive.org/stream/the-ultimate-bushcraft-survival-manual/The%20Ultimate%20Bushcraft%20Survival%20Manual_djvu.txt",
				},
			},
			FileName:         "The_Ultimate_Survial_Guide.txt",
			InitiateDownload: parts.InitiateFileDownload,
			Info:             core.ResourceInformation{},
			Status:           core.StatusQueued,
			Error:            nil,
		},

		{
			Name:        "The Ultimate Survival Guide (pdf)",
			Description: "Whether you’re lost in the woods, facing an armed insurrection, or preparing for a hurricane, the experts at Outdoor Life magazine are the people you want on your side. This book is the one you need if you want to protect your family, save yourself, and prevail over any danger.",
			Note:        "",
			Tier:        0,
			UrlGetter: core.UrlGetter{
				Key:           "tusg-pdf",
				UpdatedURLURL: []string{},
				DefaultURLs: []string{
					"https://ia803405.us.archive.org/16/items/the-ultimate-bushcraft-survival-manual/The%20Ultimate%20Bushcraft%20Survival%20Manual.pdf",
				},
			},
			FileName:         "The_Ultimate_Survial_Guide.pdf",
			InitiateDownload: parts.InitiateFileDownload,
			Info:             core.ResourceInformation{},
			Status:           core.StatusQueued,
			Error:            nil,
		},

		{
			Name:        "The Survival Medicine Handbook (pdf)",
			Description: "The Survival Medicine Handbook is written in plain English that anyone can understand. But it’s unique in that it assumes that a disaster, natural or man-made, has removed all access to hospitals or doctors for the foreseeable future; you, the average person, are now the highest medical resource left to your family.",
			Note:        "",
			Tier:        0,
			UrlGetter: core.UrlGetter{
				Key:           "tsmh-pdf",
				UpdatedURLURL: []string{},
				DefaultURLs: []string{
					"https://ia800400.us.archive.org/35/items/the-survival-medicine-handbook/The%20Survival%20Medicine%20Handbook.pdf",
				},
			},
			FileName:         "The_Survial_Medicine_Handbook.pdf",
			InitiateDownload: parts.InitiateFileDownload,
			Info:             core.ResourceInformation{},
			Status:           core.StatusQueued,
			Error:            nil,
		},

		{
			Name:        "The Survival Medicine Handbook (txt)",
			Description: "The Survival Medicine Handbook is written in plain English that anyone can understand. But it’s unique in that it assumes that a disaster, natural or man-made, has removed all access to hospitals or doctors for the foreseeable future; you, the average person, are now the highest medical resource left to your family.",
			Note:        "",
			Tier:        0,
			UrlGetter: core.UrlGetter{
				Key:           "tsmh-txt",
				UpdatedURLURL: []string{},
				DefaultURLs: []string{
					"https://archive.org/stream/the-survival-medicine-handbook/The%20Survival%20Medicine%20Handbook_djvu.txt",
				},
			},
			FileName:         "The_Survial_Medicine_Handbook.txt",
			InitiateDownload: parts.InitiateFileDownload,
			Info:             core.ResourceInformation{},
			Status:           core.StatusQueued,
			Error:            nil,
		},
	}
)
