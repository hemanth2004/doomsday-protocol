package parts

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

const dumpURL = "https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2"

func DownloadWikipediaDump(relativePath string) error {

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	fmt.Printf("Starting download from %s...\n", dumpURL)

	resp, err := client.Get(dumpURL)
	if err != nil {
		return fmt.Errorf("error fetching dump: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: HTTP %d", resp.StatusCode)
	}

	file, err := os.Create(relativePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading",
	)

	_, err = io.Copy(io.MultiWriter(file, bar), resp.Body)
	if err != nil {
		return fmt.Errorf("error saving dump: %v", err)
	}

	fmt.Println("\nDownload complete!")
	return nil
}
