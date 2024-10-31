# 🌊 Torrent RSS Downloader

[![Go Version](https://img.shields.io/github/go-mod/go-version/marcusziade/torrent-rss?style=flat-square)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](LICENSE)
[![Maintenance](https://img.shields.io/badge/maintained%3F-yes-green.svg?style=flat-square)](https://github.com/marcusziade/torrent-rss/graphs/commit-activity)

A sleek, automated torrent RSS feed monitor and downloader written in Go. This tool automatically checks your private tracker's RSS feed for new torrents matching your search terms and downloads them.

## ✨ Features

- 🔄 Automated RSS feed monitoring
- 🔍 Configurable search terms
- ⚡️ Fast and lightweight
- 🎯 File quality filters
- 🔐 Secure authentication handling
- 📁 Customizable download directory
- ⏰ Configurable check intervals

## 📋 Prerequisites

- Support is only for **TV/x264** files (for now)
- Go 1.19 or higher
- Access to TorrentDay.com (invite-only website)
- RSS feed access on TorrentDay

## 🚀 Installation

```bash
# Clone the repository
git clone https://github.com/username/torrent-rss.git

# Navigate to the project directory
cd torrent-rss

# Install dependencies
go mod download

# Copy the example environment file
cp .env.example .env
```

## ⚙️ Configuration

1. Visit TorrentDay's RSS setup page at `https://www.torrentday.com/rss`

2. Generate your RSS feed and get the following information from the generated RSS URL:
   - User ID
   - RSS Token
   - Download Token

3. Edit your `.env` file with your details:

```env
# Required configuration
TD_BASE_URL=https://www.torrentday.com
TD_USER_ID=your_user_id_here
TD_TOKEN=your_download_token_here
TD_RSS_TOKEN=your_rss_token_here
TD_SEARCH_TERMS=Formula1,UFC

# Optional configuration
TD_CHECK_INTERVAL=0 */12 * * *
TD_DOWNLOAD_PATH=/custom/path/if/needed
```

### Environment Variables Explained

- `TD_BASE_URL`: Base URL of TorrentDay (default provided)
- `TD_USER_ID`: Your TorrentDay user ID (found in RSS feed URL)
- `TD_TOKEN`: Your download authentication token
- `TD_RSS_TOKEN`: Your RSS feed token (found in RSS feed URL)
- `TD_SEARCH_TERMS`: Comma-separated list of terms to search for
- `TD_CHECK_INTERVAL`: How often to check for new torrents (cron format)
- `TD_DOWNLOAD_PATH`: Custom download directory (optional)

## 🎮 Usage

```bash
# Run the program
go run cmd/torrent-rss/main.go
```

The program will:
1. Check your RSS feed for new torrents
2. Filter for items matching your search terms
3. Download matching torrent files to your specified directory

## 🔒 Security Notes

- Keep your `.env` file secure and never commit it to version control
- Your RSS feed URL contains private tokens - never share it
- The program stores sensitive data only in environment variables

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Disclaimer

This tool is for personal use only. Ensure you comply with your tracker's rules and regulations regarding automated downloads and RSS usage.