package main

import (
	"fmt"
	"log"
	"torrent-rss/internal/config"
	"torrent-rss/internal/downloader"
	"torrent-rss/internal/parser"

	"github.com/joho/godotenv"
)

// Define ANSI color codes for a cyberpunk theme
const (
	colorReset      = "\033[0m"
	colorNeonPink   = "\033[1;35m"
	colorNeonBlue   = "\033[1;36m"
	colorNeonGreen  = "\033[1;32m"
	colorNeonYellow = "\033[1;33m"
	colorNeonRed    = "\033[1;31m"
	colorGray       = "\033[1;90m"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("%s⚠️  No .env file found, checking system environment variables...%s\n", colorNeonYellow, colorReset)
	}
}

func main() {
	cfg := config.NewConfig()
	p := parser.NewParser()

	fmt.Printf("%s⚡️>>> Searching for %s《%v》%s matches with %s1080p%s... ⚡️%s\n\n",
		colorNeonBlue, colorNeonPink, cfg.SearchTerms, colorNeonBlue, colorNeonYellow, colorNeonBlue, colorReset)

	matches, err := p.FetchAndParse(cfg.GetRSSURL(), cfg.SearchTerms)
	if err != nil {
		log.Fatalf("%s💀 Error parsing RSS feed: %v 💀%s", colorNeonRed, err, colorReset)
	}

	if len(matches) == 0 {
		fmt.Printf("%s🚫 No matches found! 🚫%s\n", colorNeonRed, colorReset)
		return
	}

	d, err := downloader.NewDownloader(cfg.DownloadPath, cfg.BaseURL, cfg.GetAuthCookie())
	if err != nil {
		log.Fatalf("%s💀 Error creating downloader: %v 💀%s", colorNeonRed, err, colorReset)
	}
	if err != nil {
		log.Fatalf("%s💀 Error creating downloader: %v 💀%s", colorNeonRed, err, colorReset)
	}

	for _, item := range matches {
		// Each match header with distinctive icons
		fmt.Printf("\n%s╔═════════════════════════════════╗%s\n", colorGray, colorReset)
		fmt.Printf("%s⚡️=== Match Found ===⚡️%s\n", colorNeonPink, colorReset)
		fmt.Printf("%sTitle:%s %s%s%s\n", colorNeonYellow, colorReset, colorNeonGreen, item.Title, colorReset)
		fmt.Printf("%sLink:%s %s%s%s\n", colorNeonYellow, colorReset, colorNeonGreen, item.Link, colorReset)
		fmt.Printf("%sDate:%s %s%s%s\n", colorNeonYellow, colorReset, colorNeonGreen, item.PubDate, colorReset)
		fmt.Printf("%sDescription:%s %s%s%s\n", colorNeonYellow, colorReset, colorNeonGreen, item.Description, colorReset)
		fmt.Printf("%s⏬ Downloading torrent file...%s\n", colorNeonBlue, colorReset)

		// Download the torrent
		if err := d.DownloadTorrent(item.Link); err != nil {
			fmt.Printf("%s💀 Error downloading torrent: %v 💀%s\n", colorNeonRed, err, colorReset)
			continue
		}

		// Success message with a futuristic divider
		fmt.Printf("%s✅ Successfully downloaded to:%s %s\n", colorNeonGreen, colorReset, cfg.DownloadPath)
		fmt.Printf("%s╚═════════════════════════════════╝%s\n", colorGray, colorReset)
	}

	// Summary message
	fmt.Printf("\n%s⚡️Total matches found: %s%d%s ⚡️\n", colorNeonYellow, colorNeonBlue, len(matches), colorReset)
}
