package resources

import (
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/core/guides"
	"github.com/hemanth2004/doomsday-protocol/dday/core/netcode"
)

var CoreResources []core.Resource = []core.Resource{

	{
		Name:                "OpenStreetMaps India",
		Description:         "An offline .osm file provides a snapshot of geographic data for a specific area, mainly including:\n1.Road Networks\n2.Points of Interest\n3.Water Features\n4.Building Footprints",
		AssociatedGuidePath: guides.Guide("Guides/DefaultResources/Core/osm.md"),

		UrlGetter: &core.UrlGetter{
			Key:           "india-osm",
			UpdatedURLURL: []string{},
			DefaultURLs: []string{
				"https://download.geofabrik.de/asia/india-latest.osm.pbf",
			},
		},
		FileName:         "india-latest.osm.pbf",
		InitiateDownload: netcode.InitiateHTTPDownload,
		Info:             core.ResourceInformation{},
		Status:           core.StatusIdle,
		Error:            nil,
		ControlChannel:   make(chan core.DownloadControl),
	},

	{
		Name:                "Wikipedia English",
		Description:         "All text of all topics of the english Wikipedia.",
		AssociatedGuidePath: guides.Guide("Guides/DefaultResources/Core/wikipedia.md"),

		UrlGetter: &core.UrlGetter{
			Key:           "simple-wikipedia",
			UpdatedURLURL: []string{},
			DefaultURLs: []string{
				"https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2",
			},
		},
		FileName:         "enwiki-latest-pages-articles.xml.bz2",
		InitiateDownload: netcode.InitiateHTTPDownload,
		Info:             core.ResourceInformation{},
		Status:           core.StatusIdle,
		Error:            nil,
		ControlChannel:   make(chan core.DownloadControl),
	},

	SurvivalBooks[0],
	SurvivalBooks[1],
	SurvivalBooks[2],
	SurvivalBooks[3],

	{
		Name:                "Sample 100MB",
		Description:         "Sample 100MB file",
		Tier:                1,
		AssociatedGuidePath: guides.Guide("Guides/DefaultResources/what-is-a-default-resource.md"),

		UrlGetter: &core.UrlGetter{
			Key:           "sample-100mb",
			UpdatedURLURL: []string{},
			DefaultURLs: []string{
				"http://speedtest.tele2.net/100MB.zip",
			},
		},
		FileName:         "100MB.zip",
		InitiateDownload: netcode.InitiateHTTPDownload,
		Info:             core.ResourceInformation{},
		Status:           core.StatusIdle,
		Error:            nil,
		ControlChannel:   make(chan core.DownloadControl),
	},

	// {
	// 	Name:        "Sample 1GB",
	// 	Description: "Sample 1GB file",
	// 	Tier:        1,
	// 	UrlGetter: core.UrlGetter{
	// 		Key:           "sample-1gb",
	// 		UpdatedURLURL: []string{},
	// 		DefaultURLs: []string{
	// 			"http://speedtest.tele2.net/1GB.zip",
	// 		},
	// 	},
	// 	FileName:         "1GB.zip",
	// 	InitiateDownload: netcode.InitiateFileDownload,
	// 	Info:             core.ResourceInformation{},
	// 	Status:           core.StatusIdle,
	// 	Error:            nil,
	// },
}
