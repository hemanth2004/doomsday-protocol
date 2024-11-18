package main

import (
	"fmt"
	"os"

	"dday-ptcl/parts"
)

func initiateWikipediaPart() {
	outputFile := "downloads/enwiki-latest-pages-articles.xml.bz2"

	if err := os.MkdirAll("downloads", os.ModePerm); err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	if err := parts.DownloadWikipediaDump(outputFile); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func main() {
	initiateWikipediaPart()
}
