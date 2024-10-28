package main

import (
	"fmt"
	"log"
	"torrent-rss/internal/config"
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

	for _, item := range matches {
		fmt.Println("=== Match Found ===")
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Date: %s\n", item.PubDate)
		fmt.Printf("Details: %s\n\n", item.Description)
	}

	fmt.Printf("Total matches found: %d\n", len(matches))
}
