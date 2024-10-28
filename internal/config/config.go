package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	SearchTerms   []string
	DownloadPath  string
	CheckInterval string
	RSSURL        string
}

func NewConfig() *Config {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Could not find home directory")
	}

	// Set downloads path to ~/Downloads/torrents
	downloadPath := filepath.Join(homeDir, "Downloads", "torrents")

	return &Config{
		SearchTerms:   []string{"Formula1", "UFC"},
		CheckInterval: "0 */12 * * *", // Twice daily
		DownloadPath:  downloadPath,
		RSSURL:        "https://www.torrentday.com/t.rss?7;u=2550949;tp=60241506062941b14d022ed0fabe8e58;GuitarIpod;private;do-not-share",
	}
}
