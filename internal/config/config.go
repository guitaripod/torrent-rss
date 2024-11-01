package config

import (
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	SearchTerms   []string
	DownloadPath  string
	CheckInterval string
	BaseURL       string
	UserID        string
	RSSToken      string // For RSS feed
	PassToken     string // For downloads
}

func NewConfig() *Config {
	// Get user's home directory for default download path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Could not find home directory")
	}

	// Get download path from env or use default
	downloadPath := os.Getenv("TD_DOWNLOAD_PATH")
	if downloadPath == "" {
		downloadPath = filepath.Join(homeDir, "Downloads", "torrents")
	}

	// Get base URL from environment
	baseURL := os.Getenv("TD_BASE_URL")
	if baseURL == "" {
		panic("TD_BASE_URL environment variable is required")
	}

	// Get search terms from environment variable
	searchTermsEnv := os.Getenv("TD_SEARCH_TERMS")
	if searchTermsEnv == "" {
		panic("TD_SEARCH_TERMS environment variable is required")
	}
	searchTerms := strings.Split(searchTermsEnv, ",")

	// Get check interval from environment
	checkInterval := os.Getenv("TD_CHECK_INTERVAL")
	if checkInterval == "" {
		checkInterval = "0 */12 * * *" // default to every 12 hours
	}

	// Get authentication tokens
	userID := os.Getenv("TD_USER_ID")
	passToken := os.Getenv("TD_TOKEN")
	rssToken := os.Getenv("TD_RSS_TOKEN")

	if userID == "" || passToken == "" || rssToken == "" {
		panic("TD_USER_ID, TD_TOKEN, and TD_RSS_TOKEN environment variables are required")
	}

	return &Config{
		SearchTerms:   searchTerms,
		DownloadPath:  downloadPath,
		CheckInterval: checkInterval,
		BaseURL:       strings.TrimRight(baseURL, "/"),
		UserID:        userID,
		RSSToken:      rssToken,
		PassToken:     passToken,
	}
}

// GetRSSURL constructs the RSS URL using the exact working format
func (c *Config) GetRSSURL() string {
	// TODO: - Improvement area:
	// The number 7 corresponds to the torrent category TV/x264. Every category has a number or a sequence of numbers for multiple categories if the RSS feed is configured as such.
	// anime(29), TV/x264(7)
	return c.BaseURL + "/t.rss?29;7;u=" + c.UserID + ";tp=" + c.RSSToken + ";GuitarIpod;private;do-not-share"
}

// GetAuthCookie returns the cookie string for downloads
func (c *Config) GetAuthCookie() string {
	return "uid=" + c.UserID + "; pass=" + c.PassToken
}
