package dday

import (
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/parts"
)

var Resources = []core.Download{
	{
		Name:             "Indian OSM",
		Description:      "Maps",
		Url:              "https://download.geofabrik.de/asia/india-latest.osm.pbf",
		FileName:         "india-latest.osm.pbf",
		InitiateDownload: parts.InitiateOSMDownload,
		Info:             core.DownloadInformation{},
		Status:           core.StatusQueued,
		Error:            nil,
	},
	{
		Name:             "Wikipedia English",
		Description:      "All text of all topics of the english Wikipedia",
		Url:              "https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2",
		FileName:         "enwiki-latest-pages-articles.xml.bz2",
		InitiateDownload: parts.InitiateWikipediaDownload,
		Info:             core.DownloadInformation{},
		Status:           core.StatusQueued,
		Error:            nil,
	},
}
