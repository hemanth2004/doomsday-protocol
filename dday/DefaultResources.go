package dday

import (
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/core/netcode"
	"github.com/hemanth2004/doomsday-protocol/dday/resources"
)

var DefaultResources []core.Resource = []core.Resource{

	{
		Name:        "OpenStreetMaps India",
		Description: "The latest map of all of India.",
		Note:        "",
		UrlGetter: core.UrlGetter{
			Key:           "india-osm",
			UpdatedURLURL: []string{},
			DefaultURLs: []string{
				"https://download.geofabrik.de/asia/india-latest.osm.pbf",
			},
		},
		FileName:         "india-latest.osm.pbf",
		InitiateDownload: netcode.InitiateHTTPDownload,
		Info:             core.ResourceInformation{},
		Status:           core.StatusQueued,
		Error:            nil,
	},

	{
		Name:        "Wikipedia English",
		Description: "All text of all topics of the english Wikipedia.",
		Note:        "",
		UrlGetter: core.UrlGetter{
			Key:           "simple-wikipedia",
			UpdatedURLURL: []string{},
			DefaultURLs: []string{
				"https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2",
			},
		},
		FileName:         "enwiki-latest-pages-articles.xml.bz2",
		InitiateDownload: netcode.InitiateHTTPDownload,
		Info:             core.ResourceInformation{},
		Status:           core.StatusQueued,
		Error:            nil,
	},

	resources.SurvivalBooks[0],
	resources.SurvivalBooks[1],
	resources.SurvivalBooks[2],
	resources.SurvivalBooks[3],

	{
		Name:        "Sample 100MB",
		Description: "Sample 100MB file",
		Note:        "",
		Tier:        1,
		UrlGetter: core.UrlGetter{
			Key:           "sample-100mb",
			UpdatedURLURL: []string{},
			DefaultURLs: []string{
				"http://speedtest.tele2.net/100MB.zip",
			},
		},
		FileName:         "100MB.zip",
		InitiateDownload: netcode.InitiateHTTPDownload,
		Info:             core.ResourceInformation{},
		Status:           core.StatusQueued,
		Error:            nil,
	},

	// {
	// 	Name:        "Sample 1GB",
	// 	Description: "Sample 1GB file",
	// 	Note:        "",
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
	// 	Status:           core.StatusQueued,
	// 	Error:            nil,
	// },
}
