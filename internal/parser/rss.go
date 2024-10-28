package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"torrent-rss/internal/models"
)

type Parser struct {
	config *http.Client
}

func NewParser() *Parser {
	return &Parser{
		config: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (p *Parser) FetchAndParse(feedURL string, searchTerms []string) ([]models.Item, error) {
	resp, err := p.config.Get(feedURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var rss models.RSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	// Filter items based on search terms AND 1080p
	var matchedItems []models.Item
	for _, item := range rss.Channel.Items {
		title := strings.ToLower(item.Title)
		// Check if title contains 1080p
		if !strings.Contains(title, "1080p") {
			continue
		}
		// Check if title contains any of our search terms
		for _, term := range searchTerms {
			if strings.Contains(title, strings.ToLower(term)) {
				matchedItems = append(matchedItems, item)
				break
			}
		}
	}

	return matchedItems, nil
}
