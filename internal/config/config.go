package config

type Config struct {
	SearchTerms   []string
	DownloadPath  string
	CheckInterval string
	// RSS URL with embedded authentication
	RSSURL string
}

func NewConfig() *Config {
	return &Config{
		SearchTerms:   []string{"Formula1", "UFC"},
		CheckInterval: "0 */12 * * *", // Twice daily
		DownloadPath:  "./downloads",  // Default download path
		RSSURL:        "https://www.torrentday.com/t.rss?7;u=2550949;tp=60241506062941b14d022ed0fabe8e58;GuitarIpod;private;do-not-share",
	}
}
