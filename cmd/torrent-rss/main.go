package main

import (
	"fmt"
	"log"

	"torrent-rss/internal/config"
	"torrent-rss/internal/downloader"
	"torrent-rss/internal/parser"
)

func main() {
	cfg := config.NewConfig()
	p := parser.NewParser()

	fmt.Printf("Searching for %v matches with 1080p...\n\n", cfg.SearchTerms)

	matches, err := p.FetchAndParse(cfg.RSSURL, cfg.SearchTerms)
	if err != nil {
		log.Fatalf("Error parsing RSS feed: %v", err)
	}

	if len(matches) == 0 {
		fmt.Println("No matches found!")
		return
	}

	// Initialize downloader with RSS URL for auth
	d, err := downloader.NewDownloader(cfg.DownloadPath, cfg.RSSURL)
	if err != nil {
		log.Fatalf("Error creating downloader: %v", err)
	}

	for _, item := range matches {
		fmt.Println("=== Match Found ===")
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Date: %s\n", item.Description)

		// Download the torrent
		fmt.Printf("Downloading torrent file...\n")
		if err := d.DownloadTorrent(item.Link); err != nil {
			fmt.Printf("Error downloading torrent: %v\n", err)
			continue
		}
	}

	fmt.Printf("Total matches found: %d\n", len(matches))
}
