package dday

import (
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/parts"
)

var DefaultResources []core.Resource = []core.Resource{

	{
		Name:        "Indian OSM",
		Description: "Maps",
		Note:        "",
		UrlGetter: core.UrlGetter{
			Key:           "india-osm",
			UpdatedURLURL: []string{},
			DefaultURLs: []string{
				"https://download.geofabrik.de/asia/india-latest.osm.pbf",
			},
		},
		FileName:         "india-latest.osm.pbf",
		InitiateDownload: parts.InitiateOSMDownload,
		Info:             core.ResourceInformation{},
		Status:           core.StatusQueued,
		Error:            nil,
	},

	{
		Name:        "Wikipedia English",
		Description: "All text of all topics of the english Wikipedia",
		Note:        "",
		UrlGetter: core.UrlGetter{
			Key:           "simple-wikipedia",
			UpdatedURLURL: []string{},
			DefaultURLs: []string{
				"https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2",
			},
		},
		FileName:         "enwiki-latest-pages-articles.xml.bz2",
		InitiateDownload: parts.InitiateWikipediaDownload,
		Info:             core.ResourceInformation{},
		Status:           core.StatusQueued,
		Error:            nil,
	},
}
